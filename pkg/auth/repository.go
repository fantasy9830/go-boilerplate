package auth

import (
	"errors"
	"go-boilerplate/pkg/models"

	"github.com/jinzhu/gorm"
)

// Error constants
var (
	ErrUserNotExist = errors.New("User Not Exist")
)

// DB DB
type DB *gorm.DB

// NewDB NewDB
func NewDB() DB {
	db := models.GetDB()

	return (DB)(db)
}

// Repository HMI Repository
type Repository struct {
	db *gorm.DB
}

// NewRepository New HMI Repository
func NewRepository(db DB) *Repository {
	return &Repository{
		db: (*gorm.DB)(db),
	}
}

// FindUser FindUser
func (r *Repository) FindUser(id uint) (*models.User, error) {
	user := new(models.User)
	if err := r.db.Preload("Roles").Preload("Roles.Permissions").Preload("Permissions").First(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// FindUserByUsername FindUserByUsername
func (r *Repository) FindUserByUsername(username string) (*models.User, error) {
	if len(username) == 0 {
		return nil, ErrUserNotExist
	}

	user := new(models.User)
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
