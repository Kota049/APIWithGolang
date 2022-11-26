package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDb() (*gorm.DB, error) {
	dsn := "root:password@tcp(mysql:3306)/api_with_golang"
	db, err := gorm.Open(mysql.Open(dsn))

	return db, err
}
