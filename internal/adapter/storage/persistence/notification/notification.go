package notification

import (
	"gorm.io/gorm"
	"template/internal/constant/errors"
	"template/internal/constant/model"
)

func (n notificationPersistence) Notifications() ([]model.Notification, error) {
	conn := n.conn
	notications := []model.Notification{}

	err := conn.Model(&model.Notification{}).Find(&notications).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrRecordNotFound
		}
		return nil, errors.ErrorUnableToFetch
	}
	return notications, nil
}
func (n notificationPersistence) NotificationByID(parm model.Notification) (*model.Notification, error) {
	conn := n.conn
	notification := &model.Notification{}

	err := conn.Model(&model.Notification{}).Where(&parm).First(notification).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrIDNotFound
		}
		return nil, errors.ErrorUnableToFetch
	}
	return notification, nil
}


func (n notificationPersistence) CreateNotification(notification model.Notification) (*model.Notification, error) {
	conn := n.conn

	err := conn.Model(&model.Notification{}).Create(&notification).Error
	if err != nil {
		if err == gorm.ErrRegistered {
			return nil,errors.ErrorUnableToCreate
		}
		return nil, errors.ErrInvalidRequest
	}
	return &notification, nil
}

func (n notificationPersistence) DeleteNotification(param model.Notification) error {
	conn := n.conn

	err := conn.Model(&model.Notification{}).Where(&param).Delete(&param).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.ErrIDNotFound
		}
		return errors.ErrUnableToDelete
	}
	return nil
}

func (n notificationPersistence) MigrateNotification() error {
	db := n.conn
	err := db.Migrator().AutoMigrate(&model.Notification{})
	if err != nil {
		return err
	}
	return nil
}

