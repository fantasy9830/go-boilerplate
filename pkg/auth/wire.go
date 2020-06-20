// +build wireinject

package auth

import (
	"go-boilerplate/pkg/user"

	"github.com/google/wire"
)

var (
	ServiceSet    = wire.NewSet(NewService, wire.Bind(new(IRepository), new(*user.Repository)), user.RepositorySet)
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
