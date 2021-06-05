package usecase

import (
	"context"
	"go-firebase-auth-server/domain/entity"
	"go-firebase-auth-server/domain/repository"
	"go-firebase-auth-server/domain/service"
	"go-firebase-auth-server/infrastructure/jwt"
	"go-firebase-auth-server/util"
	"log"
)

type AccountUsecase interface {
	SignUpWithFirebase(ctx context.Context, idToken string) (*entity.Account, error)
	SignUp(ctx context.Context, account *entity.Account) (*entity.AccessToken, error)
	Login(ctx context.Context, email, password string) (*entity.AccessToken, error)
	Me(ctx context.Context, accountID uint) (*entity.Account, error)
}

type accountUsecase struct {
	accountRepository     repository.AccountRepository
	authenticationService service.AuthenticationService
}

func NewAccountUsecase(accountRepository repository.AccountRepository, authenticationService service.AuthenticationService) AccountUsecase {
	return &accountUsecase{
		accountRepository:     accountRepository,
		authenticationService: authenticationService,
	}
}

func (u *accountUsecase) SignUpWithFirebase(ctx context.Context, idToken string) (*entity.Account, error) {
	uid, err := u.authenticationService.VerifyToken(ctx, idToken)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	user, err := u.authenticationService.GetUser(ctx, uid)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	account, err := u.accountRepository.RegisterFirebaseUser(&user)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return account, err
}

func (u *accountUsecase) SignUp(_ context.Context, account *entity.Account) (*entity.AccessToken, error) {
	if err := u.accountRepository.CreateAccount(account); err != nil {
		log.Println(err.Error())
		return nil, &entity.UnexpectedError{Err: err}
	}

	token, err := jwt.Sign(account)
	if err != nil {
		log.Println(err.Error())
		return nil, &entity.UnexpectedError{Err: err}
	}

	return token, nil
}

func (u *accountUsecase) Login(_ context.Context, email, password string) (*entity.AccessToken, error) {
	account, err := u.accountRepository.GetAccountByEmail(email)
	if err != nil {
		log.Println(err.Error())
		return nil, &entity.UnexpectedError{Err: err}
	}

	if err = util.ComparePassword(account.PasswordHash, password); err != nil {
		log.Println(err.Error())
		return nil, &entity.UnauthorizedError{
			Massage: "invalid password",
		}
	}

	token, err := jwt.Sign(account)
	if err != nil {
		log.Println(err.Error())
		return nil, &entity.UnexpectedError{Err: err}
	}

	return token, nil
}

func (u *accountUsecase) Me(_ context.Context, accountID uint) (*entity.Account, error) {
	account, err := u.accountRepository.GetAccountByID(accountID)
	if err != nil {
		log.Println(err.Error())
		return nil, &entity.UnexpectedError{Err: err}
	}

	return account, nil
}
