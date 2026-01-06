package repository

import (
	"context"
	"os"

	"github.com/hemilioaraujo/first-go-crud/src/configuration/logger"
	"github.com/hemilioaraujo/first-go-crud/src/configuration/rest_err"
	"github.com/hemilioaraujo/first-go-crud/src/model"
	"github.com/hemilioaraujo/first-go-crud/src/model/repository/entity"
	"github.com/hemilioaraujo/first-go-crud/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(userEmail string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail repository", zap.String("journey", "findUserByEmail"))
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.dbConnection.Collection(collection_name)
	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: userEmail}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "user not found with email " + userEmail
			logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "error trying to find user with email " + userEmail
		logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	logger.Info("User found by email ",
		zap.String("journey", "findUserByEmail"),
		zap.String("userEmail", userEmail),
		zap.String("userId", userEntity.ID.Hex()),
	)
	return converter.ConvertEntityToDomain(userEntity), nil
}

func (ur *userRepository) FindUserById(userId string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserById repository", zap.String("journey", "findUserById"))
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.dbConnection.Collection(collection_name)
	userEntity := &entity.UserEntity{}

	objectId, _ := bson.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "_id", Value: objectId}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "user not found with id " + userId
			logger.Error(errorMessage, err, zap.String("journey", "findUserById"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "error trying to find user with id " + userId
		logger.Error(errorMessage, err, zap.String("journey", "findUserById"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	logger.Info("User found by id ",
		zap.String("journey", "findUserById"),
		zap.String("userId", userEntity.ID.Hex()),
	)
	return converter.ConvertEntityToDomain(userEntity), nil
}

func (ur *userRepository) FindUserByEmailAndPassword(userEmail string, userPassword string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmailAndPassword repository", zap.String("journey", "findUserByEmailAndPassword"))
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.dbConnection.Collection(collection_name)
	userEntity := &entity.UserEntity{}

	filter := bson.D{
		{Key: "email", Value: userEmail},
		{Key: "password", Value: userPassword},
	}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "email or password is incorrect"
			logger.Error(errorMessage, err, zap.String("journey", "findUserByEmailAndPassword"))
			return nil, rest_err.NewForbiddenError(errorMessage)
		}
		errorMessage := "error trying to find user with email and password"
		logger.Error(errorMessage, err, zap.String("journey", "findUserByEmailAndPassword"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	logger.Info("User found by email and password",
		zap.String("journey", "findUserByEmailAndPassword"),
		zap.String("userEmail", userEmail),
		zap.String("userId", userEntity.ID.Hex()),
	)
	return converter.ConvertEntityToDomain(userEntity), nil
}
