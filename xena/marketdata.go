package xena

import (
	"fmt"
	"github.com/xenaex/client-go/xena/fixjson"
	"github.com/xenaex/client-go/xena/messages"
	"sync"
	"time"
)

const (
	wsMdURL = "wss://trading.xena.exchange/api/ws/market-data"
)

// DOMHandler called on order book updated
type DOMHandler func(md MarketDataClient, m *messages.MarketDataRefresh)

// MarketDataClient is the main interface that helps to receive market data
type MarketDataClient interface {
	Client() WsClient
	SubscribeOnDOM(symbol Symbol, handler DOMHandler, opts ...interface{}) (streamID string, err error)
	UnsubscribeOnDOM(streamID string) error
}

type marketData struct {
	client WsClient
	// subscriptions
	subscriptions    map[string]messages.MarketDataRequest
	subscribeMu      *sync.RWMutex
	domSubscriptions map[string]DOMHandler
}

// NewMarketData constructor
func NewMarketData(opts ...WsOption) MarketDataClient {
	md := marketData{
		subscriptions:    make(map[string]messages.MarketDataRequest),
		domSubscriptions: make(map[string]DOMHandler),
		subscribeMu:      &sync.RWMutex{},
	}

	defaultOpts := []WsOption{
		WithURL(wsMdURL),
		WithConnectHandler(md.onConnect),
		WithHandler(md.incomeHandler),
	}
	opts = append(defaultOpts, opts...)

	md.client = NewWsClient(opts...)
	md.client.Connect()

	return &md
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
	r := messages.MarketDataRequest{
		MsgType:                 messages.MsgType_MarketDataRequest,
		MDStreamID:              streamID,
		SubscriptionRequestType: messages.SubscriptionRequestType_SnapshotAndUpdates,
	}

	for _, o := range opts {
		switch v := o.(type) {
		case Throttle:
			d, err := time.ParseDuration(string(v))
			if err != nil {
				return "", err
			}
			r.ThrottleType = messages.ThrottleType_InboundRate
			r.ThrottleTimeInterval = d.Nanoseconds()
			r.ThrottleTimeUnit = messages.ThrottleTimeUnit_Nanoseconds
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
	mth := new(messages.MsgTypeHeader)
	err := fixjson.Unmarshal(msg, mth)
	if err != nil {
		m.client.Logger().Errorf("error: %s. on fixjson.Unmarshal(%s)", err, string(msg))
	}

	switch mth.MsgType {
	case messages.MsgType_LogonMsgType:
		v := new(messages.Logon)
		if _, err := m.unmarshal(msg, v); err == nil {
		}
	case messages.MsgType_MarketDataRequest:
		v := new(messages.MarketDataRequest)
		if _, err := m.unmarshal(msg, v); err == nil {
			//m.client.Logger().Debugf("got %#v", v)
		}

	case messages.MsgType_MarketDataIncrementalRefresh, messages.MsgType_MarketDataSnapshotFullRefresh:
		v := new(messages.MarketDataRefresh)
		if _, err := m.unmarshal(msg, v); err == nil {
			//m.client.Logger().Debugf("got %#v", v)
			m.subscribeMu.RLock()
			if handler, ok := m.domSubscriptions[v.MDStreamID]; ok {
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
