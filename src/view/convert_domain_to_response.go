package view

import (
	"github.com/hemilioaraujo/first-go-crud/src/controller/model/response"
	"github.com/hemilioaraujo/first-go-crud/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:       userDomain.GetId(),
		Email:    userDomain.GetEmail(),
		Name:     userDomain.GetName(),
		Age:      userDomain.GetAge(),
	}
}