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

func TestUserDomainService_FindUserByIdServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_user_by_id_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain(
			"teste@teste.com",
			"123456",
			"teste",
			18,
		)
		userDomain.SetID(id)
		repository.EXPECT().FindUserById(id).Return(userDomain, nil)
		userDomainFound, err := service.FindUserByIdServices(id)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainFound.GetId(), userDomain.GetId())
		assert.EqualValues(t, userDomainFound.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, userDomainFound.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainFound.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainFound.GetAge(), userDomain.GetAge())
	})

	t.Run("when_does_not_exists_user_by_id_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		repository.EXPECT().FindUserById(id).Return(nil, rest_err.NewNotFoundError("user not found"))
		userDomainFound, err := service.FindUserByIdServices(id)

		assert.NotNil(t, err)
		assert.Nil(t, userDomainFound)
		assert.EqualValues(t, err.Message, "user not found")
		assert.EqualValues(t, err.Code, http.StatusNotFound)
	})
}

func TestUserDomainService_FindUserByEmailServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_user_by_email_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "teste@teste.com"
		userDomain := model.NewUserDomain(
			email,
			"123456",
			"teste",
			18,
		)
		userDomain.SetID(id)
		repository.EXPECT().FindUserByEmail(email).Return(userDomain, nil)
		userDomainFound, err := service.FindUserByEmailServices(email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainFound.GetId(), userDomain.GetId())
		assert.EqualValues(t, userDomainFound.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, userDomainFound.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainFound.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainFound.GetAge(), userDomain.GetAge())
	})

	t.Run("when_does_not_exists_user_by_email_error", func(t *testing.T) {
		email := "teste@teste.com"

		repository.EXPECT().FindUserByEmail(email).Return(nil, rest_err.NewNotFoundError("user not found"))
		userDomainFound, err := service.FindUserByEmailServices(email)

		assert.NotNil(t, err)
		assert.Nil(t, userDomainFound)
		assert.EqualValues(t, err.Message, "user not found")
		assert.EqualValues(t, err.Code, http.StatusNotFound)
	})
}

func TestUserDomainService_FindUserByEmailAndPasswordServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := &userDomainService{repository}

	t.Run("when_exists_user_by_email_and_password_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "teste@teste.com"
		password := "123456"
		userDomain := model.NewUserDomain(
			email,
			password,
			"teste",
			18,
		)
		userDomain.SetID(id)
		repository.EXPECT().FindUserByEmailAndPassword(email, password).Return(userDomain, nil)
		userDomainFound, err := service.findUserByEmailAndPasswordServices(email, password)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainFound.GetId(), userDomain.GetId())
		assert.EqualValues(t, userDomainFound.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, userDomainFound.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainFound.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainFound.GetAge(), userDomain.GetAge())
	})

	t.Run("when_does_not_exists_user_by_email_error", func(t *testing.T) {
		email := "teste@teste.com"
		password := "123456"

		repository.EXPECT().FindUserByEmailAndPassword(email, password).Return(nil, rest_err.NewForbiddenError("user or password are invalid"))
		userDomainFound, err := service.findUserByEmailAndPasswordServices(email, password)

		assert.NotNil(t, err)
		assert.Nil(t, userDomainFound)
		assert.EqualValues(t, err.Message, "user or password are invalid")
		assert.EqualValues(t, err.Code, http.StatusForbidden)
	})
}
