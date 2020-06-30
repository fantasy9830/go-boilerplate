// +build wireinject

package services

import (
	"go-boilerplate/internal/app/repositories"

	"github.com/google/wire"
)

var (
	authServiceSet = wire.NewSet(NewAuthService, wire.Bind(new(IUserRepository), new(*repositories.UserRepository)), repositories.CreateUserRepository)
	userServiceSet = wire.NewSet(NewUserService, repositories.CreateUserRepository)
)

// CreateAuthService CreateAuthService
func CreateAuthService() *AuthService {
	wire.Build(authServiceSet)

	return nil
}

// CreateUserService CreateUserService
func CreateUserService() *UserService {
	wire.Build(userServiceSet)

	return nil
}
