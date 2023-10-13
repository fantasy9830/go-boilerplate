//go:build wireinject
// +build wireinject

package http

import (
	"go-boilerplate/internal/auth/service"

	"github.com/google/wire"
)

var (
	UserHandlerSet = wire.NewSet(NewUserHandler, service.InitialUserService)
)

func InitializeRouter() Router {
	wire.Build(
		NewRouter,
		InitialUserHandler,
	)

	return nil
}

func InitialUserHandler() *UserHandler {
	wire.Build(UserHandlerSet)

	return nil
}
