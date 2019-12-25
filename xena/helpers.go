package xena

import (
	"math/rand"
	"time"

	"github.com/xenaex/client-go/xena/xmsg"
)

const (
	maxIdLen    = 32
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
	src = rand.NewSource(time.Now().UnixNano())
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func ID(prefix string) string {
	if len(prefix) > maxIdLen {
		return prefix[:maxIdLen]
	}

	n := 32 - len(prefix)
	str := randString(n)
	str = prefix + str
	return str
}

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// CreateMarketIfTouchOrder create new market-if-touch order.
func CreateMarketIfTouchOrder(clOrdId string, symbol Symbol, side Side, orderQty string, account uint64, stopPx string) marketIfTouchOrder {
	return marketIfTouchOrder{
		order: baseOrder{
			NewOrderSingle: newOrderSingle(clOrdId, symbol, side, orderQty, account, xmsg.OrdType_MarketIfTouched, "0", stopPx),
		},
	}
}

// CreateMarketOrder create new market order.
func CreateMarketOrder(clOrdId string, symbol Symbol, side Side, orderQty string, account uint64) marketOrder {
	return marketOrder{
		order: baseOrder{
			NewOrderSingle: newOrderSingle(clOrdId, symbol, side, orderQty, account, xmsg.OrdType_Market, "", ""),
		},
	}
}

// CreateLimitOrder create new limit order.
func CreateLimitOrder(clOrdId string, symbol Symbol, side Side, orderQty string, account uint64, price string) limitOrder {
	return limitOrder{
		order: baseOrder{
			NewOrderSingle: newOrderSingle(clOrdId, symbol, side, orderQty, account, xmsg.OrdType_Limit, price, ""),
		},
	}
}

// CreateStopOrder create new stop order.
func CreateStopOrder(clOrdId string, symbol Symbol, side Side, orderQty string, account uint64, stopPx string) stopOrder {
	return stopOrder{
		order: baseOrder{
			NewOrderSingle: newOrderSingle(clOrdId, symbol, side, orderQty, account, xmsg.OrdType_Stop, "0", stopPx),
		},
	}
}
