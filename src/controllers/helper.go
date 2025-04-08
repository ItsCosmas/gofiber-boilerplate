package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// Stores Errors
var errorList []*Response

// Response object as HTTP response
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"body"`
}

// ErrorBody object
type ErrorBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// FiberErrorResponse object
type FiberErrorResponse struct {
	Error []*ErrorBody `json:"errors"`
}

// ErrorResponse object
type ErrorResponse struct {
	Error []*Response `json:"errors"`
}

// HTTPResponse normalize HTTP Response format
func HTTPResponse(httpCode int, message string, data interface{}) *Response {
	return &Response{
		Code:    httpCode,
		Message: message,
		Data:    data,
	}
}

// HTTPFiberErrorResponse normalizes error responses
func HTTPFiberErrorResponse(errorObj []*fiber.Error) *FiberErrorResponse {
	// Convert fiber.Error to ErrorBody
	// This fixes issues with swagger auto generated docs not identify fiber.Error type
	var errorSlice []*ErrorBody
	for i := 0; i < len(errorObj); i++ {
		errorSlice = append(errorSlice, mapToErrorOutput(errorObj[i]))
	}

	return &FiberErrorResponse{
		Error: errorSlice,
	}
}

// HTTPErrorResponse normalizes error responses
func HTTPErrorResponse(errorObj []*Response) *ErrorResponse {
	var errorSlice []*Response
	for i := 0; i < len(errorObj); i++ {
		errorSlice = append(errorSlice, errorObj[i])
	}

	return &ErrorResponse{
		Error: errorSlice,
	}
}

// ==================================== //
// Private Method
func mapToErrorOutput(e *fiber.Error) *ErrorBody {
	return &ErrorBody{
		Code:    e.Code,
		Message: e.Message,
	}
}
