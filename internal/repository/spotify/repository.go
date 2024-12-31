package spotify

import (
	"encoding/json"
	"fmt"
	"music_catalog/internal/config"
	"music_catalog/pkg"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Repository struct {
	cfg         *config.Config
	client      pkg.HTTPClient
	AccessToken string
	TokenType   string
	ExpiresAt   *time.Time
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func NewRepository(cfg *config.Config, client pkg.HTTPClient) *Repository {
	return &Repository{
		cfg:       cfg,
		client:    client,
		ExpiresAt: &time.Time{},
	}
}

func (r *Repository) GetToken() (string, string, error) {
	// Check if token is expired
	if r.AccessToken != "" || !time.Now().After(*r.ExpiresAt) {
		return r.AccessToken, r.TokenType, nil
	}

	// Create form
	form := url.Values{}
	form.Add("grant_type", "client_credentials")
	form.Add("client_id", r.cfg.SpotifyConfig.ClientID)
	form.Add("client_secret", r.cfg.SpotifyConfig.ClientSecret)

	accessToken, tokenType, expiresAt, err := requestToken(form.Encode(), r.client)
	if err != nil {
		return "", "", err
	}

	r.AccessToken = accessToken
	r.TokenType = tokenType
	r.ExpiresAt = expiresAt

	return accessToken, tokenType, nil
}

func requestToken(encodedUrl string, client pkg.HTTPClient) (string, string, *time.Time, error) {
	// create request
	req, err := http.NewRequest(http.MethodPost, `https://accounts.spotify.com/api/token`, strings.NewReader(encodedUrl))
	if err != nil {
		return "", "", nil, err
	}

	// set header, docs: https://developer.spotify.com/documentation/web-api
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return "", "", nil, err
	}
	defer resp.Body.Close()

	var response TokenResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", "", nil, err
	}

	expiresAt := time.Now().Add(time.Second * time.Duration(response.ExpiresIn))
	return response.AccessToken, response.TokenType, &expiresAt, nil
}

func (r *Repository) Search(q string, limit, offset int) (*ClientSearchResponse, error) {
	endpoint := url.Values{
		"type":   []string{"track"},
		"q":      []string{q},
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset)},
	}

	// use endpoint.Encode() so that it wont break when the query have spaces in it. And other things
	path := fmt.Sprintf("%s/search?%s", r.cfg.SpotifyConfig.Url, endpoint.Encode())

	req, err := http.NewRequest(http.MethodGet, path, nil) // Get request ain't needing body
	if err != nil {
		return nil, err
	}

	accessToken, tokenType, err := r.GetToken()
	if err != nil {
		return nil, err
	}

	auth := fmt.Sprintf("%s %s", tokenType, accessToken)
	req.Header.Set("Authorization", auth)

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response ClientSearchResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
