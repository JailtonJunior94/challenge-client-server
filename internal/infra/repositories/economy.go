package repositories

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/jailtonjunior94/challenge-client-server/configs"
	"github.com/jailtonjunior94/challenge-client-server/internal/domain/dtos"
	"github.com/jailtonjunior94/challenge-client-server/internal/domain/interfaces"
	"github.com/jailtonjunior94/challenge-client-server/pkg/web"

	"go.uber.org/zap"
)

type economyRepository struct {
	config      *configs.Config
	logger      *zap.SugaredLogger
	httpRequest web.IHttpRequest
}

func NewEconomy(config *configs.Config, logger *zap.SugaredLogger, httpRequest web.IHttpRequest) interfaces.EconomyRepository {
	return &economyRepository{
		config:      config,
		logger:      logger,
		httpRequest: httpRequest,
	}
}

func (r *economyRepository) USDBRL(ctx context.Context) (*dtos.USDBRLOutput, error) {
	apiUrl, err := url.Parse(r.config.EconomyBaseURL)
	if err != nil {
		r.logger.Errorw("error building apiUrl", zap.Error(err))
		return nil, err
	}

	response, err := r.httpRequest.Request(ctx, http.MethodGet, apiUrl.String(), "application/json", nil, "")
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
