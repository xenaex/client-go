package xena

import (
	"strconv"
	"time"

	"github.com/xenaex/client-go/xena/xmsg"
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

func (t *tradingREST) GetPositionsHistory(accountId uint64, request PositionsHistoryRequest) ([]*xmsg.PositionReport, error) {
	query := newQuery("accounts", strconv.FormatUint(accountId, 10), "positions-history")
	query.addQueryf("id", request.PositionId)
	query.addQueryf("parentid", request.ParentPositionId)
	query.addQueryf("symbol", request.Symbol)
	query.addQueryf("openfrom", request.OpenFrom)
	query.addQueryf("opento", request.OpenTo)
	query.addQueryf("closefrom", request.CloseFrom)
	query.addQueryf("closeto", request.CloseTo)
	query.addQueryf("page", request.PageNumber)
	query.addQueryf("limit", request.Limit)

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

func (p TradeHistoryRequest) SetPage(pageNumber, limit int) TradeHistoryRequest {
	p.Limit = &limit
	p.PageNumber = &pageNumber
	return p

}
