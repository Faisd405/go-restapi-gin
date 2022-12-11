package router

import (
	"github.com/faisd405/go-restapi-gin/app/example/examplecontroller"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()

	r.GET("/api/examples", examplecontroller.Index)
	r.GET("/api/example/:id", examplecontroller.Show)
	r.POST("/api/example", examplecontroller.Create)
	r.PUT("/api/example/:id", examplecontroller.Update)
	r.DELETE("/api/example", examplecontroller.Delete)

	return r
}
