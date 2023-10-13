package http

import (
	"go-boilerplate/internal/auth/entity"
)

type UserHandler struct {
	userService entity.UserService
}

func NewUserHandler(userService entity.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}
