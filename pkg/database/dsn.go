package database

import (
	"fmt"
	"go-boilerplate/pkg/config"
)

// Get Data Source Name
func GetDSN(driver string) (dsn string, err error) {
	switch driver {
	case "postgres":
		// host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Taipei
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
			config.Postgres.Host,
			config.Postgres.Username,
			config.Postgres.Password,
			config.Postgres.DBName,
			config.Postgres.Port,
			"Asia/Taipei",
		)
	default:
		return "", fmt.Errorf("unknown database type: %s", driver)
	}

	return
}
