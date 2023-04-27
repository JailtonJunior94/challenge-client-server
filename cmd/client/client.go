package main

import (
	"context"

	"github.com/jailtonjunior94/challenge-client-server/configs"
	"github.com/jailtonjunior94/challenge-client-server/internal/domain/interfaces"
	"github.com/jailtonjunior94/challenge-client-server/internal/infra/repositories"
	"github.com/jailtonjunior94/challenge-client-server/internal/usecases"
	"github.com/jailtonjunior94/challenge-client-server/pkg/logger"
	"github.com/jailtonjunior94/challenge-client-server/pkg/web"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	logger := logger.NewLogger(config)
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	httpClient := interfaces.NewHttpClient(config)
	httpRequest := web.NewHttpRequest(logger, httpClient)
	exchangeServer := repositories.NewExchangeServer(config, logger, httpRequest)
	createExchangeUseCase := usecases.NewCreateFileUseCase(config, logger, exchangeServer)

	if err := createExchangeUseCase.Execute(context.Background()); err != nil {
		panic(err)
	}
}
