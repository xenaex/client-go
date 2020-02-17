package xena

import (
	"fmt"
	"sync"
	"time"

	"github.com/xenaex/client-go/xena/fixjson"
	"github.com/xenaex/client-go/xena/xmsg"
)

const (
	wsTradingURL = "ws://localhost/api/ws/trading"
)

// TradingDisconnectHandler function is a disconnect handler.
type TradingDisconnectHandler func(client TradingClient, logger Logger)

// TradingClient is websocket client interface of the Xena trading api.
type TradingClient interface {
	// ListenLogon subscribes to Logon event.
	ListenLogon(handler LogonHandler)
	// ListenMarginRequirementReport subscribes to MarginRequirementReport event.
	ListenMarginRequirementReport(handler MarginRequirementReportHandler)
	// ListenExecutionReport subscribes to ExecutionReport event.
	ListenExecutionReport(handler ExecutionReportHandler)
	// ListenOrderCancelReject subscribes to OrderCancelReject event.
	ListenOrderCancelReject(handler OrderCancelRejectHandler)
	// ListenOrderMassStatusResponse subscribes to OrderMassStatusResponse event.
	ListenOrderMassStatusResponse(handler OrderMassStatusResponseHandler)
	// ListenBalanceIncrementalRefresh subscribes to BalanceIncrementalRefresh event.
	ListenBalanceIncrementalRefresh(handler BalanceIncrementalRefreshHandler)
	// ListenBalanceSnapshotRefresh subscribes to BalanceSnapshotRefresh event.
	ListenBalanceSnapshotRefresh(handler BalanceSnapshotRefreshHandler)
	// ListenPositionReport subscribes to PositionReport event.
	ListenPositionReport(handler PositionReportHandler)
	// ListenMassPositionReport subscribes to MassPositionReport event.
	ListenMassPositionReport(handler MassPositionReportHandler)
	// ListenPositionMaintenanceReport subscribes to PositionMaintenanceReport event.
	ListenPositionMaintenanceReport(handler PositionMaintenanceReportHandler)
	// ListenReject subscribes to Reject event.
	ListenReject(handler RejectHandler)
	// ListenListStatus subscribes to ListStatus event.
	ListenListStatus(handler ListStatusHandler)
	// ListenOrderMassCancelReport subscribes to OrderMassCancelReport event.
	ListenOrderMassCancelReport(handler OrderMassCancelReportHandler)
	// ListenHeartbeat subscribes to Heartbeat event.
	ListenHeartbeat(handler HeartbeatHandler)

	// ConnectAndLogon connects to websocket and sends Logon message.
	// Logon.RejectText contains reject reason.
	ConnectAndLogon() (*xmsg.Logon, error)

	// Send function sends command to server.
	Send(cmd interface{}) error

	// AccountStatusReport requests an account status report.
	// To receive a response, the client should listen for ListenBalanceSnapshotRefresh.
	AccountStatusReport(account uint64, requestId string) error

	// GetPositions requests all positions.
	// To receive a response, the client should listen for ListenMassPositionReport.
	GetPositions(account uint64, requestId string) error

	// Orders requests all orders and all fills.
	// To receive a response, the client should listen for ListenOrderMassStatusResponse.
	Orders(account uint64, requestId string) error

	// MarketOrder places new market order.
	MarketOrder(account uint64, clOrdId string, symbol string, side Side, orderQty string) error

	// LimitOrder places new limit order.
	LimitOrder(account uint64, clOrdId string, symbol string, side Side, price string, orderQty string) error

	// StopOrder places new stop order.
	StopOrder(account uint64, clOrdId string, symbol string, side Side, stopPx string, orderQty string) error

	// MarketIfTouchOrder places new market-if-touch order.
	MarketIfTouchOrder(account uint64, clOrdId string, symbol string, side Side, stopPx string, orderQty string) error

	// Cancel function cancels an existing order.
	Cancel(cmd *xmsg.OrderCancelRequest) error

	// CancelByClOrdId cancels an existing order by original client order id.
	CancelByClOrdId(account uint64, clOrdId, origClOrdId, symbol string, side Side) error

	// CancelByOrderId cancels an existing order by order id.
	CancelByOrderId(account uint64, clOrdId, orderId, symbol string, side Side) error

	// MassCancel cancels all existing orders.
	// To receive a response, the client should listen for ListenOrderMassCancelReport.
	MassCancel(account uint64, clOrdId string) error

	// Replace cancels existing order and replaces by new order.
	Replace(request xmsg.OrderCancelReplaceRequest) error

	// CollapsePositions collapses all existing positions for margin account and symbol.
	CollapsePositions(account uint64, symbol string, requestId string) error

	// SetDisconnectHandler subscribes to disconnect events.
	SetDisconnectHandler(handler TradingDisconnectHandler)

	// SendApplicationHeartbeat sends application heartbeat.
	SendApplicationHeartbeat(groupId string, heartbeatInSec int32) error
}

// DefaultTradingDisconnectHandler is a default reconnects handler.
func DefaultTradingDisconnectHandler(client TradingClient, logger Logger) {
	for {
		time.Sleep(time.Second)
		logonResponse, err := client.ConnectAndLogon()
		if err != nil {
			logger.Errorf("%s on client.ConnectAndLogon()\n", err)
		}
		if err == nil {
			logger.Debugf("GOT logonResponse ", logonResponse)
			break
		}
	}
}

// NewTradingClient creates websocket client of Xena trading.
func NewTradingClient(apiKey, apiSecret string, opts ...WsOption) TradingClient {
	t := &tradingClient{
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}

	defaultOpts := []WsOption{
		WithTradingURL(),
		WithHandler(t.incomeHandler),
		WithIgnorePingLog(true),
	}
	opts = append(defaultOpts, opts...)

	t.client = NewWsClient(opts...)

	return t
}

func (t *tradingClient) ConnectAndLogon() (*xmsg.Logon, error) {
	t.mutexLogon.Lock()
	defer t.mutexLogon.Unlock()
	logMgs := make(chan *xmsg.Logon, 1)
	t.handlers.logonInternal = func(t TradingClient, m *xmsg.Logon) {
		logMgs <- m
		close(logMgs)
	}
	defer func() { t.handlers.logonInternal = nil }()

	var loginErr error
	wg := &sync.WaitGroup{}
	wg.Add(1)
	t.client.setConnectInternalHandler(func(client WsClient) {
		defer wg.Done()
		loginCmd := t.loginCmd()
		loginErr = t.client.WriteBytes(loginCmd)
		if loginErr != nil {
			loginErr = fmt.Errorf("%s on t.client.WriteBytes()", loginErr)
			t.client.Logger().Errorf("%s", loginErr)
		}
	})

	defer func() { t.client.setConnectInternalHandler(nil) }()
	err := t.client.Connect()
	if err != nil {
		return nil, err
	}
	wg.Wait()
	if loginErr != nil {
		return nil, loginErr
	}
	select {
	case m, ok := <-logMgs:
		if ok && m != nil {
			return m, nil
		}
	case <-time.NewTimer(t.client.getConf().connectTimeoutInterval).C:
		t.client.Close()
		return nil, fmt.Errorf("timeout logon")
	}
	t.client.Close()
	return nil, fmt.Errorf("login response didn't come")
}

// LogonHandler is used to handling logon message.
type LogonHandler func(t TradingClient, m *xmsg.Logon)

// MarginRequirementReportHandler is used to handling MarginRequirementReport message.
type MarginRequirementReportHandler func(t TradingClient, m *xmsg.MarginRequirementReport)

// BalanceIncrementalRefreshHandler is used to handling BalanceIncrementalRefresh message.
type BalanceIncrementalRefreshHandler func(t TradingClient, m *xmsg.BalanceIncrementalRefresh)

// BalanceSnapshotRefreshHandler is used to handling BalanceSnapshotRefresh message.
type BalanceSnapshotRefreshHandler func(t TradingClient, m *xmsg.BalanceSnapshotRefresh)

// ExecutionReportHandler is used to handling ExecutionReport message.
type ExecutionReportHandler func(t TradingClient, m *xmsg.ExecutionReport)

// OrderCancelRejectHandler is used to handling OrderCancelReject message.
type OrderCancelRejectHandler func(t TradingClient, m *xmsg.OrderCancelReject)

// OrderMassStatusResponseHandler is used to handling OrderMassStatusResponse message.
type OrderMassStatusResponseHandler func(t TradingClient, m *xmsg.OrderMassStatusResponse)

// PositionReportHandler is used to handling PositionReport message.
type PositionReportHandler func(t TradingClient, m *xmsg.PositionReport)

// MassPositionReportHandler is used to handling MassPositionReport message.
type MassPositionReportHandler func(t TradingClient, m *xmsg.MassPositionReport)

// PositionMaintenanceReportHandler is used to handling PositionMaintenanceReport message.
type PositionMaintenanceReportHandler func(t TradingClient, m *xmsg.PositionMaintenanceReport)

// RejectHandler is used to handling Reject message.
type RejectHandler func(t TradingClient, m *xmsg.Reject)

// ListStatusHandler is used to handling ListStatus message.
type ListStatusHandler func(t TradingClient, m *xmsg.ListStatus)

// OrderMassCancelReportHandler is used to handling OrderMassCancelReport message.
type OrderMassCancelReportHandler func(t TradingClient, m *xmsg.OrderMassCancelReport)

// HeartbeatHandler is used to handling Heartbeat message.
type HeartbeatHandler func(t TradingClient, m *xmsg.Heartbeat)

type tradingClient struct {
	client    WsClient
	apiKey    string
	apiSecret string
	handlers  struct {
		heartbeat                 HeartbeatHandler
		balanceSnapshotRefresh    BalanceSnapshotRefreshHandler
		balanceIncrementalRefresh BalanceIncrementalRefreshHandler
		executionReport           ExecutionReportHandler
		listStatus                ListStatusHandler
		logon                     LogonHandler
		logonInternal             LogonHandler
		marginRequirementReport   MarginRequirementReportHandler
		massPositionReport        MassPositionReportHandler
		orderCancelReject         OrderCancelRejectHandler
		orderMassCancelReport     OrderMassCancelReportHandler
		orderMassStatus           OrderMassStatusResponseHandler
		positionMaintenanceReport PositionMaintenanceReportHandler
		positionReport            PositionReportHandler
		reject                    RejectHandler
	}
	mutexLogon sync.Mutex
}

func (t *tradingClient) ListenLogon(handler LogonHandler) {
	t.handlers.logon = handler
}

func (t *tradingClient) ListenMarginRequirementReport(handler MarginRequirementReportHandler) {
	t.handlers.marginRequirementReport = handler
}

func (t *tradingClient) ListenExecutionReport(handler ExecutionReportHandler) {
	t.handlers.executionReport = handler
}

func (t *tradingClient) ListenOrderCancelReject(handler OrderCancelRejectHandler) {
	t.handlers.orderCancelReject = handler
}

func (t *tradingClient) ListenOrderMassStatusResponse(handler OrderMassStatusResponseHandler) {
	t.handlers.orderMassStatus = handler
}

func (t *tradingClient) ListenPositionReport(handler PositionReportHandler) {
	t.handlers.positionReport = handler
}

func (t *tradingClient) ListenMassPositionReport(handler MassPositionReportHandler) {
	t.handlers.massPositionReport = handler
}

func (t *tradingClient) ListenPositionMaintenanceReport(handler PositionMaintenanceReportHandler) {
	t.handlers.positionMaintenanceReport = handler
}

func (t *tradingClient) ListenReject(handler RejectHandler) {
	t.handlers.reject = handler
}

func (t *tradingClient) ListenListStatus(handler ListStatusHandler) {
	t.handlers.listStatus = handler
}

func (t *tradingClient) ListenBalanceIncrementalRefresh(handler BalanceIncrementalRefreshHandler) {
	t.handlers.balanceIncrementalRefresh = handler
}

func (t *tradingClient) ListenBalanceSnapshotRefresh(handler BalanceSnapshotRefreshHandler) {
	t.handlers.balanceSnapshotRefresh = handler
}

func (t *tradingClient) ListenOrderMassCancelReport(handler OrderMassCancelReportHandler) {
	t.handlers.orderMassCancelReport = handler
}

func (t *tradingClient) ListenHeartbeat(handler HeartbeatHandler) {
	t.handlers.heartbeat = handler
}

func (t *tradingClient) MarketOrder(account uint64, clOrdId string, symbol string, side Side, orderQty string) error {
	return CreateMarketOrder(clOrdId, symbol, side, orderQty, account).Send(t)
}

func (t *tradingClient) LimitOrder(account uint64, clOrdId string, symbol string, side Side, price string, orderQty string) error {
	return CreateLimitOrder(clOrdId, symbol, side, orderQty, account, price).Send(t)
}

func (t *tradingClient) StopOrder(account uint64, clOrdId string, symbol string, side Side, stopPx string, orderQty string) error {
	return CreateStopOrder(clOrdId, symbol, side, orderQty, account, stopPx).Send(t)
}

func (t *tradingClient) MarketIfTouchOrder(account uint64, clOrdId string, symbol string, side Side, stopPx string, orderQty string) error {
	return CreateMarketIfTouchOrder(clOrdId, symbol, side, orderQty, account, stopPx).Send(t)
}

func (t *tradingClient) Cancel(cmd *xmsg.OrderCancelRequest) error {
	if cmd.MsgType != xmsg.MsgType_OrderCancelRequestMsgType {
		return fmt.Errorf("msgType %s. but must be %s", cmd.MsgType, xmsg.MsgType_OrderCancelRequestMsgType)
	}
	return t.Send(cmd)
}

func (t *tradingClient) CancelByClOrdId(account uint64, clOrdId, origClOrdId, symbol string, side Side) error {
	var request = xmsg.OrderCancelRequest{
		MsgType:      xmsg.MsgType_OrderCancelRequestMsgType,
		ClOrdId:      clOrdId,
		OrigClOrdId:  origClOrdId,
		Symbol:       symbol,
		Side:         string(side),
		Account:      account,
		TransactTime: time.Now().UnixNano(),
	}
	return t.Send(request)
}

func (t *tradingClient) CancelByOrderId(account uint64, clOrdId, orderId, symbol string, side Side) error {
	cmd := xmsg.OrderCancelRequest{
		MsgType:      xmsg.MsgType_OrderCancelRequestMsgType,
		ClOrdId:      clOrdId,
		OrderId:      orderId,
		Symbol:       symbol,
		Side:         string(side),
		Account:      account,
		TransactTime: time.Now().UnixNano(),
	}
	return t.Send(cmd)
}

func (t *tradingClient) Replace(cmd xmsg.OrderCancelReplaceRequest) error {
	if cmd.MsgType != xmsg.MsgType_OrderCancelReplaceRequestMsgType {
		return fmt.Errorf("msgType %s. but must be %s", cmd.MsgType, xmsg.MsgType_OrderCancelReplaceRequestMsgType)
	}
	return t.Send(cmd)
}

func (t *tradingClient) CollapsePositions(account uint64, symbol string, requestId string) error {
	request := xmsg.PositionMaintenanceRequest{
		MsgType:        xmsg.MsgType_PositionMaintenanceRequest,
		Account:        account,
		Symbol:         symbol,
		PosReqId:       requestId,
		PosTransType:   PosTransTypeCollapse,
		PosMaintAction: PosMaintActionReplace,
	}
	return t.Send(request)
}

func (t *tradingClient) AccountStatusReport(account uint64, requestId string) error {
	request := xmsg.AccountStatusReportRequest{
		MsgType:                xmsg.MsgType_AccountStatusReportRequest,
		Account:                account,
		AccountStatusRequestId: requestId,
	}
	return t.Send(request)
}

func (t *tradingClient) Orders(account uint64, requestId string) error {
	request := xmsg.OrderStatusRequest{
		MsgType:         xmsg.MsgType_OrderMassStatusRequest,
		Account:         account,
		MassStatusReqId: requestId,
	}
	return t.Send(request)
}

func (t *tradingClient) GetPositions(account uint64, requestId string) error {
	request := xmsg.PositionsRequest{
		MsgType:  xmsg.MsgType_RequestForPositions,
		Account:  account,
		PosReqId: requestId,
	}
	return t.Send(request)
}

func (t *tradingClient) MassCancel(account uint64, clOrdId string) error {
	return newOrderMassCancel(account, clOrdId).Send(t)
}

func (t *tradingClient) SendOrderCancelRequest(accountID uint64, symbol Symbol, side Side, orderID, clientOrderID string) error {
	cmd := xmsg.OrderCancelRequest{
		MsgType:      xmsg.MsgType_OrderCancelRequestMsgType,
		Account:      accountID,
		ClOrdId:      ID(""),
		Symbol:       string(symbol),
		Side:         string(side),
		OrderId:      orderID,
		OrigClOrdId:  clientOrderID,
		TransactTime: time.Now().UTC().UnixNano(),
	}
	return t.Send(cmd)
}

func (t *tradingClient) SendOrderMassStatusRequest(accountID uint64) error {
	cmd := xmsg.OrderStatusRequest{
		MsgType: xmsg.MsgType_OrderMassStatusRequest,
		Account: accountID,
	}
	return t.Send(cmd)
}

func (t *tradingClient) SendAccountStatusReportRequest(accountID uint64) error {
	cmd := xmsg.NewOrderSingle{
		MsgType: xmsg.MsgType_AccountStatusReportRequest,
		Account: accountID,
	}
	return t.Send(cmd)
}

func (t *tradingClient) SendApplicationHeartbeat(groupId string, heartbeatInSec int32) error {
	cmd := xmsg.ApplicationHeartbeat{
		MsgType:    xmsg.MsgType_ApplicationHeartbeat,
		GrpID:      groupId,
		HeartBtInt: heartbeatInSec,
	}
	return t.Send(cmd)
}

func (t *tradingClient) Send(cmd interface{}) error {
	data, err := fixjson.Marshal(cmd)
	if err != nil {
		return err
	}
	return t.client.WriteBytes(data)
}

func (t *tradingClient) incomeHandler(msg []byte) {
	mth := new(xmsg.MsgTypeHeader)
	err := fixjson.Unmarshal(msg, mth)
	if err != nil {
		t.client.Logger().Errorf("error: %s. on fixjson.Unmarshal(%s)", err, string(msg))
	}
	switch mth.MsgType {
	case xmsg.MsgType_Heartbeat:
		v := new(xmsg.Heartbeat)
		if _, err := t.unmarshal(msg, v); err == nil {
			handler := t.handlers.heartbeat
			if handler != nil {
				go handler(t, v)
			}
		}
	case xmsg.MsgType_LogonMsgType:
		v := new(xmsg.Logon)
		if _, err := t.unmarshal(msg, v); err == nil {
			if t.handlers.logon != nil {
				go t.handlers.logon(t, v)
			}
			if t.handlers.logonInternal != nil {
				go t.handlers.logonInternal(t, v)
			}

		}

	case xmsg.MsgType_MarginRequirementReport:
		v := new(xmsg.MarginRequirementReport)
		if _, err := t.unmarshal(msg, v); err == nil {
			if t.handlers.marginRequirementReport != nil {
				go t.handlers.marginRequirementReport(t, v)
			}
		}
	case xmsg.MsgType_ExecutionReportMsgType:
		v := new(xmsg.ExecutionReport)
		if _, err := t.unmarshal(msg, v); err == nil {
			if t.handlers.executionReport != nil {
				go t.handlers.executionReport(t, v)
			}
		}
	case xmsg.MsgType_OrderCancelRejectMsgType:
		v := new(xmsg.OrderCancelReject)
		if _, err := t.unmarshal(msg, v); err == nil {
			if t.handlers.orderCancelReject != nil {
				go t.handlers.orderCancelReject(t, v)
			}
		}
	case xmsg.MsgType_OrderMassStatusResponse:
		v := new(xmsg.OrderMassStatusResponse)
		if _, err := t.unmarshal(msg, v); err == nil {
			if t.handlers.orderMassStatus != nil {
				go t.handlers.orderMassStatus(t, v)
			}
		}
	case xmsg.MsgType_AccountStatusReport:
		v := new(xmsg.BalanceSnapshotRefresh)
		if _, err := t.unmarshal(msg, v); err == nil {
			if t.handlers.balanceSnapshotRefresh != nil {
				go t.handlers.balanceSnapshotRefresh(t, v)
			}
		}
	case xmsg.MsgType_AccountStatusUpdateReport:
		v := new(xmsg.BalanceIncrementalRefresh)
		if _, err := t.unmarshal(msg, v); err == nil {
			if t.handlers.balanceIncrementalRefresh != nil {
				go t.handlers.balanceIncrementalRefresh(t, v)
			}
		}
	case xmsg.MsgType_PositionReport:
		v := new(xmsg.PositionReport)
		if _, err := t.unmarshal(msg, v); err == nil {
			if t.handlers.positionReport != nil {
				go t.handlers.positionReport(t, v)
			}
		}

	case xmsg.MsgType_MassPositionReport:
		v := new(xmsg.MassPositionReport)
		if _, err := t.unmarshal(msg, v); err == nil {
			if t.handlers.massPositionReport != nil {
				go t.handlers.massPositionReport(t, v)
			}
		}
	case xmsg.MsgType_PositionMaintenanceReport:
		v := new(xmsg.PositionMaintenanceReport)
		if _, err := t.unmarshal(msg, v); err == nil {
			if t.handlers.positionMaintenanceReport != nil {
				go t.handlers.positionMaintenanceReport(t, v)
			}
		}
	case xmsg.MsgType_RejectMsgType:
		v := new(xmsg.Reject)
		if _, err := t.unmarshal(msg, v); err == nil {
			if t.handlers.reject != nil {
				go t.handlers.reject(t, v)
			}
		}
	case xmsg.MsgType_ListStatus:
		v := new(xmsg.ListStatus)
		if _, err := t.unmarshal(msg, v); err == nil {
			if t.handlers.listStatus != nil {
				go t.handlers.listStatus(t, v)
			}
		}
	case xmsg.MsgType_OrderMassCancelReport:
		v := new(xmsg.OrderMassCancelReport)
		if _, err := t.unmarshal(msg, v); err == nil {
			if t.handlers.orderMassCancelReport != nil {
				go t.handlers.orderMassCancelReport(t, v)
			}
		}
	// Not implemented yet
	default:
		t.client.Logger().Errorf("unknown message type %s", string(msg))
	}
}

func (t *tradingClient) onConnect(WsClient) {
	loginCmd := t.loginCmd()
	err := t.client.WriteBytes(loginCmd)
	if err == nil {
		t.client.Logger().Errorf("%s on t.client.WriteBytes()", err)
	}
}

func (t *tradingClient) loginCmd() []byte {
	nonce, payload, sigHex, err := sing(t.apiSecret)
	if err != nil {
		t.client.Logger().Errorf("%s on query.sing()", err)
	}

	logonCmd := xmsg.Logon{
		MsgType:     xmsg.MsgType_LogonMsgType,
		SendingTime: nonce,
		Username:    t.apiKey,
		Password:    sigHex,
		RawData:     payload,
	}
	cmd, _ := fixjson.Marshal(logonCmd)

	return cmd
}

func (t *tradingClient) unmarshal(msg []byte, v interface{}) (interface{}, error) {
	err := fixjson.Unmarshal(msg, v)
	if err != nil {
		t.client.Logger().Errorf("error: %s. on fixjson.Unmarshal(%s)", err, string(msg))
		return nil, err
	}

	return v, nil
}

func (t *tradingClient) SetDisconnectHandler(handler TradingDisconnectHandler) {
	t.client.setDisconnectHandler(func() {
		handler(t, t.client.Logger())
	})
}
