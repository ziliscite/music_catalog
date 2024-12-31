package spotify

type SearchResponse struct {
	Limit  int           `json:"limit"`
	Offset int           `json:"offset"`
	Items  []TrackObject `json:"items"`
	Total  int           `json:"total"`
}

type TrackObject struct {
	// album-related fields
	AlbumType        string   `json:"albumType"`
	AlbumTotalTracks int      `json:"totalTracks"`
	AlbumImagesURL   []string `json:"albumImagesURL"`
	AlbumName        string   `json:"albumName"`

	// artist-related fields
	ArtistsName []string `json:"artistsName"`

	// track related fields
	Explicit bool   `json:"explicit"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsLiked  *bool  `json:"isLiked"`
}

type RecommendationResponse struct {
	Items []TrackObject `json:"items"`
}
