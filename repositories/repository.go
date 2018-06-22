package repositories

import (
	"github.com/fantasy9830/go-boilerplate/database"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	db = database.GetDB()
}
