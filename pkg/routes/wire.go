// +build wireinject

package routes

import (
	"go-boilerplate/pkg/auth"
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
		auth.ControllerSet,
		RouteSet,
	)

	return nil
}
