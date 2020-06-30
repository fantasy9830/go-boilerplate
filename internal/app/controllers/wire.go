// +build wireinject

package controllers

import (
	"go-boilerplate/internal/app/services"

	"github.com/google/wire"
)

var (
	authControllerSet = wire.NewSet(NewAuthController, services.CreateAuthService)
	userControllerSet = wire.NewSet(NewUserController, services.CreateUserService)
)

// CreateAuthController CreateAuthController
func CreateAuthController() *AuthController {
	wire.Build(authControllerSet)

	return nil
}

// CreateUserController CreateUserController
func CreateUserController() *UserController {
	wire.Build(userControllerSet)

	return nil
}
