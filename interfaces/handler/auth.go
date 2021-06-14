package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"go-firebase-auth-server/application/usecase"
	"go-firebase-auth-server/interfaces/form"
	"go-firebase-auth-server/interfaces/response"
	"go-firebase-auth-server/interfaces/view"
)

type AuthHandler struct {
	db          *gorm.DB
	userUsecase usecase.UserUsecase
}

func NewAuthHandler(db *gorm.DB, userUsecase usecase.UserUsecase) *AuthHandler {
	return &AuthHandler{
		db:          db,
		userUsecase: userUsecase,
	}
}

func (h *AuthHandler) Register(r *mux.Router) {
	r.HandleFunc("/authenticate", h.Authenticate).Methods("POST")
}

func (h *AuthHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	f := &form.Authenticate{
		IDToken: r.FormValue("id_token"),
	}

	if err := f.Validate(); err != nil {
		response.Error(w, http.StatusBadRequest, "Bad Request")
		return
	}

	user, err := h.userUsecase.Authenticate(r.Context(), f.IDToken)
	if err != nil {
		response.Error(w, response.Status(err), err.Error())
		return
	}

	response.JSON(w, http.StatusOK, view.NewUser(user))
}
