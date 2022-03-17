package matchHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tinder_matches/internal/data/Infrastructure/matchRepository"
	"tinder_matches/pkg/Domain/pet"
	"tinder_matches/pkg/Domain/response"
)

type MatchHandler struct {
	Repository matchRepository.Repository
}

type Handler interface {
	CreateMatch(matcher, matched string) response.Status
	GetPossibleMatches(token string) ([]pet.Pet, response.Status)
	GetMyMatches(token string) (interface{}, response.Status)
}

func (mh *MatchHandler) CreateMatch(matcher, matched string) response.Status {

	status := mh.Repository.CheckMatch(matcher, matched)
	if status != response.NotFound {
		return status
	}
	return mh.Repository.CreateMatch(matcher, matched)
}

func (mh *MatchHandler) GetPossibleMatches(token string) ([]pet.Pet, response.Status) {

	resp, err := http.Get("http://localhost:8081/api/tinder/pets/" + token)

	if err != nil {
		fmt.Println(err)
		return nil, response.InternalServerError
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {

		var parseResp response.Response
		err = json.NewDecoder(resp.Body).Decode(&parseResp)
		if err != nil {
			return nil, response.InternalServerError
		}

		var p pet.Pet
		jsonString, _ := json.Marshal(parseResp.Data)
		err = json.Unmarshal(jsonString, &p)
		if err != nil {
			return nil, response.InternalServerError
		}

		posibbleM, status := mh.Repository.GetPossibleMatches(p)
		if status != response.SuccesfulSearch {
			return nil, status
		}

		myMatches, status := mh.Repository.GetMyMatches(p.Token)
		if status != response.SuccesfulSearch {
			return nil, status
		}

		for i := 0; i < len(posibbleM); i++ {
			for j := 0; j < len(myMatches); j++ {
				if posibbleM[i].Token == myMatches[j] {
					posibbleM = append(posibbleM[:i], posibbleM[i+1:]...)
					i--
				}
			}
		}

		return posibbleM, response.SuccesfulSearch
	}

	return nil, response.NotFound
}

func (mh *MatchHandler) GetMyMatches(token string) (interface{}, response.Status) {

	type Matches struct {
		WhoIMatched  []pet.Pet `json:"whoIMatched"`
		WhoMatchedMe []pet.Pet `json:"whoMatchedMe"`
	}

	matches := Matches{}

	whoIMatched, status := mh.Repository.GetWhoIMatched(token)
	if status != response.SuccesfulSearch {
		return nil, status
	}

	whoMatchedMe, status := mh.Repository.GetWhoMatchedMe(token)
	if status != response.SuccesfulSearch {
		return nil, status
	}

	matches.WhoIMatched = whoIMatched
	matches.WhoMatchedMe = whoMatchedMe

	return matches, response.SuccesfulSearch
}
