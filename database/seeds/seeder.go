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

		// API
		permission1 := models.Permission{
			ID:          1,
			Name:        "resource1",
			Action:      "GET",
			GuardName:   "API",
			DisplayName: "GET Resource1",
			Description: "permission description...",
		}

		// Web
		permission2 := models.Permission{
			ID:          2,
			Name:        "resource2",
			Action:      "read",
			GuardName:   "web",
			DisplayName: "Read Resource2",
			Description: "permission description...",
		}
		permission3 := models.Permission{
			ID:          3,
			Name:        "resource3",
			Action:      "read",
			GuardName:   "web",
			DisplayName: "Read Resource3",
			Description: "permission description...",
		}
		permission4 := models.Permission{
			ID:          4,
			Name:        "sidermenu1-1",
			Action:      "read",
			GuardName:   "web",
			DisplayName: "Read sidermenu1-1",
			Description: "SiderMenu1-1...",
		}
		permission5 := models.Permission{
			ID:          5,
			Name:        "sidermenu1-2",
			Action:      "read",
			GuardName:   "web",
			DisplayName: "Read sidermenu1-2",
			Description: "SiderMenu1-2...",
		}
		permission6 := models.Permission{
			ID:          6,
			Name:        "sidermenu1-3-1",
			Action:      "read",
			GuardName:   "web",
			DisplayName: "Read sidermenu1-3-1",
			Description: "SiderMenu1-3-1...",
		}
		permission7 := models.Permission{
			ID:          7,
			Name:        "sidermenu1-3-2",
			Action:      "read",
			GuardName:   "web",
			DisplayName: "Read sidermenu1-3-2",
			Description: "SiderMenu1-3-2...",
		}
		permission8 := models.Permission{
			ID:          8,
			Name:        "sidermenu2-1",
			Action:      "read",
			GuardName:   "web",
			DisplayName: "Read sidermenu2-1",
			Description: "SiderMenu2-1...",
		}
		permission9 := models.Permission{
			ID:          9,
			Name:        "sidermenu2-2",
			Action:      "read",
			GuardName:   "web",
			DisplayName: "Read sidermenu2-2",
			Description: "SiderMenu2-2...",
		}
		permission10 := models.Permission{
			ID:          10,
			Name:        "sidermenu2-3",
			Action:      "read",
			GuardName:   "web",
			DisplayName: "Read sidermenu2-3",
			Description: "SiderMenu2-3...",
		}
		permission11 := models.Permission{
			ID:          11,
			Name:        "sidermenu3",
			Action:      "read",
			GuardName:   "web",
			DisplayName: "Read sidermenu3",
			Description: "SiderMenu3...",
		}
		permission12 := models.Permission{
			ID:          12,
			Name:        "sidermenu1",
			Action:      "read",
			GuardName:   "web",
			DisplayName: "Read sidermenu1",
			Description: "SiderMenu1...",
		}
		permission13 := models.Permission{
			ID:          13,
			Name:        "sidermenu1-3",
			Action:      "read",
			GuardName:   "web",
			DisplayName: "Read sidermenu1-3",
			Description: "SiderMenu1-3...",
		}
		permission14 := models.Permission{
			ID:          14,
			Name:        "sidermenu2",
			Action:      "read",
			GuardName:   "web",
			DisplayName: "Read sidermenu2",
			Description: "SiderMenu2...",
		}

		db.Create(&models.User{
			Name:     "Demo",
			Username: "demo",
			Secret:   "8888",
			Email:    "Demo@gmail.com",
			Address:  "Demo Address",
			Roles: []models.Role{
				models.Role{
					Name:        "admin",
					GuardName:   "web",
					DisplayName: "Admin",
					Description: "role description...",
					Permissions: []models.Permission{
						permission1,
						permission2,
					},
				},
			},
			Permissions: []models.Permission{
				permission1,
				permission4,
				permission5,
				permission6,
				permission7,
				permission8,
				permission9,
				permission10,
				permission11,
				permission12,
				permission13,
				permission14,
			},
		})

		db.Create(&models.User{
			Name:     "Demo2",
			Username: "demo2",
			Secret:   "8888",
			Email:    "Demo2@gmail.com",
			Address:  "Demo2 Address",
			Roles: []models.Role{
				models.Role{
					Name:        "test",
					GuardName:   "web",
					DisplayName: "Test2",
					Description: "role description...",
					Permissions: []models.Permission{
						permission1,
						permission3,
						permission4,
						permission5,
						permission6,
						permission7,
						permission12,
						permission13,
					},
				},
			},
		})
	}
}
