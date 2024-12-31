// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -source=service.go -destination=service_mock.go -package=track
//

// Package track is a generated GoMock package.
package track

import (
	usertrack "music_catalog/internal/model/usertrack"
	spotify "music_catalog/internal/repository/spotify"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRepositorySpotify is a mock of RepositorySpotify interface.
type MockRepositorySpotify struct {
	ctrl     *gomock.Controller
	recorder *MockRepositorySpotifyMockRecorder
	isgomock struct{}
}

// MockRepositorySpotifyMockRecorder is the mock recorder for MockRepositorySpotify.
type MockRepositorySpotifyMockRecorder struct {
	mock *MockRepositorySpotify
}

// NewMockRepositorySpotify creates a new mock instance.
func NewMockRepositorySpotify(ctrl *gomock.Controller) *MockRepositorySpotify {
	mock := &MockRepositorySpotify{ctrl: ctrl}
	mock.recorder = &MockRepositorySpotifyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositorySpotify) EXPECT() *MockRepositorySpotifyMockRecorder {
	return m.recorder
}

// Search mocks base method.
func (m *MockRepositorySpotify) Search(q string, limit, offset int) (*spotify.ClientSearchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", q, limit, offset)
	ret0, _ := ret[0].(*spotify.ClientSearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockRepositorySpotifyMockRecorder) Search(q, limit, offset any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRepositorySpotify)(nil).Search), q, limit, offset)
}

// MockRepositoryUserTrack is a mock of RepositoryUserTrack interface.
type MockRepositoryUserTrack struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryUserTrackMockRecorder
	isgomock struct{}
}

// MockRepositoryUserTrackMockRecorder is the mock recorder for MockRepositoryUserTrack.
type MockRepositoryUserTrackMockRecorder struct {
	mock *MockRepositoryUserTrack
}

// NewMockRepositoryUserTrack creates a new mock instance.
func NewMockRepositoryUserTrack(ctrl *gomock.Controller) *MockRepositoryUserTrack {
	mock := &MockRepositoryUserTrack{ctrl: ctrl}
	mock.recorder = &MockRepositoryUserTrackMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryUserTrack) EXPECT() *MockRepositoryUserTrackMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepositoryUserTrack) Create(userTrack *usertrack.UserTrack) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", userTrack)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryUserTrackMockRecorder) Create(userTrack any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepositoryUserTrack)(nil).Create), userTrack)
}

// GetAllLiked mocks base method.
func (m *MockRepositoryUserTrack) GetAllLiked(userId uint, trackIds []string) (map[string]usertrack.UserTrack, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllLiked", userId, trackIds)
	ret0, _ := ret[0].(map[string]usertrack.UserTrack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllLiked indicates an expected call of GetAllLiked.
func (mr *MockRepositoryUserTrackMockRecorder) GetAllLiked(userId, trackIds any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllLiked", reflect.TypeOf((*MockRepositoryUserTrack)(nil).GetAllLiked), userId, trackIds)
}

// GetLikedById mocks base method.
func (m *MockRepositoryUserTrack) GetLikedById(userId uint, trackId string) (*usertrack.UserTrack, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLikedById", userId, trackId)
	ret0, _ := ret[0].(*usertrack.UserTrack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLikedById indicates an expected call of GetLikedById.
func (mr *MockRepositoryUserTrackMockRecorder) GetLikedById(userId, trackId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLikedById", reflect.TypeOf((*MockRepositoryUserTrack)(nil).GetLikedById), userId, trackId)
}

// Update mocks base method.
func (m *MockRepositoryUserTrack) Update(userTrack *usertrack.UserTrack) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", userTrack)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryUserTrackMockRecorder) Update(userTrack any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepositoryUserTrack)(nil).Update), userTrack)
}
