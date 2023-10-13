// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"github.com/google/wire"
	"go-boilerplate/internal/auth/entity"
	"go-boilerplate/internal/auth/repository/postgres"
	"go-boilerplate/pkg/net/email"
)

// Injectors from wire.go:

func InitialUserService() entity.UserService {
	emailer := email.New()
	userRepository := postgres.InitialUserRepository()
	entityUserService := NewUserService(emailer, userRepository)
	return entityUserService
}

// wire.go:

var (
	UserServiceSet = wire.NewSet(NewUserService, email.New, postgres.InitialUserRepository)
)
