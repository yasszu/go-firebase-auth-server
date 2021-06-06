package middleware

import (
	"go-firebase-auth-server/application/usecase"
)

type Middleware struct {
	userUsecase usecase.UserUsecase
}

func NewMiddleware(userUsecase usecase.UserUsecase) *Middleware {
	return &Middleware{userUsecase: userUsecase}
}
