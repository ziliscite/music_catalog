package track

import (
	"music_catalog/internal/config"
	model "music_catalog/internal/model/spotify"
	dto "music_catalog/internal/repository/spotify"
)

//go:generate mockgen -source=service.go -destination=service_mock.go -package=track
type Repository interface {
	Search(q string, limit, offset int) (*dto.ClientSearchResponse, error)
}

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

func (s *Service) Search(query string, pageSize, pageIndex int) (*model.SearchResponse, error) {
	offset := (pageIndex - 1) * pageSize

	tracks, err := s.r.Search(query, pageSize, offset)
	if err != nil {
		return nil, err
	}

	return tracks.Model(), nil
}
