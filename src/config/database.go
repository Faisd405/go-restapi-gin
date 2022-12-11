package config

import (
	"github.com/faisd405/go-restapi-gin/src/app/example"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_restapi_gin"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&example.Example{})

	DB = database
}
