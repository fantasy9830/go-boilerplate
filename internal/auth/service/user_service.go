package service

import (
	"go-boilerplate/internal/auth/entity"
	"go-boilerplate/pkg/net/email"
)

type userService struct {
	email   email.Emailer
	userRep entity.UserRepository
}

func NewUserService(
	email email.Emailer,
	userRep entity.UserRepository,
) entity.UserService {
	return &userService{
		email:   email,
		userRep: userRep,
	}
}
