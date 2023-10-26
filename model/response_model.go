package model

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewResponse(httpCode int, msg string) (int, Response) {
	return httpCode, Response{
		Code:    httpCode,
		Message: msg,
	}

}
