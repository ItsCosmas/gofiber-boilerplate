package controllers

// Response object as HTTP response
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"body"`
}

// HTTPResponse normalize HTTP Response format
func HTTPResponse(httpCode int, message string, data interface{}) *Response {
	return &Response{
		Code:    httpCode,
		Message: message,
		Data:    data,
	}
}
