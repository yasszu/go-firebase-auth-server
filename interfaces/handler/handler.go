package handler

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-firebase-auth-server/application/usecase"
	_middleware "go-firebase-auth-server/interfaces/middleware"
)

type Handler struct {
	middleware *_middleware.Middleware
	index      *IndexHandler
	auth       *AuthHandler
	user       *UserHandler
}

func NewHandler(indexUsecase usecase.IndexUsecase, userUsecase usecase.UserUsecase) *Handler {
	return &Handler{
		middleware: _middleware.NewMiddleware(userUsecase),
		index:      NewIndexHandler(indexUsecase),
		auth:       NewAuthHandler(userUsecase),
		user:       NewUserHandler(userUsecase),
	}
}

func (h *Handler) Register(r *mux.Router) {
	rr := r.PathPrefix("").Subrouter()
	rr.Use(h.middleware.Logging)
	rr.Use(h.middleware.CORS)
	rr.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	h.index.Register(rr)
	h.auth.Register(rr)

	v1 := r.PathPrefix("/v1").Subrouter()
	v1.Use(h.middleware.Logging)
	v1.Use(h.middleware.FirebaseAuth)
	h.user.Register(v1)
}
