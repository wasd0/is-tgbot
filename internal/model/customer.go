package model

import "time"

type CustomerResponse struct {
	ID         int64     `json:"id"`
	TelegramID *int64    `json:"telegramId,omitempty"`
	Blocked    bool      `json:"blocked"`
	CreateDate time.Time `json:"createDate"`
	CountryIso string    `json:"countryIso"`
}

type CustomerGetRequest struct {
	ID         *int64 `json:"id,omitempty"`
	TelegramID *int64 `json:"telegramId,omitempty"`
}
