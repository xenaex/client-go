package xena

import (
	"encoding/json"
	"time"

	"github.com/xenaex/client-go/xena/fixjson"
	"github.com/xenaex/client-go/xena/xmsg"
)

func NewMarketDataRpc(options ...RestOption) MarketDataRPC {
	cfg := &rpcConf{}
	for _, ots := range []RestOption{withRestDefaultLogger, WithRestMarketDataHost, withRestDefaultTimeout, WithRestUserAgent(userAgent)} {
		ots(cfg)
	}
	for _, ots := range options {
		ots(cfg)
	}
	return &marketDataRPC{
		baseRPC: newBaseRPC(cfg),
	}
}

type MarketDataRPC interface {
	GetServerTime() (time.Time, error)
	GetInstruments() ([]*xmsg.Instrument, error)
	GetTrades(symbol string, from, to time.Time, page, limit int64) (*xmsg.MarketDataRefresh, error)
	GetDom(symbol string) (*xmsg.MarketDataRefresh, error)
	GetCandles(symbol string, timeFrame string, from, to time.Time) (*xmsg.MarketDataRefresh, error)
}

type marketDataRPC struct {
	baseRPC
}

func (m *marketDataRPC) GetCandles(symbol string, timeFrame string, from, to time.Time) (*xmsg.MarketDataRefresh, error) {
	method := "candles"
	resp, body, err := m.get(NewQuery("market-data", "candles", symbol, timeFrame).AddQueryInt("from", from.UnixNano()).AddQueryInt("to", to.UnixNano()))
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

func (m *marketDataRPC) GetDom(symbol string) (*xmsg.MarketDataRefresh, error) {
	const method = "dom"
	resp, body, err := m.get(NewQuery("market-data", method, symbol))
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

func (m *marketDataRPC) GetTrades(symbol string, from, to time.Time, page, limit int64) (*xmsg.MarketDataRefresh, error) {
	const method = "trades"
	query := NewQuery("market-data", method, symbol).
		AddQueryInt("from", from.UnixNano()).
		AddQueryInt("to", to.UnixNano()).
		AddQueryInt("page", page).
		AddQueryInt("limit", limit)
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

func (m *marketDataRPC) GetServerTime() (time.Time, error) {
	const method = "server-time"
	resp, body, err := m.get(NewQuery("market-data", method))
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

func (m *marketDataRPC) GetInstruments() ([]*xmsg.Instrument, error) {
	const method = "instruments"
	resp, body, err := m.get(NewQuery("public", method))
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
