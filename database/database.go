package database

import (
	"bytes"
	"strings"
	"sync"

	"github.com/fantasy9830/go-boilerplate/config"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/jinzhu/gorm/dialects/mssql"
)

// DB連線資料
type connections map[string]string

// TODO: mysql "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
func (con *connections) mysql() string {
	return ""
}

// postgres "host=myhost port=myport user=gorm dbname=gorm password=mypassword"
func (con *connections) postgres() string {
	var buf bytes.Buffer

	for key, value := range *con {
		if key != "driver" && value != "" {
			buf.WriteString(key + "=" + value + " ")
		}
	}

	return strings.TrimSpace(buf.String())
}

// TODO: sqlite "/tmp/gorm.db"
func (con *connections) sqlite() string {
	return ""
}

// TODO: mssql "sqlserver://username:password@localhost:1433?database=dbname"
func (con *connections) mssql() string {
	return ""
}

var (
	dbs  = make(map[string]*gorm.DB)
	once = make(map[string]*sync.Once)
)

func init() {
	c := config.GetConfig()

	for dbName := range c.GetStringMap("connections") {
		if dbName != "default" && once[dbName] == nil {
			once[dbName] = new(sync.Once)
		}
	}
}

func connectDB(dbName string) func() {
	return func() {
		var (
			con connections
			dsn string
		)

		con = config.GetConfig().GetStringMapString("connections." + dbName)

		switch con["driver"] {
		case "mysql":
			dsn = con.mysql()
		case "postgres":
			dsn = con.postgres()
		case "sqlite":
			dsn = con.sqlite()
		case "mssql":
			dsn = con.mssql()
		}

		dbs[dbName], _ = gorm.Open(con["driver"], dsn)
	}
}

// GetDB gets the global db instance.
func GetDB(dbNames ...string) *gorm.DB {
	var dbName string

	if len(dbNames) > 0 {
		dbName = dbNames[0]
	} else {
		dbName = config.GetConfig().GetString("connections.default")
	}

	if once[dbName] != nil {
		once[dbName].Do(connectDB(dbName))
	}

	return dbs[dbName]
}
