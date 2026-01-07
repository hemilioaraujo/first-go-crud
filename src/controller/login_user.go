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

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	var userRequest request.UserLoginRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error binding JSON", err,
			zap.String("journey", "login_user"),
		)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)

	domainResult, token, err := uc.service.LoginUserServices(domain)
	if err != nil {
		logger.Error("Error logging in user", err,
			zap.String("journey", "login_user"),
		)
		c.JSON(err.Code, err)
		return
	}

	var tags []zap.Field
	tags = append(tags, zap.String("journey", "login_user"))
	tags = append(tags, zap.Any("user", domainResult.GetId()))
	logger.Info("User logged in", tags...)

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
