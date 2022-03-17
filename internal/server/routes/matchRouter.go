package routes

import (
	"net/http"
	"tinder_matches/pkg/Domain/response"
	"tinder_matches/pkg/Use_cases/Handlers/matchHandler"
	"tinder_matches/pkg/Use_cases/Helpers/responseHelper"

	"github.com/go-chi/chi"
)

type MatchRouter struct {
	Handler matchHandler.Handler
}

func (mr *MatchRouter) CreateMatch(w http.ResponseWriter, r *http.Request) {
	matcher := r.URL.Query().Get("matcher")
	matched := r.URL.Query().Get("matched")

	status := mr.Handler.CreateMatch(matcher, matched)
	resp, err := responseHelper.ResponseBuilder(status.Index(), status.String(), nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500: Internal Server Error"))
		return
	}

	switch status {
	case response.MatchExists:
		w.WriteHeader(http.StatusConflict)
		w.Write(resp)
		return
	case response.SuccesfulCreation:
		w.WriteHeader(http.StatusCreated)
		w.Write(resp)
		return
	case response.InternalServerError, response.DBQueryError, response.DBExecutionError, response.CreationFailure, response.DBLastRowIdError:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resp)
		return
	default:
		status = response.Unknown
		response, err := responseHelper.ResponseBuilder(status.Index(), status.String(), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500: Internal server error"))
			return
		}
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(response))
		return
	}
}

func (mr *MatchRouter) GetPossibleMatches(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")

	possibleMatches, status := mr.Handler.GetPossibleMatches(token)
	resp, err := responseHelper.ResponseBuilder(status.Index(), status.String(), possibleMatches)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500: Internal server error"))
		return
	}

	switch status {
	case response.InternalServerError, response.DBQueryError, response.DBRowsError:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resp)
		return
	case response.NotFound:
		w.WriteHeader(http.StatusNotFound)
		w.Write(resp)
		return
	case response.SuccesfulSearch:
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
		return
	default:
		status = response.Unknown
		response, err := responseHelper.ResponseBuilder(status.Index(), status.String(), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500: Internal server error"))
			return
		}
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(response))
		return
	}
}

func (mr *MatchRouter) GetMyMatches(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")

	myMatches, status := mr.Handler.GetMyMatches(token)
	resp, err := responseHelper.ResponseBuilder(status.Index(), status.String(), myMatches)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500: Internal server error"))
		return
	}

	switch status {
	case response.InternalServerError, response.DBQueryError, response.DBRowsError:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resp)
		return
	case response.SuccesfulSearch:
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
		return
	default:
		status = response.Unknown
		response, err := responseHelper.ResponseBuilder(status.Index(), status.String(), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500: Internal server error"))
			return
		}
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(response))
		return
	}
}

func (mr *MatchRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/", mr.CreateMatch)
	r.Get("/{token}", mr.GetPossibleMatches)
	r.Get("/{token}/my-matches", mr.GetMyMatches)
	return r
}
