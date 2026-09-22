package core_http_response

type ErrorResponse struct {
	Error string `json:"error" exmaple:"full error text"`
	Message string `json:"message" example:"short human-readable text"`
}