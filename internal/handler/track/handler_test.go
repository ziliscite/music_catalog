package track

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	model "music_catalog/internal/model/spotify"
	"music_catalog/pkg"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Search(t *testing.T) {
	ctrMock := gomock.NewController(t)
	defer ctrMock.Finish()

	mockService := NewMockService(ctrMock)

	tests := []struct {
		name               string
		mockFn             func()
		expectedStatusCode int
		expectedBody       model.SearchResponse
		wantErr            bool
	}{
		{
			name: "success",
			mockFn: func() {
				mockService.EXPECT().Search("A Little Death", 2, 1).Return(&model.SearchResponse{
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
				}, nil)
			},
			wantErr:            false,
			expectedStatusCode: 200,
			expectedBody: model.SearchResponse{
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
		},
		{
			name:               "failed",
			expectedStatusCode: 400,
			expectedBody:       model.SearchResponse{},
			wantErr:            true,
			mockFn: func() {
				mockService.EXPECT().Search("A Little Death", 2, 1).Return(nil, assert.AnError)
			},
		},
	}

	engine := gin.New()
	h := &Handler{
		Engine: engine,
		s:      mockService,
	}
	h.RegisterRoutes()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()

			w := httptest.NewRecorder()

			endpoint := `/track/search?q=A+Little+Death&pageSize=2&pageIndex=1`

			req, err := http.NewRequest(http.MethodGet, endpoint, nil)
			assert.NoError(t, err)
			token, err := pkg.CreateToken(1, "username", "")
			assert.NoError(t, err)
			req.Header.Set("Authorization", token)

			h.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)

			if !tt.wantErr {
				res := w.Result()
				defer res.Body.Close()

				response := model.SearchResponse{}
				err = json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody, response)
			}
		})
	}
}
