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

		permission := models.Permission{
			ID:          1,
			Name:        "resource1",
			Action:      "GET",
			GuardName:   "API",
			DisplayName: "GET Resource1",
			Description: "permission description...",
		}

		db.Create(&models.User{
			Name:     "Test",
			Username: "test",
			Secret:   "8888",
			Email:    "Test@gmail.com",
			Address:  "Test Address",
			Roles: []models.Role{
				models.Role{
					Name:        "admin",
					GuardName:   "API",
					DisplayName: "Admin",
					Description: "role description...",
					Permissions: []models.Permission{
						permission,
						models.Permission{
							ID:          2,
							Name:        "resource2",
							Action:      "Read",
							GuardName:   "Web",
							DisplayName: "Read Resource2",
							Description: "permission description...",
						},
					},
				},
			},
			Permissions: []models.Permission{
				permission,
			},
		})
	}
}
