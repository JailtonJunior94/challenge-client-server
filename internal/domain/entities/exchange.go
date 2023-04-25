package entities

import "github.com/google/uuid"

type Exchange struct {
	ID        string
	Code      string
	CodeIn    string
	Name      string
	High      string
	Low       string
	VarBid    string
	PctChange string
	Bid       string
	Ask       string
}

func NewExchange(code, codeIn, name, high, low, varBid, pctChange, bid, ask string) *Exchange {
	return &Exchange{
		ID:        uuid.New().String(),
		Code:      code,
		CodeIn:    codeIn,
		Name:      name,
		High:      high,
		Low:       low,
		VarBid:    varBid,
		PctChange: pctChange,
		Bid:       bid,
		Ask:       ask,
	}
}
