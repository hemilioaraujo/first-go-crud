package model

import (
	"fmt"

	"github.com/hemilioaraujo/first-go-crud/src/configuration/logger"
	"github.com/hemilioaraujo/first-go-crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Creating user", zap.String("journey", "create_user"))
	ud.EncryptPassword()

	fmt.Println(ud)
	return nil
}
