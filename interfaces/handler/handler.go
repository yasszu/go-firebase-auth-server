package handler

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-firebase-auth-server/application/usecase"
	"go-firebase-auth-server/interfaces/middleware"
)

type Handler struct {
	indexUsecase usecase.IndexUsecase
	userUsecase  usecase.UserUsecase
}

func NewHandler(
	indexUsecase usecase.IndexUsecase,
	userUsecase usecase.UserUsecase,
) *Handler {
	return &Handler{
		indexUsecase: indexUsecase,
		userUsecase:  userUsecase,
	}
}

func (h *Handler) Register(r *mux.Router) {
	index := NewIndexHandler(h.indexUsecase)
	auth := NewAuthHandler(h.userUsecase)
	user := NewUserHandler(h.userUsecase)

	root := r.PathPrefix("").Subrouter()
	root.Use(middleware.Logging)
	root.Use(middleware.CORS)
	root.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	index.Register(root)
	auth.Register(root)

	v1 := r.PathPrefix("/v1").Subrouter()
	v1.Use(middleware.Logging)
	v1.Use(middleware.FirebaseAuth(h.userUsecase))
	user.Register(v1)
}
