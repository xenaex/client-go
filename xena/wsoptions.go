package xena

import "time"

const (
	localWsMdURL     = "ws://localhost/api/ws/market-data"
	prodWsMdURL      = "wss://api.xena.exchange/ws/market-data"
	prodWsTradingURL = "wss://api.xena.exchange/ws/trading"
)

// WithLocalURL set localhost... url for connection.
func WithLocalURL() WsOption {
	return func(c *wsClient) {
		c.url = localWsMdURL
	}
}

// WithTradingURL set trading url for connection.
func WithTradingURL() WsOption {
	return func(c *wsClient) {
		c.url = prodWsTradingURL
	}
}

// WithMarketDataURL set market data url for connection.
func WithMarketDataURL() WsOption {
	return func(c *wsClient) {
		c.url = prodWsMdURL
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
		c.logger.debug = true
	}
}

// WithIgnorePingLog enable ignore ping logs.
func WithIgnorePingLog(ignore bool) WsOption {
	return func(c *wsClient) {
		c.logPingMessage = ignore
	}
}

// WithOutDebug disable debug logging
func WithOutDebug() WsOption {
	return func(c *wsClient) {
		c.logger.debug = false
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

// WithConnectHandler set custom connect handler
func withConnectInternalHandler(fnc ConnectHandler) WsOption {
	return func(c *wsClient) {
		if fnc == nil {
			fnc = func(client WsClient) {}
		}
		c.connectInternalHandler = fnc
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

//WithErrHandler set custom common handler
func WithErrHandler(fnc ErrHandler) WsOption {
	return func(c *wsClient) {
		if fnc == nil {
			fnc = func(err error) {}
		}
		c.errHandler = fnc
	}
}

//WithReconnectInterval set time interval between reconnects
func WithReconnectInterval(d time.Duration) WsOption {
	return func(c *wsClient) {
		c.conf.reconnectInterval = d
	}
}

//WithHeartbeatInterval set time interval between heartbeats
func WithHeartbeatInterval(d time.Duration) WsOption {
	return func(c *wsClient) {
		c.conf.heartbeatInterval = d
	}
}

//WithLogger set custom logger.
func WithLogger(logger Logger) WsOption {
	return func(c *wsClient) {
		c.logger.logger = logger
	}
}

//WithConnectTimeoutConnect set connect timeout interval.
func WithConnectTimeoutConnect(interval time.Duration) WsOption {
	return func(c *wsClient) {
		c.conf.connectTimeoutInterval = interval
	}
}

//WithDisconnectTimeoutInterval set disconnect timeout interval.
func WithDisconnectTimeoutInterval(interval time.Duration) WsOption {
	return func(c *wsClient) {
		c.conf.disconnectTimeoutInterval = interval
	}
}
