package service

import (
	"net/http"
	"testing"

	"github.com/hemilioaraujo/first-go-crud/src/configuration/rest_err"
	"github.com/hemilioaraujo/first-go-crud/src/model"
	"github.com/hemilioaraujo/first-go-crud/src/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_CreateUserServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_create_user_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain(
			"teste@teste.com",
			"123456",
			"teste",
			18,
		)
		userDomain.SetID(id)
		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)
		repository.EXPECT().CreateUser(userDomain).Return(userDomain, nil)
		userDomainCreated, err := service.CreateUserServices(userDomain)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainCreated.GetId(), userDomain.GetId())
		assert.EqualValues(t, userDomainCreated.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, userDomainCreated.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainCreated.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainCreated.GetAge(), userDomain.GetAge())
	})

	t.Run("when_user_already_exists_error", func(t *testing.T) {
		userDomain := model.NewUserDomain(
			"teste@teste.com",
			"123456",
			"teste",
			18,
		)
		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(userDomain, nil)
		userDomainCreated, err := service.CreateUserServices(userDomain)

		assert.NotNil(t, err)
		assert.Nil(t, userDomainCreated)
		assert.EqualValues(t, err.Message, "email already exists")
		assert.EqualValues(t, err.Code, http.StatusBadRequest)
	})

	t.Run("when_create_user_repository_error", func(t *testing.T) {
		userDomain := model.NewUserDomain(
			"teste@teste.com",
			"123456",
			"teste",
			18,
		)
		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)
		repository.EXPECT().CreateUser(userDomain).Return(nil, rest_err.NewInternalServerError("error trying to create user"))
		userDomainCreated, err := service.CreateUserServices(userDomain)

		assert.NotNil(t, err)
		assert.Nil(t, userDomainCreated)
		assert.EqualValues(t, err.Message, "error trying to create user")
		assert.EqualValues(t, err.Code, http.StatusInternalServerError)
	})
}
