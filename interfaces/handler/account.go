package handler

import (
	"net/http"

	"go-firebase-auth-server/domain/service"

	"go-firebase-auth-server/application/usecase"
	"go-firebase-auth-server/domain/repository"
	"go-firebase-auth-server/infrastructure/jwt"
	"go-firebase-auth-server/interfaces/form"
	"go-firebase-auth-server/interfaces/response"
	"go-firebase-auth-server/interfaces/view"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type AccountHandler struct {
	db             *gorm.DB
	accountUsecase usecase.AccountUsecase
}

func NewAccountHandler(db *gorm.DB, accountRepository repository.AccountRepository, authenticationService service.AuthenticationService) *AccountHandler {
	return &AccountHandler{
		db:             db,
		accountUsecase: usecase.NewAccountUsecase(accountRepository, authenticationService),
	}
}

func (h *AccountHandler) Register(root, v1 *mux.Router) {
	root.HandleFunc("/signup", h.Signup).Methods("POST")
	root.HandleFunc("/login", h.Login).Methods("POST")
	v1.HandleFunc("/me", h.Me).Methods("GET")
}

// Signup POST /signup
func (h *AccountHandler) Signup(w http.ResponseWriter, r *http.Request) {
	f := form.Signup{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	if err := f.Validate(); err != nil {
		response.Error(w, response.Status(err), err)
		return
	}

	account, err := f.Entity()
	if err != nil {
		response.Error(w, response.Status(err), err)
		return
	}

	token, err := h.accountUsecase.SignUp(r.Context(), &account)
	if err != nil {
		response.Error(w, response.Status(err), err)
		return
	}

	response.JSON(w, http.StatusOK, token)
}

// Login POST /login
func (h *AccountHandler) Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	token, err := h.accountUsecase.Login(r.Context(), email, password)
	if err != nil {
		response.Error(w, response.Status(err), err)
		return
	}

	response.JSON(w, http.StatusOK, token)
}

// Me  GET /v1/me
func (h *AccountHandler) Me(w http.ResponseWriter, r *http.Request) {
	accountID, ok := jwt.GetAccountID(r.Context())
	if !ok {
		response.Error(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	account, err := h.accountUsecase.Me(r.Context(), accountID)
	if err != nil {
		response.Error(w, response.Status(err), err)
		return
	}

	response.JSON(w, http.StatusOK, view.NewAccount(account))
}
