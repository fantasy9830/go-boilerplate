// +build wireinject

package auth

import (
	"github.com/google/wire"
)

var (
	RepositorySet = wire.NewSet(NewRepository, NewDB)
	ServiceSet    = wire.NewSet(NewService, wire.Bind(new(IRepository), new(*Repository)), RepositorySet)
	ControllerSet = wire.NewSet(NewController, ServiceSet)
)

// CreateController CreateController
func CreateController() *Controller {
	wire.Build(ControllerSet)

	return nil
}

// CreateService CreateService
func CreateService() *Service {
	wire.Build(ServiceSet)

	return nil
}
