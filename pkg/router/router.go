package router

import (
	players_handlers "app/pkg/router/handlers/players"
	tournaments_handlers "app/pkg/router/handlers/tournaments"
	"net/http"
)

func SetupRouter(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/players", players_handlers.GetPlayersHandler)
	mux.HandleFunc("GET /api/players/{id}", players_handlers.GetPlayerByIDHandler)
	mux.HandleFunc("POST /api/players", players_handlers.CreatePlayerHandler)

	mux.HandleFunc("GET /api/tournaments", tournaments_handlers.GetTournamentsHandler)
	mux.HandleFunc("GET /api/tournaments/{id}", tournaments_handlers.GetTournamentByIDHandler)
	mux.HandleFunc("POST /api/tournaments", tournaments_handlers.CreateTournamentHandler)
	mux.HandleFunc("POST /api/tournaments/{id}/bet", tournaments_handlers.BetOnTournamentHandler)
}
