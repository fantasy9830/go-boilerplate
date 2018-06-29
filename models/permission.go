package models

import "time"

// Permission permission model
type Permission struct {
	ID          uint `gorm:"primary_key"`
	Name        string
	Action      string
	GuardName   string
	DisplayName string
	Description string
	Users       []User `gorm:"many2many:user_permissions;"`
	Roles       []Role `gorm:"many2many:role_permissions;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
