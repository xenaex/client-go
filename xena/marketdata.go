package xena

import (
	"fmt"
	"sync"
	"time"

	"github.com/xenaex/client-go/xena/fixjson"
	"github.com/xenaex/client-go/xena/xmsg"
)

const (
	wsMdURL = "wss://trading.xena.exchange/api/ws/market-data"
)

// MarketDisconnectHandler will be called when connection will be dropped
type MarketDisconnectHandler func(client MarketDataClient)

// DOMHandler called on order book updated
type DOMHandler func(md MarketDataClient, m *xmsg.MarketDataRefresh)

// LogonHandler will be called on Logon response received
type MarketDataLogonHandler func(md MarketDataClient, m *xmsg.Logon)

// MarketDataClient is the main interface that helps to receive market data
type MarketDataClient interface {
	Client() WsClient
	SubscribeOnDOM(symbol Symbol, handler DOMHandler, opts ...interface{}) (streamID string, err error)
	UnsubscribeOnDOM(streamID string) error
	getLogger() Logger
	Connect() (*xmsg.Logon, error)
}

type marketData struct {
	client WsClient
	// subscriptions
	subscriptions    map[string]xmsg.MarketDataRequest
	subscribeMu      *sync.RWMutex
	domSubscriptions map[string]DOMHandler
	handlers         struct {
		logonInternal MarketDataLogonHandler
	}
}

func DefaultMarketDisconnectHandler(client MarketDataClient) {
	go func(client MarketDataClient) {
		l := client.getLogger()
		for {
			logonResponse, err := client.Connect()
			if err != nil {
				l.Debugf("GOT logonResponse ", logonResponse)
			}
			if err == nil {
				break
			}
			l.Errorf("%s on client.ConnectAndLogon()\n", err)
			time.Sleep(time.Minute)
		}
	}(client)
}

// NewMarketData constructor
func NewMarketData(disconnectHandler MarketDisconnectHandler, opts ...WsOption) MarketDataClient {
	md := marketData{
		subscriptions:    make(map[string]xmsg.MarketDataRequest),
		domSubscriptions: make(map[string]DOMHandler),
		subscribeMu:      &sync.RWMutex{},
	}

	defaultOpts := []WsOption{
		WithURL(wsMdURL),
		WithConnectHandler(md.onConnect),
		WithHandler(md.incomeHandler),
		WithDisconnectHandler(func() {
			if disconnectHandler != nil {
				disconnectHandler(&md)
			}
		}),
	}
	opts = append(defaultOpts, opts...)

	md.client = NewWsClient(opts...)

	return &md
}

func (m *marketData) Connect() (*xmsg.Logon, error) {
	msgs := make(chan *xmsg.Logon, 1)
	// errs := make(chan *xmsg.Logon, 1)
	m.handlers.logonInternal = func(md MarketDataClient, m *xmsg.Logon) {
		msgs <- m
		close(msgs)
	}
	defer func() { m.handlers.logonInternal = nil }()
	err := m.client.Connect()
	if err != nil {
		return nil, err
	}
	select {
	case m, ok := <-msgs:
		if ok {
			return m, nil
		}
	case <-time.NewTimer(time.Minute).C:
		m.client.Close()
		// TODO: error time out.
		return nil, fmt.Errorf("timeout logon")
	}
	return nil, fmt.Errorf("something happened")
}

func (m *marketData) Client() WsClient {
	return m.client
}

func (m *marketData) onConnect(c WsClient) {
	// do subscribe on previous streams
	m.subscribeMu.Lock()
	defer m.subscribeMu.Unlock()

	for _, r := range m.subscriptions {
		m.client.Write(r)
	}
}

// UnsubscribeOnDOM from DOM stream
func (m *marketData) UnsubscribeOnDOM(streamID string) error {
	m.subscribeMu.Lock()
	defer m.subscribeMu.Unlock()
	delete(m.domSubscriptions, streamID)
	delete(m.subscriptions, streamID)

	// TODO: send unsubscribe message

	return nil
}

// SubscribeOnDOM subscribe on order book updates. Opts can be: Throttle or AggregateBook
func (m *marketData) SubscribeOnDOM(symbol Symbol, handler DOMHandler, opts ...interface{}) (streamID string, err error) {
	m.subscribeMu.Lock()
	defer m.subscribeMu.Unlock()

	streamID = "DOM:" + string(symbol) + ":aggregated"
	r := xmsg.MarketDataRequest{
		MsgType:                 xmsg.MsgType_MarketDataRequest,
		MDStreamId:              streamID,
		SubscriptionRequestType: xmsg.SubscriptionRequestType_SnapshotAndUpdates,
	}

	for _, o := range opts {
		switch v := o.(type) {
		case Throttle:
			d, err := time.ParseDuration(string(v))
			if err != nil {
				return "", err
			}
			r.ThrottleType = xmsg.ThrottleType_InboundRate
			r.ThrottleTimeInterval = d.Nanoseconds()
			r.ThrottleTimeUnit = xmsg.ThrottleTimeUnit_Nanoseconds
		case AggregateBook:
			r.AggregatedBook = int64(v)
		default:
			return "", fmt.Errorf("unsupported type of %#v", o)
		}
	}

	// subscribe
	if _, ok := m.subscriptions[streamID]; ok {
		return streamID, nil
	}
	m.subscriptions[streamID] = r
	m.domSubscriptions[streamID] = handler

	if m.client.IsConnected() {
		err = m.client.Write(r)
	}
	return streamID, err
}

func (m *marketData) incomeHandler(msg []byte) {
	mth := new(xmsg.MsgTypeHeader)
	err := fixjson.Unmarshal(msg, mth)
	if err != nil {
		m.client.Logger().Errorf("error: %s. on fixjson.Unmarshal(%s)", err, string(msg))
	}

	switch mth.MsgType {
	case xmsg.MsgType_LogonMsgType:
		v := new(xmsg.Logon)
		_, err := m.unmarshal(msg, v)
		if err != nil {
			m.client.Logger().Errorf("unmarshal message type %s", string(msg))
			return
		}
		if m.handlers.logonInternal != nil {
			m.handlers.logonInternal(m, v)
		}

	case xmsg.MsgType_MarketDataRequest:
		v := new(xmsg.MarketDataRequest)
		if _, err := m.unmarshal(msg, v); err == nil {
			// m.client.Logger().Debugf("got %#v", v)
		}

	case xmsg.MsgType_MarketDataIncrementalRefresh, xmsg.MsgType_MarketDataSnapshotFullRefresh:
		v := new(xmsg.MarketDataRefresh)
		if _, err := m.unmarshal(msg, v); err == nil {
			// m.client.Logger().Debugf("got %#v", v)
			m.subscribeMu.RLock()
			if handler, ok := m.domSubscriptions[v.MDStreamId]; ok {
				go handler(m, v)
			}
			m.subscribeMu.RUnlock()
		}

	default:
		m.client.Logger().Errorf("unknown message type %s", string(msg))
	}
}

func (m *marketData) unmarshal(msg []byte, v interface{}) (interface{}, error) {
	err := fixjson.Unmarshal(msg, v)
	if err != nil {
		m.client.Logger().Errorf("error: %s. on fixjson.Unmarshal(%s)", err, string(msg))
		return nil, err
	}

	return v, nil
}

func (m *marketData) getLogger() Logger {
	return m.client.Logger()
}
