package response

type ErrorDetail struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Field   string `json:"field"`
}

type ErrorResponse struct {
	Errors []ErrorDetail `json:"errors"`
}
