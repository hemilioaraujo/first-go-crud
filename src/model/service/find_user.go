package service

import (
	"github.com/hemilioaraujo/first-go-crud/src/configuration/logger"
	"github.com/hemilioaraujo/first-go-crud/src/configuration/rest_err"
	"github.com/hemilioaraujo/first-go-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIdServices(userId string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Finding user by id", zap.String("journey", "findUserById"))
	return ud.userRepository.FindUserById(userId)
}

func (ud *userDomainService) FindUserByEmailServices(userEmail string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Finding user by email", zap.String("journey", "findUserByEmail"))
	return ud.userRepository.FindUserByEmail(userEmail)
}
