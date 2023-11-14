package migration

import (
	"context"
	"go-boilerplate/internal/auth/database"
	"go-boilerplate/internal/auth/entity"
)

func Init(ctx context.Context) (err error) {
	db := database.GetDB().WithContext(ctx)

	return db.AutoMigrate(
		&entity.User{},
	)
}
