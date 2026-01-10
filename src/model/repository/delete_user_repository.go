package repository

import (
	"context"
	"os"

	"github.com/hemilioaraujo/first-go-crud/src/configuration/logger"
	"github.com/hemilioaraujo/first-go-crud/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(userId string) *rest_err.RestErr {
	logger.Info("Init deleteUser repository")
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.dbConnection.Collection(collection_name)

	userIdHex, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return rest_err.NewBadRequestError("userId is not a valid hex")
	}

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}

	var tags []zap.Field
	tags = append(tags, zap.String("journey", "delete_user"))
	tags = append(tags, zap.Any("user", userId))
	logger.Info("User deleted", tags...)

	return nil
}
