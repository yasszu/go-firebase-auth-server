package middleware

import (
	"context"
	"net/http"
	"strings"

	"go-firebase-auth-server/domain/entity"
	"go-firebase-auth-server/interfaces/response"
)

func (m *Middleware) FirebaseAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		extractedToken := strings.Split(authHeader, "Bearer ")
		if len(extractedToken) == 2 {
			idToken := strings.TrimSpace(extractedToken[1])

			user, err := m.userUsecase.VerifyToken(r.Context(), idToken)
			if err != nil {
				response.Error(w, response.Status(err), err)
				return
			}

			ctx := context.WithValue(r.Context(), entity.ContextKeyUser, user)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		} else {
			response.Error(w, http.StatusBadRequest, "Bad Request")
			return
		}
	})
}
