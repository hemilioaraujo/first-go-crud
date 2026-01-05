package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hemilioaraujo/first-go-crud/src/configuration/logger"
	"github.com/hemilioaraujo/first-go-crud/src/configuration/rest_err"
	"github.com/hemilioaraujo/first-go-crud/src/configuration/validation"
	"github.com/hemilioaraujo/first-go-crud/src/controller/model/request"
	"github.com/hemilioaraujo/first-go-crud/src/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	var userRequest request.UserUpdateRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error binding JSON", err,
			zap.String("journey", "update_user"),
		)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	userId := c.Param("userId")
	if _, err := bson.ObjectIDFromHex(userId); err != nil || strings.TrimSpace(userId) == "" {
		logger.Error("Error updating user", err,
			zap.String("journey", "update_user"),
		)
		errRest := rest_err.NewBadRequestError("userId is not a valid hex")
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		logger.Error("Error updating user", err,
			zap.String("journey", "update_user"),
		)
		c.JSON(err.Code, err)
		return
	}

	var tags []zap.Field
	tags = append(tags, zap.String("journey", "create_user"))
	tags = append(tags, zap.Any("user", userId))
	logger.Info("User updated", tags...)

	c.Status(http.StatusOK)
}
