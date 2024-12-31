package usertrack

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"music_catalog/internal/model/usertrack"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormdb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

	now := time.Now()
	isLiked := true
	type args struct {
		model *usertrack.UserTrack
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{model: &usertrack.UserTrack{
				Model: gorm.Model{
					CreatedAt: now,
					UpdatedAt: now,
				},
				UserID:  1,
				TrackID: "5MOCeDoizSpQ4FnpX8VFky",
				IsLiked: &isLiked,
			}},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(`INSERT INTO "user_tracks" (.+) VALUES (.+)`).
					WithArgs(
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						args.model.UserID,
						args.model.TrackID,
						args.model.IsLiked,
					).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(uint(1)))

				mock.ExpectCommit()
			},
		},
		{
			name: "error",
			args: args{model: &usertrack.UserTrack{
				Model: gorm.Model{
					CreatedAt: now,
					UpdatedAt: now,
				},
				UserID:  1,
				TrackID: "5MOCeDoizSpQ4FnpX8VFky",
				IsLiked: &isLiked,
			}},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(`INSERT INTO "user_tracks" (.+) VALUES (.+)`).
					WithArgs(
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						args.model.UserID,
						args.model.TrackID,
						args.model.IsLiked,
					).WillReturnError(assert.AnError)

				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &Repository{
				db: gormdb,
			}
			if err := r.Create(tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_GetAllLiked(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormdb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

	now := time.Now()
	isLiked := true
	type args struct {
		UserID  uint
		TrackID []string
	}

	tests := []struct {
		name    string
		args    args
		want    map[string]usertrack.UserTrack
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				UserID:  1,
				TrackID: []string{"5MOCeDoizSpQ4FnpX8VFky"},
			},
			want: map[string]usertrack.UserTrack{
				"5MOCeDoizSpQ4FnpX8VFky": {
					Model: gorm.Model{
						ID:        1,
						CreatedAt: now,
						UpdatedAt: now,
					},
					UserID:  1,
					TrackID: "5MOCeDoizSpQ4FnpX8VFky",
					IsLiked: &isLiked,
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectQuery(`SELECT \* FROM "user_tracks" .+`).WithArgs(args.UserID, strings.Join(args.TrackID, ",")).
					WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "user_id", "track_id", "is_liked"}).
						AddRow(1, now, now, 1, "5MOCeDoizSpQ4FnpX8VFky", true))
			},
		},
		{
			name: "failed",
			args: args{
				UserID:  1,
				TrackID: []string{"5MOCeDoizSpQ4FnpX8VFky"},
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectQuery(`SELECT \* FROM "user_tracks" .+`).WithArgs(args.UserID, strings.Join(args.TrackID, ",")).
					WillReturnError(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &Repository{
				db: gormdb,
			}
			got, err := r.GetAllLiked(tt.args.UserID, tt.args.TrackID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllLiked() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllLiked() = %v, want %v", got, tt.want)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_GetLikedById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormdb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

	now := time.Now()
	isLiked := true
	type args struct {
		UserID  uint
		TrackID string
	}

	tests := []struct {
		name    string
		args    args
		want    *usertrack.UserTrack
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				UserID:  1,
				TrackID: "5MOCeDoizSpQ4FnpX8VFky",
			},
			want: &usertrack.UserTrack{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: now,
					UpdatedAt: now,
				},
				UserID:  1,
				TrackID: "5MOCeDoizSpQ4FnpX8VFky",
				IsLiked: &isLiked,
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectQuery(`SELECT \* FROM "user_tracks" .+`).WithArgs(args.UserID, args.TrackID, 1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "user_id", "track_id", "is_liked"}).
						AddRow(1, now, now, 1, "5MOCeDoizSpQ4FnpX8VFky", true))
			},
		},
		{
			name: "failed",
			args: args{
				UserID:  1,
				TrackID: "5MOCeDoizSpQ4FnpX8VFky",
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectQuery(`SELECT \* FROM "user_tracks" .+`).WithArgs(args.UserID, args.TrackID, 1).
					WillReturnError(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &Repository{
				db: gormdb,
			}
			got, err := r.GetLikedById(tt.args.UserID, tt.args.TrackID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLikedById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLikedById() = %v, want %v", got, tt.want)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormdb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

	now := time.Now()
	isLiked := true
	type args struct {
		model *usertrack.UserTrack
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{model: &usertrack.UserTrack{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: now,
					UpdatedAt: now,
				},
				UserID:  1,
				TrackID: "5MOCeDoizSpQ4FnpX8VFky",
				IsLiked: &isLiked,
			}},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectExec(`UPDATE "user_tracks" SET (.+) WHERE (.+)`).
					WithArgs(
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						args.model.UserID,
						args.model.TrackID,
						args.model.IsLiked,
						args.model.ID,
					).WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
		},
		{
			name: "error",
			args: args{model: &usertrack.UserTrack{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: now,
					UpdatedAt: now,
				},
				UserID:  1,
				TrackID: "5MOCeDoizSpQ4FnpX8VFky",
				IsLiked: &isLiked,
			}},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectExec(`UPDATE "user_tracks" SET (.+) WHERE (.+)`).
					WithArgs(
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						args.model.UserID,
						args.model.TrackID,
						args.model.IsLiked,
						args.model.ID,
					).WillReturnError(assert.AnError)

				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &Repository{
				db: gormdb,
			}
			if err := r.Update(tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
