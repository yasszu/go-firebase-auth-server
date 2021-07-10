package usecase_test

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"go-firebase-auth-server/domain/entity"
	"go-firebase-auth-server/domain/repository"
	"go-firebase-auth-server/domain/repository/mock"
)

func Test_userUsecase_GetUser(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		uid     entity.UID
		prepare func(ctrl *gomock.Controller) repository.UserRepository
		want    *entity.User
		err     error
	}{
		"it should returns user": {
			uid: entity.UID("uid123"),
			prepare: func(ctrl *gomock.Controller) repository.UserRepository {
				userRepository := mock.NewMockUserRepository(ctrl)
				userRepository.EXPECT().GetByUID(entity.UID("uid123")).Return(&entity.User{
					ID:        1,
					UID:       "uid123",
					Username:  "Chuck",
					Email:     "chuck@example.com",
					CreatedAt: time.Date(2021, 1, 2, 3, 4, 5, 6, time.UTC),
					UpdatedAt: time.Date(2021, 1, 2, 3, 4, 5, 6, time.UTC),
				}, nil)
				return userRepository
			},
			want: &entity.User{
				ID:        1,
				UID:       "uid123",
				Username:  "Chuck",
				Email:     "chuck@example.com",
				CreatedAt: time.Date(2021, 1, 2, 3, 4, 5, 6, time.UTC),
				UpdatedAt: time.Date(2021, 1, 2, 3, 4, 5, 6, time.UTC),
			},
			err: nil,
		},
		"it should returns NotFountError when user is nil": {
			uid: entity.UID("uid000"),
			prepare: func(ctrl *gomock.Controller) repository.UserRepository {
				userRepository := mock.NewMockUserRepository(ctrl)
				userRepository.EXPECT().GetByUID(entity.UID("uid000")).Return(nil, nil)
				return userRepository
			},
			want: nil,
			err:  &entity.NotFoundError{Name: "user"},
		},
		"it should returns UnexpectedError when err occurs on DB": {
			uid: entity.UID("uid000"),
			prepare: func(ctrl *gomock.Controller) repository.UserRepository {
				userRepository := mock.NewMockUserRepository(ctrl)
				userRepository.EXPECT().GetByUID(entity.UID("uid000")).Return(nil, errors.New("test"))
				return userRepository
			},
			want: nil,
			err:  &entity.UnexpectedError{Err: errors.New("test")},
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			r := tt.prepare(ctrl)
			got, err := r.GetByUID(tt.uid)
			if err != nil {
				assert.Error(t, tt.err, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
