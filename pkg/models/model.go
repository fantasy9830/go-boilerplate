package models

import (
	"fmt"
	"go-boilerplate/pkg/config"

	// mssql
	_ "github.com/jinzhu/gorm/dialects/mssql"
	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Init Init
func Init() (err error) {
	db, err = ConnectDB()
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %v", err)
	}

	db.LogMode(config.App.Debug)

	influx = NewInfluxClient()

	return nil
}
