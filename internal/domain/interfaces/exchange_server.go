package interfaces

import (
	"context"

	"github.com/jailtonjunior94/challenge-client-server/internal/domain/dtos"
)

type ExchangeServer interface {
	Exchange(ctx context.Context) (*dtos.USDBRLOutput, error)
}
