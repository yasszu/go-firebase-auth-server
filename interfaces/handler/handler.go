package handler

import (
	"net/http"

	"go-firebase-auth-server/registry"

	"github.com/gorilla/mux"

	"go-firebase-auth-server/application/usecase"
	"go-firebase-auth-server/interfaces/middleware"
)

type Handler struct {
	indexUsecase usecase.IndexUsecase
	userUsecase  usecase.UserUsecase
}

func NewHandler(r registry.Usecase) *Handler {
	return &Handler{
		indexUsecase: r.NewIndex(),
		userUsecase:  r.NewUser(),
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
	root.HandleFunc("/", index.Index).Methods("GET")
	root.HandleFunc("/healthy", index.Healthy).Methods("GET")
	root.HandleFunc("/ready", index.Ready).Methods("GET")
	root.HandleFunc("/authenticate", auth.Authenticate).Methods("POST")

	v1 := r.PathPrefix("/v1").Subrouter()
	v1.Use(middleware.Logging)
	v1.Use(middleware.FirebaseAuth(h.userUsecase))
	v1.HandleFunc("/me", user.Me).Methods("GET")
}
