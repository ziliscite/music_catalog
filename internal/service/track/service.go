package track

import (
	"errors"
	"gorm.io/gorm"
	"music_catalog/internal/config"
	model "music_catalog/internal/model/spotify"
	"music_catalog/internal/model/usertrack"
	dto "music_catalog/internal/repository/spotify"
)

//go:generate mockgen -source=service.go -destination=service_mock.go -package=track
type RepositorySpotify interface {
	Search(q string, limit, offset int) (*dto.ClientSearchResponse, error)
}

//go:generate mockgen -source=service.go -destination=service_mock.go -package=track
type RepositoryUserTrack interface {
	Create(userTrack *usertrack.UserTrack) error
	GetAllLiked(userId uint, trackIds []string) (map[string]usertrack.UserTrack, error)
	GetLikedById(userId uint, trackId string) (*usertrack.UserTrack, error)
	Update(userTrack *usertrack.UserTrack) error
}

type Service struct {
	cfg *config.Config
	rs  RepositorySpotify
	rut RepositoryUserTrack
}

func NewService(cfg *config.Config, rs RepositorySpotify, rut RepositoryUserTrack) *Service {
	return &Service{
		cfg,
		rs,
		rut,
	}
}

func (s *Service) Search(query string, pageSize, pageIndex int) (*model.SearchResponse, error) {
	offset := (pageIndex - 1) * pageSize

	tracks, err := s.rs.Search(query, pageSize, offset)
	if err != nil {
		return nil, err
	}

	return tracks.Model(), nil
}

func (s *Service) Upsert(userId uint, request *usertrack.LikeRequest) (bool, error) {
	// check if user track is already in the database
	track, err := s.rut.GetLikedById(userId, request.TrackID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}

	// update user track if it is already in the database
	if track != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		track.IsLiked = request.IsLiked
		return false, s.rut.Update(track)
	}

	// create a new user track otherwise
	return true, s.rut.Create(request.Model(userId))
}
