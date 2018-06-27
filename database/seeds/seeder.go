package seeds

import (
	"github.com/fantasy9830/go-boilerplate/database"
	"github.com/fantasy9830/go-boilerplate/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	db = database.GetDB()
}

// Run run the seeds.
func Run() {
	if db.HasTable(&models.User{}) {
		db.Create(&models.User{
			Name:     "Test",
			Username: "admin",
			Secret:   "8888",
			Email:    "Test@gmail.com",
			Address:  "Test Address",
		})
	}
}
