package matchRepository

import (
	"tinder_matches/pkg/Domain/pet"
	"tinder_matches/pkg/Domain/response"
	"tinder_matches/pkg/Use_cases/Helpers/dbHelper"
)

type MatchRepository struct {
}

type Repository interface {
	CreateMatch(matcher, matched string) response.Status
	CheckMatch(matcher, matched string) response.Status
	GetPossibleMatches(p pet.Pet) ([]pet.Pet, response.Status)
	GetMyMatches(token string) ([]string, response.Status)
	GetWhoIMatched(token string) ([]pet.Pet, response.Status)
	GetWhoMatchedMe(token string) ([]pet.Pet, response.Status)
}

func (mr *MatchRepository) CreateMatch(matcher, matched string) response.Status {
	sqlCon, err := dbHelper.GetDB()
	if err != nil {
		return response.InternalServerError
	}
	defer sqlCon.Close()

	insForm, err := sqlCon.Prepare("INSERT INTO matches(matcher, matched) VALUES(?, ?)")
	if err != nil {
		return response.DBQueryError
	}
	defer insForm.Close()

	_, err = insForm.Exec(matcher, matched)
	if err != nil {
		return response.DBExecutionError
	}

	return response.SuccesfulCreation
}

func (mr *MatchRepository) CheckMatch(matcher, matched string) response.Status {
	sqlCon, err := dbHelper.GetDB()
	if err != nil {
		return response.InternalServerError
	}
	defer sqlCon.Close()

	selForm, err := sqlCon.Prepare("SELECT * FROM matches WHERE matcher = ? AND matched = ?")
	if err != nil {
		return response.DBQueryError
	}
	defer selForm.Close()

	rows, err := selForm.Query(matcher, matched)
	if err != nil {
		return response.DBRowsError
	}
	defer rows.Close()

	for rows.Next() {
		return response.MatchExists
	}

	return response.NotFound
}

func (mr *MatchRepository) GetPossibleMatches(p pet.Pet) ([]pet.Pet, response.Status) {
	sqlCon, err := dbHelper.GetDB()
	if err != nil {
		return nil, response.InternalServerError
	}
	defer sqlCon.Close()

	selForm, err := sqlCon.Prepare("SELECT token, owner_token, name, type, sex, image FROM pets WHERE token != ? AND owner_token != ? AND sex != ? AND type = ?")
	if err != nil {
		return nil, response.DBQueryError
	}
	defer selForm.Close()

	rows, err := selForm.Query(
		p.Token,
		p.OwnerToken,
		p.Sex,
		p.Type,
	)
	if err != nil {
		return nil, response.DBRowsError
	}
	defer rows.Close()

	var results []pet.Pet
	for rows.Next() {
		var matchedP pet.Pet
		err = rows.Scan(
			&matchedP.Token,
			&matchedP.OwnerToken,
			&matchedP.Name,
			&matchedP.Type,
			&matchedP.Sex,
			&matchedP.Image,
		)

		if err != nil {
			return nil, response.DBScanError
		}
		results = append(results, matchedP)
	}

	return results, response.SuccesfulSearch
}

func (mr *MatchRepository) GetMyMatches(token string) ([]string, response.Status) {
	sqlCon, err := dbHelper.GetDB()
	if err != nil {
		return nil, response.InternalServerError
	}
	defer sqlCon.Close()

	selForm, err := sqlCon.Prepare("SELECT matched FROM matches WHERE matcher = ?")
	if err != nil {
		return nil, response.DBQueryError
	}
	defer selForm.Close()

	rows, err := selForm.Query(token)
	if err != nil {
		return nil, response.DBRowsError
	}

	var results []string
	for rows.Next() {
		var matched string
		err = rows.Scan(&matched)
		if err != nil {
			return nil, response.DBScanError
		}
		results = append(results, matched)
	}

	return results, response.SuccesfulSearch
}

func (mr *MatchRepository) GetWhoIMatched(token string) ([]pet.Pet, response.Status) {
	sqlCon, err := dbHelper.GetDB()
	if err != nil {
		return nil, response.InternalServerError
	}
	defer sqlCon.Close()

	selForm, err := sqlCon.Prepare("SELECT P.token, P.owner_token, P.name, P.type, P.sex, P.image FROM pets P INNER JOIN matches M on P.token = M.matched WHERE M.matcher = ?")
	if err != nil {
		return nil, response.DBQueryError
	}
	defer selForm.Close()

	rows, err := selForm.Query(token)
	if err != nil {
		return nil, response.DBRowsError
	}

	var pets []pet.Pet
	for rows.Next() {
		var p pet.Pet
		err = rows.Scan(
			&p.Token,
			&p.OwnerToken,
			&p.Name,
			&p.Type,
			&p.Sex,
			&p.Image,
		)
		if err != nil {
			return nil, response.DBScanError
		}
		pets = append(pets, p)
	}

	return pets, response.SuccesfulSearch
}

func (mr *MatchRepository) GetWhoMatchedMe(token string) ([]pet.Pet, response.Status) {
	sqlCon, err := dbHelper.GetDB()
	if err != nil {
		return nil, response.InternalServerError
	}
	defer sqlCon.Close()

	selForm, err := sqlCon.Prepare("SELECT P.token, P.owner_token, P.name, P.type, P.sex, P.image FROM pets P INNER JOIN matches M on P.token = M.matcher WHERE M.matched = ?")
	if err != nil {
		return nil, response.DBQueryError
	}
	defer selForm.Close()

	rows, err := selForm.Query(token)
	if err != nil {
		return nil, response.DBRowsError
	}

	var pets []pet.Pet
	for rows.Next() {
		var p pet.Pet
		err = rows.Scan(
			&p.Token,
			&p.OwnerToken,
			&p.Name,
			&p.Type,
			&p.Sex,
			&p.Image,
		)
		if err != nil {
			return nil, response.DBScanError
		}
		pets = append(pets, p)
	}

	return pets, response.SuccesfulSearch
}
