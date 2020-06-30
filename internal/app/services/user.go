package services

import "go-boilerplate/internal/app/repositories"

// UserService UserService
type UserService struct {
	rep *repositories.UserRepository
}

// NewUserService New User Service
func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{
		rep: repository,
	}
}
