package web

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/jailtonjunior94/challenge-client-server/internal/domain/interfaces"
	"go.uber.org/zap"
)

type (
	IHttpRequest interface {
		Request(ctx context.Context, method, url, contentType string, body io.Reader, token string) ([]byte, error)
	}

	httpRequest struct {
		logger     *zap.SugaredLogger
		httpClient interfaces.HttpClient
	}
)

func NewHttpRequest(logger *zap.SugaredLogger, httpClient interfaces.HttpClient) IHttpRequest {
	return &httpRequest{
		logger:     logger,
		httpClient: httpClient,
	}
}

func (r *httpRequest) Request(ctx context.Context, method, url, contentType string, body io.Reader, token string) ([]byte, error) {
	request, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		r.logger.Errorw("error building request", zap.Error(err))
		return nil, err
	}

	request.Header.Set("Content-Type", contentType)
	if token != "" {
		request.Header.Set("Authorization", token)
	}

	response, err := r.httpClient.Do(request)
	if err != nil {
		r.logger.Errorw("error sending request", zap.Error(err))
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		r.logger.Errorw("error readAll", zap.Error(err))
		return nil, err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return nil, fmt.Errorf("requestURL: %s statusCode: %d message: %s", url, response.StatusCode, string(responseBody))
	}
	return responseBody, nil
}
