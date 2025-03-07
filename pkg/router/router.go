package router

import (
	users_handlers "app/pkg/router/handlers/users"
	"net/http"
)

func SetupRouter(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/users", users_handlers.GetUsersHandler)
	mux.HandleFunc("GET /api/users/{id}", users_handlers.GetUserByIDHandler)
	mux.HandleFunc("POST /api/users", users_handlers.CreateUserHandler)
}
