package xena

import (
	"strconv"
	"time"

	"github.com/xenaex/client-go/xena/xmsg"
)

type PositionsHistoryRequest struct {
	// ID of a concrete position to get an only one.
	PositionId *uint64

	// Parent position id to get all positions with specified parent.
	ParentPositionId *uint64

	// Symbol.
	Symbol *string

	// Position open date "From".
	OpenFrom *time.Time

	// Position open date "To".
	OpenTo *time.Time

	// Position close date "From" (filters by SettlDate).
	CloseFrom *time.Time

	// Position close date "To" (filters by SettlDate).
	CloseTo *time.Time
	// Page number (pagination).
	PageNumber *int

	// Number of positions to get (pagination).
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

func (t *tradingREST) GetPositionsHistory(accountId uint64, request PositionsHistoryRequest) ([]*xmsg.PositionReport, error) {
	query := NewQuery("accounts", strconv.FormatUint(accountId, 10), "positions-history")
	query.AddQueryf("id", request.PositionId)
	query.AddQueryf("parentid", request.ParentPositionId)
	query.AddQueryf("symbol", request.Symbol)
	query.AddQueryf("openfrom", request.OpenFrom)
	query.AddQueryf("opento", request.OpenTo)
	query.AddQueryf("closefrom", request.CloseFrom)
	query.AddQueryf("closeto", request.CloseTo)
	query.AddQueryf("page", request.PageNumber)
	query.AddQueryf("limit", request.Limit)

	resp := make([]*xmsg.PositionReport, 0)
	err := t.sendGet(query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
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

func (p TradeHistoryRequest) SetPage(limit, pageNumber int) TradeHistoryRequest {
	p.Limit = &limit
	p.PageNumber = &pageNumber
	return p

}
