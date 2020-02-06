package xena

//Symbol is an instrument.
type Symbol string
type Side string
type DOMThrottle string
type CandlesThrottle string
type TradesThrottle string
type AggregateBook int64
type PositionEffect string

//MarketDepth is level depth limit in dom.
type MarketDepth int64

func (s Symbol) String() string {
	return string(s)
}

const (
	BTCUSDT Symbol = "BTC/USDT"
	ETHSDT  Symbol = "ETH/USDT"
	XBTUSD  Symbol = "XBTUSD"
	ETHUSD  Symbol = "ETHUSD"

	SideBuy  Side = "1"
	SideSell Side = "2"

	ThrottleDOMDefault DOMThrottle = ""
	ThrottleDOM500ms   DOMThrottle = "500ms"
	ThrottleDOM5s      DOMThrottle = "5s"

	ThrottleCandlesDefault CandlesThrottle = ""
	ThrottleCandles250ms   CandlesThrottle = "250ms"
	ThrottleCandles1s      CandlesThrottle = "1s"

	ThrottleTradesDefault TradesThrottle = ""
	ThrottleTrades500ms   TradesThrottle = "500ms"
	ThrottleTrades5s      TradesThrottle = "5s"

	CancelOrdersForASecurity = "1"
	CancelAllOrders          = "7"

	PosTransTypeCollapse  = "20"
	PosMaintActionReplace = "2"

	PositionEffectClose   PositionEffect = "C"
	PositionEffectOpen    PositionEffect = "O"
	PositionEffectDefault PositionEffect = "D"

	AggregateBookDefault AggregateBook = 0
	AggregateBook5       AggregateBook = 5
	AggregateBook10      AggregateBook = 10
	AggregateBook25      AggregateBook = 25
	AggregateBook50      AggregateBook = 50
	AggregateBook100     AggregateBook = 100
	AggregateBook250     AggregateBook = 250

	MarketDepthDefault MarketDepth = 0
	MarketDepth10      MarketDepth = 10
	MarketDepth20      MarketDepth = 20
)
