package model

type BalanceGetRequest struct {
	CustomerId   *int64  `json:"customerId"`
	TelegramId   *int64  `json:"telegramId"`
	CurrencyCode *string `json:"currencyCode"`
}

type BalanceGetResponse struct {
	BalanceId int64   `json:"balanceId"`
	Currency  string  `json:"currency"`
	Sum       *string `json:"sum"`
}
