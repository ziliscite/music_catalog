package membership

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"music_catalog/internal/config"
	"music_catalog/internal/models/membership"
)

// Repository is an interface for user repository.
type Repository interface {
	Create(user *membership.User) error
	GetByID(id uint) (*membership.User, error)
	GetByEmail(email string) (*membership.User, error)
	GetByUsername(username string) (*membership.User, error)
}

// Interfaces in Go are already reference types,
// meaning they inherently hold a reference to the underlying concrete type that implements them.

type Service struct {
	cfg *config.Config
	r   Repository
}

func NewService(cfg *config.Config, r Repository) *Service {
	return &Service{
		cfg,
		r,
	}
}

func (s *Service) SignUp(request *membership.SignUpRequest) error {
	// check if email already exists
	email, err := s.r.GetByEmail(request.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if email != nil {
		return membership.ErrUserAlreadyExists
	}

	// check if a username already exists
	username, err := s.r.GetByUsername(request.Username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if username != nil {
		return membership.ErrUserAlreadyExists
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.r.Create(&membership.User{
		Email:    request.Email,
		Username: request.Username,
		Password: string(hashed),
	})
}
