package usecases

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jailtonjunior94/challenge-client-server/configs"
	"github.com/jailtonjunior94/challenge-client-server/internal/domain/interfaces"

	"go.uber.org/zap"
)

type CreateFileUseCase struct {
	config         *configs.Config
	logger         *zap.SugaredLogger
	exchangeServer interfaces.ExchangeServer
}

func NewCreateFileUseCase(
	config *configs.Config,
	logger *zap.SugaredLogger,
	exchangeServer interfaces.ExchangeServer,
) *CreateFileUseCase {
	return &CreateFileUseCase{
		config:         config,
		logger:         logger,
		exchangeServer: exchangeServer,
	}
}

func (u *CreateFileUseCase) Execute(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(u.config.ServerTimeoutMs*int(time.Millisecond)))
	defer cancel()

	exchange, err := u.exchangeServer.Exchange(ctx)
	if err != nil {
		u.logger.Errorw(err.Error(), zap.Error(err))
		return err
	}

	file, err := os.Create("cotacao.txt")
	if err != nil {
		u.logger.Errorw(err.Error(), zap.Error(err))
		return err
	}
	defer file.Close()

	size, err := file.Write([]byte(fmt.Sprintf("Dólar: %s", exchange.USDBRL.Bid)))
	if err != nil {
		u.logger.Errorw(err.Error(), zap.Error(err))
		return err
	}

	u.logger.Infof("successfully generated file with the size of %d", size)
	return nil
}
