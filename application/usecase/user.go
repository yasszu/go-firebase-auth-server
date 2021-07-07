package usecase

import (
	"context"
	"log"

	"go-firebase-auth-server/domain/entity"
	"go-firebase-auth-server/domain/repository"
	"go-firebase-auth-server/domain/service"
)

type UserUsecase interface {
	Authenticate(ctx context.Context, idToken string) (*entity.User, error)
	VerifyToken(ctx context.Context, idToken string) (*entity.User, error)
	GetUser(ctx context.Context, uid string) (*entity.User, error)
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

func (u userUsecase) Authenticate(ctx context.Context, idToken string) (*entity.User, error) {
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

func (u userUsecase) VerifyToken(ctx context.Context, idToken string) (*entity.User, error) {
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

func (u userUsecase) GetUser(_ context.Context, uid string) (*entity.User, error) {
	user, err := u.userRepository.GetByUID(uid)
	if err != nil {
		log.Println("Error: ", err)
		return nil, &entity.UnexpectedError{Err: err}
	}

	if user == nil {
		return nil, &entity.NotFoundError{Name: "user"}
	}

	return user, nil
}
