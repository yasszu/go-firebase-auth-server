package handler_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yasszu/go-firebase-auth-server/interfaces/response"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/yasszu/go-firebase-auth-server/application/usecase/mock"
	"github.com/yasszu/go-firebase-auth-server/domain/entity"
	"github.com/yasszu/go-firebase-auth-server/interfaces/handler"
	"github.com/yasszu/go-firebase-auth-server/registry"
)

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkNodWNrIiwiaWF0IjoxNTE2MjM5MDIyfQ.Gsc5-cGTqp0XXIlHzJTixnubgnna4zdi1aq_wIzTWpQ"

func TestUserHandler_Me(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		token   string
		prepare func(ctx context.Context, ctrl *gomock.Controller) registry.Usecase
		want    *response.User
		code    int
	}{
		{
			name:  "it should returns user",
			token: token,
			prepare: func(ctx context.Context, ctrl *gomock.Controller) registry.Usecase {
				r := registry.NewMockUsecase(ctrl)
				r.UserUsecase.(*mock.MockUserUsecase).EXPECT().VerifyToken(gomock.Any(), entity.IDToken(token)).Return(
					&entity.User{
						ID:       1,
						UID:      "DCHfBC88grC3vwmdqsQwVvWJQBPR96kA",
						Username: "Chuck",
						Email:    "chuck@example.com",
					}, nil)
				return r
			},
			want: &response.User{
				UserID:   1,
				Username: "Chuck",
				Email:    "chuck@example.com",
			},
			code: http.StatusOK,
		},
		{
			name:  "it should returns unauthorized error",
			token: token,
			prepare: func(ctx context.Context, ctrl *gomock.Controller) registry.Usecase {
				err := &entity.UnauthorizedError{Massage: "Unauthorized"}
				r := registry.NewMockUsecase(ctrl)
				r.UserUsecase.(*mock.MockUserUsecase).EXPECT().VerifyToken(gomock.Any(), entity.IDToken(token)).Return(nil, err)
				return r
			},
			want: &response.User{},
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
			h := handler.NewHandler(u)
			r := mux.NewRouter()
			h.Register(r)

			req, err := http.NewRequest(http.MethodGet, "/v1/me", nil)
			assert.NoError(t, err)
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tt.token))

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)
			res := rr.Result()

			var user *response.User
			err = json.Unmarshal(rr.Body.Bytes(), &user)
			assert.NoError(t, err)
			assert.Equal(t, tt.code, res.StatusCode)
			assert.Empty(t, cmp.Diff(tt.want, user))
		})
	}
}
