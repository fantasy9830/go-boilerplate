package database

import (
	"go-boilerplate/internal/auth/config"
	"go-boilerplate/pkg/database/orm"
	"sync"

	"gorm.io/gorm"
)

var (
	onceORM        sync.Once
	postgresConfig *orm.Config
)

func GetDB() *gorm.DB {
	onceORM.Do(func() {
		postgresConfig = &orm.Config{
			Driver:   "postgres",
			Host:     config.Postgres.Host,
			Port:     config.Postgres.Port,
			Username: config.Postgres.Username,
			Password: config.Postgres.Password,
			DBName:   config.Postgres.DBName,
		}
	})

	return orm.GetDB(postgresConfig)
}
