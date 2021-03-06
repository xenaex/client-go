package xena

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/xenaex/client-go/xena/xmsg"
)

// NewMarketDataREST creates a rest client of Xena market data.
func NewMarketDataREST(options ...RestOption) MarketDataREST {
	cfg := &restConf{}
	for _, ots := range []RestOption{withRestDefaultLogger, WithRestMarketDataHost, withRestDefaultTimeout, WithRestUserAgent(userAgent)} {
		ots(cfg)
	}
	for _, ots := range options {
		ots(cfg)
	}
	return &marketDataREST{
		baseREST: newBaseREST(cfg),
	}
}

// MarketDataREST is a rest client interface of the Xena market data.
type MarketDataREST interface {
	// GetServerTime returns server time.
	GetServerTime() (time.Time, error)
	// GetInstruments returns instruments.
	GetInstruments() ([]*xmsg.Instrument, error)
	// GetTrades returns trades.
	GetTrades(symbol string, from, to time.Time, page, limit int64) (*xmsg.MarketDataRefresh, error)
	// GetDom returns dom.
	GetDom(symbol string, opts ...interface{}) (*xmsg.MarketDataRefresh, error)
	// GetCandles returns candles.
	GetCandles(symbol string, timeFrame string, from, to time.Time) (*xmsg.MarketDataRefresh, error)
}

type marketDataREST struct {
	baseREST
}

func (m *marketDataREST) GetCandles(symbol string, timeFrame string, from, to time.Time) (*xmsg.MarketDataRefresh, error) {
	method := "candles"
	resp, body, err := m.get(newQuery("market-data/v2", "candles", symbol, timeFrame).addQueryInt("from", from.UnixNano()).addQueryInt("to", to.UnixNano()))
	if err != nil {
		m.config.logger.Errorf("%s on m.get(%s, %s, %s, %s)", err, method, symbol, from, to)
		return nil, err
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	msg := &xmsg.MarketDataRefresh{}
	err = json.Unmarshal(body, msg)
	if err != nil {
		m.config.logger.Errorf("%s on json.Unmarshal()", err)
		return nil, err
	}
	return msg, nil
}

func (m *marketDataREST) GetDom(symbol string, opts ...interface{}) (*xmsg.MarketDataRefresh, error) {
	const method = "dom"
	aggregatedBook := int64(AggregateBookDefault)
	marketDepth := int64(MarketDepthDefault)
	throttling := int64(RestThrottleDOMDefault)

	for _, o := range opts {
		switch v := o.(type) {
		case AggregateBook:
			aggregatedBook = int64(v)
		case MarketDepth:
			marketDepth = int64(v)
		case RestDOMThrottle:
			throttling = int64(v)
		default:
			return nil, fmt.Errorf("unsupported type of %#v", o)
		}
	}

	query := newQuery("market-data/v2", method, symbol).
		addQueryf("aggr", &aggregatedBook).
		addQueryf("depth", &marketDepth).
		addQueryf("throtling", &throttling)

	resp, body, err := m.get(query)
	if err != nil {
		m.config.logger.Errorf("%s on m.get(%s, %s, %d, %d)", err, method, symbol, aggregatedBook, marketDepth)
		return nil, err
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	msg := &xmsg.MarketDataRefresh{}
	err = json.Unmarshal(body, msg)
	if err != nil {
		m.config.logger.Errorf("%s on json.Unmarshal()", err)
		return nil, err
	}
	return msg, nil
}

func (m *marketDataREST) GetTrades(symbol string, from, to time.Time, page, limit int64) (*xmsg.MarketDataRefresh, error) {
	const method = "trades"
	query := newQuery("market-data/v2", method, symbol).
		addQueryInt("from", from.UnixNano()).
		addQueryInt("to", to.UnixNano()).
		addQueryInt("page", page).
		addQueryInt("limit", limit)
	resp, body, err := m.get(query)
	if err != nil {
		m.config.logger.Errorf("%s on m.get(%s, %s, %s, %s, %d, %d)", err, method, symbol, from, to, page, limit)
		return nil, err
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	msg := &xmsg.MarketDataRefresh{}
	err = json.Unmarshal(body, msg)
	if err != nil {
		m.config.logger.Errorf("%s on json.Unmarshal()", err)
		return nil, err
	}
	return msg, nil
}

func (m *marketDataREST) GetServerTime() (time.Time, error) {
	const method = "server-time"
	resp, body, err := m.get(newQuery("market-data/v2", method))
	if err != nil {
		m.config.logger.Errorf("%s on m.get(%s, %s, %s, %s, %d, %d)", err, method)
		return time.Time{}, err
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	msg := &xmsg.Heartbeat{}
	err = json.Unmarshal(body, msg)
	if err != nil {
		m.config.logger.Errorf("%s on json.Unmarshal()", err)
		return time.Time{}, err
	}
	return time.Unix(0, msg.TransactTime), nil
}

func (m *marketDataREST) GetInstruments() ([]*xmsg.Instrument, error) {
	const method = "instruments"
	resp, body, err := m.get(newQuery("public", method))
	if err != nil {
		m.config.logger.Errorf("%s on m.get(%s)", err, method)
		return []*xmsg.Instrument{}, err
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	msg := make([]*xmsg.Instrument, 0)
	err = json.Unmarshal(body, &msg)
	if err != nil {
		m.config.logger.Errorf("%s on fixjson.Unmarshal()", err)
		return []*xmsg.Instrument{}, err
	}
	return msg, nil
}
