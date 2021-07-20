package usecase

import (
	"context"
	"log"

	"github.com/yasszu/go-firebase-auth-server/domain/entity"
	"github.com/yasszu/go-firebase-auth-server/domain/repository"
	"github.com/yasszu/go-firebase-auth-server/domain/service"
)

//go:generate mockgen -source=./user.go -destination=./mock/user.go -package=mock
type UserUsecase interface {
	Authenticate(ctx context.Context, idToken entity.IDToken) (*entity.User, error)
	VerifyToken(ctx context.Context, idToken entity.IDToken) (*entity.User, error)
}

type userUsecase struct {
	userRepository        repository.UserRepository
	authenticationService service.AuthenticationService
}

func NewUserUsecase(
	userRepository repository.UserRepository,
	authenticationService service.AuthenticationService,
) UserUsecase {
	return &userUsecase{
		userRepository:        userRepository,
		authenticationService: authenticationService,
	}
}

func (u userUsecase) Authenticate(ctx context.Context, idToken entity.IDToken) (*entity.User, error) {
	uid, err := u.authenticationService.VerifyToken(ctx, idToken)
	if err != nil {
		log.Println("Error: ", err)
		return nil, &entity.UnauthorizedError{Massage: "failed verifying idToken"}
	}

	user, err := u.userRepository.GetByUID(uid)
	if err != nil {
		log.Println("Error: ", err)
		return nil, &entity.UnexpectedError{Err: err}
	}

	if user == nil {
		user, err = u.authenticationService.GetFirebaseUser(ctx, uid)
		if err != nil {
			log.Println("Error: ", err)
			return nil, &entity.UnexpectedError{Err: err}
		}
		if err = u.userRepository.Crete(user); err != nil {
			log.Println("Error: ", err)
			return nil, &entity.UnexpectedError{Err: err}
		}
		return user, nil
	}

	return user, nil
}

func (u userUsecase) VerifyToken(ctx context.Context, idToken entity.IDToken) (*entity.User, error) {
	uid, err := u.authenticationService.VerifyToken(ctx, idToken)
	if err != nil {
		log.Println("Error: ", err)
		return nil, &entity.UnauthorizedError{Massage: "error verifying idToken"}
	}

	user, err := u.userRepository.GetByUID(uid)
	if err != nil {
		log.Println("Error: ", err)
		return nil, &entity.UnexpectedError{Err: err}
	}

	if user == nil {
		return nil, &entity.UnauthorizedError{Massage: "not signup"}
	}

	return user, nil
}
