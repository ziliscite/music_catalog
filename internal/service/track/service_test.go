package track

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"music_catalog/internal/config"
	model "music_catalog/internal/model/spotify"
	"music_catalog/internal/repository/spotify"
	"reflect"
	"testing"
)

func TestService_Search(t *testing.T) {
	ctrMock := gomock.NewController(t)
	defer ctrMock.Finish()

	mockRepository := NewMockRepository(ctrMock)
	next := "https://api.spotify.com/v1/search?offset=2&limit=2&query=A%20Little%20Death&type=track&market=US&locale=en-US,en;q%3D0.9,id;q%3D0.8"

	type args struct {
		q         string
		pageSize  int
		pageIndex int
	}

	tests := []struct {
		name    string
		args    args
		want    *model.SearchResponse
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				q:         "A Little Death",
				pageSize:  2,
				pageIndex: 1,
			},
			want: &model.SearchResponse{
				Limit:  2,
				Offset: 0,
				Total:  900,
				Items: []model.TrackObject{
					{
						AlbumType:        "single",
						AlbumTotalTracks: 2,
						AlbumImagesURL:   []string{"https://i.scdn.co/image/ab67616d0000b2736492453ee238cd8546c6850e", "https://i.scdn.co/image/ab67616d00001e026492453ee238cd8546c6850e", "https://i.scdn.co/image/ab67616d000048516492453ee238cd8546c6850e"},
						AlbumName:        "Thank You,",
						ArtistsName:      []string{"The Neighbourhood"},
						Explicit:         false,
						ID:               "0Ot6e3wYVQQ1Us9PM977jE",
						Name:             "A Little Death",
						IsLiked:          nil,
					},
					{
						AlbumType:        "album",
						AlbumTotalTracks: 11,
						AlbumImagesURL:   []string{"https://i.scdn.co/image/ab67616d0000b2736019b1ddb28634421cc291a0", "https://i.scdn.co/image/ab67616d00001e026019b1ddb28634421cc291a0", "https://i.scdn.co/image/ab67616d000048516019b1ddb28634421cc291a0"},
						AlbumName:        "What Do You Really Know?",
						ArtistsName:      []string{"Reality Club"},
						Explicit:         false,
						ID:               "6ZGgaShxOimGDfRz1T1zje",
						Name:             "Alexandra",
						IsLiked:          nil,
					},
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepository.EXPECT().Search(args.q, 2, 0).Return(&spotify.ClientSearchResponse{
					Tracks: spotify.Tracks{
						Href:   "https://api.spotify.com/v1/search?offset=0&limit=2&query=A%20Little%20Death&type=track&market=US&locale=en-US,en;q%3D0.9,id;q%3D0.8",
						Limit:  2,
						Next:   &next,
						Offset: 0,
						Total:  900,
						Items: []spotify.TrackObject{
							{
								Album: spotify.AlbumObject{
									AlbumType:   "single",
									TotalTracks: 2,
									Images: []spotify.AlbumImage{
										{URL: "https://i.scdn.co/image/ab67616d0000b2736492453ee238cd8546c6850e"},
										{URL: "https://i.scdn.co/image/ab67616d00001e026492453ee238cd8546c6850e"},
										{URL: "https://i.scdn.co/image/ab67616d000048516492453ee238cd8546c6850e"},
									},
									Name: "Thank You,",
								},
								Artists: []spotify.ArtistObject{
									{
										Href: "https://api.spotify.com/v1/artists/77SW9BnxLY8rJ0RciFqkHh",
										Name: "The Neighbourhood",
									},
								},
								Explicit: false,
								Href:     "https://api.spotify.com/v1/tracks/0Ot6e3wYVQQ1Us9PM977jE",
								ID:       "0Ot6e3wYVQQ1Us9PM977jE",
								Name:     "A Little Death",
							},
							{
								Album: spotify.AlbumObject{
									AlbumType:   "album",
									TotalTracks: 11,
									Images: []spotify.AlbumImage{
										{URL: "https://i.scdn.co/image/ab67616d0000b2736019b1ddb28634421cc291a0"},
										{URL: "https://i.scdn.co/image/ab67616d00001e026019b1ddb28634421cc291a0"},
										{URL: "https://i.scdn.co/image/ab67616d000048516019b1ddb28634421cc291a0"},
									},
									Name: "What Do You Really Know?",
								},
								Artists: []spotify.ArtistObject{
									{
										Href: "https://api.spotify.com/v1/artists/1DjZI46mVZZZYmmmygRnTw",
										Name: "Reality Club",
									},
								},
								Explicit: false,
								Href:     "https://api.spotify.com/v1/tracks/6ZGgaShxOimGDfRz1T1zje",
								ID:       "6ZGgaShxOimGDfRz1T1zje",
								Name:     "Alexandra",
							},
						},
					},
				}, nil)
			},
		},
		{
			name: "error",
			args: args{
				q:         "A Little Death",
				pageSize:  2,
				pageIndex: 1,
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockRepository.EXPECT().Search(args.q, 2, 0).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			s := &Service{
				cfg: &config.Config{
					SpotifyConfig: config.SpotifyConfig{
						Url: "https://api.spotify.com/v1",
					},
				},
				r: mockRepository,
			}

			got, err := s.Search(tt.args.q, tt.args.pageSize, tt.args.pageIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
		})
	}
}
