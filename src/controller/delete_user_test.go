package controller

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hemilioaraujo/first-go-crud/src/configuration/rest_err"
	"github.com/hemilioaraujo/first-go-crud/src/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"

	"go.uber.org/mock/gomock"
)

func TestUserControllerInterface_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)

	t.Run("when user is deleted successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		userId := bson.NewObjectID().Hex()
		param := gin.Params{
			{Key: "userId", Value: userId},
		}
		MakeRequest(ctx, param, url.Values{}, "DELETE", nil)
		service.EXPECT().DeleteUser(userId).Return(nil)
		controller.DeleteUser(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})

	t.Run("when userId is invalid", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		userId := "invalid"
		param := gin.Params{
			{Key: "userId", Value: userId},
		}
		MakeRequest(ctx, param, url.Values{}, "DELETE", nil)
		controller.DeleteUser(ctx)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("when error deleting user occurs", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		userId := bson.NewObjectID().Hex()
		param := gin.Params{
			{Key: "userId", Value: userId},
		}
		MakeRequest(ctx, param, url.Values{}, "DELETE", nil)
		service.EXPECT().DeleteUser(userId).Return(rest_err.NewInternalServerError("error deleting user"))
		controller.DeleteUser(ctx)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})
}
