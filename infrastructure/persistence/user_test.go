package persistence_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/yasszu/go-firebase-auth-server/domain/entity"
	"github.com/yasszu/go-firebase-auth-server/infrastructure/persistence"
)

func TestUserRepository_Crete(t *testing.T) {
	tests := map[string]struct {
		user *entity.User
		want bool
	}{
		"it should create new user": {
			user: &entity.User{
				UID:      "abcdefg",
				Username: "name",
				Email:    "name@example.com",
			},
			want: true,
		},
		"it should not create new user when uid is duplicate": {
			user: &entity.User{
				UID:      "uid1",
				Username: "name",
				Email:    "name@example.com",
			},
			want: false,
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			testDB := openTestDB()
			txDB, err := testDB.DB()
			assert.NoError(t, err)
			defer txDB.Close()

			testDB.Create(&entity.User{ID: 1, UID: "uid1", Username: "Tom", Email: "tom@test.com"})

			userRepository := persistence.NewUserRepository(testDB)
			err = userRepository.Crete(tt.user)
			assert.Equal(t, tt.want, err == nil)
		})
	}
}

func TestUserRepository_GetByUID(t *testing.T) {
	testDB := openTestDB()
	testDB.Create(&entity.User{
		ID:        1,
		UID:       "uid1",
		Username:  "Tom",
		Email:     "tom@test.com",
		CreatedAt: now(),
		UpdatedAt: now(),
	})

	tests := map[string]struct {
		uid  entity.UID
		want *entity.User
	}{
		"success": {
			uid: "uid1",
			want: &entity.User{
				ID:        1,
				UID:       "uid1",
				Username:  "Tom",
				Email:     "tom@test.com",
				CreatedAt: now(),
				UpdatedAt: now(),
			},
		},
		"not found": {
			uid:  "illegal-uid",
			want: nil,
		},
		"request empty": {
			uid:  "",
			want: nil,
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			userRepository := persistence.NewUserRepository(testDB)
			got, err := userRepository.GetByUID(tt.uid)
			assert.NoError(t, err)
			assert.Empty(t, cmp.Diff(tt.want, got))
		})
	}
}
