package usecase

import (
	"context"
	"errors"
	"log"

	"gorm.io/gorm"

	"go-firebase-auth-server/domain/entity"
	"go-firebase-auth-server/domain/repository"
	"go-firebase-auth-server/domain/service"
)

type UserUsecase interface {
	SignUp(ctx context.Context, idToken string) (*entity.User, error)
	Me(ctx context.Context, uid string) (*entity.User, error)
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

func (u userUsecase) SignUp(ctx context.Context, idToken string) (*entity.User, error) {
	uid, err := u.authenticationService.VerifyToken(ctx, idToken)
	if err != nil {
		log.Println("Error: ", err)
		return nil, &entity.UnexpectedError{Err: err}
	}

	user, err := u.authenticationService.GetFirebaseUser(ctx, uid)
	if err != nil {
		log.Println("Error: ", err)
		return nil, &entity.UnexpectedError{Err: err}
	}

	if err = u.userRepository.Crete(user); err != nil {
		log.Println("Error: ", err)
		return nil, err
	}

	return user, nil
}

func (u userUsecase) Me(ctx context.Context, uid string) (*entity.User, error) {
	user, err := u.userRepository.GetByUID(uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &entity.NotFoundError{Name: "user"}
		}

		log.Println("Error: ", err)
		return nil, &entity.UnexpectedError{Err: err}
	}

	return user, nil
}
