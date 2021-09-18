package notification

import (
    "gorm.io/gorm"
	"template/internal/constant/model"
)

type NotificationPersistence interface {
	Notifications() ([]model.Notification, error)
	NotificationByID(parm model.Notification) (*model.Notification, error)
	CreateNotification(activity model.Notification) (*model.Notification, error)
	DeleteNotification(param model.Notification) error
	MigrateNotification() error
}

type notificationPersistence struct {
	conn *gorm.DB
}

func NotificationInit(conn *gorm.DB) NotificationPersistence {
	return &notificationPersistence{
		conn: conn,
	}
}