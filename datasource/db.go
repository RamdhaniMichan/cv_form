package datasource

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

const DbMaxIdleConn = 2
const DbMaxOpenConn = 500
const DbMaxLifetimeConn = 2 * time.Minute

// OpenDB ..
func OpenDB() *gorm.DB {
	var err error
	if db == nil {
		db, err = gorm.Open(os.Getenv("DATABASE"), os.Getenv("CONNECT_DB"))
		if err != nil {
			panic(fmt.Sprintf("failed to connect database: %s", err.Error()))
		}
		db.DB().SetMaxIdleConns(DbMaxIdleConn)
		db.DB().SetMaxOpenConns(DbMaxOpenConn)
		db.DB().SetConnMaxLifetime(DbMaxLifetimeConn)
		return db
	}

	if err = db.DB().Ping(); err != nil {
		db.Close()
		db, err = gorm.Open(os.Getenv("DATABASE"), os.Getenv("CONNECT_DB"))
		if err != nil {
			panic(fmt.Sprintf("failed to connect database: %s", err.Error()))
		}
		db.DB().SetMaxIdleConns(DbMaxIdleConn)
		db.DB().SetMaxOpenConns(DbMaxOpenConn)
		db.DB().SetConnMaxLifetime(DbMaxLifetimeConn)
		return db
	}
	return db
}
