package repository

import (
	"github.com/hemilioaraujo/first-go-crud/src/configuration/rest_err"
	"github.com/hemilioaraujo/first-go-crud/src/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func NewUserRepository(dbConnection *mongo.Database) UserRepository {
	return &userRepository{dbConnection}
}

type userRepository struct {
	dbConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(userEmail string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserById(userId string) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(userId string) *rest_err.RestErr
}
