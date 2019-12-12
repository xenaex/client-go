package xena

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"github.com/xenaex/client-go/xena/fixjson"
	"github.com/xenaex/client-go/xena/xmsg"
	"time"
)

const (
	wsTradingURL = "ws://localhost/api/ws/trading"
)

// TradingClient interface
type TradingClient interface {
	Client() WsClient
	SendLimitOrder(accountID uint64, clientOrderID string, symbol Symbol, side Side, price, qty string) error
	SendAccountStatusReportRequest(accountID uint64) error
	SendOrderMassStatusRequest(accountID uint64) error
	SendOrderCancelRequest(accountID uint64, symbol Symbol, side Side, orderID, clientOrderID string) error
	ListenLogon(handler LogonHandler)
	ListenMarginRequirementReport(handler MarginRequirementReportHandler)
	ListenExecutionReport(handler ExecutionReportHandler)
	ListenOrderCancelReject(handler OrderCancelRejectHandler)
	ListenOrderMassStatusResponse(handler OrderMassStatusResponseHandler)
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
	t.client.Connect()

	return &t
}

// LogonHandler will be called on Logon response received
type LogonHandler func(t TradingClient, m *xmsg.Logon)

// MarginRequirementReportHandler will be called on MarginRequirementReport received
type MarginRequirementReportHandler func(t TradingClient, m *xmsg.MarginRequirementReport)

// ExecutionReportHandler will be called on ExecutionReport received
type ExecutionReportHandler func(t TradingClient, m *xmsg.ExecutionReport)

// OrderCancelRejectHandler will be called on OrderCancelReject received
type OrderCancelRejectHandler func(t TradingClient, m *xmsg.OrderCancelReject)

// OrderMassStatusResponseHandler will be called on OrderMassStatusResponse received
type OrderMassStatusResponseHandler func(t TradingClient, m *xmsg.OrderMassStatusResponse)

type tradingClient struct {
	client    WsClient
	apiKey    string
	apiSecret string
	handlers  struct {
		logon                   LogonHandler
		marginRequirementReport MarginRequirementReportHandler
		executionReport         ExecutionReportHandler
		orderMassStatus         OrderMassStatusResponseHandler
		orderCancelReject       OrderCancelRejectHandler
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
	return t.send(cmd)
}

func (t *tradingClient) SendOrderMassStatusRequest(accountID uint64) error {
	cmd := xmsg.NewOrderSingle{
		MsgType: xmsg.MsgType_OrderMassStatusRequest,
		Account: accountID,
	}
	return t.send(cmd)
}

func (t *tradingClient) SendAccountStatusReportRequest(accountID uint64) error {
	cmd := xmsg.NewOrderSingle{
		MsgType: xmsg.MsgType_AccountStatusReportRequest,
		Account: accountID,
	}
	return t.send(cmd)
}

func (t *tradingClient) send(cmd interface{}) error {
	data, err := fixjson.Marshal(cmd)
	if err != nil {
		return err
	}
	return t.client.WriteBytes(data)
}

func (t *tradingClient) SendLimitOrder(accountID uint64, clientOrderID string, symbol Symbol, side Side, price, qty string) error {
	cmd := xmsg.NewOrderSingle{
		MsgType:      xmsg.MsgType_NewOrderSingleMsgType,
		TransactTime: time.Now().UTC().UnixNano(),
		OrdType:      xmsg.OrdType_Limit,
		ClOrdId:      clientOrderID,
		Symbol:       string(symbol),
		Side:         string(side),
		Account:      accountID,
		OrderQty:     qty,
		Price:        price,
	}
	return t.send(cmd)
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
		// Not implemented yet
	default:
		t.client.Logger().Errorf("unknown message type %s", string(msg))
	}
}

func (t *tradingClient) Client() WsClient {
	return t.client
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
