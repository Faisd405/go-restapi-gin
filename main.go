package main

import (
	"github.com/faisd405/go-restapi-gin/src/config"
	"github.com/faisd405/go-restapi-gin/src/router"
	// "github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := router.Routes()

	r.Run()
}
