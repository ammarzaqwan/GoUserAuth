package routes

import (
	"example/apigo/config"
	"example/apigo/controllers"
	"example/apigo/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	db := config.DB
	// Public routes

	reg := r.Group("/")
	reg.POST("/register", controllers.Register)
	reg.POST("/login", controllers.Login)

	// Protected routes (require authentication)
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	auth.Use(middleware.LoggingMiddleware(db))
	auth.PUT("/user/:username", controllers.UpdateUser)
	auth.POST("/user/logout", controllers.LogOut)

	return r
}
