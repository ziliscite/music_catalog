package usertrack

import "gorm.io/gorm"

type UserTrack struct {
	gorm.Model
	UserID  uint   `gorm:"not null"`
	TrackID string `gorm:"not null"`
	IsLiked *bool
}

type LikeRequest struct {
	TrackID string `json:"track_id"`
	IsLiked *bool  `json:"is_liked"` // true = liked, false = dislike, null = neutral
}

func (l *LikeRequest) Model(userId uint) *UserTrack {
	return &UserTrack{
		UserID:  userId,
		TrackID: l.TrackID,
		IsLiked: l.IsLiked,
	}
}
