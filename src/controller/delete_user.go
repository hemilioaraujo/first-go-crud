package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hemilioaraujo/first-go-crud/src/configuration/logger"
	"github.com/hemilioaraujo/first-go-crud/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	if _, err := bson.ObjectIDFromHex(userId); err != nil || strings.TrimSpace(userId) == "" {
		logger.Error("Error deleting user", err,
			zap.String("journey", "delete_user"),
		)
		errRest := rest_err.NewBadRequestError("userId is not a valid hex")
		c.JSON(errRest.Code, errRest)
		return
	}

	err := uc.service.DeleteUser(userId)
	if err != nil {
		logger.Error("Error deleting user", err,
			zap.String("journey", "delete_user"),
		)
		c.JSON(err.Code, err)
		return
	}

	var tags []zap.Field
	tags = append(tags, zap.String("journey", "delete_user"))
	tags = append(tags, zap.Any("user", userId))
	logger.Info("User deleted", tags...)

	c.Status(http.StatusOK)
}
