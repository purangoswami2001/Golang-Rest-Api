package v1

import (
	"api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userHandler *handlers.UserHandler) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/users", userHandler.CreateUser)
		v1.PUT("/users/:id", userHandler.UpdateUser)
		v1.GET("/users/:id", userHandler.GetUser)
		v1.GET("/users", userHandler.GetAllUsers)
		v1.DELETE("/users/:id", userHandler.DeleteUser)
	}
}
