package routes

import (
	"tinder_matches/internal/data/Infrastructure/matchRepository"
	"tinder_matches/pkg/Use_cases/Handlers/matchHandler"

	"net/http"

	"github.com/go-chi/chi"
)

// Instanciamos los handlers de los endpoints
func New() http.Handler {
	r := chi.NewRouter()

	mr := &MatchRouter{
		Handler: &matchHandler.MatchHandler{
			Repository: &matchRepository.MatchRepository{},
		},
	}
	r.Mount("/matches", mr.Routes())

	//Retornamos la api ya construida
	return r
}
