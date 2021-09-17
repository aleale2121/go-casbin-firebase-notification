package errors

import "strconv"

type ErrorModel struct {
	ErrorCode        string `json:"errorCode"`
	ErrorMessage     string `json:"errorMessage"`
	ErrorDescription string `json:"errorDescription"`
}
func NewErrorResponse(err error) ErrorModel {
	return ErrorModel{
		ErrorMessage:       err.Error(),
		ErrorDescription: Descriptions[err],
		ErrorCode:  strconv.Itoa(ErrCodes[err]),
	}
}
	