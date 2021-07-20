package service

import (
	"context"

	"github.com/yasszu/go-firebase-auth-server/domain/entity"
)

type AuthenticationService interface {
	VerifyToken(ctx context.Context, idToken entity.IDToken) (entity.UID, error)
	SetClaims(ctx context.Context, uid entity.UID, claims map[string]interface{}) error
	GetFirebaseUser(ctx context.Context, uid entity.UID) (*entity.User, error)
}
