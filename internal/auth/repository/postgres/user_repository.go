package postgres

import (
	"go-boilerplate/internal/auth/entity"
	"go-boilerplate/pkg/database/orm"

	"gorm.io/gorm"
)

type userRepository struct {
	*orm.BaseRepository[entity.User]
}

func NewUserRepository(db *gorm.DB) entity.UserRepository {
	return &userRepository{
		BaseRepository: &orm.BaseRepository[entity.User]{DB: db},
	}
}
