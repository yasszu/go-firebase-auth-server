package handler

import (
	"net/http"

	"go-firebase-auth-server/application/usecase"
	"go-firebase-auth-server/domain/entity"
	"go-firebase-auth-server/interfaces/response"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	user, err := entity.GetCurrentUser(r.Context())
	if err != nil {
		response.Error(w, response.Status(err), err)
		return
	}

	response.JSON(w, http.StatusOK, user.Response())
}
