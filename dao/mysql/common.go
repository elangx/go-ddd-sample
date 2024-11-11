package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dbTmp, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = dbTmp
}

func GetDB() *gorm.DB {
	return db
}
