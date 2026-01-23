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

func TestUserControllerInterface_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)

	t.Run("when user is created successfully", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		userRequest := request.UserRequest{
			Email:    "teste@teste.com",
			Password: "123456!",
			Name:     "joaquim",
			Age:      15,
		}
		domain := model.NewUserDomain(
			userRequest.Email,
			userRequest.Password,
			userRequest.Name,
			userRequest.Age,
		)
		jsonBody, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(jsonBody)))
		MakeRequest(ctx, gin.Params{}, url.Values{}, "POST", stringReader)
		service.EXPECT().CreateUserServices(domain).Return(domain, nil)
		controller.CreateUser(ctx)

		assert.EqualValues(t, http.StatusCreated, recorder.Code)
	})

	t.Run("when validation returns error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		userRequest := request.UserRequest{
			Email:    "testeteste.com",
			Password: "123456!",
			Name:     "joaquim",
			Age:      15,
		}
		jsonBody, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(jsonBody)))
		MakeRequest(ctx, gin.Params{}, url.Values{}, "POST", stringReader)

		controller.CreateUser(ctx)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("when service returns error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		userRequest := request.UserRequest{
			Email:    "teste@teste.com",
			Password: "123456!",
			Name:     "joaquim",
			Age:      15,
		}
		domain := model.NewUserDomain(
			userRequest.Email,
			userRequest.Password,
			userRequest.Name,
			userRequest.Age,
		)
		jsonBody, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(jsonBody)))
		MakeRequest(ctx, gin.Params{}, url.Values{}, "POST", stringReader)
		service.EXPECT().CreateUserServices(domain).Return(nil, rest_err.NewBadRequestError("invalid user"))
		controller.CreateUser(ctx)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

}
