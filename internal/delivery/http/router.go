package http

import (
	"net/http"

	"webproject/internal/delivery/http/handler"
	"webproject/internal/delivery/http/middleware"
)

func NewRouter(userHandler *handler.UserHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", userHandler.Create)
	mux.HandleFunc("GET /users/{id}", userHandler.GetByID)

	return middleware.Logging(mux)
}
