package track

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
	"music_catalog/internal/config"
	model "music_catalog/internal/model/spotify"
	"music_catalog/internal/model/usertrack"
	"music_catalog/internal/repository/spotify"
	"reflect"
	"testing"
)

func TestService_Search(t *testing.T) {
	ctrMock := gomock.NewController(t)
	defer ctrMock.Finish()

	mockSpotifyRepository := NewMockRepositorySpotify(ctrMock)
	mockUserTrackRepository := NewMockRepositoryUserTrack(ctrMock)

	next := "https://api.spotify.com/v1/search?offset=2&limit=2&query=A%20Little%20Death&type=track&market=US&locale=en-US,en;q%3D0.9,id;q%3D0.8"
	isLikedTrue := true
	isLikedFalse := false

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
						IsLiked:          &isLikedTrue,
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
						IsLiked:          &isLikedFalse,
					},
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockSpotifyRepository.EXPECT().Search(args.q, 2, 0).Return(&spotify.ClientSearchResponse{
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

				mockUserTrackRepository.EXPECT().GetAllLiked(uint(1), []string{"0Ot6e3wYVQQ1Us9PM977jE", "6ZGgaShxOimGDfRz1T1zje"}).
					Return(map[string]usertrack.UserTrack{
						"0Ot6e3wYVQQ1Us9PM977jE": {
							IsLiked: &isLikedTrue,
						},
						"6ZGgaShxOimGDfRz1T1zje": {
							IsLiked: &isLikedFalse,
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
				mockSpotifyRepository.EXPECT().Search(args.q, 2, 0).Return(nil, assert.AnError)
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
				rs:  mockSpotifyRepository,
				rut: mockUserTrackRepository,
			}

			got, err := s.Search(tt.args.q, tt.args.pageSize, tt.args.pageIndex, 1)
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

func TestService_Upsert(t *testing.T) {
	ctrMock := gomock.NewController(t)
	defer ctrMock.Finish()

	mockUserTrackRepository := NewMockRepositoryUserTrack(ctrMock)

	isLikedTrue := true
	isLikedFalse := false

	type args struct {
		userId  uint
		request *usertrack.LikeRequest
	}

	tests := []struct {
		name      string
		args      args
		wantErr   bool
		isCreated bool
		mockFn    func(args args)
	}{
		{
			name: "success: create",
			args: args{
				userId: 1,
				request: &usertrack.LikeRequest{
					TrackID: "5MOCeDoizSpQ4FnpX8VFky",
					IsLiked: &isLikedTrue,
				},
			},
			wantErr:   false,
			isCreated: true,
			mockFn: func(args args) {
				mockUserTrackRepository.EXPECT().GetLikedById(args.userId, args.request.TrackID).Return(nil, gorm.ErrRecordNotFound)
				mockUserTrackRepository.EXPECT().Create(args.request.Model(args.userId)).Return(nil)
			},
		},
		{
			name: "success: update from disliked to liked",
			args: args{
				userId: 1,
				request: &usertrack.LikeRequest{
					TrackID: "5MOCeDoizSpQ4FnpX8VFky",
					IsLiked: &isLikedFalse,
				},
			},
			wantErr:   false,
			isCreated: false,
			mockFn: func(args args) {
				userTracker := &usertrack.UserTrack{
					UserID:  args.userId,
					TrackID: args.request.TrackID,
					IsLiked: &isLikedFalse,
				}
				// expect existing user track with like false
				mockUserTrackRepository.EXPECT().GetLikedById(args.userId, args.request.TrackID).Return(userTracker, nil)

				// update user track
				userTracker.IsLiked = args.request.IsLiked
				mockUserTrackRepository.EXPECT().Update(userTracker).Return(nil)
			},
		},
		{
			name: "failed: get",
			args: args{
				userId: 1,
				request: &usertrack.LikeRequest{
					TrackID: "5MOCeDoizSpQ4FnpX8VFky",
					IsLiked: &isLikedTrue,
				},
			},
			wantErr:   true,
			isCreated: false,
			mockFn: func(args args) {
				mockUserTrackRepository.EXPECT().GetLikedById(args.userId, args.request.TrackID).Return(nil, assert.AnError)
			},
		},
		{
			name: "failed: create",
			args: args{
				userId: 1,
				request: &usertrack.LikeRequest{
					TrackID: "5MOCeDoizSpQ4FnpX8VFky",
					IsLiked: &isLikedTrue,
				},
			},
			wantErr:   true,
			isCreated: true,
			mockFn: func(args args) {
				mockUserTrackRepository.EXPECT().GetLikedById(args.userId, args.request.TrackID).Return(nil, gorm.ErrRecordNotFound)
				mockUserTrackRepository.EXPECT().Create(args.request.Model(args.userId)).Return(assert.AnError)
			},
		},
		{
			name: "failed: update",
			args: args{
				userId: 1,
				request: &usertrack.LikeRequest{
					TrackID: "5MOCeDoizSpQ4FnpX8VFky",
					IsLiked: &isLikedTrue,
				},
			},
			wantErr:   true,
			isCreated: false,
			mockFn: func(args args) {
				userTracker := &usertrack.UserTrack{
					UserID:  args.userId,
					TrackID: args.request.TrackID,
					IsLiked: &isLikedFalse,
				}
				mockUserTrackRepository.EXPECT().GetLikedById(args.userId, args.request.TrackID).Return(userTracker, nil)
				mockUserTrackRepository.EXPECT().Update(userTracker).Return(assert.AnError)
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
				rut: mockUserTrackRepository,
			}

			created, err := s.Upsert(tt.args.userId, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Upsert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if created != tt.isCreated {
				t.Errorf("Upsert() created = %v, want %v", created, tt.isCreated)
			}
		})
	}
}
