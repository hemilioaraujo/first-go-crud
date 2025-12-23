package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hemilioaraujo/first-go-crud/src/configuration/database/mongodb"
	"github.com/hemilioaraujo/first-go-crud/src/configuration/logger"
	"github.com/hemilioaraujo/first-go-crud/src/controller"
	"github.com/hemilioaraujo/first-go-crud/src/controller/routes"
	"github.com/hemilioaraujo/first-go-crud/src/model/repository"
	"github.com/hemilioaraujo/first-go-crud/src/model/service"
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
	database, err := mongodb.NewMongoDBConnection()
	if err != nil {
		log.Fatalf("Error connecting to MongoDB, error=%s \n", err.Error())
	}
	userRepository := repository.NewUserRepository(database)
	service := service.NewUserDomainService(userRepository)
	userController := controller.NewUserControllerInterface(service)
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		logger.Error("Error running server", err, zap.String("journey", "main"))
	}
}
