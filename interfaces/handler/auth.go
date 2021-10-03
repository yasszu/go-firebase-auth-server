package handler

import (
	"net/http"

	"github.com/yasszu/go-firebase-auth-server/application/usecase"
	"github.com/yasszu/go-firebase-auth-server/domain/entity"
	"github.com/yasszu/go-firebase-auth-server/interfaces/response"
)

type AuthHandler struct {
	userUsecase usecase.UserUsecase
}

func NewAuthHandler(userUsecase usecase.UserUsecase) *AuthHandler {
	return &AuthHandler{
		userUsecase: userUsecase,
	}
}

func (h *AuthHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	f := &entity.IDTokenForm{
		IDToken: r.FormValue("id_token"),
	}

	if err := f.Validate(); err != nil {
		response.Error(w, http.StatusBadRequest, "Bad Request")
		return
	}

	user, err := h.userUsecase.Authenticate(r.Context(), f.Entity())
	if err != nil {
		response.Error(w, response.Status(err), err.Error())
		return
	}

	response.JSON(w, http.StatusOK, response.NewUser(user))
}
