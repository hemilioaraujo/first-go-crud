package service

import (
	"github.com/hemilioaraujo/first-go-crud/src/configuration/logger"
	"github.com/hemilioaraujo/first-go-crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(userId string) *rest_err.RestErr {
	logger.Info("Deleting user", zap.String("journey", "delete_user"))
	if err := ud.userRepository.DeleteUser(userId); err != nil {
		return err
	}
	return nil
}
