package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	shop "github.com/zangar-tm/online_shop"
	"github.com/zangar-tm/online_shop/pkg/service"
	mock_service "github.com/zangar-tm/online_shop/pkg/service/mocks"
)

func TestHandler_signUp(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *mock_service.MockAuthorization, user shop.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            shop.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"username": "username", "name": "Test Name", "password": "secret"}`,
			inputUser: shop.User{
				Username: "username",
				Name:     "Test Name",
				Password: "secret",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user shop.User) {
				r.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{"username": "username"}`,
			inputUser:            shop.User{},
			mockBehavior:         func(r *mock_service.MockAuthorization, user shop.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"username": "username", "name": "Test Name", "password": "secret"}`,
			inputUser: shop.User{
				Username: "username",
				Name:     "Test Name",
				Password: "secret",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user shop.User) {
				r.EXPECT().CreateUser(user).Return(0, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
