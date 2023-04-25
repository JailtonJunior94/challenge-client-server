package repositories

import (
	"context"
	"database/sql"

	"github.com/jailtonjunior94/challenge-client-server/internal/domain/entities"
	"github.com/jailtonjunior94/challenge-client-server/internal/domain/interfaces"
)

type exchangeRepository struct {
	DB *sql.DB
}

func NewExchangeRepository(db *sql.DB) interfaces.ExchangeRepository {
	return &exchangeRepository{DB: db}
}

func (r *exchangeRepository) Save(ctx context.Context, e *entities.Exchange) error {
	query := "INSERT INTO exchange (id, code, code_in, name, high, low, var_bid, pct_change, bid, ask) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	stmt, err := r.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, e.ID, e.Code, e.CodeIn, e.Name, e.High, e.Low, e.VarBid, e.PctChange, e.Bid, e.Ask)
	if err != nil {
		return err
	}
	return nil
}
