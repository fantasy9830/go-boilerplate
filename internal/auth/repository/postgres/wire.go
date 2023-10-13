//go:build wireinject
// +build wireinject

package postgres

import (
	"go-boilerplate/internal/auth/entity"
	"go-boilerplate/pkg/database/orm"

	"github.com/google/wire"
)

var (
	UserRepositorySet = wire.NewSet(NewUserRepository, orm.GetDB, wire.Value("postgres"))
)

func InitialUserRepository() entity.UserRepository {
	wire.Build(UserRepositorySet)

	return nil
}
