package xena

import (
	"errors"
	"github.com/gorilla/websocket"
	"github.com/xenaex/client-go/xena/fixjson"
	"sync"
	"time"
)

const (
	heartbeatMsg        = `{"35":"0"}`
	wsReconnectInterval = time.Second
	wsHeartbeatInterval = 3 * time.Second
)

// Handler handle raw websocket message
type Handler func(msg []byte)

// ErrHandler handles errors
type ErrHandler func(err error)

// ConnectHandler will be called when connection will be established
type ConnectHandler func(client WsClient)

// DisconnectHandler will be called when connection will be dropped
type DisconnectHandler func()

// WsOption is a function that can alter behavior
type WsOption func(s *wsClient)

// WsClient websocket client interface
type WsClient interface {
	IsConnected() bool
	//OnConnect(ConnectHandler)
	Connect()
	With(opt WsOption)
	Close()
	Write(m interface{}) error
	WriteString(msg string) error
	WriteBytes(data []byte) error
	Logger() Logger
}

type wsClient struct {
	// settings
	url               string
	reconnectInterval time.Duration
	heartbeatInterval time.Duration

	// handlers
	handler           Handler
	errHandler        ErrHandler
	connectHandler    ConnectHandler
	disconnectHandler DisconnectHandler

	logger Logger

	// no customizable elements
	conn     *websocket.Conn
	mu       sync.Mutex
	stopChan chan struct{}
	close    bool
}

// NewWsClient websocket client constructor
func NewWsClient(opts ...WsOption) WsClient {
	c := wsClient{
		reconnectInterval: wsReconnectInterval,
		heartbeatInterval: wsHeartbeatInterval,
		logger:            newLogger(),
		stopChan:          make(chan struct{}),
	}
	c.initDefaultHandlers()

	// Apply options
	for _, optionFunc := range opts {
		optionFunc(&c)
	}

	return &c
}

// With applies modification of behavior
func (c *wsClient) With(opt WsOption) {
	opt(c)
}

// Connect start connecting
func (c *wsClient) Connect() {
	go func() {
		if err := c.connect(); err != nil {
			c.logger.Errorf("Failed to connect to %s: %s", c.url, err)
		}
	}()
}

/*
// OnConnect register the main handler that will be called after establishing the connection
func (c *wsClient) OnConnect(fnc ConnectHandler) {
	c.With(WithConnectHandler(fnc))
	go func() {
		_ = c.connect()
	}()
}

*/

// IsConnected return true if connection was established
func (c *wsClient) IsConnected() bool {
	return c.conn != nil
}

// Logger return logger
func (c *wsClient) Logger() Logger {
	return c.logger
}

// Write marshal and send message to socket
func (c *wsClient) Write(v interface{}) error {
	data, err := fixjson.Marshal(v)
	if err != nil {
		return err
	}
	return c.WriteBytes(data)
}

// WriteString send message to socket
func (c *wsClient) WriteString(msg string) error {
	return c.WriteBytes([]byte(msg))
}

// WriteBytes send message to socket
func (c *wsClient) WriteBytes(data []byte) error {
	if c.conn == nil {
		return errors.New("ws connection not established yet")
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	err := c.conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		c.handleError(err)
		go c.stop()
	}

	c.logger.Debugf("ws. sent: %s", string(data))

	return err
}

// Close connection permanent
func (c *wsClient) Close() {
	c.close = true
	go c.stop()
}

func (c *wsClient) initDefaultHandlers() {
	c.handler = func(message []byte) {}
	c.errHandler = func(err error) {}
	c.connectHandler = func(client WsClient) {}
	c.disconnectHandler = func() {}
}

func (c *wsClient) connect() error {
	c.mu.Lock()
	if c.conn != nil {
		c.mu.Unlock()
		return nil
	}
	c.mu.Unlock()

	err := c.connectAndListen()
	if err != nil {
		return err
	}
	go c.heartbeats()

	return nil
}

func (c *wsClient) connectAndListen() error {
	if c.close {
		err := errors.New("connection was closed by user")
		c.handleError(err)
		return err
	}

	// Connection loop
	for {
		conn, _, err := websocket.DefaultDialer.Dial(c.url, nil)
		if err != nil {
			c.handleError(err)
			time.Sleep(c.reconnectInterval)
			continue
		}
		// OK
		c.mu.Lock()
		c.conn = conn
		c.stopChan = make(chan struct{})
		c.mu.Unlock()
		break
	}

	c.connectHandler(c)
	c.logger.Debugf("ws. connected")

	go c.listen()
	return nil
}

func (c *wsClient) listen() {
	defer func() {
		c.mu.Lock()
		err := c.conn.Close()
		c.mu.Unlock()
		if err != nil {
			c.handleError(err)
		}
		c.disconnectHandler()
		c.logger.Debugf("ws. disconnected")

		// reconnect
		go c.connectAndListen()
	}()

	// Listen
	for {
		select {
		case <-c.stopChan:
			return
		default:
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				c.handleError(err)
				go c.stop()
				return
			}
			c.handleMsg(message)
		}
	}
}

func (c *wsClient) stop() {
	select {
	case c.stopChan <- struct{}{}:
	default:
	}
}

func (c *wsClient) heartbeats() {
	ticker := time.NewTicker(c.heartbeatInterval)
	defer ticker.Stop()
	for {
		if c.conn != nil {
			_ = c.WriteString(heartbeatMsg)
		}
		<-ticker.C
	}
}

func (c *wsClient) handleMsg(msg []byte) {
	str := string(msg)
	c.logger.Debugf("ws. received: %s", str)

	switch str {
	case heartbeatMsg:
		// do nothing
	default:
		c.handler(msg)
	}
}

func (c *wsClient) handleError(err error) {
	c.logger.Debugf("ws. error: %s", err)
	c.errHandler(err)
}
