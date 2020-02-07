package xena

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"github.com/xenaex/client-go/xena/fixjson"
)

const (
	heartbeatMsg        = `{"35":"0"}`
	wsHeartbeatInterval = 3 * time.Second
	wsTimeoutInterval   = 15 * time.Second
)

type wsConf struct {
	heartbeatInterval         time.Duration
	connectTimeoutInterval    time.Duration
	disconnectTimeoutInterval time.Duration
}

// NewWsClient creates instance of WsClient.
func NewWsClient(opts ...WsOption) WsClient {
	c := wsClient{
		conf: wsConf{
			heartbeatInterval:         wsHeartbeatInterval,
			connectTimeoutInterval:    wsTimeoutInterval,
			disconnectTimeoutInterval: 2 * wsHeartbeatInterval,
		},
		logger:   newLogger(false),
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

// WsClient is base websocket client interface of Xena.
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
	setDisconnectHandler(handler DisconnectHandler)
	setConnectInternalHandler(handler ConnectHandler)
}

type wsClient struct {
	// settings
	url            string
	disablePingLog bool
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

// Handler function is raw message handler.
type Handler func(msg []byte)

// ErrHandler function is a error handler.
type ErrHandler func(err error)

// ConnectHandler function is a connection handler.
type ConnectHandler func(client WsClient)

// DisconnectHandler function is a disconnect handler.
type DisconnectHandler func()

// WsOption is function that alter behavior.
type WsOption func(s *wsClient)

// With applies modification of behavior.
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

// Connect connects to the server
func (c *wsClient) Connect() error {
	err := c.connect()
	if err != nil {
		err = fmt.Errorf("failed to connect to %s: %s", c.url, err)
		return err
	}
	return nil
}

// IsConnected returns true if the connection is established.
func (c *wsClient) IsConnected() bool {
	return c.conn != nil
}

// Logger returns logger.
func (c *wsClient) Logger() Logger {
	return c.logger
}

// Write marshals and sends a message to the socket.
func (c *wsClient) Write(v interface{}) error {
	data, err := fixjson.Marshal(v)
	if err != nil {
		return err
	}
	return c.WriteBytes(data)
}

// WriteString sends a message to the socket.
func (c *wsClient) WriteString(msg string) error {
	return c.WriteBytes([]byte(msg))
}

// WriteBytes sends a message to socket.
func (c *wsClient) WriteBytes(data []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn == nil {
		return errors.New("ws connection not established yet")
	}
	err := c.conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		c.handleError(err)
		go c.stop()
	}

	if string(data) != heartbeatMsg || !c.disablePingLog {
		c.logger.Debugf("ws. sent: %s", string(data))
	}

	return err
}

// Close connection permanent
func (c *wsClient) Close() {
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
		c.logger.Errorf("%s on c.connectAndListen()")
		return err
	}

	return nil
}

func (c *wsClient) connectAndListen() error {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(c.conf.connectTimeoutInterval))
	conn, _, err := websocket.DefaultDialer.DialContext(ctx, c.url, nil)
	if err != nil {
		c.logger.Errorf("%s on websocket.DefaultDialer.Dial(%s)", err, c.url)
		c.handleError(err)
		c.disconnectHandler()
		return err
	}
	// OK
	c.mu.Lock()
	c.conn = conn
	c.stopChan = make(chan struct{})
	c.mu.Unlock()

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

	msg := make(chan []byte)
	go func() {
		m := msg
		defer close(m)
		for {
			conn := c.conn
			if conn == nil {
				return
			}
			_, message, err := conn.ReadMessage()
			if err != nil {
				c.logger.Errorf("%s on conn.ReadMessage()", err)
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
		case message, ok := <-msg:
			if !ok {
				msg = nil
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
		c.logger.Debugf("wsClient.stop is called")
	default:
	}
}

func (c *wsClient) heartbeats() {
	ticker := time.NewTicker(c.conf.heartbeatInterval)
	defer ticker.Stop()
	for {
		c.mu.Lock()
		conn := c.conn
		c.mu.Unlock()
		if conn != nil {
			_ = c.WriteString(heartbeatMsg)
		}
		<-ticker.C
	}
}

func (c *wsClient) handleMsg(msg []byte) {
	str := string(msg)

	if str != heartbeatMsg || !c.disablePingLog {
		c.logger.Debugf("ws. received: %s", str)
	}
	c.handler(msg)
}

func (c *wsClient) handleError(err error) {
	c.logger.Errorf("ws. error: %s", err)
	c.errHandler(err)
}

// setDisconnectHandler subscribes to disconnect event.
func (c *wsClient) setDisconnectHandler(handler DisconnectHandler) {
	c.disconnectHandler = handler
}

// setConnectInternalHandler subscribes to connection event.
func (c *wsClient) setConnectInternalHandler(handler ConnectHandler) {
	c.connectInternalHandler = handler
}
