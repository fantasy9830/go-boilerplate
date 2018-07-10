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
func (r *UserRepository) LookupID(userID uint) *models.User {
	user := models.User{}

	db.First(&user, userID)

	return &user
}

// Roles ...
func (r *UserRepository) Roles(userID uint, guardName string) []models.Role {
	user := models.User{}

	db.Preload("Roles", "guard_name = ?", guardName).First(&user, userID)

	return user.Roles
}

// Permissions ...
func (r *UserRepository) Permissions(userID uint, guardName string) []models.Permission {
	user := models.User{}

	db.Preload(
		"Permissions",
		"guard_name = ?",
		guardName,
	).Preload(
		"Roles.Permissions",
		"guard_name = ?",
		guardName,
	).First(&user, userID)

	result := make([]models.Permission, len(user.Permissions))
	copy(result, user.Permissions)
	for _, role := range user.Roles {
		for _, permission := range role.Permissions {
			result = append(result, permission)
		}
	}

	unique := make([]models.Permission, 0, len(result))
	flag := make(map[uint]bool)
	for _, permission := range result {
		if _, ok := flag[permission.ID]; !ok {
			flag[permission.ID] = true
			unique = append(unique, permission)
		}
	}

	return unique
}
