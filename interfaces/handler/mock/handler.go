package mock

import (
	"go-firebase-auth-server/interfaces/handler"
	"go-firebase-auth-server/registry"
)

func NewHandler(r *registry.MockUsecase) *handler.Handler {
	return &handler.Handler{
		IndexUsecase: r.Index,
		UserUsecase:  r.User,
	}
}
