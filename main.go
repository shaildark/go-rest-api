package main

import (
	"os"

	"example.com/go-api/request/validation"
	"example.com/go-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	validation.RegisterCustomValidation()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":" + port)
}
