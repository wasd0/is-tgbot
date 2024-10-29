package model

import "time"

type Response[T any] struct {
	Data *T `json:"data"`
}

type ErrResponse struct {
	Err error `json:"-"`

	Message string    `json:"message"`
	ErrCode int       `json:"code"`
	Time    time.Time `json:"time"`
	ErrDesc string    `json:"description,omitempty"`
}
