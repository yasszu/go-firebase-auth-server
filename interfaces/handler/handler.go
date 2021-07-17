package handler

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-firebase-auth-server/interfaces/middleware"
	"go-firebase-auth-server/registry"
)

type Handler struct {
	index                  *IndexHandler
	auth                   *AuthHandler
	user                   *UserHandler
	firebaseAuthMiddleware *middleware.FirebaseAuthMiddleware
}

func NewHandler(r registry.Usecase) *Handler {
	return &Handler{
		index:                  NewIndexHandler(r.NewIndex()),
		auth:                   NewAuthHandler(r.NewUser()),
		user:                   NewUserHandler(r.NewUser()),
		firebaseAuthMiddleware: middleware.NewFirebaseAuthMiddleware(r.NewUser()),
	}
}

func (h *Handler) Register(r *mux.Router) {
	root := r.PathPrefix("").Subrouter()
	root.Use(middleware.Logging)
	root.Use(middleware.CORS)
	root.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	root.HandleFunc("/", h.index.Index).Methods("GET")
	root.HandleFunc("/healthy", h.index.Healthy).Methods("GET")
	root.HandleFunc("/ready", h.index.Ready).Methods("GET")
	root.HandleFunc("/authenticate", h.auth.Authenticate).Methods("POST")

	v1 := r.PathPrefix("/v1").Subrouter()
	v1.Use(middleware.Logging)
	v1.Use(h.firebaseAuthMiddleware.FirebaseAuth())
	v1.HandleFunc("/me", h.user.Me).Methods("GET")
}
