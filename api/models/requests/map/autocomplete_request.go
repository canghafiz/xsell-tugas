package _map

type AutocompleteRequest struct {
	Q     string `json:"q"`
	Limit int    `json:"limit"`
}
