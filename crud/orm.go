package crud

import (
	"github.com/cdfmlr/crud/orm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

type DBType int

const (
	MySQL DBType = iota + 1
	PostgresSQL
	SQLite
)

func getDriver(dbType DBType) orm.DBOpener {
	switch dbType {
	case MySQL:
		return mysql.Open
	case PostgresSQL:
		return postgres.Open
	case SQLite:
		return sqlite.Open
	default:
		log.Fatal("db type not support")
	}
	return nil
}

func ConnectDB(dbType DBType, dsn string) {
	openDriver := getDriver(dbType)
	db, err := gorm.Open(openDriver(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}

func AutoMigrate(models ...interface{}) {
	for _, model := range models {
		// 迁移 schema
		err := DB.AutoMigrate(model)
		if err != nil {
			return
		}
	}
}
