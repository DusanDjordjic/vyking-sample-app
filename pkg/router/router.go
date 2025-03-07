package router

import (
	players_handlers "app/pkg/router/handlers/players"
	"net/http"
)

func SetupRouter(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/players", players_handlers.GetPlayersHandler)
	mux.HandleFunc("GET /api/players/{id}", players_handlers.GetPlayerByIDHandler)
	mux.HandleFunc("POST /api/players", players_handlers.CreatePlayerHandler)
}
