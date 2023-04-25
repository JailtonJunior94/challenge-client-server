package interfaces

import (
	"context"

	"github.com/jailtonjunior94/challenge-client-server/internal/domain/dtos"
)

type EconomyRepository interface {
	USDBRL(ctx context.Context) (*dtos.USDBRLOutput, error)
}
