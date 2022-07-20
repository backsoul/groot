package services

import (
	"fmt"

	"github.com/backsoul/groot/cmd/groot/migrations"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDb() {
	dsn := "root:password@tcp(mysqldb:3306)/groot?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	fmt.Println("Paso")
	migrations.All(db)
}
