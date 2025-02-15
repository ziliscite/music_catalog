package usertrack

import (
	"gorm.io/gorm"
	"music_catalog/internal/model/usertrack"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db,
	}
}

func (r *Repository) Create(userTrack *usertrack.UserTrack) error {
	return r.db.Create(&userTrack).Error
}

func (r *Repository) GetAllLiked(userId uint, trackIds []string) (map[string]usertrack.UserTrack, error) {
	var userTracks []usertrack.UserTrack
	if err := r.db.Where("user_id = ? AND track_id IN ?", userId, trackIds).Find(&userTracks).Error; err != nil {
		return nil, err
	}

	response := make(map[string]usertrack.UserTrack)
	for _, userTrack := range userTracks {
		response[userTrack.TrackID] = userTrack
	}

	return response, nil
}

func (r *Repository) GetLikedById(userId uint, trackId string) (*usertrack.UserTrack, error) {
	var userTrack usertrack.UserTrack
	if err := r.db.Where("user_id = ? AND track_id = ?", userId, trackId).First(&userTrack).Error; err != nil {
		return nil, err
	}

	return &userTrack, nil
}

func (r *Repository) Update(userTrack *usertrack.UserTrack) error {
	return r.db.Save(&userTrack).Error
}
