package migration

import (
	"context"
	"go-boilerplate/internal/auth/entity"
	"go-boilerplate/pkg/database/orm"
)

func Init(ctx context.Context) (err error) {
	db := orm.GetDB("postgres").WithContext(ctx)

	return db.AutoMigrate(
		&entity.User{},
	)
}
