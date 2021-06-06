package service

import (
	"context"

	"go-firebase-auth-server/domain/entity"
)

type AuthenticationService interface {
	VerifyToken(ctx context.Context, token string) (string, error)
	SetClaims(ctx context.Context, uid string, claims map[string]interface{}) error
	GetFirebaseUser(ctx context.Context, uid string) (*entity.User, error)
}
