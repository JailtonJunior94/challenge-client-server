package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/jailtonjunior94/challenge-client-server/configs"
	"github.com/jailtonjunior94/challenge-client-server/internal/domain/dtos"
	"github.com/jailtonjunior94/challenge-client-server/internal/domain/interfaces"

	"go.uber.org/zap"
)

type economyRepository struct {
	config     *configs.Config
	logger     *zap.SugaredLogger
	httpClient interfaces.HttpClient
}

func NewEconomy(config *configs.Config, logger *zap.SugaredLogger, httpClient interfaces.HttpClient) interfaces.EconomyRepository {
	return &economyRepository{
		config:     config,
		logger:     logger,
		httpClient: httpClient,
	}
}

func (r *economyRepository) USDBRL(ctx context.Context) (*dtos.USDBRLOutput, error) {
	apiUrl, err := url.Parse(r.config.EconomyBaseURL)
	if err != nil {
		r.logger.Errorw("error building apiUrl", zap.Error(err))
		return nil, err
	}

	response, err := r.request(ctx, http.MethodGet, apiUrl.String(), "application/json", nil, "")
	if err != nil {
		r.logger.Errorw("error fetch usd brl in external api", zap.Error(err))
		return nil, err
	}

	var output *dtos.USDBRLOutput
	err = json.Unmarshal(response, &output)
	if err != nil {
		r.logger.Errorw("error unmarshal response", zap.Error(err))
		return nil, err
	}
	return output, nil
}

func (r *economyRepository) request(ctx context.Context, method, url, contentType string, body io.Reader, token string) ([]byte, error) {
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
