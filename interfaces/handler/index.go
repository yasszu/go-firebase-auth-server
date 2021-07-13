package handler

import (
	"net/http"

	"go-firebase-auth-server/application/usecase"
	"go-firebase-auth-server/interfaces/response"
)

type IndexHandler struct {
	indexUsecase usecase.IndexUsecase
}

func NewIndexHandler(indexUsecase usecase.IndexUsecase) *IndexHandler {
	return &IndexHandler{indexUsecase: indexUsecase}
}

// Index AccountHandler
func (h *IndexHandler) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := h.indexUsecase.Index(r.Context())
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
func (h *IndexHandler) Ready(w http.ResponseWriter, r *http.Request) {
	if err := h.indexUsecase.Ready(r.Context()); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	response.OK(w)
}
