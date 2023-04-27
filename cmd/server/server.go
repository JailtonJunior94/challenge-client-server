package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/jailtonjunior94/challenge-client-server/configs"
	"github.com/jailtonjunior94/challenge-client-server/internal/domain/interfaces"
	"github.com/jailtonjunior94/challenge-client-server/internal/infra/repositories"
	"github.com/jailtonjunior94/challenge-client-server/internal/infra/web/handlers"
	"github.com/jailtonjunior94/challenge-client-server/internal/usecases"
	"github.com/jailtonjunior94/challenge-client-server/pkg/logger"
	"github.com/jailtonjunior94/challenge-client-server/pkg/web"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
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

	db, err := sql.Open(config.DriverNameDB, config.ConnectionStringDB)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	httpClient := interfaces.NewHttpClient(config)
	httpRequest := web.NewHttpRequest(logger, httpClient)
	exchangeRepository := repositories.NewExchangeRepository(db)
	economyRepository := repositories.NewEconomy(config, logger, httpRequest)
	createExchangeUseCase := usecases.NewCreateExchangeUseCase(config, logger, economyRepository, exchangeRepository)
	exchangeHandler := handlers.NewExchangeHandler(createExchangeUseCase)

	router := chi.NewRouter()
	router.Use(middleware.Heartbeat("/health"))
	router.Post("/cotacao", exchangeHandler.CreateExchange)

	server := http.Server{
		ReadTimeout:       time.Duration(config.ServerTimeout) * time.Second,
		ReadHeaderTimeout: time.Duration(config.ServerTimeout) * time.Second,
		Handler:           router,
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", config.ServerPort))
	if err != nil {
		panic(err)
	}
	server.Serve(listener)
}
