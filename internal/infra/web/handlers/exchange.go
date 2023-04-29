package handlers

import (
	"context"
	"net/http"

	"github.com/jailtonjunior94/challenge-client-server/internal/usecases"
	"github.com/jailtonjunior94/challenge-client-server/pkg/responses"
)

type ExchangeHandler struct {
	createExchange *usecases.CreateExchangeUseCase
}

func NewExchangeHandler(createExchange *usecases.CreateExchangeUseCase) *ExchangeHandler {
	return &ExchangeHandler{
		createExchange: createExchange,
	}
}

func (h *ExchangeHandler) CreateExchange(w http.ResponseWriter, r *http.Request) {
	economy, err := h.createExchange.Execute(context.Background())
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	responses.JSON(w, http.StatusOK, economy)
}
