package notification

import (
	"net/http"
	"template/internal/constant"
	"template/internal/constant/errors"
	"template/internal/constant/model"
)

func (s service) Notifications() (*constant.SuccessData,*errors.ErrorModel){
	data, err := s.notificationPersistance.Notifications()
	if err != nil {
		errorData:=errors.NewErrorResponse(err)
		return nil,&errorData
	}
	return &constant.SuccessData{
		Code: http.StatusOK,
		Data: data,
	}, nil

}

func (s service) CreateNotification(notification model.Notification) (*constant.SuccessData, *errors.ErrorModel) {

	if notification.ApiKey == "" {
		errorData:=errors.NewErrorResponse(errors.ErrInvalidAPIKey)
		return nil, &errorData
	}
	if notification.Token == nil {
		errorData:=errors.NewErrorResponse(errors.ErrInvalidToken)
		return nil, &errorData
	}
	_, err := s.notificationPersistance.NotificationByID(notification)
	if err != nil {
		errorData:=errors.NewErrorResponse(errors.ErrDataAlreayExist)
		return nil, &errorData
	}


	newnotification, err := s.notificationPersistance.CreateNotification(notification)
	if err != nil {
		errorData:=errors.NewErrorResponse(err)
		return nil, &errorData
	}
	return &constant.SuccessData{
		Code: http.StatusOK,
		Data: newnotification,
	}, nil
}


func (s service) DeleteNotification(param model.Notification) (*constant.SuccessData, *errors.ErrorModel) {
	_, err := s.notificationPersistance.NotificationByID(param)
	if err != nil {
		errorData:=errors.NewErrorResponse(err)
			return nil, &errorData
	}
	err = s.notificationPersistance.DeleteNotification(param)
	if err != nil {
		errorData:=errors.NewErrorResponse(err)
		return nil, &errorData
	}
	return &constant.SuccessData{
		Code: http.StatusOK,
		Data: "Notification Deleted",
	}, nil


}




