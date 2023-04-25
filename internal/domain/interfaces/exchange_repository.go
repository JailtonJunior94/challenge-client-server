package interfaces

import (
	"context"

	"github.com/jailtonjunior94/challenge-client-server/internal/domain/entities"
)

type ExchangeRepository interface {
	Save(ctx context.Context, exchange *entities.Exchange) error
}
