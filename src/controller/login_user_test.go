package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hemilioaraujo/first-go-crud/src/configuration/rest_err"
	"github.com/hemilioaraujo/first-go-crud/src/controller/model/request"
	"github.com/hemilioaraujo/first-go-crud/src/model"
	"github.com/hemilioaraujo/first-go-crud/src/tests/mocks"
	"github.com/stretchr/testify/assert"

	"go.uber.org/mock/gomock"
)

func TestUserControllerInterface_LoginUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)

	t.Run("when user is logged in successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		userRequest := request.UserLoginRequest{
			Email:    "teste@teste.com",
			Password: "123456!",
		}
		domain := model.NewUserLoginDomain(
			userRequest.Email,
			userRequest.Password,
		)
		jsonBody, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(jsonBody)))
		MakeRequest(ctx, gin.Params{}, url.Values{}, "POST", stringReader)
		service.EXPECT().LoginUserServices(domain).Return(domain, "token123", nil)

		controller.LoginUser(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
		assert.EqualValues(t, "token123", recorder.Header().Get("Authorization"))
	})

	t.Run("when validation returns error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		userRequest := request.UserLoginRequest{
			Email:    "testeteste.com",
			Password: "123456!",
		}
		jsonBody, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(jsonBody)))
		MakeRequest(ctx, gin.Params{}, url.Values{}, "POST", stringReader)

		controller.LoginUser(ctx)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("when service returns error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		userRequest := request.UserLoginRequest{
			Email:    "teste@teste.com",
			Password: "123456!",
		}
		domain := model.NewUserLoginDomain(
			userRequest.Email,
			userRequest.Password,
		)
		jsonBody, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(jsonBody)))
		MakeRequest(ctx, gin.Params{}, url.Values{}, "POST", stringReader)
		service.EXPECT().LoginUserServices(domain).Return(nil, "", rest_err.NewBadRequestError("invalid user"))
		controller.LoginUser(ctx)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

}
