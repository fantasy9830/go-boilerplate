package models

import "time"

// Permission permission model
type Permission struct {
	ID          uint      `json:"id,omitempty" gorm:"primary_key"`
	Name        string    `json:"name,omitempty" gorm:"unique"`
	DisplayName string    `json:"display_name,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	Users       Users     `json:"users,omitempty" gorm:"many2many:user_permissions;"`
	Roles       Roles     `json:"roles,omitempty" gorm:"many2many:role_permissions;"`
}

// Permissions Permissions
type Permissions []*Permission
