package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGORM() *gorm.DB {
	if true { // todo move to env
		return InitSqliteGORM()
	}
	dsn := ""
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db
}
