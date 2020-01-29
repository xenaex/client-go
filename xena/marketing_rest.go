package xena

import (
	"encoding/json"
	"time"

	"github.com/xenaex/client-go/xena/fixjson"
	"github.com/xenaex/client-go/xena/xmsg"
)

//NewMarketDataREST creates rest client of Xena market data.
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

//MarketDataREST is rest client interface of the Xena market data.
type MarketDataREST interface {
	//GetServerTime returns server time.
	GetServerTime() (time.Time, error)
	//GetInstruments returns instruments.
	GetInstruments() ([]*xmsg.Instrument, error)
	//GetTrades returns traders.
	GetTrades(symbol string, from, to time.Time, page, limit int64) (*xmsg.MarketDataRefresh, error)
	//GetDom returns dom.
	GetDom(symbol string) (*xmsg.MarketDataRefresh, error)
	//GetCandles returns candles.
	GetCandles(symbol string, timeFrame string, from, to time.Time) (*xmsg.MarketDataRefresh, error)
}

type marketDataREST struct {
	baseREST
}

func (m *marketDataREST) GetCandles(symbol string, timeFrame string, from, to time.Time) (*xmsg.MarketDataRefresh, error) {
	method := "candles"
	resp, body, err := m.get(newQuery("market-data", "candles", symbol, timeFrame).addQueryInt("from", from.UnixNano()).addQueryInt("to", to.UnixNano()))
	if err != nil {
		m.config.logger.Errorf("%s on m.get(%s, %s, %s, %s)", err, method, symbol, from, to)
		return nil, err
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	msg := &xmsg.MarketDataRefresh{}
	err = fixjson.Unmarshal(body, msg)
	if err != nil {
		m.config.logger.Errorf("%s on fixjson.Unmarshal()", err)
		return nil, err
	}
	return msg, nil
}

func (m *marketDataREST) GetDom(symbol string) (*xmsg.MarketDataRefresh, error) {
	const method = "dom"
	resp, body, err := m.get(newQuery("market-data", method, symbol))
	if err != nil {
		m.config.logger.Errorf("%s on m.get(%s, %s)", err, method, symbol)
		return nil, err
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	msg := &xmsg.MarketDataRefresh{}
	err = fixjson.Unmarshal(body, msg)
	m.config.logger.Debugf("code %s, body %s", resp.Status, body)
	if err != nil {
		m.config.logger.Errorf("%s on fixjson.Unmarshal()", err)
		return nil, err
	}
	return msg, nil
}

func (m *marketDataREST) GetTrades(symbol string, from, to time.Time, page, limit int64) (*xmsg.MarketDataRefresh, error) {
	const method = "trades"
	query := newQuery("market-data", method, symbol).
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
	err = fixjson.Unmarshal(body, msg)
	if err != nil {
		m.config.logger.Errorf("%s on fixjson.Unmarshal()", err)
		return nil, err
	}
	return msg, nil
}

func (m *marketDataREST) GetServerTime() (time.Time, error) {
	const method = "server-time"
	resp, body, err := m.get(newQuery("market-data", method))
	if err != nil {
		m.config.logger.Errorf("%s on m.get(%s, %s, %s, %s, %d, %d)", err, method)
		return time.Time{}, err
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	msg := &xmsg.Heartbeat{}
	err = fixjson.Unmarshal(body, msg)
	if err != nil {
		m.config.logger.Errorf("%s on fixjson.Unmarshal()", err)
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
	// m.config.logger.Debugf("body %s", body)
	// fmt.Printf("body %s", body)
	err = json.Unmarshal(body, &msg)
	if err != nil {
		m.config.logger.Errorf("%s on fixjson.Unmarshal()", err)
		return []*xmsg.Instrument{}, err
	}
	return msg, nil
}
