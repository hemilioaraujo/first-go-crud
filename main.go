package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hemilioaraujo/first-go-crud/src/configuration/logger"
	"github.com/hemilioaraujo/first-go-crud/src/controller/routes"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	logger.Info("Starting application...")
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", err, zap.String("journey", "main"))
	}

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		logger.Error("Error running server", err, zap.String("journey", "main"))
	}
}
