package handler

import (
	"net/http"

	"go-firebase-auth-server/interfaces/form"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"go-firebase-auth-server/application/usecase"
	"go-firebase-auth-server/domain/entity"
	"go-firebase-auth-server/interfaces/response"
	"go-firebase-auth-server/interfaces/view"
)

type UserHandler struct {
	db          *gorm.DB
	userUsecase usecase.UserUsecase
}

func NewUserHandler(db *gorm.DB, userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		db:          db,
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) Register(root, v1 *mux.Router) {
	root.HandleFunc("/authenticate", h.Authenticate).Methods("POST")
	v1.HandleFunc("/me", h.Me).Methods("GET")
}

func (h *UserHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
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

func (h *UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	user, err := entity.GetCurrentUser(r.Context())
	if err != nil {
		response.Error(w, response.Status(err), err.Error())
		return
	}

	response.JSON(w, http.StatusOK, view.NewUser(user))
}
