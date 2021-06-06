package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"go-firebase-auth-server/domain/entity"
)

func JSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

func Error(w http.ResponseWriter, code int, message interface{}) {
	JSON(w, code, map[string]string{
		"error": fmt.Sprint(message),
	})
}

func OK(w http.ResponseWriter) {
	JSON(w, http.StatusOK, map[string]string{
		"message": "OK",
	})
}

func Status(err error) int {
	if errors.Is(err, &entity.UnexpectedError{}) {
		return http.StatusInternalServerError
	}
	if errors.Is(err, &entity.NotFoundError{}) {
		return http.StatusNotFound
	}
	if errors.Is(err, &entity.UnauthorizedError{}) {
		return http.StatusUnauthorized
	}

	return http.StatusInternalServerError
}
