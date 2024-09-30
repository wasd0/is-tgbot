package client

import (
	"io"
	"is-tgbot/pkg/logger"
	"net/http"
)

func closeResponse(response *http.Response) {
	func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Log().Errorf(err, "failed to close response body")
		}
	}(response.Body)
}
