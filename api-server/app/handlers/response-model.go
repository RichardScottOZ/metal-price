package handlers

// response is a model for the response.
type response struct {
	Metal    string  `json:"metal"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
	Unit     string  `json:"unit"`
}
