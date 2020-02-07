package xena

import (
	"time"
)

type PositionsHistoryRequest struct {
	PositionId *uint64

	ParentPositionId *uint64

	Symbol *string

	OpenFrom *time.Time

	OpenTo *time.Time

	CloseFrom *time.Time

	CloseTo *time.Time

	PageNumber *int

	Limit *int
}

func (p PositionsHistoryRequest) SetPositionId(id uint64) PositionsHistoryRequest {
	p.PositionId = &id
	return p
}

func (p PositionsHistoryRequest) SetParentPositionId(id uint64) PositionsHistoryRequest {
	p.ParentPositionId = &id
	return p
}

func (p PositionsHistoryRequest) SetSymbol(symbol string) PositionsHistoryRequest {
	if len(symbol) > 0 {
		p.Symbol = &symbol
	}
	return p
}

func (p PositionsHistoryRequest) SetOpenFrom(openFrom time.Time) PositionsHistoryRequest {
	if !openFrom.IsZero() {
		p.OpenFrom = &openFrom
	}
	return p
}

func (p PositionsHistoryRequest) SetOpenTo(openTo time.Time) PositionsHistoryRequest {
	if !openTo.IsZero() {
		p.OpenTo = &openTo
	}
	return p
}

func (p PositionsHistoryRequest) SetCloseFrom(closeFrom time.Time) PositionsHistoryRequest {
	if !closeFrom.IsZero() {
		p.CloseFrom = &closeFrom
	}
	return p
}

func (p PositionsHistoryRequest) SetCloseTo(closeTo time.Time) PositionsHistoryRequest {
	if !closeTo.IsZero() {
		p.CloseTo = &closeTo
	}
	return p
}

func (p PositionsHistoryRequest) SetPage(limit, pageNumber int) PositionsHistoryRequest {
	p.Limit = &limit
	p.PageNumber = &pageNumber
	return p

}

type TradeHistoryRequest struct {
	TradeId    *string
	ClOrdId    *string
	Symbol     *string
	From       *time.Time
	To         *time.Time
	PageNumber *int
	Limit      *int
}

func (p TradeHistoryRequest) SetTradeId(tradeId string) TradeHistoryRequest {
	p.TradeId = &tradeId
	return p

}

func (p TradeHistoryRequest) SetClOrdId(clOrdId string) TradeHistoryRequest {
	p.ClOrdId = &clOrdId
	return p

}

func (p TradeHistoryRequest) SetSymbol(symbol string) TradeHistoryRequest {
	p.Symbol = &symbol
	return p
}

func (p TradeHistoryRequest) SetFromTo(from, to time.Time) TradeHistoryRequest {
	p.From = &from
	p.To = &to
	return p

}

func (p TradeHistoryRequest) SetPage(pageNumber, limit int) TradeHistoryRequest {
	p.Limit = &limit
	p.PageNumber = &pageNumber
	return p

}
