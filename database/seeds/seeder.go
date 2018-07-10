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
				models.Permission{
					ID:          4,
					Name:        "resource4",
					Action:      "Read",
					GuardName:   "Web",
					DisplayName: "Read Resource4",
					Description: "permission description...",
				},
			},
		})

		db.Create(&models.User{
			Name:     "Test2",
			Username: "test2",
			Secret:   "8888",
			Email:    "Test2@gmail.com",
			Address:  "Test2 Address",
			Roles: []models.Role{
				models.Role{
					Name:        "test",
					GuardName:   "API",
					DisplayName: "Test2",
					Description: "role description...",
					Permissions: []models.Permission{
						permission,
						models.Permission{
							ID:          3,
							Name:        "resource3",
							Action:      "Read",
							GuardName:   "Web",
							DisplayName: "Read Resource3",
							Description: "permission description...",
						},
					},
				},
			},
		})
	}
}
