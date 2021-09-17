package model

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Notification struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
    Type      string  `json:"type"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type PushedNotification struct {
	UserID         uuid.UUID `json:"user_id"`
	NotificationID uuid.UUID `json:"notification_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Status         string    `json:"status"`
}
