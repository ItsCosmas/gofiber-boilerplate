package controllers

import "github.com/gofiber/fiber/v2"

// Response object as HTTP response
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"body"`
}

// ErrorResponse object
type ErrorResponse struct {
	Error []*fiber.Error `json:"errors"`
}

// HTTPResponse normalize HTTP Response format
func HTTPResponse(httpCode int, message string, data interface{}) *Response {
	return &Response{
		Code:    httpCode,
		Message: message,
		Data:    data,
	}
}

// HTTPErrorResponse normalizes error responses
func HTTPErrorResponse(errorObj []*fiber.Error) *ErrorResponse {
	return &ErrorResponse{
		Error: errorObj,
	}
}
