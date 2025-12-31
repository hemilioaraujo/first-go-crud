package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hemilioaraujo/first-go-crud/src/configuration/logger"
	"github.com/hemilioaraujo/first-go-crud/src/configuration/validation"
	"github.com/hemilioaraujo/first-go-crud/src/controller/model/request"
	"github.com/hemilioaraujo/first-go-crud/src/model"
	"github.com/hemilioaraujo/first-go-crud/src/view"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error binding JSON", err,
			zap.String("journey", "create_user"),
		)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	domainResult, err := uc.service.CreateUserServices(domain)
	if err != nil {
		logger.Error("Error creating user", err,
			zap.String("journey", "create_user"),
		)
		c.JSON(err.Code, err)
		return
	}

	var tags []zap.Field
	tags = append(tags, zap.String("journey", "create_user"))
	tags = append(tags, zap.Any("user", view.ConvertDomainToResponse(domainResult)))
	logger.Info("User created", tags...)

	c.JSON(http.StatusCreated, view.ConvertDomainToResponse(domainResult))
}
