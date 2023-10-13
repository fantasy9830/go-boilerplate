package postgres

import (
	"go-boilerplate/internal/auth/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	*BaseRepository[entity.User]
}

func NewUserRepository(db *gorm.DB) entity.UserRepository {
	return &userRepository{
		BaseRepository: &BaseRepository[entity.User]{db},
	}
}
