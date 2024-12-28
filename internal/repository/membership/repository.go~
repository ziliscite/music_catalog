package membership

import (
	"gorm.io/gorm"
	"music_catalog/internal/models/membership"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db,
	}
}

func (r *Repository) Create(user *membership.User) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetByID(id uint) (*membership.User, error) {
	var user membership.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) GetByEmail(email string) (*membership.User, error) {
	var user membership.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) GetByUsername(username string) (*membership.User, error) {
	var user membership.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
