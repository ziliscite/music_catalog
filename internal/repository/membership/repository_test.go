package membership

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"music_catalog/internal/model/membership"
	"testing"
	"time"
)

func TestRepository_Create(t *testing.T) {
	// Add go mock -- https://github.com/DATA-DOG/go-sqlmock
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	gormdb, err := gorm.Open(
		postgres.New(
			postgres.Config{
				Conn: db,
			},
		),
	)
	assert.NoError(t, err)

	type args struct {
		user *membership.User
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFnc func(args args)
	}{
		{
			name: "success",
			args: args{
				user: &membership.User{
					Email:    "test@gmail.com",
					Username: "username",
					Password: "password",
				},
			},
			wantErr: false,
			mockFnc: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(`INSERT INTO "users" \("created_at","updated_at","deleted_at","email","username","password"\) VALUES \(\$1,\$2,\$3,\$4,\$5,\$6\) RETURNING "id"`).
					WithArgs(
						sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
						args.user.Email, args.user.Username, args.user.Password,
					).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectCommit()
			},
		},
		{
			name: "error",
			args: args{
				user: &membership.User{
					Email:    "test@gmail.com",
					Username: "username",
					Password: "password",
				},
			},
			wantErr: true,
			mockFnc: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(`INSERT INTO "users" \("created_at","updated_at","deleted_at","email","username","password"\) VALUES \(\$1,\$2,\$3,\$4,\$5,\$6\) RETURNING "id"`).
					WithArgs(
						sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
						args.user.Email, args.user.Username, args.user.Password,
					).
					WillReturnError(assert.AnError)

				mock.ExpectRollback()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFnc(tt.args)
			r := &Repository{
				db: gormdb,
			}
			if err := r.Create(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_GetByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	gormdb, err := gorm.Open(
		postgres.New(
			postgres.Config{
				Conn: db,
			},
		),
	)
	assert.NoError(t, err)
	now := time.Now()

	type args struct {
		email string
	}

	tests := []struct {
		name    string
		args    args
		want    *membership.User
		wantErr bool
		mockFnc func(args args)
	}{
		{
			name: "success",
			args: args{
				email: "test@gmail.com",
			},
			want: &membership.User{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: now,
					UpdatedAt: now,
				},
				Email:    "test@gmail.com",
				Username: "testusername",
				Password: "password",
			},
			wantErr: false,
			mockFnc: func(args args) {
				mock.ExpectQuery(`SELECT \* FROM "users" WHERE email = \$1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT \$2`).
					WithArgs(args.email, 1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "email", "username", "password"}).
						AddRow(1, now, now, "test@gmail.com", "testusername", "password"))
			},
		},
		{
			name: "failed",
			args: args{
				email: "test@gmail.com",
			},
			want:    nil,
			wantErr: true,
			mockFnc: func(args args) {
				mock.ExpectQuery(`SELECT \* FROM "users" WHERE email = \$1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT \$2`).
					WithArgs(args.email, 1).
					WillReturnError(assert.AnError)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFnc(tt.args)
			r := &Repository{
				db: gormdb,
			}
			got, err := r.GetByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByEmail(%v), wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equalf(t, tt.want, got, "GetByEmail(%v)", tt.args.email)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_GetByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	gormdb, err := gorm.Open(
		postgres.New(
			postgres.Config{
				Conn: db,
			},
		),
	)
	assert.NoError(t, err)
	now := time.Now()

	type args struct {
		username string
	}

	tests := []struct {
		name    string
		args    args
		want    *membership.User
		wantErr bool
		mockFnc func(args args)
	}{
		{
			name: "success",
			args: args{
				username: "testusername",
			},
			want: &membership.User{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: now,
					UpdatedAt: now,
				},
				Email:    "test@gmail.com",
				Username: "testusername",
				Password: "password",
			},
			wantErr: false,
			mockFnc: func(args args) {
				mock.ExpectQuery(`SELECT \* FROM "users" WHERE email = \$1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT \$2`).
					WithArgs(args.username, 1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "email", "username", "password"}).
						AddRow(1, now, now, "test@gmail.com", "testusername", "password"))
			},
		},
		{
			name: "failed",
			args: args{
				username: "testusername",
			},
			want:    nil,
			wantErr: true,
			mockFnc: func(args args) {
				mock.ExpectQuery(`SELECT \* FROM "users" WHERE email = \$1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT \$2`).
					WithArgs(args.username, 1).
					WillReturnError(assert.AnError)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFnc(tt.args)
			r := &Repository{
				db: gormdb,
			}
			got, err := r.GetByEmail(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByUsername(%v), wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equalf(t, tt.want, got, "GetByUsername(%v)", tt.args.username)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
