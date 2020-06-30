// +build wireinject

package repositories

import (
	"go-boilerplate/internal/pkg/database"

	"github.com/google/wire"
)

var (
	userRepositorySet = wire.NewSet(NewUserRepository, database.GetDB)
)

// CreateUserRepository CreateUserRepository
func CreateUserRepository() *UserRepository {
	wire.Build(userRepositorySet)

	return nil
}
