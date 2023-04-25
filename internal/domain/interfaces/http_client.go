package interfaces

import (
	"net/http"
	"time"

	"github.com/jailtonjunior94/challenge-client-server/configs"
)

type HttpClient interface {
	Do(re *http.Request) (*http.Response, error)
}

func NewHttpClient(config *configs.Config) HttpClient {
	client := &http.Client{
		Timeout: time.Duration(config.HttpClientTimeout) * time.Second,
	}
	return client
}
