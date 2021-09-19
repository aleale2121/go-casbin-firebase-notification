package email

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
	"template/internal/constant"
	"template/internal/constant/errors"
	"template/internal/constant/model"
	"template/internal/module/notification/email"
)
//EmailHandler contains all email handler interfaces
type EmailHandler interface {
	MiddleWareValidateEmailMessage(c *gin.Context)
	SendEmailMessage(c *gin.Context) 
	GetCountUnreadEmailMessages(c *gin.Context)
}
//emailHandler implements notification service and golang validator object
type emailHandler struct {
	notificationUseCase email.Usecase
	validate            *validator.Validate
}
//NewEmailHandler  initializes notification services and golang validator
func NewEmailHandler(em email.Usecase, valid *validator.Validate) EmailHandler {
	return &emailHandler{
		notificationUseCase: em,
		validate:            valid,
	}
}
//MiddleWareValidateEmailMessage binds pushed notification data as json
func (n emailHandler) MiddleWareValidateEmailMessage(c *gin.Context) {
	email := model.EmailNotification{}
	err := c.Bind(&email)
	if err != nil {
		errValue := errors.ErrorModel{
			ErrorCode:        strconv.Itoa(errors.StatusCodes[errors.ErrInvalidRequest]),
			ErrorDescription: errors.Descriptions[errors.ErrInvalidRequest],
			ErrorMessage:     errors.ErrInvalidRequest.Error(),
		}
		constant.ResponseJson(c, errValue, errors.StatusCodes[errors.ErrInvalidRequest])
	}
	errV := constant.StructValidator(email, n.validate)
	if errV != nil {
		constant.ResponseJson(c, errV, errors.StatusCodes[errors.ErrorUnableToBindJsonToStruct])
	}
	c.Set("x-email", email)
	c.Next()
}
//SendEmailMessage send email message via valid email
func (n emailHandler) SendEmailMessage(c *gin.Context) {
	panic("implement me")
}
//GetCountUnreadEmailMessages return the number of unread message
func (n emailHandler) GetCountUnreadEmailMessages(c *gin.Context) {
	panic("implement me")
}
