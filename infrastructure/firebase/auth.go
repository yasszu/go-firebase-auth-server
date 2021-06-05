package firebase

import (
	"context"
	"log"

	"go-firebase-auth-server/domain/entity"
	"go-firebase-auth-server/domain/service"

	_firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

var _ service.AuthenticationService = &AuthenticationService{}

type AuthenticationService struct {
	client *auth.Client
}

func NewAuthenticationService() AuthenticationService {
	app, err := _firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	return AuthenticationService{
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

func (s AuthenticationService) GetUser(ctx context.Context, uid string) (entity.FirebaseUser, error) {
	// Lookup the user associated with the specified uid.
	user, err := s.client.GetUser(ctx, uid)
	if err != nil {
		log.Fatal(err)
	}
	// The claims can be accessed on the user record.
	if admin, ok := user.CustomClaims["admin"]; ok {
		if admin.(bool) {
			log.Println(admin)
		}
	}

	return entity.FirebaseUser{}, err
}
