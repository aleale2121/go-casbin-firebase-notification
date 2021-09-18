package notification

import (
	"template/internal/adapter/storage/persistence/notification"
	"template/internal/constant"
	"template/internal/constant/errors"
	"template/internal/constant/model"
)

// Usecase interface contains function of business logic for domain Notification
type Usecase interface {
	Notifications() (*constant.SuccessData,*errors.ErrorModel)
	CreateNotification(notification model.Notification) (*constant.SuccessData, *errors.ErrorModel)
	DeleteNotification(param model.Notification) (*constant.SuccessData, *errors.ErrorModel)
}
//service defines all necessary service for the domain Notification
type service struct {
	notificationPersistance  notification.NotificationPersistence
}

// Initialize creates a new object with UseCase type
func Initialize(notificationPersistance notification.NotificationPersistence) Usecase {
	return &service{
		notificationPersistance: notificationPersistance,
	}
}
