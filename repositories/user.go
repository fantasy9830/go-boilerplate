package repositories

import (
	"github.com/fantasy9830/go-boilerplate/models"
)

// UserRepository ...
type UserRepository struct{}

// NewUserRepository constructor
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Lookup looks up a user by username.
func (r *UserRepository) Lookup(username string) *models.User {
	user := models.User{}

	db.Where("username = ?", username).First(&user)

	return &user
}

// LookupID looks up a user by userid.
func (r *UserRepository) LookupID(uid uint) *models.User {
	user := models.User{}

	db.First(&user, uid)

	return &user
}
