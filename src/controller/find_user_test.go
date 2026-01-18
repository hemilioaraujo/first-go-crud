package controller

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hemilioaraujo/first-go-crud/src/configuration/rest_err"
	"github.com/hemilioaraujo/first-go-crud/src/model"
	"github.com/hemilioaraujo/first-go-crud/src/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserControllerInterface_FindUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)

	t.Run("when_find_user_by_email_is_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		email := "teste@teste.com"
		param := gin.Params{
			{Key: "userEmail", Value: email},
		}
		MakeRequest(ctx, param, url.Values{}, "GET", nil)

		userDomain := model.NewUserDomain(email, "teste", "teste", 18)

		service.EXPECT().FindUserByEmailServices(email).Return(userDomain, nil)
		controller.FindUserByEmail(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})

	t.Run("when_find_user_by_email_has_a_invalid_email", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		param := gin.Params{
			{Key: "userEmail", Value: "invalid_email"},
		}
		MakeRequest(ctx, param, url.Values{}, "GET", nil)

		controller.FindUserByEmail(ctx)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("when_find_user_by_email_has_a_valid_email_and_service_returns_an_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		email := "teste@teste.com"
		param := gin.Params{
			{Key: "userEmail", Value: email},
		}
		MakeRequest(ctx, param, url.Values{}, "GET", nil)

		service.EXPECT().FindUserByEmailServices(email).Return(nil, rest_err.NewInternalServerError("internal server error"))
		controller.FindUserByEmail(ctx)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})
}

func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(recorder)

	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func MakeRequest(
	c *gin.Context,
	param gin.Params,
	u url.Values,
	method string,
	body io.ReadCloser,
) {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = param
	c.Request.Body = body
	c.Request.URL.RawQuery = u.Encode()
}
