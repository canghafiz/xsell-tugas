package page

type ProductDetailLayoutRequest struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
	ExceptId  int    `json:"except_id"`
	Limit     int    `json:"limit"`
}
