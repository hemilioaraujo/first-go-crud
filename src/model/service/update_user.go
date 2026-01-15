package service

import (
	"github.com/hemilioaraujo/first-go-crud/src/configuration/logger"
	"github.com/hemilioaraujo/first-go-crud/src/configuration/rest_err"
	"github.com/hemilioaraujo/first-go-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Updating user", zap.String("journey", "update_user"))

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		return err
	}
	return nil
}
