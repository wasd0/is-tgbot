package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wasd0/is-common/pkg/logger"
	"io"
	"is-tgbot/internal/keys"
	"is-tgbot/internal/model"
	"is-tgbot/internal/utils"
	"net/http"
)

func GetBalance(request model.BalanceGetRequest) (model.BalanceGetResponse, error) {
	url := utils.NewUrl(getServerUrl() + keys.PathBalance)

	if request.TelegramId != nil {
		url.AddParam("telegramId", fmt.Sprintf("%d", *request.TelegramId))
	}
	if request.CustomerId != nil {
		url.AddParam("customerId", fmt.Sprintf("%d", *request.CustomerId))
	}
	if request.CurrencyCode != nil {
		url.AddParam("currencyCode", *request.CurrencyCode)
	}

	response, err := http.Get(url.Build())

	if err != nil {
		return model.BalanceGetResponse{}, err
	}

	defer closeResponse(response)

	byteData, readError := io.ReadAll(response.Body)

	if readError != nil {
		return model.BalanceGetResponse{}, readError
	}

	logger.Log().Infof("%s: response", keys.PathCustomer)

	var balance model.Response[model.BalanceGetResponse]

	unmarshallErr := json.Unmarshal(byteData, &balance)

	if unmarshallErr != nil {
		return model.BalanceGetResponse{}, unmarshallErr
	}
	if balance.Data == nil {
		var errResp model.ErrResponse
		if jsonErr := json.Unmarshal(byteData, &errResp); jsonErr != nil {
			return model.BalanceGetResponse{}, jsonErr
		}
		return model.BalanceGetResponse{}, errors.New(errResp.Message)
	}

	return *balance.Data, nil
}
