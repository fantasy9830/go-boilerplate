//go:build wireinject
// +build wireinject

package service

import (
	"go-boilerplate/internal/auth/entity"
	"go-boilerplate/internal/auth/repository/postgres"
	"go-boilerplate/pkg/net/email"

	"github.com/google/wire"
)

var (
	UserServiceSet = wire.NewSet(NewUserService, email.New, postgres.InitialUserRepository)
)

func InitialUserService() entity.UserService {
	wire.Build(UserServiceSet)

	return nil
}
