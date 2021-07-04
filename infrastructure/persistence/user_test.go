package persistence_test

import (
	"testing"
	"time"

	"go-firebase-auth-server/domain/entity"
	"go-firebase-auth-server/infrastructure/persistence"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUserRepository_GetByUID(t *testing.T) {
	t.Parallel()

	now := time.Date(2021, 7, 5, 12, 13, 24, 0, time.UTC)

	tests := map[string]struct {
		uid  string
		want *entity.User
		err  error
	}{
		"success": {
			uid: "uid1",
			want: &entity.User{
				ID:        1,
				UID:       "uid1",
				Username:  "Tom",
				Email:     "tom@test.com",
				CreatedAt: now,
				UpdatedAt: now,
			},
			err: nil,
		},
		"failure": {
			uid:  "failure",
			want: nil,
			err:  gorm.ErrRecordNotFound,
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			testDB := openTestDB()
			txDB, err := testDB.DB()
			assert.NoError(t, err)
			defer txDB.Close()

			testDB.Create(&entity.User{ID: 1, UID: "uid1", Username: "Tom", Email: "tom@test.com", CreatedAt: now, UpdatedAt: now})

			userRepository := persistence.NewUserRepository(testDB)
			got, err := userRepository.GetByUID(tt.uid)
			assert.Equal(t, tt.err, err)
			assert.Empty(t, cmp.Diff(tt.want, got))
		})
	}
}
