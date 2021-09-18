package constant

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
	"template/internal/constant/errors"
)

type SuccessData struct {
	Code int
	Data interface{}
}

func ResponseJson(c *gin.Context, responseData interface{} ,statusCode int) {
	c.JSON(statusCode, responseData)
}
func StructValidator(structName interface{},validate *validator.Validate) *errors.ErrorModel {
	err := validate.Struct(structName)
	if err != nil {
		return &errors.ErrorModel{
			ErrorCode:strconv.Itoa(errors.StatusCodes[errors.ErrorUnableToBindJsonToStruct]),
			ErrorDescription:errors.Descriptions[errors.ErrorUnableToBindJsonToStruct],
			ErrorMessage: errors.ErrorUnableToBindJsonToStruct.Error(),
		}
	}
	return nil
}

func  ValidateVariable(parm interface{},validate *validator.Validate) error {
	errs := validate.Var(parm, "required")
	if errs != nil {
		return errs
	}
	return nil
}
