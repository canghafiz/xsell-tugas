package impl

import (
	"be/models/domains"
	"be/models/requests/member/product"
	"errors"
	"fmt"
	"log"
	"sort"

	geo "github.com/kellydunn/golang-geo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepoMemberImpl struct{}

func NewProductRepoMemberImpl() *ProductRepoMemberImpl {
	return &ProductRepoMemberImpl{}
}

func (repo *ProductRepoMemberImpl) Create(db *gorm.DB, product domains.Products) error {
	return db.Create(&product).Error
}

func (repo *ProductRepoMemberImpl) Update(db *gorm.DB, product domains.Products) error {
	if product.ProductId == 0 {
		return gorm.ErrInvalidValue
	}

	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Recover from panic and rollback transaction
	defer func() {
		if r := recover(); r != nil {
			// Log the panic for debugging (replace with your logger if needed)
			log.Printf("Panic during product update (ID: %d): %v", product.ProductId, r)
			tx.Rollback()
		}
	}()

	// Check if product exists
	var exists bool
	if err := tx.Model(&domains.Products{}).
		Select("1").
		Where("product_id = ?", product.ProductId).
		Limit(1).
		Scan(&exists).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to check product existence: %w", err)
	}
	if !exists {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	// --- 1. Update Product Fields (excluding slug) ---
	updates := make(map[string]interface{})

	if product.Title != "" {
		updates["title"] = product.Title
		// ⚠️ DO NOT UPDATE product_slug — it should be immutable after creation for SEO
	}
	if product.Description != "" {
		updates["description"] = product.Description
	}
	if product.Price != 0 {
		updates["price"] = product.Price
	}
	if product.Condition != "" {
		updates["condition"] = product.Condition
	}
	if product.Status != "" {
		updates["status"] = product.Status
	}
	if product.SubCategoryId != 0 {
		updates["sub_category_id"] = product.SubCategoryId
	}

	if len(updates) > 0 {
		if err := tx.Model(&domains.Products{}).
			Where("product_id = ?", product.ProductId).
			Updates(updates).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update product fields: %w", err)
		}
	}

	// --- 2. Handle Images: Full Replacement ---
	if product.ProductImages != nil {
		// Delete all existing images
		if err := tx.Where("product_id = ?", product.ProductId).Delete(&domains.ProductImages{}).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to delete old product images: %w", err)
		}

		// Insert new images if any
		if len(product.ProductImages) > 0 {
			for i := range product.ProductImages {
				product.ProductImages[i].ProductId = product.ProductId
				product.ProductImages[i].ImageId = 0 // Let DB assign new ID
			}
			if err := tx.Create(&product.ProductImages).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to create new product images: %w", err)
			}
		}
	}

	// --- 3. Handle Specs: Partial UPSERT ---
	if product.ProductSpecs != nil {
		for _, spec := range product.ProductSpecs {
			// Skip empty spec values if your business logic requires it
			if spec.SpecValue == "" {
				continue
			}
			spec.ProductId = product.ProductId
			spec.ProductSpecId = 0 // Ensure new ID isn't used

			err := tx.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "product_id"}, {Name: "category_product_spec_id"}},
				DoUpdates: clause.Assignments(map[string]interface{}{"spec_value": spec.SpecValue}),
			}).Create(&spec).Error

			if err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to upsert product spec: %w", err)
			}
		}
	}

	// --- 4. Handle Location: Upsert only if location data is provided ---
	// Compare against zero value to detect if location was intentionally sent
	emptyLocation := domains.Location{}
	if product.Location != emptyLocation {
		var existingLocation domains.Location
		err := tx.Where("product_id = ?", product.ProductId).First(&existingLocation).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Insert new location
			product.Location.ProductId = product.ProductId
			if err := tx.Create(&product.Location).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to create product location: %w", err)
			}
		} else if err != nil {
			// Unexpected DB error
			tx.Rollback()
			return fmt.Errorf("failed to fetch product location: %w", err)
		} else {
			// Update existing location
			updatesLoc := make(map[string]interface{})
			if product.Location.Latitude != 0 {
				updatesLoc["latitude"] = product.Location.Latitude
			}
			if product.Location.Longitude != 0 {
				updatesLoc["longitude"] = product.Location.Longitude
			}
			if product.Location.Address != "" {
				updatesLoc["address"] = product.Location.Address
			}

			if len(updatesLoc) > 0 {
				if err := tx.Model(&existingLocation).Updates(updatesLoc).Error; err != nil {
					tx.Rollback()
					return fmt.Errorf("failed to update product location: %w", err)
				}
			}
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit product update transaction: %w", err)
	}

	return nil
}

func (repo *ProductRepoMemberImpl) GetSingleBySlug(db *gorm.DB, slug string) (*domains.Products, error) {
	if slug == "" {
		return nil, gorm.ErrInvalidValue
	}

	var product domains.Products
	err := db.Preload("SubCategory.Category").
		Preload("ProductImages").
		Preload("ProductSpecs.CategoryProductSpec").
		Preload("Location").
		Preload("Listing").
		Where("product_slug = ?", slug).
		First(&product).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}

func (repo *ProductRepoMemberImpl) GetRelatedByCategory(
	db *gorm.DB,
	categoryIds []int,
	exceptProductId, limit int,
	lat, lng float64,
) ([]domains.Products, error) {

	const maxRadiusKm = 100.0

	// Return empty if no categories or invalid location
	if len(categoryIds) == 0 || (lat == 0 && lng == 0) {
		return []domains.Products{}, nil
	}

	// 1️⃣ Fetch products under the given main categories
	var products []domains.Products
	err := db.
		Joins("JOIN subcategories sc ON sc.sub_category_id = products.sub_category_id").
		Where("sc.parent_category_id IN ?", categoryIds).
		Where("products.product_id != ?", exceptProductId).
		Where("products.status = ?", "Available").
		Preload("ProductImages").
		Preload("Location").
		Preload("Listing").
		Find(&products).Error

	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return []domains.Products{}, nil
	}

	// 2️⃣ Filter by 100km radius
	userPoint := geo.NewPoint(lat, lng)
	filtered := make([]domains.Products, 0, len(products))
	distances := make(map[int]float64)

	for _, p := range products {
		if p.Location.Latitude == 0 && p.Location.Longitude == 0 {
			continue
		}
		productPoint := geo.NewPoint(p.Location.Latitude, p.Location.Longitude)
		distance := userPoint.GreatCircleDistance(productPoint)
		if distance <= maxRadiusKm {
			filtered = append(filtered, p)
			distances[p.ProductId] = distance
		}
	}

	if len(filtered) == 0 {
		return []domains.Products{}, nil
	}

	// 3️⃣ Sort by nearest
	sort.SliceStable(filtered, func(i, j int) bool {
		return distances[filtered[i].ProductId] < distances[filtered[j].ProductId]
	})

	// 4️⃣ Apply limit
	if len(filtered) > limit {
		filtered = filtered[:limit]
	}

	return filtered, nil
}

func (repo *ProductRepoMemberImpl) GetByCategory(db *gorm.DB, filter product.FilterModel) ([]domains.Products, error) {
	const maxRadiusKm = 100.0

	// Return empty if lat/lng not set
	if filter.Lat == 0 && filter.Lng == 0 {
		return []domains.Products{}, nil
	}

	// Apply default values for price range
	minPrice := filter.MinPrice
	maxPrice := filter.MaxPrice
	if minPrice < 0 {
		minPrice = 0
	}
	if maxPrice <= 0 {
		maxPrice = 999999999
	}

	// Apply default limit (max 100, default 20)
	limit := filter.Limit
	if limit <= 0 || limit > 100 {
		limit = 20
	}

	// Apply default sort
	sortBy := filter.SortBy
	if sortBy == "" {
		sortBy = "latest"
	}

	// Step 1: Fetch parent category
	var parentCategory domains.Categories
	if err := db.Where("category_slug = ?", filter.CategorySlug).First(&parentCategory).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []domains.Products{}, nil
		}
		return nil, err
	}

	// Step 2: Determine subcategories
	var subCategoryIds []int
	if len(filter.SubCategorySlug) > 0 {
		if err := db.Model(&domains.SubCategories{}).
			Where("parent_category_id = ?", parentCategory.CategoryId).
			Where("sub_category_slug IN ?", filter.SubCategorySlug).
			Pluck("sub_category_id", &subCategoryIds).Error; err != nil {
			return nil, err
		}
		if len(subCategoryIds) == 0 {
			return []domains.Products{}, nil
		}
	} else {
		if err := db.Model(&domains.SubCategories{}).
			Where("parent_category_id = ?", parentCategory.CategoryId).
			Pluck("sub_category_id", &subCategoryIds).Error; err != nil {
			return nil, err
		}
		if len(subCategoryIds) == 0 {
			return []domains.Products{}, nil
		}
	}

	// Step 3: Fetch products (with buffer for radius filtering)
	var products []domains.Products
	query := db.Where("sub_category_id IN ?", subCategoryIds).
		Where("price >= ? AND price <= ?", minPrice, maxPrice).
		Where("status = ?", "Available").
		Preload("SubCategory").
		Preload("ProductImages").
		Preload("Location").
		Preload("Listing")

	// Apply sorting from request — this will be the FINAL order
	switch sortBy {
	case "price_asc":
		query = query.Order("products.price ASC")
	case "price_desc":
		query = query.Order("products.price DESC")
	case "oldest":
		query = query.Order("products.created_at ASC")
	case "latest":
		query = query.Order("products.created_at DESC")
	default:
		query = query.Order("products.created_at DESC")
	}

	// Fetch extra to compensate for radius filtering
	query = query.Limit(limit * 3)

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return []domains.Products{}, nil
	}

	// Step 4: FILTER by 100km radius (do NOT change order!)
	userPoint := geo.NewPoint(filter.Lat, filter.Lng)
	filtered := make([]domains.Products, 0, len(products))

	for _, p := range products {
		if p.Location.Latitude == 0 && p.Location.Longitude == 0 {
			continue
		}
		productPoint := geo.NewPoint(p.Location.Latitude, p.Location.Longitude)
		distance := userPoint.GreatCircleDistance(productPoint)
		if distance <= maxRadiusKm {
			filtered = append(filtered, p)
			// Stop early if we already have enough items
			if len(filtered) >= limit {
				break
			}
		}
	}

	// Truncate to final limit (preserve original order from DB)
	if len(filtered) > limit {
		filtered = filtered[:limit]
	}

	return filtered, nil
}

func (repo *ProductRepoMemberImpl) GetBySectionKey(
	db *gorm.DB,
	key string,
	model product.FilterModel,
) ([]domains.Products, error) {

	const maxRadiusKm = 100.0

	if key == "" {
		return []domains.Products{}, nil
	}

	// Fetch section
	var section domains.ContentSection
	if err := db.
		Where("section_key = ? AND is_active = true", key).
		First(&section).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []domains.Products{}, nil
		}
		return nil, err
	}

	cfg := section.Config

	// Set limit
	limit := model.Limit
	if limit <= 0 {
		limit = cfg.Limit
	}
	if limit <= 0 || limit > 100 {
		limit = 12
	}

	// Base query
	query := db.
		Model(&domains.Products{}).
		Where("products.status = ?", "Available").
		Preload("ProductImages").
		Preload("Location").
		Preload("Listing")

	needSubCategoryJoin := model.CategorySlug != "" || len(model.SubCategorySlug) > 0

	if needSubCategoryJoin {
		query = query.Joins("JOIN subcategories sc ON sc.sub_category_id = products.sub_category_id")
	}

	if model.CategorySlug != "" {
		query = query.
			Joins("JOIN categories c ON c.category_id = sc.parent_category_id").
			Where("c.category_slug = ?", model.CategorySlug)
	}

	if len(model.SubCategorySlug) > 0 {
		query = query.Where("sc.sub_category_slug IN ?", model.SubCategorySlug)
	}

	// Price filter
	if model.MinPrice > 0 {
		query = query.Where("products.price >= ?", model.MinPrice)
	}
	if model.MaxPrice > 0 && model.MaxPrice < 999999999 {
		query = query.Where("products.price <= ?", model.MaxPrice)
	}

	// Section config filters
	if cfg.Filters != nil {
		if v, ok := cfg.Filters["status"].(string); ok && v != "" {
			query = query.Where("products.status = ?", v)
		}
		if v, ok := cfg.Filters["condition"].(string); ok && v != "" {
			query = query.Where("products.condition = ?", v)
		}
	}

	// === SORTING LOGIC (FINAL VERSION) ===
	if model.SortBy == "" || model.SortBy == "default" {
		// No user sort → use section strategy
		switch cfg.Strategy {
		case "trending":
			query = query.Order("products.view_count DESC")
		case "newly_listed":
			query = query.Order("products.created_at DESC")
		default:
			query = query.Order("products.created_at DESC")
		}
	} else {
		// User explicitly chose a sort → honor it
		switch model.SortBy {
		case "price_asc":
			query = query.Order("products.price ASC")
		case "price_desc":
			query = query.Order("products.price DESC")
		case "oldest":
			query = query.Order("products.created_at ASC")
		case "latest":
			query = query.Order("products.created_at DESC")
		default:
			query = query.Order("products.created_at DESC")
		}
	}

	// Fetch with buffer
	query = query.Limit(limit * 2)

	var products []domains.Products
	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return []domains.Products{}, nil
	}

	// Radius filter (skip for trending/newly_listed)
	useRadius :=
		model.Lat != 0 &&
			model.Lng != 0 &&
			key != "newly_listed" &&
			key != "trending"

	if !useRadius {
		if len(products) > limit {
			products = products[:limit]
		}
		return products, nil
	}

	userPoint := geo.NewPoint(model.Lat, model.Lng)
	filtered := make([]domains.Products, 0, len(products))
	distances := make(map[int]float64)

	for _, p := range products {
		if p.Location.Latitude == 0 || p.Location.Longitude == 0 {
			continue
		}

		productPoint := geo.NewPoint(p.Location.Latitude, p.Location.Longitude)
		d := userPoint.GreatCircleDistance(productPoint)

		if d <= maxRadiusKm {
			filtered = append(filtered, p)
			distances[p.ProductId] = d
		}
	}

	if len(filtered) == 0 {
		return []domains.Products{}, nil
	}

	sort.SliceStable(filtered, func(i, j int) bool {
		return distances[filtered[i].ProductId] < distances[filtered[j].ProductId]
	})

	if len(filtered) > limit {
		filtered = filtered[:limit]
	}

	return filtered, nil
}

func (repo *ProductRepoMemberImpl) Delete(
	db *gorm.DB,
	product domains.Products,
) ([]string, error) {

	if product.ProductId == 0 {
		return nil, gorm.ErrInvalidValue
	}

	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// 1. Cek product exist
	var count int64
	if err := tx.Model(&domains.Products{}).
		Where("product_id = ?", product.ProductId).
		Count(&count).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if count == 0 {
		tx.Rollback()
		return nil, gorm.ErrRecordNotFound
	}

	// 2. Ambil image urls
	var images []domains.ProductImages
	if err := tx.
		Where("product_id = ?", product.ProductId).
		Find(&images).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	imageUrls := make([]string, 0, len(images))
	for _, img := range images {
		imageUrls = append(imageUrls, img.ImageUrl)
	}

	// 3. Delete product
	if err := tx.Delete(&domains.Products{}, product.ProductId).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// 4. Commit
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return imageUrls, nil
}

func (repo *ProductRepoMemberImpl) Search(db *gorm.DB, title string, filter *product.FilterModel) ([]domains.Products, error) {
	const maxRadiusKm = 100.0

	// Return empty if lat/lng not set
	if filter.Lat == 0 && filter.Lng == 0 {
		return []domains.Products{}, nil
	}

	// Apply default values for price range
	minPrice := filter.MinPrice
	maxPrice := filter.MaxPrice
	if minPrice < 0 {
		minPrice = 0
	}
	if maxPrice <= 0 {
		maxPrice = 999999999
	}

	// Apply default limit (max 100, default 20)
	limit := filter.Limit
	if limit <= 0 || limit > 100 {
		limit = 20
	}

	// Apply default sort
	sortBy := filter.SortBy
	if sortBy == "" {
		sortBy = "latest"
	}

	// Step 1: Build base query with title search
	query := db.Model(&domains.Products{}).
		Where("title ILIKE ?", "%"+title+"%").
		Where("price >= ? AND price <= ?", minPrice, maxPrice).
		Where("status = ?", "Available")

	// Step 2: Apply category filter if provided
	if filter.CategorySlug != "" {
		var parentCategory domains.Categories
		if err := db.Where("category_slug = ?", filter.CategorySlug).First(&parentCategory).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return []domains.Products{}, nil
			}
			return nil, err
		}

		// Determine subcategories
		var subCategoryIds []int
		if len(filter.SubCategorySlug) > 0 {
			if err := db.Model(&domains.SubCategories{}).
				Where("parent_category_id = ?", parentCategory.CategoryId).
				Where("sub_category_slug IN ?", filter.SubCategorySlug).
				Pluck("sub_category_id", &subCategoryIds).Error; err != nil {
				return nil, err
			}
			if len(subCategoryIds) == 0 {
				return []domains.Products{}, nil
			}
		} else {
			if err := db.Model(&domains.SubCategories{}).
				Where("parent_category_id = ?", parentCategory.CategoryId).
				Pluck("sub_category_id", &subCategoryIds).Error; err != nil {
				return nil, err
			}
			if len(subCategoryIds) == 0 {
				return []domains.Products{}, nil
			}
		}

		query = query.Where("sub_category_id IN ?", subCategoryIds)
	}

	// Step 3: Apply sorting
	switch sortBy {
	case "price_asc":
		query = query.Order("products.price ASC")
	case "price_desc":
		query = query.Order("products.price DESC")
	case "oldest":
		query = query.Order("products.created_at ASC")
	case "latest":
		query = query.Order("products.created_at DESC")
	default:
		query = query.Order("products.created_at DESC")
	}

	// Step 4: Preload associations and fetch with buffer for radius filtering
	query = query.
		Preload("ProductImages").
		Preload("Location").
		Preload("Listing").
		Limit(limit * 3)

	var products []domains.Products
	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return []domains.Products{}, nil
	}

	// Step 5: Filter by 100km radius (preserve order from DB)
	userPoint := geo.NewPoint(filter.Lat, filter.Lng)
	filtered := make([]domains.Products, 0, len(products))

	for _, p := range products {
		if p.Location.Latitude == 0 && p.Location.Longitude == 0 {
			continue
		}
		productPoint := geo.NewPoint(p.Location.Latitude, p.Location.Longitude)
		distance := userPoint.GreatCircleDistance(productPoint)
		if distance <= maxRadiusKm {
			filtered = append(filtered, p)
			// Stop early if we already have enough items
			if len(filtered) >= limit {
				break
			}
		}
	}

	// Truncate to final limit (preserve original order from DB)
	if len(filtered) > limit {
		filtered = filtered[:limit]
	}

	return filtered, nil
}

func (repo *ProductRepoMemberImpl) GetProductsByUserId(
	db *gorm.DB,
	userId int,
	filter product.FilterMyAds,
) ([]domains.ProductWithWishlist, error) {

	var results []domains.ProductWithWishlist

	// base query
	query := db.
		Table("products p").
		Select(`
			p.*,
			COUNT(w.product_id) AS total_wishlist
		`).
		Joins("LEFT JOIN wishlists w ON w.product_id = p.product_id").
		Where("p.listing_user_id = ?", userId).
		Group("p.product_id")

	// sorting
	switch filter.SortBy {
	case "oldest_to_new":
		query = query.Order("p.created_at ASC")
	case "most_liked":
		query = query.Order("total_wishlist DESC, p.created_at DESC")
	default:
		query = query.Order("p.created_at DESC")
	}

	// execute
	if err := query.Scan(&results).Error; err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return []domains.ProductWithWishlist{}, nil
	}

	// collect product IDs
	productIDs := make([]int, 0, len(results))
	for _, r := range results {
		productIDs = append(productIDs, r.Product.ProductId)
	}

	// preload images
	var products []domains.Products
	if err := db.
		Preload("ProductImages").
		Where("product_id IN ?", productIDs).
		Find(&products).Error; err != nil {
		return nil, err
	}

	// mapping
	productMap := make(map[int]domains.Products)
	for _, p := range products {
		productMap[p.ProductId] = p
	}

	for i := range results {
		if p, ok := productMap[results[i].Product.ProductId]; ok {
			results[i].Product = p
		}
	}

	return results, nil
}

func (repo *ProductRepoMemberImpl) UpdateViewCount(db *gorm.DB, productId int) error {
	err := db.
		Model(&domains.Products{}).
		Where("product_id = ?", productId).
		Updates(map[string]interface{}{
			"view_count": gorm.Expr("view_count + 1"),
		}).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *ProductRepoMemberImpl) UpdateStatus(db *gorm.DB, product domains.Products) error {
	err := db.
		Model(&domains.Products{}).
		Where("product_id = ?", product.ProductId).
		Updates(map[string]interface{}{
			"status": product.Status,
		}).Error

	if err != nil {
		return err
	}

	return nil
}
