// +build wireinject

package routes

import (
	"go-boilerplate/internal/app/controllers"
	"net/http"

	"github.com/google/wire"
)

var (
	RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))
	RouteSet  = wire.NewSet(NewRoute, RouterSet)
)

// InitRoute InitRoute
func InitRoute() http.Handler {
	wire.Build(
		controllers.CreateAuthController,
		controllers.CreateUserController,
		RouteSet,
	)

	return nil
}
