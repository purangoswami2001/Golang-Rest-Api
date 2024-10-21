package v1

import (
	"api/internal/handlers"
	"api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userHandler *handlers.UserHandler) {
	v1 := r.Group("/api/v1")
	{
		// Public route for user login
		v1.POST("/user/login", userHandler.Login)

		// Protected routes with the JWT middleware
		protected := v1.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.POST("/users", userHandler.CreateUser)
			protected.PUT("/users/:id", userHandler.UpdateUser)
			protected.GET("/users/:id", userHandler.GetUser)
			protected.GET("/users", userHandler.GetAllUsers)
			protected.DELETE("/users/:id", userHandler.DeleteUser)
		}
	}
}
