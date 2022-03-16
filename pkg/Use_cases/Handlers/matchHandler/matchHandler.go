package matchHandler

import "tinder_matches/internal/data/Infrastructure/matchRepository"

type MatchHandler struct {
	Repository matchRepository.Repository
}

type Handler interface{}
