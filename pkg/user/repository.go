package user

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

// Find Find
func (r *Repository) Find(id uint) (*models.User, error) {
	user := new(models.User)
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// FindByUsername FindByUsername
func (r *Repository) FindByUsername(username string) (*models.User, error) {
	if len(username) == 0 {
		return nil, ErrUserNotExist
	}

	user := new(models.User)
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
