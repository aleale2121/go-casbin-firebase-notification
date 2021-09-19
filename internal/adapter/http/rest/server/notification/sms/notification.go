package sms

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
	"template/internal/constant"
	"template/internal/constant/errors"
	"template/internal/constant/model"
	"template/internal/module/notification/sms"
)
//SmsHandler contains all handler interfaces
type SmsHandler interface {
	MiddleWareValidateSmsMessage(c *gin.Context)
	SendSmsMessage(c *gin.Context)
	GetCountUnreadSMsMessages(c *gin.Context)
}
//smsHandler implements sms servicea and golang validator object
type smsHandler struct {
	smsUseCase        sms.Usecase
	validate            *validator.Validate
}

//NewSmsHandler  initializes notification services and golang validator
func NewSmsHandler(s sms.Usecase, valid *validator.Validate) SmsHandler {
	return &smsHandler{smsUseCase: s, validate:    valid,}
}
//MiddleWareValidateSmsMessage binds sms data SMS struct
func (n smsHandler) MiddleWareValidateSmsMessage(c *gin.Context) {
	sms := model.SMS{}
	err := c.Bind(&sms)
	if err != nil {
		errValue := errors.ErrorModel{
			ErrorCode:        strconv.Itoa(errors.StatusCodes[errors.ErrInvalidRequest]),
			ErrorDescription: errors.Descriptions[errors.ErrInvalidRequest],
			ErrorMessage:     errors.ErrInvalidRequest.Error(),
		}
		constant.ResponseJson(c, errValue, errors.StatusCodes[errors.ErrInvalidRequest])
	}
	errV := constant.StructValidator(sms, n.validate)
	if errV != nil {
		constant.ResponseJson(c, errV, errors.StatusCodes[errors.ErrorUnableToBindJsonToStruct])
	}
	c.Set("x-sms", sms)
	c.Next()
}
//SendSmsMessage  sends sms message to a user via phone number
func (n smsHandler) SendSmsMessage(c *gin.Context) {
	panic("implement me")
}
//GetCountUnreadSMsMessages counts unread sms message
func (n smsHandler) GetCountUnreadSMsMessages(c *gin.Context) {
	panic("implement me")
}
