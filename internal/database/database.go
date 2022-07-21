package database

import (
	"fmt"

	"github.com/backsoul/groot/cmd/groot/migrations"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db = &gorm.DB{}

func InitializeDb() {
	dsn := "root:password@tcp(mysqldb:3306)/groot?charset=utf8mb4&parseTime=True&loc=Local"
	cb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = cb
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	migrations.All(cb)
}

func DB() *gorm.DB {
	return db
}
