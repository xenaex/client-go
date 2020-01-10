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
type MarketDisconnectHandler func(client MarketDataClient, logger Logger)

// DOMHandler called on order book updated
type DOMHandler func(md MarketDataClient, m *xmsg.MarketDataRefresh)

type MDHandler func(md MarketDataClient, r *xmsg.MarketDataRequestReject, m *xmsg.MarketDataRefresh)

//MarketDataLogonHandler will be called on Logon response received.
type MarketDataLogonHandler func(md MarketDataClient, m *xmsg.Logon)

// MarketDataRejectHandler will be called on Reject received.
type MarketDataRejectHandler func(md MarketDataClient, m *xmsg.MarketDataRequestReject)

// MarketDataClient is the main interface that helps to receive market data
type MarketDataClient interface {
	Client() WsClient
	SubscribeOnDOM(symbol Symbol, handler DOMHandler, opts ...interface{}) (streamID string, err error)
	UnsubscribeOnDOM(streamID string) error
	Connect() (*xmsg.Logon, error)
	SetDisconnectHandler(handler MarketDisconnectHandler)

	SubscribeOnCandles(symbol, timeframe string, handler MDHandler, opts ...interface{}) (streamID string, err error)
	SubscribeOnDom(symbol string, handler MDHandler, opts ...interface{}) (streamID string, err error)
	SubscribeOnTrades(symbol string, handler MDHandler, opts ...interface{}) (streamID string, err error)
	SubscribeOnMarketWatch(symbol string, handler MDHandler) (streamID string, err error)
}
type marketData struct {
	client WsClient
	// subscriptions
	subscriptions    map[string]xmsg.MarketDataRequest
	subscribeMu      *sync.RWMutex
	domSubscriptions map[string]DOMHandler
	userHandlers     map[string]MDHandler
	handlers         struct {
		logonInternal MarketDataLogonHandler
		reject        MarketDataRejectHandler
	}
	mutexLogon sync.Mutex
}

//DefaultMarketDisconnectHandler default Disconnect handler.
func DefaultMarketDisconnectHandler(client MarketDataClient, logger Logger) {
	go func(client MarketDataClient) {
		for {
			logonResponse, err := client.Connect()
			if err != nil {
				logger.Debugf("GOT logonResponse ", logonResponse)
			}
			if err == nil {
				break
			}
			logger.Errorf("%s on client.ConnectAndLogon()\n", err)
			time.Sleep(time.Minute)
		}
	}(client)
}

// NewMarketData constructor
func NewMarketData(opts ...WsOption) MarketDataClient {
	md := marketData{
		subscriptions:    make(map[string]xmsg.MarketDataRequest),
		domSubscriptions: make(map[string]DOMHandler),
		userHandlers:     make(map[string]MDHandler),
		subscribeMu:      &sync.RWMutex{},
	}

	defaultOpts := []WsOption{
		WithURL(wsMdURL),
		WithConnectHandler(md.onConnect),
		WithHandler(md.incomeHandler),
	}
	opts = append(defaultOpts, opts...)

	md.client = NewWsClient(opts...)

	return &md
}

func (m *marketData) Connect() (*xmsg.Logon, error) {
	m.mutexLogon.Lock()
	defer m.mutexLogon.Unlock()
	logMsg := make(chan *xmsg.Logon, 1)
	m.handlers.logonInternal = func(md MarketDataClient, m *xmsg.Logon) {
		logMsg <- m
		close(logMsg)
	}
	defer func() { m.handlers.logonInternal = nil }()
	err := m.client.Connect()
	if err != nil {
		return nil, err
	}
	select {
	case m, ok := <-logMsg:
		if ok && m != nil {
			if len(m.RejectText) != 0 {
				return nil, fmt.Errorf("reject reason: %s", m.RejectText)
			}
			return m, nil
		}
	case <-time.NewTimer(m.client.getConf().connectTimeoutInterval).C:
		m.client.Close()
		return nil, fmt.Errorf("timeout logon")
	}
	return nil, fmt.Errorf("login response didn't come")
}

func (m *marketData) Client() WsClient {
	return m.client
}

func (m *marketData) onConnect(WsClient) {
	// do subscribe on previous streams
	m.subscribeMu.Lock()
	defer m.subscribeMu.Unlock()

	for _, r := range m.subscriptions {
		err := m.client.Write(r)
		if err == nil {
			m.client.Logger().Errorf("%s on m.client.Write(%s)", err, &r)
		}

	}
}

func (m *marketData) ListenReject(handler MarketDataRejectHandler) {
	m.handlers.reject = handler
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
		case DOMThrottle:
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

func (m *marketData) subscribe(req xmsg.MarketDataRequest, handler MDHandler) (err error) {
	m.subscribeMu.Lock()
	defer m.subscribeMu.Unlock()

	// subscribe
	if _, ok := m.subscriptions[req.MDStreamId]; ok {
		return nil
	}
	m.subscriptions[req.MDStreamId] = req
	m.userHandlers[req.MDStreamId] = handler

	if m.client.IsConnected() {
		err = m.client.Write(req)
		return err
	}
	return nil
}

func (m *marketData) createRequest(MsgType, streamName, symbol, streamPostfix string) (req xmsg.MarketDataRequest, err error) {
	streamId := streamName
	if len(symbol) > 0 {
		streamId += fmt.Sprintf(":%s", symbol)
	}
	if len(streamPostfix) > 0 {
		streamId += fmt.Sprintf(":%s", streamPostfix)
	}
	req.MsgType = MsgType
	req.MDStreamId = streamId
	req.SubscriptionRequestType = xmsg.SubscriptionRequestType_SnapshotAndUpdates
	req.ThrottleType = xmsg.ThrottleType_OutstandingRequests
	req.ThrottleTimeUnit = xmsg.ThrottleTimeUnit_Nanoseconds

	return req, nil
}

func (m *marketData) SubscribeOnCandles(symbol, timeframe string, handler MDHandler, opts ...interface{}) (streamID string, err error) {
	aggregatedBook := int64(0)
	throttleInterval := int64(250)
	if len(timeframe) == 0 {
		timeframe = "1m"
	}
	for _, o := range opts {
		switch v := o.(type) {
		case CandlesThrottle:
			d, err := time.ParseDuration(string(v))
			if err != nil {
				return "", err
			}
			throttleInterval = d.Nanoseconds()
		case AggregateBook:
			aggregatedBook = int64(v)
		default:
			return "", fmt.Errorf("unsupported type of %#v", o)
		}
	}
	req := xmsg.MarketDataRequest{}
	req, err = m.createRequest(xmsg.MsgType_MarketDataRequest, "candles", symbol, timeframe)
	if err != nil {
		return "", err
	}
	req.AggregatedBook = aggregatedBook

	req.ThrottleType = xmsg.ThrottleType_OutstandingRequests
	req.ThrottleTimeInterval = throttleInterval
	req.ThrottleTimeUnit = xmsg.ThrottleTimeUnit_Nanoseconds

	err = m.subscribe(req, handler)
	if err != nil {
		return "", err
	}
	return req.MDStreamId, nil
}
func (m *marketData) SubscribeOnDom(symbol string, handler MDHandler, opts ...interface{}) (streamID string, err error) {
	aggregatedBook := int64(0)
	throttleInterval := int64(500)
	req := xmsg.MarketDataRequest{}
	for _, o := range opts {
		switch v := o.(type) {
		case DOMThrottle:
			d, err := time.ParseDuration(string(v))
			if err != nil {
				return "", err
			}
			throttleInterval = d.Nanoseconds()
		case AggregateBook:
			aggregatedBook = int64(v)
		default:
			return "", fmt.Errorf("unsupported type of %#v", o)
		}
	}

	req, err = m.createRequest(xmsg.MsgType_MarketDataRequest, "DOM", symbol, "aggregated")
	if err != nil {
		return "", err
	}
	req.AggregatedBook = aggregatedBook
	req.ThrottleType = xmsg.ThrottleType_OutstandingRequests
	req.ThrottleTimeInterval = throttleInterval
	req.ThrottleTimeUnit = xmsg.ThrottleTimeUnit_Nanoseconds

	err = m.subscribe(req, handler)
	if err != nil {
		return "", err
	}
	return req.MDStreamId, nil
}
func (m *marketData) SubscribeOnTrades(symbol string, handler MDHandler, opts ...interface{}) (streamID string, err error) {
	throttleInterval := int64(500)
	req := xmsg.MarketDataRequest{}
	for _, o := range opts {
		switch v := o.(type) {
		case TradesThrottle:
			if len(v) != 0 {
				d, err := time.ParseDuration(string(v))
				if err != nil {
					return "", err
				}
				throttleInterval = d.Nanoseconds()
			} else {
				throttleInterval = 0
			}
		default:
			return "", fmt.Errorf("unsupported type of %#v", o)
		}
	}

	req, err = m.createRequest(xmsg.MsgType_MarketDataRequest, "trades", symbol, "")
	req.ThrottleTimeInterval = throttleInterval
	if err != nil {
		return "", err
	}
	err = m.subscribe(req, handler)
	if err != nil {
		return "", err
	}
	return req.MDStreamId, nil
}
func (m *marketData) SubscribeOnMarketWatch(symbol string, handler MDHandler) (streamID string, err error) {
	req := xmsg.MarketDataRequest{}
	req, err = m.createRequest(xmsg.MsgType_MarketDataRequest, "market-watch", "", "")
	if err != nil {
		return "", err
	}

	err = m.subscribe(req, handler)
	if err != nil {
		return "", err
	}

	return req.MDStreamId, nil
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

	case xmsg.MsgType_RejectMsgType:
		v := new(xmsg.MarketDataRequestReject)
		if _, err := m.unmarshal(msg, v); err == nil {
			if m.handlers.reject != nil {
				go m.handlers.reject(m, v)
			}
			if h, ok := m.userHandlers[v.MDStreamId]; ok {
				h(m, v, nil)
			}
		}

	case xmsg.MsgType_MarketDataRequest:
		v := new(xmsg.MarketDataRefresh)
		if _, err := m.unmarshal(msg, v); err == nil {
			if h, ok := m.userHandlers[v.MDStreamId]; ok {
				h(m, nil, v)
			}
		}

	case xmsg.MsgType_MarketDataIncrementalRefresh, xmsg.MsgType_MarketDataSnapshotFullRefresh:
		v := new(xmsg.MarketDataRefresh)
		if _, err := m.unmarshal(msg, v); err == nil {
			// m.client.Logger().Debugf("got %#v", v)
			m.subscribeMu.RLock()
			if handler, ok := m.domSubscriptions[v.MDStreamId]; ok {
				go handler(m, v)
			}
			if h, ok := m.userHandlers[v.MDStreamId]; ok {
				h(m, nil, v)
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

// SetDisconnectHandler set disconnect handler.
func (m *marketData) SetDisconnectHandler(handler MarketDisconnectHandler) {
	m.client.SetDisconnectHandler(func() {
		handler(m, m.client.Logger())
	})
}
