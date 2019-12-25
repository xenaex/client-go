package xena

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/xenaex/client-go/xena/fixjson"
	"github.com/xenaex/client-go/xena/xmsg"
)

const (
	wsTradingURL = "ws://localhost/api/ws/trading"
)

// OrderMassCancel string symbol = null, string side = null, string positionEffect = PositionEffect.Default;
type OrderMassCancel struct {
	client TradingClient
}

// TradingClient Xena Trading websocket client interface.
type TradingClient interface {
	ListenLogon(handler LogonHandler)
	ListenMarginRequirementReport(handler MarginRequirementReportHandler)
	ListenExecutionReport(handler ExecutionReportHandler)
	ListenOrderCancelReject(handler OrderCancelRejectHandler)
	ListenOrderMassStatusResponse(handler OrderMassStatusResponseHandler)
	ListenBalanceIncrementalRefresh(handler BalanceIncrementalRefreshHandler)
	ListenBalanceSnapshotRefresh(handler BalanceSnapshotRefreshHandler)
	ListenPositionReport(handler PositionReportHandler)
	ListenMassPositionReport(handler MassPositionReportHandler)
	ListenPositionMaintenanceReport(handler PositionMaintenanceReportHandler)
	ListenReject(handler RejectHandler)
	ListenListStatus(handler ListStatusHandler)
	// ListenOrderMassCancelReport(handler OrderMassCancelReportHandler)

	// ConnectAndLogon connects to websocket and if connection was successful sends Logon message with provided authorization data.
	// Logon response. If logon is rejected Logon.RejectText will contain the reject reason.
	ConnectAndLogon() (*xmsg.Logon, error)

	Send(cmd interface{}) error

	// MarketOrder place new market order.
	MarketOrder(clOrdId string, symbol Symbol, side Side, orderQty string, account uint64) error

	// LimitOrder place new limit order.
	LimitOrder(clOrdId string, symbol Symbol, side Side, orderQty string, account uint64, price string) error

	// StopOrder place new stop order.
	StopOrder(clOrdId string, symbol Symbol, side Side, orderQty string, account uint64, stopPx string) error

	// MarketIfTouchOrder place new market-if-touch order.
	MarketIfTouchOrder(clOrdId string, symbol Symbol, side Side, orderQty string, account uint64, stopPx string) error

	// CancelOrderByClOrdId cancels an existing order by original client order id.
	CancelOrderByClOrdId(clOrdId, origClOrdId string, symbol Symbol, side Side, account uint64) error

	// CancelOrderByOrderId cancel an existing order by order id.
	CancelOrderByOrderId(clOrdId, orderId string, symbol Symbol, side Side, account uint64) error

	// CancelReplaceOrder cancel an existing order and replace.
	CancelReplaceOrder(request xmsg.OrderCancelReplaceRequest) error

	// CollapsePositions collapse all existing positions for margin account and symbol.
	CollapsePositions(account uint64, symbol Symbol, posReqId string) error

	// AccountStatusReport request account status report.
	// To receive response, client has to listen ListenBalanceSnapshotRefresh.
	AccountStatusReport(account uint64, requestId string) error

	// GetOrdersAndFills request all orders and fills for account.
	// To receive response, client has to listen Listen???.
	GetOrdersAndFills(account uint64, requestId string) error

	// GetPositions request all positions for account.
	// To receive response, client has to listen ListenMassPositionReport.
	GetPositions(account uint64, requestId string) error

	// OrderMassCancel send OrderMassCancelRequest request.
	// To receive response, client has to listen ListenOrderMassCancelReport.
	OrderMassCancel(account uint64, clOrdId string, symbol Symbol, side Side, positionEffect string) error
}

// NewTradingClient constructor
func NewTradingClient(apiKey, apiSecret string, opts ...WsOption) TradingClient {
	t := tradingClient{
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}

	defaultOpts := []WsOption{
		WithURL(wsTradingURL),
		WithConnectHandler(t.onConnect),
		WithHandler(t.incomeHandler),
	}
	opts = append(defaultOpts, opts...)

	t.client = NewWsClient(opts...)

	return &t
}

func (t *tradingClient) ConnectAndLogon() (*xmsg.Logon, error) {
	msgs := make(chan *xmsg.Logon, 1)
	// errs := make(chan *xmsg.Logon, 1)
	t.handlers.logonInternal = func(t TradingClient, m *xmsg.Logon) {
		msgs <- m
		close(msgs)
	}
	defer func() { t.handlers.logonInternal = nil }()
	err := t.client.ConnectOnly()
	if err != nil {
		return nil, err
	}
	select {
	case m, ok := <-msgs:
		if ok {
			return m, nil
		}
	case <-time.NewTimer(time.Minute).C:
		// TODO: error time out.
		return nil, fmt.Errorf("timeout logon")
	}
	return nil, fmt.Errorf("something happened")
}

// LogonHandler will be called on Logon response received
type LogonHandler func(t TradingClient, m *xmsg.Logon)

// MarginRequirementReportHandler will be called on MarginRequirementReport received
type MarginRequirementReportHandler func(t TradingClient, m *xmsg.MarginRequirementReport)

// BalanceIncrementalRefreshHandler will be called on BalanceIncrementalRefresh received
type BalanceIncrementalRefreshHandler func(t TradingClient, m *xmsg.BalanceIncrementalRefresh)

// BalanceSnapshotRefreshHandler will be called on BalanceSnapshotRefresh received
type BalanceSnapshotRefreshHandler func(t TradingClient, m *xmsg.BalanceSnapshotRefresh)

// ExecutionReportHandler will be called on ExecutionReport received
type ExecutionReportHandler func(t TradingClient, m *xmsg.ExecutionReport)

// OrderCancelRejectHandler will be called on OrderCancelReject received
type OrderCancelRejectHandler func(t TradingClient, m *xmsg.OrderCancelReject)

// OrderMassStatusResponseHandler will be called on OrderMassStatusResponse received
type OrderMassStatusResponseHandler func(t TradingClient, m *xmsg.OrderMassStatusResponse)

// PositionReportHandler will be called on PositionReport received
type PositionReportHandler func(t TradingClient, m *xmsg.PositionReport)

// MassPositionReportHandler will be called on MassPositionReport received
type MassPositionReportHandler func(t TradingClient, m *xmsg.MassPositionReport)

// PositionMaintenanceReportHandler will be called on PositionMaintenanceReport received
type PositionMaintenanceReportHandler func(t TradingClient, m *xmsg.PositionMaintenanceReport)

// RejectHandler will be called on Reject received
type RejectHandler func(t TradingClient, m *xmsg.Reject)

// ListStatusHandler will be called on ListStatus received
type ListStatusHandler func(t TradingClient, m *xmsg.ListStatus)

// OrderMassCancelReportHandler will be called on OrderMassCancelReport received
// type OrderMassCancelReportHandler func(t TradingClient, m *xmsg.OrderMassCancelReport)

type tradingClient struct {
	client    WsClient
	apiKey    string
	apiSecret string
	handlers  struct {
		balanceSnapshotRefresh    BalanceSnapshotRefreshHandler
		balanceIncrementalRefresh BalanceIncrementalRefreshHandler
		executionReport           ExecutionReportHandler
		listStatus                ListStatusHandler
		logon                     LogonHandler
		logonInternal             LogonHandler
		marginRequirementReport   MarginRequirementReportHandler
		massPositionReport        MassPositionReportHandler
		orderCancelReject         OrderCancelRejectHandler
		// orderMassCancelReport     OrderMassCancelReportHandler
		orderMassStatus           OrderMassStatusResponseHandler
		positionMaintenanceReport PositionMaintenanceReportHandler
		positionReport            PositionReportHandler
		reject                    RejectHandler
	}
}

// ListenLogon subscribe on Logon messages
func (t *tradingClient) ListenLogon(handler LogonHandler) {
	t.handlers.logon = handler
}

// ListenMarginRequirementReport subscribe on MarginRequirementReport messages
func (t *tradingClient) ListenMarginRequirementReport(handler MarginRequirementReportHandler) {
	t.handlers.marginRequirementReport = handler
}

// ListenExecutionReport subscribe on ExecutionReport messages
func (t *tradingClient) ListenExecutionReport(handler ExecutionReportHandler) {
	t.handlers.executionReport = handler
}

// ListenOrderCancelReject subscribe on OrderCancelReject messages
func (t *tradingClient) ListenOrderCancelReject(handler OrderCancelRejectHandler) {
	t.handlers.orderCancelReject = handler
}

// ListenOrderMassStatusResponse subscribe on OrderMassStatusResponse messages
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

// MarketOrder place new market order.
func (t *tradingClient) MarketOrder(clOrdId string, symbol Symbol, side Side, orderQty string, account uint64) error {
	return CreateMarketOrder(clOrdId, symbol, side, orderQty, account).Send(t)
}

// LimitOrder place new limit order.
func (t *tradingClient) LimitOrder(clOrdId string, symbol Symbol, side Side, orderQty string, account uint64, price string) error {
	return CreateLimitOrder(clOrdId, symbol, side, orderQty, account, price).Send(t)
}

// StopOrder place new stop order.
func (t *tradingClient) StopOrder(clOrdId string, symbol Symbol, side Side, orderQty string, account uint64, stopPx string) error {
	return CreateStopOrder(clOrdId, symbol, side, orderQty, account, stopPx).Send(t)
}

// MarketIfTouchOrder place new market-if-touch order.
func (t *tradingClient) MarketIfTouchOrder(clOrdId string, symbol Symbol, side Side, orderQty string, account uint64, stopPx string) error {
	return CreateMarketIfTouchOrder(clOrdId, symbol, side, orderQty, account, stopPx).Send(t)
}

// CancelOrderByClOrdId cancels an existing order by original client order id.
func (t *tradingClient) CancelOrderByClOrdId(clOrdId, origClOrdId string, symbol Symbol, side Side, account uint64) error {
	var request = xmsg.OrderCancelRequest{
		MsgType:      xmsg.MsgType_OrderCancelRequestMsgType,
		ClOrdId:      clOrdId,
		OrigClOrdId:  origClOrdId,
		Symbol:       string(symbol),
		Side:         string(side),
		Account:      account,
		TransactTime: time.Now().UnixNano(),
	}
	return t.Send(request)
}

// CancelOrderByOrderId cancel an existing order by order id.
func (t *tradingClient) CancelOrderByOrderId(clOrdId, orderId string, symbol Symbol, side Side, account uint64) error {
	request := xmsg.OrderCancelRequest{
		MsgType:      xmsg.MsgType_OrderCancelRequestMsgType,
		ClOrdId:      clOrdId,
		OrderId:      orderId,
		Symbol:       string(symbol),
		Side:         string(side),
		Account:      account,
		TransactTime: time.Now().UnixNano(),
	}
	return t.Send(request)
}

// CancelReplaceOrder cancel an existing order and replace.
func (t *tradingClient) CancelReplaceOrder(request xmsg.OrderCancelReplaceRequest) error {
	return t.Send(request)
}

// CollapsePositions collapse all existing positions for margin account and symbol.
func (t *tradingClient) CollapsePositions(account uint64, symbol Symbol, posReqId string) error {
	request := xmsg.PositionMaintenanceRequest{
		MsgType:        xmsg.MsgType_PositionMaintenanceRequest,
		Account:        account,
		Symbol:         string(symbol),
		PosReqID:       posReqId,
		PosTransType:   "20",
		PosMaintAction: "2",
	}
	return t.Send(request)
}

// AccountStatusReport request account status report.
// To receive response, client has to listen ListenBalanceSnapshotRefresh.
func (t *tradingClient) AccountStatusReport(account uint64, requestId string) error {
	request := xmsg.AccountStatusReportRequest{
		MsgType: xmsg.MsgType_AccountStatusReportRequest,
		Account: account,
		// AccountStatusRequestId: requestId,
	}
	return t.Send(request)
}

// GetOrdersAndFills request all orders and fills for account.
func (t *tradingClient) GetOrdersAndFills(account uint64, requestId string) error {
	request := xmsg.OrderStatusRequest{
		MsgType: xmsg.MsgType_OrderMassStatusRequest,
		Account: account,
		// MassStatusReqId: requestId,
	}
	return t.Send(request)
}

// GetPositions request all positions for account.
// To receive response, client has to listen ListenMassPositionReport.
func (t *tradingClient) GetPositions(account uint64, requestId string) error {
	request := xmsg.PositionsRequest{
		MsgType: xmsg.MsgType_RequestForPositions,
		Account: account,
		// PosReqId: requestId,
	}
	return t.Send(request)
}

func (t *tradingClient) OrderMassCancel(account uint64, clOrdId string, symbol Symbol, side Side, positionEffect string) error {
	panic("implement me")
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
	cmd := xmsg.NewOrderSingle{
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
		// case xmsg.MsgType_OrderMassCancelReport:
		//	v := new(xmsg.OrderMassCancelReport)
		//	if _, err := t.unmarshal(msg, v); err == nil {
		//		//	if t.handlers.orderMassCancelReport != nil {
		//		//		go t.handlers.orderMassCancelReport(t, v)
		//		//	}
		//	}
	// Not implemented yet
	default:
		t.client.Logger().Errorf("unknown message type %s", string(msg))
	}

	// case xmsg.MsgType_Heartbeat: -> xmsg.Heartbeat
	// it is market data only
	// case xmsg.MsgType_MarketDataRequestReject: -> xmsg.MarketDataRequestReject
	// case xmsg.MsgType_MarketDataIncrementalRefresh: -> xmsg.MarketDataRefresh
	// case xmsg.MsgType_MarketDataSnapshotFullRefresh: -> xmsg.MarketDataRefresh
}

func (t *tradingClient) onConnect(c WsClient) {
	loginCmd := t.loginCmd()
	t.client.WriteBytes(loginCmd)
}

func (t *tradingClient) loginCmd() []byte {
	nonce := time.Now().UnixNano()
	payload := fmt.Sprintf("AUTH%d", nonce)

	// Signature - SHA512 + ECDSA
	privKeyData, err := hex.DecodeString(t.apiSecret)
	if err != nil {
		t.client.Logger().Errorf("error: %s on DecodeString", err)
		return nil
	}

	privKey, err := x509.ParseECPrivateKey(privKeyData)
	if err != nil {
		t.client.Logger().Errorf("error: %s on ParseECPrivateKey", err)
		return nil
	}

	digest := sha256.Sum256([]byte(payload))
	r, s, err := ecdsa.Sign(rand.Reader, privKey, digest[:])
	signature := append(r.Bytes(), s.Bytes()...)
	sigHex := hex.EncodeToString(signature)

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
