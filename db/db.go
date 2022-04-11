package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Connect() *gorm.DB {
	dsn := "root:@tcp(localhost:3306)/yuk_bayar?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
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
