// +build wireinject

package user

import (
	"go-boilerplate/pkg/models"

	"github.com/google/wire"
)

var (
	RepositorySet = wire.NewSet(NewRepository, models.GetDB)
)

// CreateService CreateService
func CreateRepository() *Repository {
	wire.Build(RepositorySet)

	return nil
}
