package router

import (
	examplecontroller "github.com/faisd405/go-restapi-gin/src/app/example/controller"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()

	r.GET("/api/examples", examplecontroller.Index)
	r.GET("/api/example/:id", examplecontroller.Show)
	r.POST("/api/example", examplecontroller.Create)
	r.PUT("/api/example/:id", examplecontroller.Update)
	r.DELETE("/api/example/:id", examplecontroller.Delete)

	return r
}
