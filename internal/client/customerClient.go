package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"is-tgbot/internal/keys"
	"is-tgbot/internal/model"
	"is-tgbot/pkg/logger"
	"net/http"
	"os"
)

func GetCustomer(request model.CustomerGetRequest) (model.CustomerResponse, error) {
	url := os.Getenv(keys.Server)

	if len(url) == 0 {
		return model.CustomerResponse{}, errors.New("env error")
	}

	url += keys.Customer

	body, jsonError := json.Marshal(request)

	if jsonError != nil {
		return model.CustomerResponse{}, jsonError
	}

	response, err := http.Post(url, keys.Json, bytes.NewBuffer(body))

	if err != nil {
		return model.CustomerResponse{}, err
	}

	defer closeResponse(response)

	byteData, readError := io.ReadAll(response.Body)

	if readError != nil {
		return model.CustomerResponse{}, readError
	}

	logger.Log().Infof("%s: response", keys.Customer)

	var customer model.Response[model.CustomerResponse]

	unmarshallErr := json.Unmarshal(byteData, &customer)

	if unmarshallErr != nil {
		return model.CustomerResponse{}, unmarshallErr
	}

	return customer.Data, nil
}
