package handler

import (
	"net/http"

	"github.com/yasszu/go-firebase-auth-server/application/usecase"
	"github.com/yasszu/go-firebase-auth-server/domain/entity"
	"github.com/yasszu/go-firebase-auth-server/interfaces/response"
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
