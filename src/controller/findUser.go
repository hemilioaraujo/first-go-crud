package controller

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/hemilioaraujo/first-go-crud/src/configuration/logger"
	"github.com/hemilioaraujo/first-go-crud/src/configuration/rest_err"
	"github.com/hemilioaraujo/first-go-crud/src/model"
	"github.com/hemilioaraujo/first-go-crud/src/view"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	logger.Info("Init FindUserById controller", zap.String("journey", "findUserById"))
	user, err := model.VerifyToken(c.GetHeader("Authorization"))
	if err != nil {
		logger.Error("Error verifying token", err,
			zap.String("journey", "login_user"),
		)
		c.JSON(err.Code, err)
		return
	}
	logger.Info(fmt.Sprintf("User authenticated: %#v", user), zap.String("journey", "login_user"))
	userId := c.Param("userId")
	fmt.Println("User id: " + userId)
	if _, err := bson.ObjectIDFromHex(userId); err != nil {
		errorMessage := rest_err.NewBadRequestError("Invalid user id" + userId)
		logger.Error("Error trying to validate userId", err, zap.String("journey", "findUserById"))
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIdServices(userId)
	if err != nil {
		logger.Error("Error trying to call findUserById service", err, zap.String("journey", "findUserById"))
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init FindUserByEmail controller", zap.String("journey", "findUserByEmail"))
	user, err := model.VerifyToken(c.GetHeader("Authorization"))
	if err != nil {
		logger.Error("Error verifying token", err,
			zap.String("journey", "login_user"),
		)
		c.JSON(err.Code, err)
		return
	}
	logger.Info(fmt.Sprintf("User authenticated: %#v", user), zap.String("journey", "login_user"))
	userEmail := c.Param("userEmail")
	if _, err := mail.ParseAddress(userEmail); err != nil {
		errorMessage := rest_err.NewBadRequestError("Invalid user email")
		logger.Error("Error trying to validate userEmail", err, zap.String("journey", "findUserByEmail"))
		c.JSON(errorMessage.Code, errorMessage)
		return
	}
	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error trying to call findUserByEmail service", err, zap.String("journey", "findUserByEmail"))
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
