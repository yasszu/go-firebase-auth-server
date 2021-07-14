package handler

import (
	"net/http"

	"go-firebase-auth-server/registry"

	"github.com/gorilla/mux"

	"go-firebase-auth-server/application/usecase"
	"go-firebase-auth-server/interfaces/middleware"
)

type Handler struct {
	IndexUsecase usecase.IndexUsecase
	UserUsecase  usecase.UserUsecase
}

func NewHandler(r registry.Usecase) *Handler {
	return &Handler{
		IndexUsecase: r.NewIndex(),
		UserUsecase:  r.NewUser(),
	}
}

func (h *Handler) Register(r *mux.Router) {
	index := NewIndexHandler(h.IndexUsecase)
	auth := NewAuthHandler(h.UserUsecase)
	user := NewUserHandler(h.UserUsecase)

	root := r.PathPrefix("").Subrouter()
	root.Use(middleware.Logging)
	root.Use(middleware.CORS)
	root.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	root.HandleFunc("/", index.Index).Methods("GET")
	root.HandleFunc("/healthy", index.Healthy).Methods("GET")
	root.HandleFunc("/ready", index.Ready).Methods("GET")
	root.HandleFunc("/authenticate", auth.Authenticate).Methods("POST")

	v1 := r.PathPrefix("/v1").Subrouter()
	v1.Use(middleware.Logging)
	v1.Use(middleware.FirebaseAuth(h.UserUsecase))
	v1.HandleFunc("/me", user.Me).Methods("GET")
}
