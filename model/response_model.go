package model

// Response struct for response
// This aims to standardize the error response format.
type Response struct {
	// Code is the http status code
	Code int `json:"code"`
	// Message is the error message
	Message string `json:"message"`
}

// NewResponse function for creating new response with httpCode and msg
func NewResponse(httpCode int, msg string) (int, Response) {
	return httpCode, Response{
		Code:    httpCode,
		Message: msg,
	}

}
