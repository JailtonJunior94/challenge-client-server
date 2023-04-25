package usecases

import (
	"context"
	"time"

	"github.com/jailtonjunior94/challenge-client-server/configs"
	"github.com/jailtonjunior94/challenge-client-server/internal/domain/dtos"
	"github.com/jailtonjunior94/challenge-client-server/internal/domain/entities"
	"github.com/jailtonjunior94/challenge-client-server/internal/domain/interfaces"
	"go.uber.org/zap"
)

type CreateExchangeUseCase struct {
	config             *configs.Config
	logger             *zap.SugaredLogger
	economyRepository  interfaces.EconomyRepository
	exchangeRepository interfaces.ExchangeRepository
}

func NewCreateExchangeUseCase(
	config *configs.Config,
	logger *zap.SugaredLogger,
	economyRepository interfaces.EconomyRepository,
	exchangeRepository interfaces.ExchangeRepository,
) *CreateExchangeUseCase {
	return &CreateExchangeUseCase{
		config:             config,
		logger:             logger,
		exchangeRepository: exchangeRepository,
		economyRepository:  economyRepository,
	}
}

func (u *CreateExchangeUseCase) Execute(ctx context.Context) (*dtos.USDBRLOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(u.config.HttpClientTimeout*int(time.Millisecond)))
	defer cancel()

	economy, err := u.economyRepository.USDBRL(ctx)
	if err != nil {
		u.logger.Errorw(err.Error())
		return nil, err
	}

	exchange := entities.NewExchange(
		economy.USDBRL.Code,
		economy.USDBRL.Codein,
		economy.USDBRL.Name,
		economy.USDBRL.High,
		economy.USDBRL.Low,
		economy.USDBRL.VarBid,
		economy.USDBRL.PctChange,
		economy.USDBRL.Bid,
		economy.USDBRL.Ask,
	)

	ctxDatabase, cancelDatabase := context.WithTimeout(ctx, time.Duration(u.config.DBTimeout*int(time.Millisecond)))
	defer cancelDatabase()
	if err := u.exchangeRepository.Save(ctxDatabase, exchange); err != nil {
		u.logger.Errorw(err.Error())
		return nil, err
	}
	return economy, nil
}
