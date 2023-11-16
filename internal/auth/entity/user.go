package entity

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Password string

func (p Password) Hash() ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
}

func (p Password) Verify(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p), []byte(password)) == nil
}

func (p Password) Empty() bool {
	return len(p) == 0
}

type User struct {
	ID              uint       `json:"id" gorm:"primaryKey"`
	Username        string     `json:"username" gorm:"uniqueIndex:idx_user_username;type:varchar(64)"`
	Password        Password   `json:"-" gorm:"type:varchar(72)"`
	Fullname        string     `json:"fullname" gorm:"type:varchar(64)"`
	Description     string     `json:"description"`
	Email           string     `json:"email" gorm:"type:varchar(320)"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

type Users []*User

type UserService interface{}

type UserRepository interface {
	FindFirst(ctx context.Context, field string, value any) (*User, error)
}

func (u User) ValidatePassword(password string) bool {
	if !u.Password.Empty() && len(password) > 0 {
		return u.Password.Verify(password)
	}

	return false
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	if len(u.Username) > 0 {
		tx.Statement.SetColumn("Username", strings.ToLower(u.Username))
	}

	if len(u.Email) > 0 {
		tx.Statement.SetColumn("Email", strings.ToLower(u.Email))
	}

	if !u.Password.Empty() {
		cost, err := bcrypt.Cost([]byte(u.Password))
		var password []byte
		if cost == 0 || err != nil {
			password, err = u.Password.Hash()
			if err != nil {
				return err
			}
		} else {
			password = []byte(u.Password)
		}

		tx.Statement.SetColumn("Password", password)
	}

	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Password.Empty() {
		return errors.New("password is empty")
	}

	cost, err := bcrypt.Cost([]byte(u.Password))
	var password []byte
	if cost == 0 || err != nil {
		password, err = u.Password.Hash()
		if err != nil {
			return err
		}
	} else {
		password = []byte(u.Password)
	}

	tx.Statement.SetColumn("Password", password)

	return nil
}

// MarshalBinary MarshalBinary
func (u User) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}

// UnmarshalBinary UnmarshalBinary
func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
