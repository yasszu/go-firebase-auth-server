package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/yasszu/go-firebase-auth-server/domain/entity"
)

var (
	UnexpectedError   *entity.UnexpectedError
	NotFoundError     *entity.NotFoundError
	UnauthorizedError *entity.UnauthorizedError
)

func Status(err error) int {
	if errors.As(err, &UnexpectedError) {
		return http.StatusInternalServerError
	}
	if errors.As(err, &NotFoundError) {
		return http.StatusNotFound
	}
	if errors.As(err, &UnauthorizedError) {
		return http.StatusUnauthorized
	}

	return http.StatusInternalServerError
}

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
