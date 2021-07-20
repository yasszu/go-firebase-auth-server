package firebase

import (
	"context"
	"log"

	_firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/yasszu/go-firebase-auth-server/domain/entity"
	"github.com/yasszu/go-firebase-auth-server/domain/service"
)

type AuthenticationService struct {
	client *auth.Client
}

func NewAuthenticationService() service.AuthenticationService {
	app, err := _firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	return &AuthenticationService{
		client: client,
	}
}

func (s AuthenticationService) VerifyToken(ctx context.Context, idToken entity.IDToken) (entity.UID, error) {
	token, err := s.client.VerifyIDToken(ctx, string(idToken))
	if err != nil {
		return "", err
	}

	return entity.UID(token.UID), nil
}

func (s AuthenticationService) SetClaims(ctx context.Context, uid entity.UID, claims map[string]interface{}) error {
	err := s.client.SetCustomUserClaims(ctx, string(uid), claims)
	if err != nil {
		return err
	}
	// The new custom claims will propagate to the user's ID token the
	// next time a new one is issued.auth.go
	return nil
}

func (s AuthenticationService) GetFirebaseUser(ctx context.Context, uid entity.UID) (*entity.User, error) {
	userRecord, err := s.client.GetUser(ctx, string(uid))
	if err != nil {
		return nil, err
	}
	return &entity.User{
		UID:      userRecord.UID,
		Username: userRecord.DisplayName,
		Email:    userRecord.Email,
	}, nil
}
