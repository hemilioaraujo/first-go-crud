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

func TestUserDomainService_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_update_user_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain(
			"teste@teste.com",
			"123456",
			"teste",
			18,
		)
		userDomain.SetID(id)
		repository.EXPECT().UpdateUser(userDomain.GetId(), userDomain).Return(nil)
		err := service.UpdateUser(userDomain.GetId(), userDomain)

		assert.Nil(t, err)
	})

	t.Run("when_update_user_repository_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain(
			"teste@teste.com",
			"123456",
			"teste",
			18,
		)
		userDomain.SetID(id)
		repository.EXPECT().UpdateUser(userDomain.GetId(), userDomain).Return(rest_err.NewInternalServerError("error trying to update user"))
		err := service.UpdateUser(userDomain.GetId(), userDomain)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to update user")
		assert.EqualValues(t, err.Code, http.StatusInternalServerError)
	})
}
