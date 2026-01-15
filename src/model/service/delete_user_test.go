package service

import (
	"net/http"
	"testing"

	"github.com/hemilioaraujo/first-go-crud/src/configuration/rest_err"
	"github.com/hemilioaraujo/first-go-crud/src/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_delete_user_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		repository.EXPECT().DeleteUser(id).Return(nil)
		err := service.DeleteUser(id)

		assert.Nil(t, err)
	})

	t.Run("when_user_id_is_invalid_returns_error", func(t *testing.T) {
		id := "invalid_id"
		repository.EXPECT().DeleteUser(id).Return(rest_err.NewBadRequestError("userId is not a valid hex"))
		err := service.DeleteUser(id)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "userId is not a valid hex")
		assert.EqualValues(t, err.Code, http.StatusBadRequest)
	})

	t.Run("when_delete_user_repository_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		repository.EXPECT().DeleteUser(id).Return(rest_err.NewInternalServerError("error trying to delete user"))
		err := service.DeleteUser(id)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to delete user")
		assert.EqualValues(t, err.Code, http.StatusInternalServerError)
	})
}
