package routes

import (
	"go-react-mvc/backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupGinRoutes() *gin.Engine {
	router := gin.Default()
	// Auth routes
	router.POST("/api/auth/register", controllers.Register)
	router.POST("/api/auth/login", controllers.Login)

	// User routes
	router.GET("/api/users", controllers.GetUsers)
	router.POST("/api/users", controllers.CreateUser)
	router.PUT("/api/users/:id", controllers.UpdateUser)
	router.DELETE("/api/users/:id", controllers.DeleteUser)

	return router
}
