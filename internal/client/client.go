package client

import (
	"github.com/wasd0/is-common/pkg/logger"
	"io"
	"is-tgbot/internal/keys"
	"net/http"
	"os"
)

func getServerUrl() string {
	var serverUrl = os.Getenv(keys.EnvServer)

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
