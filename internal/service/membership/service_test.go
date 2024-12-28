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
			if err := s.SignUp(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
