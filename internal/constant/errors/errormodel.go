package errors

import "strconv"

type ErrorModel struct {
	ErrorCode        string `json:"errorCode"`
	ErrorMessage     string `json:"errorMessage"`
	ErrorDescription string `json:"errorDescription"`
}

type ValErrorModel struct {
	ErrorCode string            `json:"errorCode"`
	ValError  map[string]string `json:"validationErrors"`
}

func NewErrorResponse(err error) ErrorModel {
	return ErrorModel{
		ErrorMessage:     err.Error(),
		ErrorDescription: Descriptions[err],
		ErrorCode:        strconv.Itoa(ErrCodes[err]),
	}
}
// <<<<<<< user-fix

// func NewValidationError(err validator.ValidationErrorsTranslations) ValErrorModel {
// 	return ValErrorModel{
// 		ErrorCode: strconv.Itoa(ErrCodes[ErrUnknown]),
// 		ValError:  err,
// 	}
// }

// func (v ValErrorModel) Error() string {
// 	return "i am an error"
// }
// =======
// func GetStatusCode(err error) int {
// 	return StatusCodes[err]
// }
// >>>>>>> main
