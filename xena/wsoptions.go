package xena

import (
	"time"
)

const (
	localWsMdURL     = "ws://localhost/api/ws/market-data"
	prodWsMdURL      = "wss://api.xena.exchange/ws/market-data"
	prodWsTradingURL = "wss://api.xena.exchange/ws/trading"
)

//WithTradingURL sets the url for connection to the local market data.
func WithLocalURL() WsOption {
	return func(c *wsClient) {
		c.url = localWsMdURL
	}
}

//WithTradingURL sets the url for connection to the trading api.
func WithTradingURL() WsOption {
	return func(c *wsClient) {
		c.url = prodWsTradingURL
	}
}

//WithLocalURL sets the url for connection to the market data.
func WithMarketDataURL() WsOption {
	return func(c *wsClient) {
		c.url = prodWsMdURL
	}
}

//WithURL sets the url for connection.
func WithURL(url string) WsOption {
	return func(c *wsClient) {
		c.url = url
	}
}

//WithDebug enables debug logging.
func WithDebug() WsOption {
	return func(c *wsClient) {
		c.logger.debug = true
	}
}

//WithIgnorePingLog disables ping logs.
func WithIgnorePingLog(disable bool) WsOption {
	return func(c *wsClient) {
		c.disablePingLog = disable
	}
}

//WithDebug disables debug logging.
func WithOutDebug() WsOption {
	return func(c *wsClient) {
		c.logger.debug = false
	}
}

//WithConnectHandler sets custom connection handler.
func WithConnectHandler(fnc ConnectHandler) WsOption {
	return func(c *wsClient) {
		if fnc == nil {
			fnc = func(client WsClient) {}
		}
		c.connectHandler = fnc
	}
}

//WithDisconnectHandler sets custom disconnection handler.
func WithDisconnectHandler(fnc DisconnectHandler) WsOption {
	return func(c *wsClient) {
		if fnc == nil {
			fnc = func() {}
		}
		c.disconnectHandler = fnc
	}
}

//WithHandler sets custom common message handler.
func WithHandler(fnc Handler) WsOption {
	return func(c *wsClient) {
		if fnc == nil {
			fnc = func(msg []byte) {}
		}
		c.handler = fnc
	}
}

//WithErrHandler sets custom common error handler.
func WithErrHandler(fnc ErrHandler) WsOption {
	return func(c *wsClient) {
		if fnc == nil {
			fnc = func(err error) {}
		}
		c.errHandler = fnc
	}
}

//WithHeartbeatInterval sets time interval between heartbeats.
func WithHeartbeatInterval(d time.Duration) WsOption {
	return func(c *wsClient) {
		c.conf.heartbeatInterval = d
	}
}

//WithLogger sets custom logger.
func WithLogger(logger Logger) WsOption {
	return func(c *wsClient) {
		c.logger.logger = logger
	}
}

//WithConnectTimeoutConnect sets timeout interval connection.
func WithConnectTimeoutConnect(interval time.Duration) WsOption {
	return func(c *wsClient) {
		c.conf.connectTimeoutInterval = interval
	}
}

//WithConnectTimeoutConnect sets timeout interval disconnect.
func WithDisconnectTimeoutInterval(interval time.Duration) WsOption {
	return func(c *wsClient) {
		c.conf.disconnectTimeoutInterval = interval
	}
}
