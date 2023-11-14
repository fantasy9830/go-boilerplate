//go:build wireinject
// +build wireinject

package service

import (
	"go-boilerplate/internal/auth/email"
	"go-boilerplate/internal/auth/entity"
	"go-boilerplate/internal/auth/repository/postgres"

	"github.com/google/wire"
)

var (
	UserServiceSet = wire.NewSet(NewUserService, email.Get, postgres.InitialUserRepository)
)

func InitialUserService() entity.UserService {
	wire.Build(UserServiceSet)

	return nil
}
