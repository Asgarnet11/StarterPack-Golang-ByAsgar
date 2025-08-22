package handler

import (
	"net/http"
	"starter-golang/pkg/response"
)


func Health(w http.ResponseWriter, r *http.Request) {
response.JSON(w, http.StatusOK, map[string]any{
"status": "ok",
"message": "service healthy",
})
}