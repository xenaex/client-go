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

//IsMargin returns true if accountId is margin account.
func IsMargin(accountId uint64) bool {
	if accountId > 1000000000 {
		return true
	}
	return false
}

//ID function generates new random id.
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

//CreateMarketIfTouchOrder creates new builder of market-if-touch order.
func CreateMarketIfTouchOrder(clOrdId string, symbol string, side Side, orderQty string, account uint64, stopPx string) marketIfTouchOrder {
	return marketIfTouchOrder{
		order: baseOrder{
			NewOrderSingle: newOrderSingle(clOrdId, symbol, side, orderQty, account, xmsg.OrdType_MarketIfTouched, "", stopPx),
		},
	}
}

//CreateMarketOrder creates new builder of market order.
func CreateMarketOrder(clOrdId string, symbol string, side Side, orderQty string, account uint64) marketOrder {
	return marketOrder{
		order: baseOrder{
			NewOrderSingle: newOrderSingle(clOrdId, symbol, side, orderQty, account, xmsg.OrdType_Market, "", ""),
		},
	}
}

//CreateLimitOrder creates new builder of limit order.
func CreateLimitOrder(clOrdId string, symbol string, side Side, orderQty string, account uint64, price string) limitOrder {
	return limitOrder{
		order: baseOrder{
			NewOrderSingle: newOrderSingle(clOrdId, symbol, side, orderQty, account, xmsg.OrdType_Limit, price, ""),
		},
	}
}

//CreateStopOrder creates new builder of stop order.
func CreateStopOrder(clOrdId string, symbol string, side Side, orderQty string, account uint64, stopPx string) stopOrder {
	return stopOrder{
		order: baseOrder{
			NewOrderSingle: newOrderSingle(clOrdId, symbol, side, orderQty, account, xmsg.OrdType_Stop, "", stopPx),
		},
	}
}

//CreateOrderMassCancel creates new builder of mass cancel order.
func CreateOrderMassCancel(account uint64, clOrdId string) orderMassCancel {
	return newOrderMassCancel(account, clOrdId)
}

//CreateReplace creates new cancel request and replace request of order, based on execution report.
func CreateReplace(replaceId string, executionReport *xmsg.ExecutionReport) xmsg.OrderCancelReplaceRequest {
	cmd := xmsg.OrderCancelReplaceRequest{}
	cmd.MsgType = xmsg.MsgType_OrderCancelReplaceRequestMsgType
	cmd.ClOrdId = replaceId
	cmd.OrigClOrdId = executionReport.ClOrdId
	cmd.Symbol = executionReport.Symbol
	cmd.Side = executionReport.Side
	cmd.TransactTime = time.Now().UnixNano()
	cmd.Account = executionReport.Account
	cmd.Price = executionReport.Price
	cmd.StopPx = executionReport.StopPx
	cmd.CapPrice = executionReport.CapPrice
	cmd.OrderQty = executionReport.OrderQty
	cmd.PegPriceType = executionReport.PegPriceType
	cmd.PegOffsetType = executionReport.PegOffsetType
	cmd.PegOffsetValue = executionReport.PegOffsetValue

	for _, s := range executionReport.SLTP {
		sltp := &xmsg.SLTP{}
		sltp.OrdType = s.OrdType
		sltp.Price = s.Price
		sltp.StopPx = s.StopPx
		sltp.CapPrice = s.CapPrice
		sltp.PegPriceType = s.PegPriceType
		sltp.PegOffsetType = s.PegOffsetType
		sltp.PegOffsetValue = s.PegOffsetValue
		cmd.SLTP = append(cmd.SLTP, sltp)
	}

	return cmd
}

//CreateCancelRequestFromExecutionReport creates new cancel request of order, based on execution report.
func CreateCancelRequestFromExecutionReport(cancelID string, executionReport *xmsg.ExecutionReport) *xmsg.OrderCancelRequest {
	cmd := &xmsg.OrderCancelRequest{}
	cmd.MsgType = xmsg.MsgType_OrderCancelRequestMsgType
	cmd.ClOrdId = cancelID
	cmd.OrigClOrdId = executionReport.ClOrdId
	cmd.Symbol = executionReport.Symbol
	cmd.Side = executionReport.Side
	cmd.TransactTime = time.Now().UnixNano()
	cmd.Account = executionReport.Account
	return cmd
}
