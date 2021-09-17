package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserName   string         `json:"username" binding:"required"`
	Password   string         `json:"password"`
	Phone      string         `json:"phone"`
	FirstName  string         `json:"first_name"`
	MiddleName string         `json:"middle_name"`
	LastName   string         `json:"last_name"`
	Email      string         `json:"email"`
	RoleName   string         `json:"role_name"`
	CreatedAt  time.Time      `json:"created_at,omitempty"`
	UpdatedAt  time.Time      `json:"updated_at,omitempty"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}

type UserCompanyRole struct {
	UserID    uuid.UUID `json:"user_id,omitempty"`
	User      *User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
	CompanyID uuid.UUID `json:"company_id,omitempty"`
	Company   *Company  `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
	RoleName  string    `json:"role_name"`
	Role      *Role     `json:"role,omitempty" gorm:"foreignKey:RoleName"`
}
