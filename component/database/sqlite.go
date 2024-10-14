package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSqliteGORM() *gorm.DB {
	dsn := "component/database/under_stream.db"
	db, _ := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	return db
}
