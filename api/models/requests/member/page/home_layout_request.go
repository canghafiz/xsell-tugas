package page

type HomeLayoutRequest struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
	Limit     int    `json:"limit"`
}
