package xena

// Symbol is an instrument
type Symbol string
type Side string
type Throttle string
type AggregateBook int64
type PositionEffect string

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

	ThrottleDOMDefault Throttle = ""
	ThrottleDOM500ms   Throttle = "500ms"
	ThrottleDOM5s      Throttle = "5s"

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
)
