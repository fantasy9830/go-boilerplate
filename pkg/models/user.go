package models

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User user model
type User struct {
	ID          uint        `json:"id,omitempty" gorm:"primary_key"`
	Name        string      `json:"name,omitempty"`
	Username    string      `json:"username,omitempty" gorm:"unique"`
	Password    string      `json:"password,omitempty"`
	Email       string      `json:"email,omitempty"`
	CreatedAt   time.Time   `json:"created_at,omitempty"`
	UpdatedAt   time.Time   `json:"updated_at,omitempty"`
	Roles       Roles       `json:"roles,omitempty" gorm:"many2many:user_roles;"`
	Permissions Permissions `json:"permissions,omitempty" gorm:"many2many:user_permissions;"`
}

// Users Users
type Users []*User

// BeforeCreate 密碼用 bcrypt 儲存起來
func (u *User) BeforeCreate(scope *gorm.Scope) (err error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err == nil {
		scope.SetColumn("Password", hashPassword)
	}

	return
}

// BeforeSave BeforeSave
func (u *User) BeforeSave(scope *gorm.Scope) (err error) {
	if len(u.Username) > 0 {
		scope.SetColumn("Username", strings.ToLower(u.Username))
	}

	if len(u.Email) > 0 {
		scope.SetColumn("Email", strings.ToLower(u.Email))
	}

	return
}
