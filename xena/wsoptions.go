package xena

import "time"

const (
	localWsMdURL = "ws://localhost/api/ws/market-data"
)

// WithURL set localhost... url for connection
func WithLocalURL() WsOption {
	return func(c *wsClient) {
		c.url = localWsMdURL
	}
}

// WithURL set URL for connection
func WithURL(url string) WsOption {
	return func(c *wsClient) {
		c.url = url
	}
}

// WithDebug enable debug logging
func WithDebug() WsOption {
	return func(c *wsClient) {
		c.Logger().SetDebug(true)
	}
}

// WithOutDebug disable debug logging
func WithOutDebug() WsOption {
	return func(c *wsClient) {
		c.Logger().SetDebug(false)
	}
}

// WithConnectHandler set custom connect handler
func WithConnectHandler(fnc ConnectHandler) WsOption {
	return func(c *wsClient) {
		if fnc == nil {
			fnc = func(client WsClient) {}
		}
		c.connectHandler = fnc
	}
}

// WithDisconnectHandler set custom disconnect handler
func WithDisconnectHandler(fnc DisconnectHandler) WsOption {
	return func(c *wsClient) {
		if fnc == nil {
			fnc = func() {}
		}
		c.disconnectHandler = fnc
	}
}

// WithHandler set custom common handler
func WithHandler(fnc Handler) WsOption {
	return func(c *wsClient) {
		if fnc == nil {
			fnc = func(msg []byte) {}
		}
		c.handler = fnc
	}
}

// WithErrHandler set custom common handler
func WithErrHandler(fnc ErrHandler) WsOption {
	return func(c *wsClient) {
		if fnc == nil {
			fnc = func(err error) {}
		}
		c.errHandler = fnc
	}
}

// WithReconnectInterval set time interval between reconnects
func WithReconnectInterval(d time.Duration) WsOption {
	return func(c *wsClient) {
		c.reconnectInterval = d
	}
}

// WithHeartbeatInterval set time interval between heartbeats
func WithHeartbeatInterval(d time.Duration) WsOption {
	return func(c *wsClient) {
		c.heartbeatInterval = d
	}
}
