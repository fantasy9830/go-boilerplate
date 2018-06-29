package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User user model
type User struct {
	ID          uint `gorm:"primary_key"`
	Name        string
	Username    string `gorm:"unique"`
	Secret      string `gorm:"unique"`
	Email       string `gorm:"type:varchar(100);unique"`
	Address     string
	Roles       []Role       `gorm:"many2many:user_roles;"`
	Permissions []Permission `gorm:"many2many:user_permissions;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// BeforeSave 密碼用bcrypt儲存起來
func (user *User) BeforeSave(scope *gorm.Scope) (err error) {
	if pw, err := bcrypt.GenerateFromPassword([]byte(user.Secret), bcrypt.DefaultCost); err == nil {
		scope.SetColumn("Secret", pw)
	}

	return
}
