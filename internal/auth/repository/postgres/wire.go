//go:build wireinject
// +build wireinject

package postgres

import (
	"go-boilerplate/internal/auth/database"
	"go-boilerplate/internal/auth/entity"

	"github.com/google/wire"
)

var (
	UserRepositorySet = wire.NewSet(NewUserRepository, database.GetDB)
)

func InitialUserRepository() entity.UserRepository {
	wire.Build(UserRepositorySet)

	return nil
}
