package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	// d, err := gorm.Open("mysql", "root:root/simplerest?charset=urf&parseTime=True&loc=local")
	d, err := gorm.Open("mysql", "root:root@/simplerest?charset=utf8&parseTime=True&loc=UTC")

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
