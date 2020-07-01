// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package repositories

import (
	"github.com/google/wire"
	"go-boilerplate/internal/pkg/database"
)

// Injectors from wire.go:

func CreateUserRepository() *UserRepository {
	db := database.GetDB()
	userRepository := NewUserRepository(db)
	return userRepository
}

// wire.go:

var (
	userRepositorySet = wire.NewSet(NewUserRepository, database.GetDB)
)