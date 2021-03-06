// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package services

import (
	"github.com/google/wire"
	"go-boilerplate/internal/app/repositories"
)

// Injectors from wire.go:

func CreateAuthService() *AuthService {
	userRepository := repositories.CreateUserRepository()
	authService := NewAuthService(userRepository)
	return authService
}

func CreateUserService() *UserService {
	userRepository := repositories.CreateUserRepository()
	userService := NewUserService(userRepository)
	return userService
}

// wire.go:

var (
	authServiceSet = wire.NewSet(NewAuthService, wire.Bind(new(IUserRepository), new(*repositories.UserRepository)), repositories.CreateUserRepository)
	userServiceSet = wire.NewSet(NewUserService, repositories.CreateUserRepository)
)
