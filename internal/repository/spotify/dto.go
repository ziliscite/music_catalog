package spotify

import (
	"music_catalog/internal/model/spotify"
	"music_catalog/internal/model/usertrack"
)

type ClientSearchResponse struct {
	Tracks Tracks `json:"tracks"`
}

func (c *ClientSearchResponse) Model(userTracks map[string]usertrack.UserTrack) *spotify.SearchResponse {
	tracks := make([]spotify.TrackObject, 0)
	for _, item := range c.Tracks.Items {
		artistsName := make([]string, len(item.Artists))
		for idx, artist := range item.Artists {
			artistsName[idx] = artist.Name
		}

		imageUrls := make([]string, len(item.Album.Images))
		for idx, image := range item.Album.Images {
			imageUrls[idx] = image.URL
		}

		tracks = append(tracks, spotify.TrackObject{
			// album related fields
			AlbumType:        item.Album.AlbumType,
			AlbumTotalTracks: item.Album.TotalTracks,
			AlbumImagesURL:   imageUrls,
			AlbumName:        item.Album.Name,
			// artist related fields
			ArtistsName: artistsName,
			// track related fields
			Explicit: item.Explicit,
			ID:       item.ID,
			Name:     item.Name,

			IsLiked: userTracks[item.ID].IsLiked,
		})
	}

	return &spotify.SearchResponse{
		Limit:  c.Tracks.Limit,
		Offset: c.Tracks.Offset,
		Items:  tracks,
		Total:  c.Tracks.Total,
	}
}

type RecommendationResponse struct {
	Items []TrackObject `json:"items"`
}

func (r *RecommendationResponse) Model(userTracks map[string]usertrack.UserTrack) *spotify.RecommendationResponse {
	tracks := make([]spotify.TrackObject, 0)
	for _, item := range r.Items {
		artistsName := make([]string, len(item.Artists))
		for idx, artist := range item.Artists {
			artistsName[idx] = artist.Name
		}

		imageUrls := make([]string, len(item.Album.Images))
		for idx, image := range item.Album.Images {
			imageUrls[idx] = image.URL
		}

		tracks = append(tracks, spotify.TrackObject{
			// album related fields
			AlbumType:        item.Album.AlbumType,
			AlbumTotalTracks: item.Album.TotalTracks,
			AlbumImagesURL:   imageUrls,
			AlbumName:        item.Album.Name,
			// artist related fields
			ArtistsName: artistsName,
			// track related fields
			Explicit: item.Explicit,
			ID:       item.ID,
			Name:     item.Name,

			IsLiked: userTracks[item.ID].IsLiked,
		})
	}

	return &spotify.RecommendationResponse{
		Items: tracks,
	}
}

type Tracks struct {
	Href     string        `json:"href"`
	Limit    int           `json:"limit"`
	Next     *string       `json:"next"`
	Offset   int           `json:"offset"`
	Previous *string       `json:"previous"`
	Total    int           `json:"total"`
	Items    []TrackObject `json:"items"`
}

type TrackObject struct {
	Album    AlbumObject    `json:"album"`
	Artists  []ArtistObject `json:"artists"`
	Explicit bool           `json:"explicit"`
	Href     string         `json:"href"`
	ID       string         `json:"id"`
	Name     string         `json:"name"`
}

type AlbumObject struct {
	AlbumType   string       `json:"album_type"`
	TotalTracks int          `json:"total_tracks"`
	Images      []AlbumImage `json:"images"`
	Name        string       `json:"name"`
}

type AlbumImage struct {
	URL string `json:"url"`
}

type ArtistObject struct {
	Href string `json:"href"`
	Name string `json:"name"`
}
