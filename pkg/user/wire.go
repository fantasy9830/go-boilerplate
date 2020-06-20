// +build wireinject

package user

import (
	"github.com/google/wire"
)

var (
	RepositorySet = wire.NewSet(NewRepository, NewDB)
)

// CreateService CreateService
func CreateRepository() *Repository {
	wire.Build(RepositorySet)

	return nil
}
