package app

import (
	"api/config"
	"api/internal/handlers"
	"api/internal/models"
	"api/internal/repositories"
	"api/internal/services"
	"api/internal/utils"
	v1 "api/routes/v1"
	"log"

	"github.com/gin-gonic/gin"
)

func Run(config *config.Config) {
	db, err := utils.ConnectDB(config)
	if err != nil {
		log.Fatalf("Unable to connect to the database")
	}
	r := gin.Default()

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Migration failed: %v", err) // Log if migration fails
	} else {
		log.Println("Database migration completed successfully.")
	}
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	v1.SetupRoutes(r, userHandler)
	r.Run("localhost:8080")

}
