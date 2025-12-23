package service

import (
	"fmt"

	"github.com/hemilioaraujo/first-go-crud/src/configuration/logger"
	"github.com/hemilioaraujo/first-go-crud/src/configuration/rest_err"
	"github.com/hemilioaraujo/first-go-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Creating user", zap.String("journey", "create_user"))
	userDomain.EncryptPassword()

	fmt.Println(userDomain)
	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		return nil, err
	}
	return userDomainRepository, nil
}
