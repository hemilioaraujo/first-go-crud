package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hemilioaraujo/first-go-crud/src/model/service"
)
func NewUserControllerInterface(serviceInterface service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{service: serviceInterface}
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}