package repositories

import (
	"errors"
	"go-boilerplate/internal/app/models"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Error constants
var (
	ErrUserNotExist = errors.New("User Not Exist")
)

// UserRepository User Repository
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository New User Repository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Find Find
func (r *UserRepository) Find(id uint) (*models.User, error) {
	user := new(models.User)
	if err := r.db.Preload("Roles").Preload("Roles.Permissions").Preload("Permissions").First(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// FindByUsername FindByUsername
func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	if len(username) == 0 {
		return nil, ErrUserNotExist
	}

	user := new(models.User)
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Create Create
func (r *UserRepository) Create(user models.User) (*models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Update Update
func (r *UserRepository) Update(id uint, user models.User) (*models.User, error) {
	u, err := r.Find(id)
	if err != nil {
		return nil, err
	}

	if user.Password != "" {
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashPassword)
	}

	r.db.Model(u).Updates(user)

	return u, nil
}
