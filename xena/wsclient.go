package xena

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"github.com/xenaex/client-go/xena/fixjson"
)

const (
	heartbeatMsg        = `{"35":"0"}`
	wsReconnectInterval = time.Second
	wsHeartbeatInterval = 3 * time.Second
	wsTimeoutInterval   = 15 * time.Second
)

type wsConf struct {
	reconnectInterval         time.Duration
	heartbeatInterval         time.Duration
	connectTimeoutInterval    time.Duration
	disconnectTimeoutInterval time.Duration
}

// NewWsClient websocket client constructor
func NewWsClient(opts ...WsOption) WsClient {
	c := wsClient{
		conf: wsConf{
			reconnectInterval:         wsReconnectInterval,
			heartbeatInterval:         wsHeartbeatInterval,
			connectTimeoutInterval:    wsTimeoutInterval,
			disconnectTimeoutInterval: 2 * wsHeartbeatInterval,
		},
		logger:   newLogger(),
		stopChan: make(chan struct{}),
	}
	c.initDefaultHandlers()

	// Apply options
	for _, optionFunc := range opts {
		optionFunc(&c)
	}
	go c.heartbeats()

	return &c
}

// WsClient websocket client interface
type WsClient interface {
	IsConnected() bool
	Connect() error
	With(opt WsOption)
	Close()
	Write(m interface{}) error
	WriteString(msg string) error
	WriteBytes(data []byte) error
	Logger() Logger
	getConf() wsConf
	SetDisconnectHandler(handler DisconnectHandler)
	setConnectInternalHandler(handler ConnectHandler)
}

type wsClient struct {
	// settings
	url            string
	logPingMessage bool
	conf           wsConf

	// handlers
	handler                Handler
	errHandler             ErrHandler
	connectHandler         ConnectHandler
	connectInternalHandler ConnectHandler
	disconnectHandler      DisconnectHandler

	logger *logger

	// no customizable elements
	conn     *websocket.Conn
	mu       sync.Mutex
	stopChan chan struct{}
	close    bool
}

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

// With applies modification of behavior
func (c *wsClient) With(opt WsOption) {
	opt(c)
}

func (c *wsClient) getConf() wsConf {
	return c.conf
}

func (c *wsClient) initDefaultHandlers() {
	c.handler = func(message []byte) {}
	c.errHandler = func(err error) {}
	c.connectHandler = func(client WsClient) {}
	c.connectInternalHandler = func(client WsClient) {}
	c.disconnectHandler = func() {}
}

// Connect start connecting
func (c *wsClient) Connect() error {
	err := c.connect()
	if err != nil {
		err = fmt.Errorf("failed to connect to %s: %s", c.url, err)
		return err
	}
	return nil
}

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

	if string(data) != heartbeatMsg || c.logPingMessage {
		c.logger.Debugf("ws. sent: %s", string(data))
	}

	return err
}

// Close connection permanent
func (c *wsClient) Close() {
	c.close = true
	go c.stop()
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

	return nil
}

func (c *wsClient) connectAndListen() error {
	// Connection loop
	for {
		if c.close {
			err := errors.New("connection was closed by user")
			c.handleError(err)
			return err
		}
		conn, _, err := websocket.DefaultDialer.Dial(c.url, nil)
		if err != nil {
			c.logger.Errorf("%s on websocket.DefaultDialer.Dial(%s)", err, c.url)
			c.handleError(err)
			time.Sleep(c.conf.reconnectInterval)
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
	c.connectInternalHandler(c)
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
		c.conn = nil
		c.disconnectHandler()
		c.logger.Debugf("ws. disconnected")
	}()

	c.mu.Lock()
	c.close = false
	c.mu.Unlock()

	mgs := make(chan []byte)
	go func() {
		m := mgs
		defer close(m)
		for {
			conn := c.conn
			if conn == nil {
				return
			}
			_, message, err := conn.ReadMessage()
			if err != nil {
				c.handleError(err)
				go c.stop()
				return
			}
			m <- message
		}
	}()

	// Listen
	for {
		select {
		case <-c.stopChan:
			return
		case message, ok := <-mgs:
			if !ok {
				mgs = nil
				continue
			}
			c.handleMsg(message)
		case <-time.NewTimer(c.conf.disconnectTimeoutInterval).C:
			return
		}
	}
}

func (c *wsClient) stop() {
	c.mu.Lock()
	defer c.mu.Unlock()
	select {
	case c.stopChan <- struct{}{}:
	default:
	}
}

func (c *wsClient) heartbeats() {
	ticker := time.NewTicker(c.conf.heartbeatInterval)
	defer ticker.Stop()
	for {
		c.mu.Lock()
		ok := !c.close
		conn := c.conn
		c.mu.Unlock()
		if ok && conn != nil {
			_ = c.WriteString(heartbeatMsg)
		}
		<-ticker.C
	}
}

func (c *wsClient) handleMsg(msg []byte) {
	str := string(msg)

	if str != heartbeatMsg || c.logPingMessage {
		c.logger.Debugf("ws. received: %s", str)
	}

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

func (c *wsClient) SetDisconnectHandler(handler DisconnectHandler) {
	c.disconnectHandler = handler
}

func (c *wsClient) setConnectInternalHandler(handler ConnectHandler) {
	c.connectInternalHandler = handler
}
