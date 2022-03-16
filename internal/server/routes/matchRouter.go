package routes

import (
	"net/http"
	"tinder_matches/pkg/Use_cases/Handlers/matchHandler"

	"github.com/go-chi/chi"
)

type MatchRouter struct {
	Handler matchHandler.Handler
}

func (mr *MatchRouter) Routes() http.Handler {
	r := chi.NewRouter()

	return r
}
