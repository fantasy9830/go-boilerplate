// +build wireinject

package auth

import (
	"go-boilerplate/pkg/user"

	"github.com/google/wire"
)

var (
	ServiceSet    = wire.NewSet(NewService, wire.Bind(new(IRepository), new(*user.Repository)), user.RepositorySet)
	ControllerSet = wire.NewSet(NewController, ServiceSet)
	RouteSet      = wire.NewSet(NewRoute, ControllerSet)
)

// InitRoute InitRoute
func InitRoute() Router {
	wire.Build(RouteSet)

	return nil
}

// CreateService CreateService
func CreateService() *Service {
	wire.Build(ServiceSet)

	return nil
}
