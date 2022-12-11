package config

import (
	example "github.com/faisd405/go-restapi-gin/src/app/example/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/golang"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&example.Example{})

	DB = database
}
