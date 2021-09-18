package notification

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	 "github.com/satori/go.uuid"
	"strconv"
	"template/internal/constant"
	"template/internal/constant/errors"
	"template/internal/constant/model"
	"template/internal/module/notification"
)

type NotificationHandler interface {
	MiddleWareValidateNotification(c *gin.Context)
	GetNotifications(c *gin.Context)
	CreateNotification(c *gin.Context)
	DeleteNotification(c *gin.Context)
}
type notificationHandler struct {
	notificationUseCase notification.Usecase
	validate            *validator.Validate
}

func NewRoleHandler(notfCase notification.Usecase, valid *validator.Validate) NotificationHandler {
	return &notificationHandler{
		notificationUseCase: notfCase,
		validate:            valid,
	}
}
func (n notificationHandler) MiddleWareValidateNotification(c *gin.Context) {
	notification := model.Notification{}
	err := c.Bind(&notification)
	if err != nil {
		errValue:=errors.ErrorModel{
			ErrorCode:strconv.Itoa(errors.StatusCodes[errors.ErrInvalidRequest]),
			ErrorDescription:errors.Descriptions[errors.ErrInvalidRequest],
			ErrorMessage: errors.ErrInvalidRequest.Error(),
		}
	    constant.ResponseJson(c,errValue,errors.StatusCodes[errors.ErrInvalidRequest])
	}
	errV := constant.StructValidator(notification, n.validate)
	if errV != nil {
		constant.ResponseJson(c,errV,errors.StatusCodes[errors.ErrorUnableToBindJsonToStruct])
	}
	c.Set("x-notification", notification)
	c.Next()
}

func (n notificationHandler) GetNotifications(c *gin.Context) {
	successData, errData := n.notificationUseCase.Notifications()
	if errData != nil {
		code,err:=strconv.Atoi(errData.ErrorCode)
		if err != nil {
			errValue:=errors.ErrorModel{
				ErrorCode:strconv.Itoa(errors.StatusCodes[errors.ErrorUnableToConvert]),
				ErrorDescription:errors.Descriptions[errors.ErrorUnableToConvert],
				ErrorMessage: errors.ErrorUnableToConvert.Error(),
			}
			constant.ResponseJson(c,errValue,errors.StatusCodes[errors.ErrorUnableToConvert])
		}
		constant.ResponseJson(c, *errData,code)
	}
	constant.ResponseJson(c, *successData,successData.Code)

}
func (n notificationHandler) CreateNotification(c *gin.Context) {
	notification := c.MustGet("x-notification").(model.Notification)
	successData, errData := n.notificationUseCase.CreateNotification(notification)
	if errData != nil {
		code,err:=strconv.Atoi(errData.ErrorCode)
		if err != nil {
			errValue:=errors.ErrorModel{
				ErrorCode:strconv.Itoa(errors.StatusCodes[errors.ErrorUnableToConvert]),
				ErrorDescription:errors.Descriptions[errors.ErrorUnableToConvert],
				ErrorMessage: errors.ErrorUnableToConvert.Error(),
			}
			constant.ResponseJson(c,errValue,errors.StatusCodes[errors.ErrorUnableToConvert])
		}
		constant.ResponseJson(c, *errData,code)
	}
	// TODO:01 push notification code put here
	// TODO:02 store push notification here
	constant.ResponseJson(c, *successData,successData.Code)

}

func (n notificationHandler) DeleteNotification(c *gin.Context) {

	id := c.Param("id")
	u_id,err:=uuid.FromString(id)
	if err != nil {
		errValue:=errors.ErrorModel{
			ErrorCode:strconv.Itoa(errors.StatusCodes[errors.ErrorUnableToConvert]),
			ErrorDescription:errors.Descriptions[errors.ErrorUnableToConvert],
			ErrorMessage: errors.ErrorUnableToConvert.Error(),
		}
		constant.ResponseJson(c,errValue,errors.StatusCodes[errors.ErrorUnableToConvert])
	}



	err = constant.ValidateVariable(u_id, n.validate)
	if err != nil {
		errValue:=errors.ErrorModel{
			ErrorCode:strconv.Itoa(errors.StatusCodes[errors.ErrInvalidVariable]),
			ErrorDescription:errors.Descriptions[errors.ErrInvalidVariable],
			ErrorMessage: errors.ErrInvalidVariable.Error(),
		}
		constant.ResponseJson(c,errValue,errors.StatusCodes[errors.ErrInvalidVariable])
	}

	successData, errData := n.notificationUseCase.DeleteNotification(model.Notification{ID: u_id})
	if errData != nil {
		code,err:=strconv.Atoi(errData.ErrorCode)
		if err != nil {
			errValue:=errors.ErrorModel{
				ErrorCode:strconv.Itoa(errors.StatusCodes[errors.ErrorUnableToConvert]),
				ErrorDescription:errors.Descriptions[errors.ErrorUnableToConvert],
				ErrorMessage: errors.ErrorUnableToConvert.Error(),
			}
			constant.ResponseJson(c,errValue,code)
		}
	}
	constant.ResponseJson(c, *successData,successData.Code)
}



