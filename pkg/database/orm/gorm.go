package orm

import (
	"fmt"
	"log/slog"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var connStore sync.Map

type Config struct {
	Driver      string
	Host        string
	Port        int
	Username    string
	Password    string
	DBName      string
	TablePrefix string
}

// Get Data Source Name
func (c *Config) DSN() string {
	switch c.Driver {
	case "postgres":
		// host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Taipei
		return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
			c.Host,
			c.Username,
			c.Password,
			c.DBName,
			c.Port,
			"Asia/Taipei",
		)
	default:
		return ""
	}
}

// Connect Connect Database
func (c *Config) Connect() (*gorm.DB, error) {
	dsn := c.DSN()

	opts := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: c.TablePrefix,
		},
	}

	switch c.Driver {
	case "postgres":
		return gorm.Open(postgres.Open(dsn), opts)
	default:
		return nil, fmt.Errorf("unknown database type: %s", c.Driver)
	}
}

// GetDB GetDB
func GetDB(c *Config) *gorm.DB {
	if db, ok := connStore.Load(c); ok {
		return db.(*gorm.DB)
	}

	db, err := c.Connect()
	if err != nil {
		slog.Error("failed to connect to database", "err", err)
		os.Exit(1)
	}

	connStore.Store(c, db)

	if db, ok := connStore.Load(c); ok {
		return db.(*gorm.DB)
	}

	return db
}
