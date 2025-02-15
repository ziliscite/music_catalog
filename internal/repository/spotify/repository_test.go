package spotify

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"io"
	"music_catalog/internal/config"
	"music_catalog/pkg"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestRepository_Search(t *testing.T) {
	ctrMock := gomock.NewController(t)
	defer ctrMock.Finish()

	mockClient := pkg.NewMockHTTPClient(ctrMock)

	next := "https://api.spotify.com/v1/search?offset=2&limit=2&query=A%20Little%20Death&type=track&market=US&locale=en-US,en;q%3D0.9,id;q%3D0.8"
	type args struct {
		q      string
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		args    args
		want    *ClientSearchResponse
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				q:      "A Little Death",
				limit:  2,
				offset: 0,
			},
			want: &ClientSearchResponse{
				Tracks: Tracks{
					Href:   "https://api.spotify.com/v1/search?offset=0&limit=2&query=A%20Little%20Death&type=track&market=US&locale=en-US,en;q%3D0.9,id;q%3D0.8",
					Limit:  2,
					Next:   &next,
					Offset: 0,
					Total:  900,
					Items: []TrackObject{
						{
							Album: AlbumObject{
								AlbumType:   "single",
								TotalTracks: 2,
								Images: []AlbumImage{
									{URL: "https://i.scdn.co/image/ab67616d0000b2736492453ee238cd8546c6850e"},
									{URL: "https://i.scdn.co/image/ab67616d00001e026492453ee238cd8546c6850e"},
									{URL: "https://i.scdn.co/image/ab67616d000048516492453ee238cd8546c6850e"},
								},
								Name: "Thank You,",
							},
							Artists: []ArtistObject{
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
							Album: AlbumObject{
								AlbumType:   "album",
								TotalTracks: 11,
								Images: []AlbumImage{
									{URL: "https://i.scdn.co/image/ab67616d0000b2736019b1ddb28634421cc291a0"},
									{URL: "https://i.scdn.co/image/ab67616d00001e026019b1ddb28634421cc291a0"},
									{URL: "https://i.scdn.co/image/ab67616d000048516019b1ddb28634421cc291a0"},
								},
								Name: "What Do You Really Know?",
							},
							Artists: []ArtistObject{
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
			},
			wantErr: false,
			mockFn: func(args args) {
				params := url.Values{}
				params.Set("q", args.q)
				params.Set("type", "track")
				params.Set("limit", strconv.Itoa(args.limit))
				params.Set("offset", strconv.Itoa(args.offset))

				basePath := `https://api.spotify.com/v1/search`
				urlPath := fmt.Sprintf("%s?%s", basePath, params.Encode())
				req, err := http.NewRequest(http.MethodGet, urlPath, nil)
				assert.NoError(t, err)

				req.Header.Set("Authorization", "Bearer accessToken")
				mockClient.EXPECT().Do(req).Return(&http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewBufferString(searchResponse)),
				}, nil)
			},
		},
		{
			name: "error",
			args: args{
				q:      "A Little Death",
				limit:  2,
				offset: 0,
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				params := url.Values{}
				params.Set("q", args.q)
				params.Set("type", "track")
				params.Set("limit", strconv.Itoa(args.limit))
				params.Set("offset", strconv.Itoa(args.offset))

				basePath := `https://api.spotify.com/v1/search`
				urlPath := fmt.Sprintf("%s?%s", basePath, params.Encode())
				req, err := http.NewRequest(http.MethodGet, urlPath, nil)
				assert.NoError(t, err)

				req.Header.Set("Authorization", "Bearer accessToken")
				mockClient.EXPECT().Do(req).Return(&http.Response{
					StatusCode: 500,
					Body:       io.NopCloser(bytes.NewBufferString(`Internal Server Error`)),
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			now := time.Now().Add(1 * time.Hour)
			r := &Repository{
				cfg: &config.Config{
					SpotifyConfig: config.SpotifyConfig{
						Url: "https://api.spotify.com/v1",
					},
				},
				client:      mockClient,
				AccessToken: "accessToken",
				TokenType:   "Bearer",
				ExpiresAt:   &now,
			}

			got, err := r.Search(tt.args.q, tt.args.limit, tt.args.offset)
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
