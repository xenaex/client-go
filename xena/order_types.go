package xena

import (
	"time"

	"github.com/xenaex/client-go/xena/xmsg"
)

func newOrderSingle(clOrdId string, symbol string, side Side, orderQty string, account uint64, ordType string, price string, stopPx string) *xmsg.NewOrderSingle {
	cmd := &xmsg.NewOrderSingle{
		MsgType:      xmsg.MsgType_NewOrderSingleMsgType,
		ClOrdId:      clOrdId,
		Symbol:       symbol,
		Side:         string(side),
		OrderQty:     orderQty,
		Price:        price,
		Account:      account,
		OrdType:      ordType,
		StopPx:       stopPx,
		TransactTime: time.Now().UnixNano(),
	}

	return cmd
}

type baseOrder struct {
	client TradingClient
	*xmsg.NewOrderSingle
}

func (b baseOrder) setTimeInForce(timeInForce string) {
	b.NewOrderSingle.TimeInForce = timeInForce
}

func (b baseOrder) addExecInst(execInst ...string) {
	b.NewOrderSingle.ExecInst = append(b.NewOrderSingle.ExecInst, execInst...)
}

func (b baseOrder) setPositionId(positionId uint64) {
	b.NewOrderSingle.PositionId = positionId
}

func (b baseOrder) setStopLossPrice(stopLossPrice string) {
	if len(stopLossPrice) > 0 {
		b.NewOrderSingle.SLTP = append(b.NewOrderSingle.SLTP, &xmsg.SLTP{
			OrdType: xmsg.OrdType_Stop,
			StopPx:  stopLossPrice,
		})
	}
}

func (b baseOrder) setTakeProfitPrice(takeProfitPrice string) {
	if len(takeProfitPrice) > 0 {
		b.NewOrderSingle.SLTP = append(b.NewOrderSingle.SLTP, &xmsg.SLTP{
			OrdType: xmsg.OrdType_Limit,
			StopPx:  takeProfitPrice,
		})
	}
}

func (b baseOrder) setText(text string) {
	b.NewOrderSingle.Text = text
}

func (b baseOrder) setTrailingOffset(trailingOffset string, capPrice string) {
	if len(trailingOffset) > 0 {
		b.NewOrderSingle.SLTP = append(b.NewOrderSingle.SLTP, &xmsg.SLTP{
			OrdType:        xmsg.OrdType_Stop,
			CapPrice:       capPrice,
			PegPriceType:   xmsg.PegPriceType_TrailingStopPeg,
			PegOffsetType:  xmsg.PegOffsetType_BasisPoints,
			PegOffsetValue: trailingOffset,
		})
	}
}

func (b baseOrder) build() *xmsg.NewOrderSingle {
	return b.NewOrderSingle
}

func (b baseOrder) send(client TradingClient) error {
	return client.Send(b.NewOrderSingle)
}

type marketOrder struct {
	order baseOrder
}

func (m marketOrder) SetTimeInForce(timeInForce string) marketOrder {
	m.order.setTimeInForce(timeInForce)
	return m
}

func (m marketOrder) AddExecInst(execInst ...string) marketOrder {
	m.order.addExecInst(execInst...)
	return m
}

func (m marketOrder) SetPositionId(positionId uint64) marketOrder {
	m.order.setPositionId(positionId)
	return m
}

func (m marketOrder) SetStopLossPrice(stopLossPrice string) marketOrder {
	m.order.setStopLossPrice(stopLossPrice)
	return m
}

func (m marketOrder) SetTakeProfitPrice(takeProfitPrice string) marketOrder {
	m.order.setTakeProfitPrice(takeProfitPrice)
	return m
}

func (m marketOrder) SetText(text string) marketOrder {
	m.order.setText(text)
	return m
}

func (m marketOrder) Build() *xmsg.NewOrderSingle {
	return m.order.build()
}

func (m marketOrder) Send(client TradingClient) error {
	return m.order.send(client)
}

type limitOrder struct {
	order baseOrder
}

func (l limitOrder) SetTimeInForce(timeInForce string) limitOrder {
	l.order.setTimeInForce(timeInForce)
	return l
}

func (l limitOrder) AddExecInst(execInst ...string) limitOrder {
	l.order.addExecInst(execInst...)
	return l
}

func (l limitOrder) SetPositionId(positionId uint64) limitOrder {
	l.order.setPositionId(positionId)
	return l
}

func (l limitOrder) SetStopLossPrice(stopLossPrice string) limitOrder {
	l.order.setStopLossPrice(stopLossPrice)
	return l
}

func (l limitOrder) SetTakeProfitPrice(takeProfitPrice string) limitOrder {
	l.order.setTakeProfitPrice(takeProfitPrice)
	return l
}

func (l limitOrder) SetText(text string) limitOrder {
	l.order.setText(text)
	return l
}

func (l limitOrder) SetTrailingOffset(trailingOffset string) limitOrder {
	l.order.setTrailingOffset(trailingOffset, "")
	return l
}

func (l limitOrder) SetTrailingOffsetAndCapPrice(trailingOffset string, capPrice string) limitOrder {
	l.order.setTrailingOffset(trailingOffset, capPrice)
	return l
}

func (l limitOrder) Build() *xmsg.NewOrderSingle {
	return l.order.build()
}

func (l limitOrder) Send(client TradingClient) error {
	return l.order.send(client)
}

type stopOrder struct {
	order baseOrder
}

func (s stopOrder) SetTimeInForce(timeInForce string) stopOrder {
	s.order.setTimeInForce(timeInForce)
	return s
}

func (s stopOrder) AddExecInst(execInst ...string) stopOrder {
	s.order.addExecInst(execInst...)
	return s
}

func (s stopOrder) SetPositionId(positionId uint64) stopOrder {
	s.order.setPositionId(positionId)
	return s
}

func (s stopOrder) SetStopLossPrice(stopLossPrice string) stopOrder {
	s.order.setStopLossPrice(stopLossPrice)
	return s
}

func (s stopOrder) SetTakeProfitPrice(takeProfitPrice string) stopOrder {
	s.order.setTakeProfitPrice(takeProfitPrice)
	return s
}

func (s stopOrder) SetText(text string) stopOrder {
	s.order.setText(text)
	return s
}

func (s stopOrder) SetTrailingOffset(trailingOffset string) stopOrder {
	s.order.setTrailingOffset(trailingOffset, "")
	return s
}

func (s stopOrder) SetTrailingOffsetAndCapPrice(trailingOffset string, capPrice string) stopOrder {
	s.order.setTrailingOffset(trailingOffset, capPrice)
	return s
}

func (s stopOrder) Build() *xmsg.NewOrderSingle {
	return s.order.build()
}

func (s stopOrder) Send(client TradingClient) error {
	return s.order.send(client)
}

type marketIfTouchOrder struct {
	client TradingClient
	order  baseOrder
}

func (m marketIfTouchOrder) SetTimeInForce(timeInForce string) marketIfTouchOrder {
	m.order.setTimeInForce(timeInForce)
	return m
}

func (m marketIfTouchOrder) AddExecInst(execInst ...string) marketIfTouchOrder {
	m.order.addExecInst(execInst...)
	return m
}

func (m marketIfTouchOrder) SetPositionId(positionId uint64) marketIfTouchOrder {
	m.order.setPositionId(positionId)
	return m
}

func (m marketIfTouchOrder) SetStopLossPrice(stopLossPrice string) marketIfTouchOrder {
	m.order.setStopLossPrice(stopLossPrice)
	return m
}

func (m marketIfTouchOrder) SetTakeProfitPrice(takeProfitPrice string) marketIfTouchOrder {
	m.order.setTakeProfitPrice(takeProfitPrice)
	return m
}

func (m marketIfTouchOrder) SetText(text string) marketIfTouchOrder {
	m.order.setText(text)
	return m
}

func (m marketIfTouchOrder) SetTrailingOffset(trailingOffset string) marketIfTouchOrder {
	m.order.setTrailingOffset(trailingOffset, "")
	return m
}

func (m marketIfTouchOrder) SetTrailingOffsetAndCapPrice(trailingOffset string, capPrice string) marketIfTouchOrder {
	m.order.setTrailingOffset(trailingOffset, capPrice)
	return m
}

func (m marketIfTouchOrder) Build() *xmsg.NewOrderSingle {
	return m.order.build()
}

func (m marketIfTouchOrder) Send(client TradingClient) error {
	return m.order.send(client)
}

type orderMassCancel struct {
	*xmsg.OrderMassCancelRequest
}

func newOrderMassCancel(account uint64, clOrdId string) orderMassCancel {
	return orderMassCancel{
		OrderMassCancelRequest: &xmsg.OrderMassCancelRequest{
			MsgType:               xmsg.MsgType_OrderMassCancelRequest,
			Account:               account,
			ClOrdId:               clOrdId,
			MassCancelRequestType: CancelAllOrders,
		},
	}
}

func (m orderMassCancel) SetSide(side Side) orderMassCancel {
	m.Side = string(side)
	return m
}

func (m orderMassCancel) SetSymbol(symbol string) orderMassCancel {
	m.Symbol = symbol
	if len(symbol) == 0 {
		m.MassCancelRequestType = CancelAllOrders
	} else {
		m.MassCancelRequestType = CancelOrdersForASecurity
	}
	return m
}

func (m orderMassCancel) SetPositionEffect(positionEffect PositionEffect) orderMassCancel {
	m.PositionEffect = string(positionEffect)
	return m
}

func (m orderMassCancel) Build() *xmsg.OrderMassCancelRequest {
	return m.OrderMassCancelRequest
}

func (m orderMassCancel) Send(client TradingClient) error {
	return client.Send(m.Build())
}
