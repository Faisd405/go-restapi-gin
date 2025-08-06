package router

import (
	examplecontroller "github.com/faisd405/go-restapi-gin/src/app/example/controller"
	usercontroller "github.com/faisd405/go-restapi-gin/src/app/user/controller"
	userrepository "github.com/faisd405/go-restapi-gin/src/app/user/repository"
	userservice "github.com/faisd405/go-restapi-gin/src/app/user/service"
	"github.com/faisd405/go-restapi-gin/src/config"
	"github.com/faisd405/go-restapi-gin/src/middleware"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()

	// Add middleware
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.LoggerMiddleware())

	// Initialize user dependencies
	userRepo := userrepository.NewUserRepository(config.GetDB())
	userSvc := userservice.NewUserService(userRepo)
	userCtrl := usercontroller.NewUserController(userSvc)

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// Auth routes (public)
		auth := v1.Group("/auth")
		{
			auth.POST("/register", userCtrl.Register)
			auth.POST("/login", userCtrl.Login)
		}

		// User routes (protected)
		users := v1.Group("/users")
		users.Use(middleware.AuthMiddleware())
		{
			users.GET("/profile", userCtrl.GetProfile)
			users.PUT("/profile", userCtrl.UpdateProfile)
			users.PUT("/change-password", userCtrl.ChangePassword)
		}

		// Admin routes (protected + admin only)
		admin := v1.Group("/admin")
		admin.Use(middleware.AuthMiddleware())
		admin.Use(middleware.AdminMiddleware())
		{
			admin.GET("/users", userCtrl.GetAllUsers)
			admin.DELETE("/users/:id", userCtrl.DeleteUser)
		}

		// Example routes (for backward compatibility)
		examples := v1.Group("/examples")
		{
			examples.GET("", examplecontroller.Index)
			examples.GET("/:id", examplecontroller.Show)
			examples.POST("", examplecontroller.Create)
			examples.PUT("/:id", examplecontroller.Update)
			examples.DELETE("/:id", examplecontroller.Delete)
		}
	}

	// Health check route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "Restaurant API",
		})
	})

	return r
}
