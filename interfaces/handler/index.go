package handler

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"go-firebase-auth-server/interfaces/response"
)

type IndexHandler struct {
	db *gorm.DB
}

func NewIndexHandler(db *gorm.DB) *IndexHandler {
	return &IndexHandler{db: db}
}

func (h IndexHandler) Register(root *mux.Router) {
	root.HandleFunc("/", h.Index).Methods("GET")
	root.HandleFunc("/healthy", h.Healthy).Methods("GET")
	root.HandleFunc("/ready", h.Ready).Methods("GET")
	root.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
}

// Index AccountHandler
func (h *IndexHandler) Index(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("web/index.html")
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = tmpl.Execute(w, nil); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
}

// Healthy is used for liveness probes
func (h *IndexHandler) Healthy(w http.ResponseWriter, _ *http.Request) {
	response.OK(w)
}

// Ready is used for readiness probes
func (h *IndexHandler) Ready(w http.ResponseWriter, _ *http.Request) {
	var i int
	if err := h.db.Raw("SELECT 1").Scan(&i).Error; err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	response.OK(w)
}
