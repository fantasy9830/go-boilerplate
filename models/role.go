package models

import (
	"time"
)

// Role role model
type Role struct {
	ID          uint `gorm:"primary_key"`
	Name        string
	GuardName   string
	DisplayName string
	Description string
	Users       []User       `gorm:"many2many:user_roles;"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
