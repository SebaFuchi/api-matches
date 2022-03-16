package responseHelper

import (
	"encoding/json"
	"tinder_users/pkg/Domain/response"
)

func ResponseBuilder(status int, message string, data interface{}) ([]byte, error) {
	response := response.Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	marshalResponse, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return marshalResponse, nil

}
