package product

type FilterModel struct {
	CategorySlug                 string
	SubCategorySlug              []string
	SortBy                       string
	MinPrice, MaxPrice, Lat, Lng float64
	Limit                        int
}
