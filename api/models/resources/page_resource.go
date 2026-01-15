package resources

type PageResource struct {
	PageKey string      `json:"page_key"`
	Data    interface{} `json:"data,omitempty"`
}
