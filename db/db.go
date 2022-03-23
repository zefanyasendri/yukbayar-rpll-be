package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "root:@tcp(localhost:3306)/yuk_bayar?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Disconnect(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic(err.Error())
	}
	dbSQL.Close()
}
