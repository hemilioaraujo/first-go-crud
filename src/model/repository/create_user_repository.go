package repository

import (
	"context"
	"os"

	"github.com/hemilioaraujo/first-go-crud/src/configuration/logger"
	"github.com/hemilioaraujo/first-go-crud/src/configuration/rest_err"
	"github.com/hemilioaraujo/first-go-crud/src/model"
	"github.com/hemilioaraujo/first-go-crud/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository")
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.dbConnection.Collection(collection_name)

	jsonValue := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), jsonValue)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	jsonValue.ID = result.InsertedID.(bson.ObjectID)

	return converter.ConvertEntityToDomain(jsonValue), nil
}
