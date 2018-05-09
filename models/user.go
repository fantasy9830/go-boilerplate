package models

import (
	"github.com/jinzhu/gorm"
)

// User user model
type User struct {
	gorm.Model
	Name string
}
