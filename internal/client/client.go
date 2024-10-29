package client

import (
	"io"
	"is-tgbot/internal/keys"
	"is-tgbot/pkg/logger"
	"net/http"
	"os"
)

func getServerUrl() string {
	var serverUrl = os.Getenv(keys.Server)

	if len(serverUrl) == 0 {
		panic("server environment variable not set")
	}

	return serverUrl
}

func closeResponse(response *http.Response) {
	func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Log().Errorf(err, "failed to close response body")
		}
	}(response.Body)
}
