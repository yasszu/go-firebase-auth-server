package service

import (
	"context"
	"go-firebase-auth-server/domain/entity"
)

type AuthenticationService interface {
	VerifyToken(ctx context.Context, token string) (string, error)
	SetCustomClaims(ctx context.Context, uid string, claims map[string]interface{}) error
	GetUser(ctx context.Context, uid string) (entity.FirebaseUser, error)
}
