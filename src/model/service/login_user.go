package service

import (
	"github.com/hemilioaraujo/first-go-crud/src/configuration/logger"
	"github.com/hemilioaraujo/first-go-crud/src/configuration/rest_err"
	"github.com/hemilioaraujo/first-go-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr) {
	logger.Info("Init loggin user", zap.String("journey", "login_user"))
	userDomain.EncryptPassword()

	user, err := ud.userRepository.FindUserByEmailAndPassword(
		userDomain.GetEmail(),
		userDomain.GetPassword(),
	)
	if err != nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	logger.Info(
		"User logged in",
		zap.String("journey", "login_user"),
		zap.String("userId", user.GetId()),
	)
	return user, token, nil
}
