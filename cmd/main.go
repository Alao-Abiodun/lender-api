package main

import (
	"fmt"
	"os"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/Alao-Abiodun/lender-api/internal/application"
	"github.com/Alao-Abiodun/lender-api/internal/infrastructure/mongo"
	"github.com/Alao-Abiodun/lender-api/internal/interfaces/http"
)

func main() {

	_ = godotenv.Load()

	dbURI := os.Getenv("MONGO_URI")
	dbClient, err := mongo.Connect(dbURI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	userRepo := mongo.NewUserRepository(dbClient.Database("lenderdb"))
	userService := application.NewUserService(userRepo)
	userHandler := http.NewUserHandler(userService)

	router := gin.Default()
	router.POST("/api/v1/users", userHandler.Register)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(fmt.Sprintf(":%s", port))
}