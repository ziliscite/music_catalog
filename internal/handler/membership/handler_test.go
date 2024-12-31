package membership

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"music_catalog/internal/model/membership"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_SignUp(t *testing.T) {
	ctrMock := gomock.NewController(t)
	defer ctrMock.Finish()

	mockService := NewMockService(ctrMock)

	tests := []struct {
		name               string
		mockFn             func()
		expectedStatusCode int
	}{
		{
			name: "success",
			mockFn: func() {
				mockService.EXPECT().SignUp(&membership.SignUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "password",
				}).Return(nil)
			},
			expectedStatusCode: 201,
		},
		{
			name: "failed when user exists",
			mockFn: func() {
				mockService.EXPECT().SignUp(&membership.SignUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "password",
				}).Return(membership.ErrUserAlreadyExists)
			},
			expectedStatusCode: 409,
		},
		{
			name: "failed when unknown error",
			mockFn: func() {
				mockService.EXPECT().SignUp(&membership.SignUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "password",
				}).Return(assert.AnError)
			},
			expectedStatusCode: 500,
		},
	}

	engine := gin.New()
	h := &Handler{
		Engine: engine,
		s:      mockService,
	}
	h.RegisterRoutes()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()

			w := httptest.NewRecorder()

			endpoint := `/membership/signup`
			model := membership.SignUpRequest{
				Email:    "test@gmail.com",
				Username: "testusername",
				Password: "password",
			}

			val, err := json.Marshal(model)
			assert.NoError(t, err)

			body := bytes.NewReader(val)
			req, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)

			h.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func TestHandler_SignIn(t *testing.T) {
	ctrMock := gomock.NewController(t)
	defer ctrMock.Finish()

	mockService := NewMockService(ctrMock)

	tests := []struct {
		name               string
		mockFn             func()
		expectedStatusCode int
		expectedBody       membership.SignInResponse
		wantErr            bool
	}{
		{
			name: "success",
			mockFn: func() {
				mockService.EXPECT().SignIn(&membership.SignInRequest{
					Email:    "test@gmail.com",
					Password: "password",
				}).Return("accessToken", nil)
			},
			expectedStatusCode: 200,
			expectedBody: membership.SignInResponse{
				AccessToken: "accessToken",
			},
			wantErr: false,
		},
		{
			name: "failed when user not found",
			mockFn: func() {
				mockService.EXPECT().SignIn(&membership.SignInRequest{
					Email:    "test@gmail.com",
					Password: "password",
				}).Return("", membership.ErrUserNotFound)
			},
			expectedStatusCode: 404,
			expectedBody:       membership.SignInResponse{},
			wantErr:            true,
		},
		{
			name: "failed when invalid credentials",
			mockFn: func() {
				mockService.EXPECT().SignIn(&membership.SignInRequest{
					Email:    "test@gmail.com",
					Password: "password",
				}).Return("", membership.ErrInvalidCredentials)
			},
			expectedStatusCode: 401,
			expectedBody:       membership.SignInResponse{},
			wantErr:            true,
		},
		{
			name: "failed when unknown error",
			mockFn: func() {
				mockService.EXPECT().SignIn(&membership.SignInRequest{
					Email:    "test@gmail.com",
					Password: "password",
				}).Return("", assert.AnError)
			},
			expectedStatusCode: 500,
			expectedBody:       membership.SignInResponse{},
			wantErr:            true,
		},
	}

	engine := gin.New()
	h := &Handler{
		Engine: engine,
		s:      mockService,
	}
	h.RegisterRoutes()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()

			w := httptest.NewRecorder()

			endpoint := `/membership/signin`
			model := membership.SignInRequest{
				Email:    "test@gmail.com",
				Password: "password",
			}

			val, err := json.Marshal(model)
			assert.NoError(t, err)

			body := bytes.NewReader(val)
			req, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)
			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			if !tt.wantErr {
				res := w.Result()
				defer res.Body.Close()

				response := membership.SignInResponse{}
				err = json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody, response)
			}
		})
	}
}
