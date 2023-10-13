package orm

import (
	"fmt"
	"go-boilerplate/pkg/database"
	"log/slog"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connMap sync.Map

// ConnectDB Connect Database
func ConnectDB(driver string) (*gorm.DB, error) {
	dsn, err := database.GetDSN(driver)
	if err != nil {
		return nil, err
	}

	switch driver {
	case "postgres":
		return gorm.Open(postgres.Open(dsn), &gorm.Config{})
	default:
		return nil, fmt.Errorf("unknown database type: %s", driver)
	}
}

// GetDB GetDB
func GetDB(driver string) *gorm.DB {
	if db, ok := connMap.Load(driver); ok {
		return db.(*gorm.DB)
	}

	db, err := ConnectDB(driver)
	if err != nil {
		slog.Error("failed to connect to database", "err", err)
		os.Exit(1)
	}

	connMap.Store(driver, db)

	return db
}
