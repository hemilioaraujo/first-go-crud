package service

import (
	"testing"

	"github.com/hemilioaraujo/first-go-crud/src/configuration/rest_err"
	"github.com/hemilioaraujo/first-go-crud/src/model"
	"github.com/hemilioaraujo/first-go-crud/src/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_LoginUserServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := &userDomainService{repository}

	t.Run("when_login_user_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain(
			"teste@teste.com",
			"123456",
			"teste",
			18,
		)
		userDomain.SetID(id)

		userDomainMock := model.NewUserDomain(
			userDomain.GetEmail(),
			userDomain.GetPassword(),
			userDomain.GetName(),
			userDomain.GetAge(),
		)
		userDomainMock.SetID(id)
		userDomainMock.EncryptPassword()

		repository.EXPECT().FindUserByEmailAndPassword(userDomainMock.GetEmail(), userDomainMock.GetPassword()).Return(userDomain, nil)
		userDomainLogged, token, err := service.LoginUserServices(userDomain)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainLogged.GetId(), userDomain.GetId())
		assert.EqualValues(t, userDomainLogged.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, userDomainLogged.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainLogged.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainLogged.GetAge(), userDomain.GetAge())
		assert.NotEmpty(t, token)
	})

	t.Run("when_repository_return_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain(
			"teste@teste.com",
			"123456",
			"teste",
			18,
		)
		userDomain.SetID(id)

		userDomainFake := model.NewUserDomain(
			userDomain.GetEmail(),
			userDomain.GetPassword(),
			userDomain.GetName(),
			userDomain.GetAge(),
		)
		userDomainFake.SetID(id)

		userDomainFake.EncryptPassword()
		repository.EXPECT().FindUserByEmailAndPassword(userDomainFake.GetEmail(), userDomainFake.GetPassword()).Return(nil, rest_err.NewInternalServerError("error trying to find user"))
		userDomainLogged, token, err := service.LoginUserServices(userDomain)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to find user")
		assert.Nil(t, userDomainLogged)
		assert.Empty(t, token)
	})

	t.Run("when_generate_token_return_error", func(t *testing.T) {
		userDomainMock := mocks.NewMockUserDomainInterface(ctrl)
		userDomainMock.EXPECT().GetEmail().Return("teste@teste.com")
		userDomainMock.EXPECT().GetPassword().Return("123456")
		userDomainMock.EXPECT().EncryptPassword()
		userDomainMock.EXPECT().GenerateToken().Return("", rest_err.NewInternalServerError("error trying to generate token"))

		repository.EXPECT().FindUserByEmailAndPassword("teste@teste.com", "123456").Return(userDomainMock, nil)

		userDomainLogged, token, err := service.LoginUserServices(userDomainMock)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to generate token")
		assert.Nil(t, userDomainLogged)
		assert.Empty(t, token)
	})
}
