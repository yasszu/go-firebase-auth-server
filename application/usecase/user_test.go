package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yasszu/go-firebase-auth-server/application/usecase"
	"github.com/yasszu/go-firebase-auth-server/domain/entity"
	"github.com/yasszu/go-firebase-auth-server/domain/repository/mock"
	repository "github.com/yasszu/go-firebase-auth-server/domain/repository/mock"
	service "github.com/yasszu/go-firebase-auth-server/domain/service/mock"
)

func Test_userUsecase_VerifyToken(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		idToken entity.IDToken
		prepare func(ctx context.Context, ctrl *gomock.Controller) usecase.UserUsecase
		want    *entity.User
		err     error
	}{
		"it should returns user": {
			idToken: entity.IDToken("idToken123"),
			prepare: func(ctx context.Context, ctrl *gomock.Controller) usecase.UserUsecase {
				authenticationService := service.NewMockAuthenticationService(ctrl)
				authenticationService.EXPECT().VerifyToken(ctx, entity.IDToken("idToken123")).Return(entity.UID("uid123"), nil)

				userRepository := repository.NewMockUserRepository(ctrl)
				userRepository.EXPECT().GetByUID(entity.UID("uid123")).Return(&entity.User{
					ID:       1,
					UID:      "uid123",
					Username: "Chuck",
					Email:    "chuck@example.com",
				}, nil)

				return usecase.NewUserUsecase(userRepository, authenticationService)
			},
			want: &entity.User{
				ID:       1,
				UID:      "uid123",
				Username: "Chuck",
				Email:    "chuck@example.com",
			},
			err: nil,
		},
		"it should returns NotFountError when user is nil": {
			idToken: entity.IDToken("idToken000"),
			prepare: func(ctx context.Context, ctrl *gomock.Controller) usecase.UserUsecase {
				authenticationService := service.NewMockAuthenticationService(ctrl)
				authenticationService.EXPECT().VerifyToken(ctx, entity.IDToken("idToken000")).Return(entity.UID("uid000"), nil)

				userRepository := mock.NewMockUserRepository(ctrl)
				userRepository.EXPECT().GetByUID(entity.UID("uid000")).Return(nil, nil)

				return usecase.NewUserUsecase(userRepository, authenticationService)
			},
			want: nil,
			err:  &entity.NotFoundError{Name: "user"},
		},
		"it should returns UnexpectedError when err occurs on DB": {
			idToken: entity.IDToken("idToken000"),
			prepare: func(ctx context.Context, ctrl *gomock.Controller) usecase.UserUsecase {
				authenticationService := service.NewMockAuthenticationService(ctrl)
				authenticationService.EXPECT().VerifyToken(ctx, entity.IDToken("idToken000")).Return(entity.UID("uid000"), nil)

				userRepository := mock.NewMockUserRepository(ctrl)
				userRepository.EXPECT().GetByUID(entity.UID("uid000")).Return(nil, errors.New("test"))

				return usecase.NewUserUsecase(userRepository, authenticationService)
			},
			want: nil,
			err:  &entity.UnexpectedError{Err: errors.New("test")},
		},
		"it should returns UnauthorizedError when user is not sign in": {
			idToken: entity.IDToken("idToken000"),
			prepare: func(ctx context.Context, ctrl *gomock.Controller) usecase.UserUsecase {
				authenticationService := service.NewMockAuthenticationService(ctrl)
				authenticationService.EXPECT().VerifyToken(ctx, entity.IDToken("idToken000")).Return(entity.UID("uid000"), nil)

				userRepository := mock.NewMockUserRepository(ctrl)
				userRepository.EXPECT().GetByUID(entity.UID("uid000")).Return(nil, nil)

				return usecase.NewUserUsecase(userRepository, authenticationService)
			},
			want: nil,
			err:  &entity.UnauthorizedError{Massage: "not signup"},
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.Background()
			u := tt.prepare(ctx, ctrl)
			got, err := u.VerifyToken(ctx, tt.idToken)
			if err != nil {
				assert.Error(t, tt.err, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
