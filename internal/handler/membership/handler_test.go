package membership

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"music_catalog/internal/models/membership"
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
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Engine: tt.fields.Engine,
				s:      tt.fields.s,
			}
			h.SignIn(tt.args.c)
		})
	}
}
