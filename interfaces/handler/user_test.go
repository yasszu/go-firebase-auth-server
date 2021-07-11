package handler_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"go-firebase-auth-server/application/usecase"
	"go-firebase-auth-server/application/usecase/mock"
	"go-firebase-auth-server/domain/entity"
	"go-firebase-auth-server/interfaces/handler"
	_middleware "go-firebase-auth-server/interfaces/middleware"
)

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkNodWNrIiwiaWF0IjoxNTE2MjM5MDIyfQ.Gsc5-cGTqp0XXIlHzJTixnubgnna4zdi1aq_wIzTWpQ"

func TestUserHandler_Me(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		token   string
		prepare func(ctx context.Context, ctrl *gomock.Controller) usecase.UserUsecase
		want    *entity.UserResponse
		code    int
	}{
		{
			name:  "it should returns user",
			token: token,
			prepare: func(ctx context.Context, ctrl *gomock.Controller) usecase.UserUsecase {
				u := mock.NewMockUserUsecase(ctrl)
				u.EXPECT().VerifyToken(gomock.Any(), entity.IDToken(token)).Return(
					&entity.User{
						ID:        1,
						UID:       "DCHfBC88grC3vwmdqsQwVvWJQBPR96kA",
						Username:  "Chuck",
						Email:     "chuck@example.com",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					}, nil)
				return u
			},
			want: &entity.UserResponse{
				UserID:   1,
				Username: "Chuck",
				Email:    "chuck@example.com",
			},
			code: http.StatusOK,
		},
		{
			name:  "it should returns unauthorized error",
			token: token,
			prepare: func(ctx context.Context, ctrl *gomock.Controller) usecase.UserUsecase {
				u := mock.NewMockUserUsecase(ctrl)
				err := &entity.UnauthorizedError{Massage: "Unauthorized"}
				u.EXPECT().VerifyToken(gomock.Any(), entity.IDToken(token)).Return(nil, err)
				return u
			},
			want: &entity.UserResponse{},
			code: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			ctrl, ctx := gomock.WithContext(ctx, t)
			defer ctrl.Finish()

			u := tt.prepare(ctx, ctrl)
			h := handler.NewUserHandler(u)
			m := _middleware.NewMiddleware(u)

			r := mux.NewRouter()
			v1 := r.PathPrefix("/v1").Subrouter()
			v1.Use(m.FirebaseAuth)
			h.Register(v1)

			req, err := http.NewRequest(http.MethodGet, "/v1/me", nil)
			assert.NoError(t, err)
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tt.token))

			rr := httptest.NewRecorder()
			v1.ServeHTTP(rr, req)
			res := rr.Result()

			var user *entity.UserResponse
			err = json.Unmarshal(rr.Body.Bytes(), &user)
			assert.NoError(t, err)
			assert.Equal(t, tt.code, res.StatusCode)
			assert.Empty(t, cmp.Diff(tt.want, user))
		})
	}
}
