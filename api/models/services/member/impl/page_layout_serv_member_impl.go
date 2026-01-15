package impl

import (
	"be/helpers"
	"be/models/domains"
	"be/models/repositories"
	"be/models/requests/member/page"
	"be/models/resources"
	"be/models/resources/section"
	"be/models/services"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/kellydunn/golang-geo"
	"gorm.io/gorm"
)

type PageLayoutServMemberImpl struct {
	Db             *gorm.DB
	RedisServ      services.RedisService
	PageLayoutRepo repositories.PageLayoutRepo
}

func NewPageLayoutServMemberImpl(
	db *gorm.DB,
	redisServ services.RedisService,
	pageLayoutRepo repositories.PageLayoutRepo,
) *PageLayoutServMemberImpl {
	return &PageLayoutServMemberImpl{
		Db:             db,
		RedisServ:      redisServ,
		PageLayoutRepo: pageLayoutRepo,
	}
}

/* ================= PUBLIC METHODS ================= */

// GetHomeLayouts fetches the home page layout with dynamic sections.
// It supports optional user location (latitude/longitude) for geo-aware sections.
func (serv *PageLayoutServMemberImpl) GetHomeLayouts(
	request page.HomeLayoutRequest,
) (resources.PageResource, error) {
	cacheKey := fmt.Sprintf("page:home:lat:%s:lng:%s:limit:%d", request.Latitude, request.Longitude, request.Limit)

	// Attempt to retrieve from cache
	var cached resources.PageResource
	cachedData, err := helpers.GetFromCache(serv.RedisServ.GetData, cacheKey, &cached)
	if err == nil && cachedData != nil {
		if result, ok := cachedData.(resources.PageResource); ok {
			return result, nil
		}
	}

	// Fetch fresh data from database
	pageRes, err := serv.getPageLayout("home", 0, request.Limit, request.Latitude, request.Longitude)
	if err != nil {
		return resources.PageResource{}, err
	}

	// Cache the result for 60 seconds
	if errCache := helpers.SetToCache(serv.RedisServ.SetData, cacheKey, pageRes, 60); errCache != nil {
		log.Printf("Failed to cache home layouts: %v", errCache)
	}

	return pageRes, nil
}

// GetProductDetailLayouts fetches the product detail page layout.
// It excludes the current product and supports location-based recommendations.
func (serv *PageLayoutServMemberImpl) GetProductDetailLayouts(
	request page.ProductDetailLayoutRequest,
) (resources.PageResource, error) {
	cacheKey := fmt.Sprintf("page:product_detail:exceptId:%d:lat:%s:lng:%s:limit:%d",
		request.ExceptId, request.Latitude, request.Longitude, request.Limit)

	// Attempt to retrieve from cache
	var cached resources.PageResource
	cachedData, err := helpers.GetFromCache(serv.RedisServ.GetData, cacheKey, &cached)
	if err == nil && cachedData != nil {
		if result, ok := cachedData.(resources.PageResource); ok {
			return result, nil
		}
	}

	// Fetch fresh data from database
	pageRes, errRes := serv.getPageLayout("product_detail", request.ExceptId, request.Limit, request.Latitude, request.Longitude)
	if errRes != nil {
		return resources.PageResource{}, errRes
	}

	// Cache the result for 60 seconds
	if errCache := helpers.SetToCache(serv.RedisServ.SetData, cacheKey, pageRes, 60); errCache != nil {
		log.Printf("Failed to cache product_detail layouts: %v", errCache)
	}

	return pageRes, nil
}

/* ================= CORE LOGIC ================= */

// getPageLayout resolves all sections for a given page key (e.g., "home", "product_detail").
// It applies user location (if provided) and enforces per-section product limits.
func (serv *PageLayoutServMemberImpl) getPageLayout(
	pageKey string,
	exceptId int,
	fallbackLimit int,
	latStr string,
	lngStr string,
) (resources.PageResource, error) {

	var userLat, userLng float64
	if latStr != "" && lngStr != "" {
		if lat, err1 := strconv.ParseFloat(latStr, 64); err1 == nil {
			userLat = lat
		}
		if lng, err2 := strconv.ParseFloat(lngStr, 64); err2 == nil {
			userLng = lng
		}
	}

	// Load page layout configuration
	layouts, err := serv.PageLayoutRepo.GetPageLayout(serv.Db, pageKey)
	if err != nil {
		return resources.PageResource{}, err
	}

	sections := make([]section.Resource, 0, len(layouts))

	for _, layout := range layouts {
		if layout.Section == nil {
			continue
		}

		// Determine per-section limit: prefer config, fallback to request, then default
		limitToUse := fallbackLimit
		if layout.Section.Config.Limit > 0 {
			limitToUse = layout.Section.Config.Limit
		}
		if limitToUse <= 0 {
			limitToUse = 10 // final safety fallback
		}

		// Resolve products based on section configuration and user context
		products, err := serv.getProductsByConfig(
			layout.Section.Config,
			limitToUse,
			userLat,
			userLng,
			exceptId,
		)
		if err != nil {
			log.Printf("Failed to get products for section %s: %v", layout.Section.SectionKey, err)
			// Continue rendering other sections even if one fails
			products = []domains.Products{}
		}

		sections = append(
			sections,
			*section.ToResource(*layout.Section, products),
		)
	}

	return resources.PageResource{
		PageKey: pageKey,
		Data:    sections,
	}, nil
}

/* ================= PRODUCT RESOLUTION ================= */

// getProductsByConfig resolves products based on section configuration, strategy, and user context.
func (serv *PageLayoutServMemberImpl) getProductsByConfig(
	config domains.ContentConfig,
	overrideLimit int,
	lat, lng float64,
	exceptId int,
) ([]domains.Products, error) {

	query := serv.Db.Model(&domains.Products{}).
		Where("products.status = ?", domains.Available).
		Preload("ProductImages").
		Preload("SubCategory").
		Preload("SubCategory.Category").
		Preload("Location").
		Preload("Listing")

	// Exclude a specific product (e.g., current product on detail page)
	if exceptId > 0 {
		query = query.Where("products.product_id != ?", exceptId)
	}

	// STEP 1: Enable radius filtering if valid coordinates are provided
	useRadius := lat != 0 && lng != 0
	if useRadius {
		query = query.
			Joins("INNER JOIN location loc ON loc.product_id = products.product_id").
			Where("loc.latitude != 0 AND loc.longitude != 0")
	}

	// STEP 2: Apply static filters from section config (condition, category, price range)
	if config.Filters != nil {
		if v, ok := config.Filters["condition"].(string); ok && v != "" {
			query = query.Where("products.condition = ?", v)
		}
		if v, ok := config.Filters["category_slug"].(string); ok && v != "" {
			query = query.
				Joins("JOIN subcategories sc ON sc.sub_category_id = products.sub_category_id").
				Joins("JOIN categories c ON c.category_id = sc.parent_category_id").
				Where("c.category_slug = ?", v)
		}
		if v, ok := config.Filters["max_price"].(float64); ok && v > 0 {
			query = query.Where("products.price <= ?", v)
		}
		if v, ok := config.Filters["min_price"].(float64); ok && v > 0 {
			query = query.Where("products.price >= ?", v)
		}
	}

	var products []domains.Products
	strategy := strings.TrimSpace(config.Strategy)

	// STEP 3: Apply sorting strategy
	switch strategy {
	case "fixed_ids":
		if len(config.ProductIDs) > 0 {
			query = query.Where("products.product_id IN ?", config.ProductIDs)
			query = query.Limit(overrideLimit)
			if err := query.Find(&products).Error; err != nil {
				return nil, err
			}
			// Apply radius filtering if location is provided
			if useRadius {
				return serv.filterByRadius(products, lat, lng, overrideLimit), nil
			}
			return products, nil
		}
		// Fall back to newly listed if no product IDs are defined
		fallthrough

	case "newly_listed", "":
		query = query.Order("products.created_at DESC")

	case "trending":
		query = query.Order("products.view_count DESC")

	case "nearby":
		// For "nearby", sort by distance (closest first)
		if useRadius {
			query = query.Limit(overrideLimit * 3) // Fetch extra for radius filtering
			if err := query.Find(&products).Error; err != nil {
				return nil, err
			}
			return serv.sortByDistanceWithLimit(products, lat, lng, overrideLimit), nil
		}
		// Fallback to newest if no location is provided
		query = query.Order("products.created_at DESC")

	default:
		// Unknown strategy → default to newest
		query = query.Order("products.created_at DESC")
	}

	// STEP 4: Fetch with buffer if radius filtering will be applied later
	fetchLimit := overrideLimit
	if useRadius {
		fetchLimit = overrideLimit * 3
	}
	query = query.Limit(fetchLimit)

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	// STEP 5: Apply radius filtering for non-"nearby" strategies (preserve original order)
	if useRadius && strategy != "nearby" {
		return serv.filterByRadius(products, lat, lng, overrideLimit), nil
	}

	// Truncate to final limit if no radius filtering was applied
	if len(products) > overrideLimit {
		products = products[:overrideLimit]
	}
	return products, nil
}

/* ================= GEO-SPATIAL FILTERING ================= */

// filterByRadius filters products within 100km of the user's location.
// It preserves the original order of the input slice.
func (serv *PageLayoutServMemberImpl) filterByRadius(
	products []domains.Products,
	lat, lng float64,
	limit int,
) []domains.Products {

	const maxRadiusKm = 100.0

	if len(products) == 0 {
		return []domains.Products{}
	}

	userPoint := geo.NewPoint(lat, lng)
	var filtered []domains.Products

	for _, p := range products {
		// Skip products without valid coordinates
		if p.Location.Latitude == 0 && p.Location.Longitude == 0 {
			continue
		}

		productPoint := geo.NewPoint(p.Location.Latitude, p.Location.Longitude)
		distance := userPoint.GreatCircleDistance(productPoint)

		if distance <= maxRadiusKm {
			filtered = append(filtered, p)
			if len(filtered) >= limit {
				break // Early exit once limit is reached
			}
		}
	}

	return filtered
}

// sortByDistanceWithLimit filters and sorts products by distance (closest first).
// Used exclusively for the "nearby" strategy.
func (serv *PageLayoutServMemberImpl) sortByDistanceWithLimit(
	products []domains.Products,
	lat, lng float64,
	limit int,
) []domains.Products {

	const maxRadiusKm = 100.0

	if len(products) == 0 {
		return []domains.Products{}
	}

	userPoint := geo.NewPoint(lat, lng)
	type prodWithDist struct {
		Product  domains.Products
		Distance float64
	}
	var filtered []prodWithDist

	for _, p := range products {
		if p.Location.Latitude == 0 && p.Location.Longitude == 0 {
			continue
		}

		productPoint := geo.NewPoint(p.Location.Latitude, p.Location.Longitude)
		distance := userPoint.GreatCircleDistance(productPoint)

		if distance <= maxRadiusKm {
			filtered = append(filtered, prodWithDist{Product: p, Distance: distance})
		}
	}

	if len(filtered) == 0 {
		return []domains.Products{}
	}

	// Sort by distance (ascending)
	sort.SliceStable(filtered, func(i, j int) bool {
		return filtered[i].Distance < filtered[j].Distance
	})

	// Apply final limit
	if len(filtered) > limit {
		filtered = filtered[:limit]
	}

	// Extract products only
	result := make([]domains.Products, len(filtered))
	for i, f := range filtered {
		result[i] = f.Product
	}

	return result
}

// findNearestProducts is a legacy helper (not currently used in main flow).
// It returns products within 100km, sorted by distance, with a distance map.
func (serv *PageLayoutServMemberImpl) findNearestProducts(
	lat, lng float64,
	limit int,
	exceptId int,
) ([]domains.Products, map[int]float64, error) {

	const maxRadiusKm = 100.0

	var products []domains.Products
	query := serv.Db.
		Model(&domains.Products{}).
		Joins("INNER JOIN location loc ON loc.product_id = products.product_id").
		Where("products.status = ?", domains.Available).
		Where("loc.latitude != 0 AND loc.longitude != 0").
		Preload("ProductImages").
		Preload("SubCategory").
		Preload("SubCategory.Category").
		Preload("Location").
		Preload("Listing")

	if exceptId > 0 {
		query = query.Where("products.product_id != ?", exceptId)
	}

	query = query.Limit(limit * 3)

	if err := query.Find(&products).Error; err != nil {
		return nil, nil, err
	}

	if len(products) == 0 {
		return []domains.Products{}, nil, nil
	}

	userPoint := geo.NewPoint(lat, lng)
	type prodWithDist struct {
		Product  domains.Products
		Distance float64
	}
	var filtered []prodWithDist

	for _, p := range products {
		if p.Location.Latitude == 0 && p.Location.Longitude == 0 {
			continue
		}

		productPoint := geo.NewPoint(p.Location.Latitude, p.Location.Longitude)
		d := userPoint.GreatCircleDistance(productPoint)

		if d <= maxRadiusKm {
			filtered = append(filtered, prodWithDist{Product: p, Distance: d})
		}
	}

	if len(filtered) == 0 {
		return []domains.Products{}, nil, nil
	}

	sort.SliceStable(filtered, func(i, j int) bool {
		return filtered[i].Distance < filtered[j].Distance
	})

	if len(filtered) > limit {
		filtered = filtered[:limit]
	}

	distMap := make(map[int]float64)
	result := make([]domains.Products, len(filtered))
	for i, f := range filtered {
		result[i] = f.Product
		distMap[f.Product.ProductId] = f.Distance
	}

	return result, distMap, nil
}
