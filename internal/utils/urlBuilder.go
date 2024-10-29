package utils

import (
	"fmt"
	"strings"
)

type UrlBuilder struct {
	url    string
	params map[string]string
}

func NewUrl(url string) *UrlBuilder {
	return &UrlBuilder{url: url, params: make(map[string]string)}
}

func (u *UrlBuilder) AddParam(key string, value string) *UrlBuilder {
	u.params[key] = value
	return u
}

func (u *UrlBuilder) Build() string {
	params := make([]string, 0, len(u.params))
	if len(u.params) > 0 {
		u.url += "?"

		for k, v := range u.params {
			params = append(params, fmt.Sprintf("%s=%s", k, v))
		}
	}
	return u.url + strings.Join(params, "&")
}
