package models

const ErrorTypeValidation = 1
const ErrorTypeUnauthorized = 2
const ErrorTypeError = 2

type LoginRequest struct {
	UserName string
	Password string
}

type ErrorDetail struct {
	ErrorType    int
	ErrorMessage string
}

type Response struct {
	Error   []ErrorDetail
	Data    interface{}
	Status  int
	Message string
}
