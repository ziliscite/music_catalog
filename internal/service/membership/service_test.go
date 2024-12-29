package membership

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
	"music_catalog/internal/config"
	"music_catalog/internal/models/membership"
	"testing"
)

func TestService_SignUp(t *testing.T) {
	ctrMock := gomock.NewController(t) // create controller
	defer ctrMock.Finish()

	mockRepository := NewMockRepository(ctrMock) // the generated mock repo

	type args struct {
		request membership.SignUpRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				request: membership.SignUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "password",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepository.EXPECT().GetByEmail(args.request.Email).Return(nil, gorm.ErrRecordNotFound)
				mockRepository.EXPECT().GetByUsername(args.request.Username).Return(nil, gorm.ErrRecordNotFound)
				mockRepository.EXPECT().Create(gomock.Any()).Return(nil)
			},
		},
		{
			name: "failed when GetByEmail",
			args: args{
				request: membership.SignUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepository.EXPECT().GetByEmail(args.request.Email).Return(nil, assert.AnError)
			},
		},
		{
			name: "failed when GetByUsername",
			args: args{
				request: membership.SignUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepository.EXPECT().GetByEmail(args.request.Email).Return(nil, gorm.ErrRecordNotFound)
				mockRepository.EXPECT().GetByUsername(args.request.Username).Return(nil, assert.AnError)
			},
		},
		{
			name: "failed when CreateUser",
			args: args{
				request: membership.SignUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepository.EXPECT().GetByEmail(args.request.Email).Return(nil, gorm.ErrRecordNotFound)
				mockRepository.EXPECT().GetByUsername(args.request.Username).Return(nil, gorm.ErrRecordNotFound)
				mockRepository.EXPECT().Create(gomock.Any()).Return(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &Service{
				cfg: &config.Config{},
				r:   mockRepository,
			}
			if err := s.SignUp(&tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_SignIn(t *testing.T) {
	ctrMock := gomock.NewController(t)
	defer ctrMock.Finish()

	mockRepository := NewMockRepository(ctrMock)

	type args struct {
		request *membership.SignInRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				request: &membership.SignInRequest{
					Email:    "manzil.akbar@gmail.com",
					Password: "zilzil123",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepository.EXPECT().GetByEmail(args.request.Email).Return(&membership.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:    "manzil.akbar@gmail.com",
					Username: "manzil",
					Password: "$2a$10$5vhUElXxK3uM3PpTF2HkMekGz5BHXvXf/4FL29u733Z9F.7.bJPGu",
				}, nil)
			},
		},
		{
			name: "user not found",
			args: args{
				request: &membership.SignInRequest{
					Email:    "failed@gmail.com",
					Password: "zilzil123",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepository.EXPECT().GetByEmail(args.request.Email).Return(nil, membership.ErrUserNotFound)
			},
		},
		{
			name: "password not match",
			args: args{
				request: &membership.SignInRequest{
					Email:    "manzil.akbar@gmail.com",
					Password: "zilzil11111",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepository.EXPECT().GetByEmail(args.request.Email).Return(&membership.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:    "manzil.akbar@gmail.com",
					Username: "manzil",
					Password: "zilzil11111",
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			s := &Service{
				cfg: &config.Config{
					Service: config.Service{
						SecretKey: "abc",
					},
				},
				r: mockRepository,
			}

			got, err := s.SignIn(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotEmpty(t, got)
			} else {
				assert.Empty(t, got)
			}
		})
	}
}
