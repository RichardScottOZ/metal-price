package handlers

// Response is a model for the response.
type Response struct {
	// metal element
	Metal string `json:"metal" example:"rhodium"`
	// value
	Price float64 `json:"price" example:"8200"`
	// money system
	Currency string `json:"currency" example:"USD"`
	// weight unit
	Unit string `json:"unit" example:"oz"`
}

// HTTPError is a model for the error response.
type HTTPError struct {
	// error response
	Message string `json:"message" example:"call metal service: material ssilver not found"`
}
