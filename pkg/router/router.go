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
	mux.HandleFunc("PATCH /api/players/{id}", players_handlers.UpdatePlayerAccountBalanceHandler)

	mux.HandleFunc("GET /api/tournaments", tournaments_handlers.GetTournamentsHandler)
	mux.HandleFunc("POST /api/tournaments", tournaments_handlers.CreateTournamentHandler)
	mux.HandleFunc("GET /api/tournaments/rankings", tournaments_handlers.GetRankingForAllTournaments)
	mux.HandleFunc("GET /api/tournaments/{id}", tournaments_handlers.GetTournamentByIDHandler)
	mux.HandleFunc("POST /api/tournaments/{id}/bets", tournaments_handlers.BetOnTournamentHandler)
	mux.HandleFunc("GET /api/tournaments/{id}/rankings", tournaments_handlers.GetRankingForSingleTournament)
}
