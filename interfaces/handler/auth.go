package handler

import (
	"net/http"

	"go-firebase-auth-server/domain/entity"

	"go-firebase-auth-server/application/usecase"
	"go-firebase-auth-server/interfaces/response"

	"github.com/gorilla/mux"
)

type AuthHandler struct {
	userUsecase usecase.UserUsecase
}

func NewAuthHandler(userUsecase usecase.UserUsecase) *AuthHandler {
	return &AuthHandler{
		userUsecase: userUsecase,
	}
}

func (h *AuthHandler) Register(r *mux.Router) {
	r.HandleFunc("/authenticate", h.Authenticate).Methods("POST")
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

	response.JSON(w, http.StatusOK, user.Response())
}
