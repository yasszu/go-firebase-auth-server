package firebase

import (
	"context"
	"log"

	_firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"

	"go-firebase-auth-server/domain/entity"
	"go-firebase-auth-server/domain/service"
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

func (s AuthenticationService) VerifyToken(ctx context.Context, idToken string) (string, error) {
	token, err := s.client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return "", err
	}

	return token.UID, nil
}

func (s AuthenticationService) SetClaims(ctx context.Context, uid string, claims map[string]interface{}) error {
	err := s.client.SetCustomUserClaims(ctx, uid, claims)
	if err != nil {
		return err
	}
	// The new custom claims will propagate to the user's ID token the
	// next time a new one is issued.auth.go
	return nil
}

func (s AuthenticationService) GetFirebaseUser(ctx context.Context, uid string) (*entity.User, error) {
	userRecord, err := s.client.GetUser(ctx, uid)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &entity.User{
		UID:      userRecord.UID,
		Username: userRecord.DisplayName,
		Email:    userRecord.Email,
	}, nil
}
