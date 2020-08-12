package models

import "fmt"

const (
	ErrorUnkownMessage         = "Unknown error."
	ErrorUnkown                = "UNKNOWN"
	ErrorNotFound              = "NOT_FOUND"
	ErrorBadRequest            = "BAD_REQUEST"
	ErrorUserNotConfirmedEmail = "USER_NOT_CONFIRMED_EMAIL"
	ErrorUserWrongCredentials  = "USER_WRONG_CREDENTIALS"
	Forbidden                  = "Forbidden"
)

var knownStatusCodes map[string]int = map[string]int{
	ErrorUnkown:                500,
	ErrorNotFound:              404,
	ErrorBadRequest:            400,
	ErrorUserNotConfirmedEmail: 400,
	Forbidden:                  401,
}

// AppError Custom error handling
type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (me *AppError) Error() string {
	return fmt.Sprintf("%s %s", me.Code, me.Message)
}

func (me *AppError) GetErrorStatusCode() int {
	statusCode := 500

	if c, ok := knownStatusCodes[me.Code]; ok {
		statusCode = c
	}

	return statusCode
}
