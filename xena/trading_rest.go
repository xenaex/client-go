package xena

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/xenaex/client-go/xena/xmsg"
)

func NewTradingREST(apiKey, apiSecret string, options ...RestOption) TradingREST {
	cfg := &restConf{}
	for _, ots := range []RestOption{withRestDefaultLogger, WithRestTradingHost, withRestDefaultTimeout, WithRestUserAgent(userAgent), withRestHeader("X-Auth-Api-Key", apiKey)} {
		ots(cfg)
	}
	for _, ots := range options {
		ots(cfg)
	}
	return &tradingREST{
		apiSecret: apiSecret,
		baseREST:  newBaseREST(cfg),
	}
}

type TradingREST interface {
	// NewOrder places new order.
	NewOrder(cmd *xmsg.NewOrderSingle) (*xmsg.ExecutionReport, error)

	//SendMarketOrder places new market order.
	SendMarketOrder(accountId uint64, clOrdId string, symbol string, side Side, orderQty string) (*xmsg.ExecutionReport, error)
	//SendLimitOrder places new limit order.
	SendLimitOrder(accountId uint64, clOrdId string, symbol string, side Side, price string, orderQty string) (*xmsg.ExecutionReport, error)
	//SendStopOrder places new stop order.
	SendStopOrder(accountId uint64, clOrdId string, symbol string, side Side, stopPx string, orderQty string) (*xmsg.ExecutionReport, error)
	//SendMarketIfTouchOrder place new market-if-touch order.
	SendMarketIfTouchOrder(accountId uint64, clOrdId string, symbol string, side Side, stopPx string, orderQty string) (*xmsg.ExecutionReport, error)
	// SendCancel Cancels an existing order.
	SendCancel(cmd *xmsg.OrderCancelRequest) (*xmsg.ExecutionReport, error)
	SendMassCancelCmd(cmd *xmsg.OrderMassCancelRequest) (*xmsg.ExecutionReport, error)

	//SendCancelByClOrdId cancels an existing order by original client order id.
	SendCancelByClOrdId(account uint64, clOrdId, origClOrdId, symbol string, side Side) (*xmsg.ExecutionReport, error)

	//SendCancelByOrderId cancel an existing order by order id.
	SendCancelByOrderId(account uint64, clOrdId, orderId, symbol string, side Side) (*xmsg.ExecutionReport, error)
	// SendMassCancel creates OrderMassCancelRequest and send request.
	SendMassCancel(account uint64, clOrdId string) (*xmsg.ExecutionReport, error)
	// SendReplace cancel an existing order and replace.
	SendReplace(request xmsg.OrderCancelReplaceRequest) (*xmsg.ExecutionReport, error)

	// SendCollapsePositions collapse all existing positions for margin account and symbol.
	SendCollapsePositions(account uint64, symbol string, requestId string) (*xmsg.PositionMaintenanceReport, error)
	// GetPositions request all positions for account.
	// lists account's open positions.
	GetPositions(account uint64, requestId string) ([]*xmsg.PositionReport, error)
	GetPositionsHistory(accountId uint64, request PositionsHistoryRequest) ([]*xmsg.PositionReport, error)
	GetOrders(account uint64, requestId string) ([]*xmsg.ExecutionReport, error)
	GetTradeHistory(accountId uint64, request TradeHistoryRequest) ([]*xmsg.ExecutionReport, error)
	GetBalance(accountId uint64) (*xmsg.BalanceSnapshotRefresh, error)
	GetMarginRequirements(accountId uint64) (*xmsg.MarginRequirementReport, error)
	GetAccounts() ([]*xmsg.AccountInfo, error)
}

type tradingREST struct {
	baseREST
	apiSecret string
}

// GetOrders request all orders and fills for account.
// To receive response.
func (t *tradingREST) NewOrder(cmd *xmsg.NewOrderSingle) (*xmsg.ExecutionReport, error) {
	query := NewQuery("order", "new")
	er := new(xmsg.ExecutionReport)
	err := t.sendPost(query, cmd, er)
	if err != nil {
		return nil, err
	}
	return er, nil
}

//SendMarketOrder place new market order.
func (t *tradingREST) SendMarketOrder(accountId uint64, clOrdId string, symbol string, side Side, orderQty string) (*xmsg.ExecutionReport, error) {
	order := CreateMarketOrder(clOrdId, symbol, side, orderQty, accountId).Build()
	return t.NewOrder(order)
}

//SendLimitOrder place new limit order.
func (t *tradingREST) SendLimitOrder(accountId uint64, clOrdId string, symbol string, side Side, price string, orderQty string) (*xmsg.ExecutionReport, error) {
	order := CreateLimitOrder(clOrdId, symbol, side, orderQty, accountId, price).Build()
	return t.NewOrder(order)
}

//SendStopOrder place new stop order.
func (t *tradingREST) SendStopOrder(accountId uint64, clOrdId string, symbol string, side Side, stopPx string, orderQty string) (*xmsg.ExecutionReport, error) {
	order := CreateStopOrder(clOrdId, symbol, side, orderQty, accountId, stopPx).Build()
	return t.NewOrder(order)
}

//SendMarketIfTouchOrder place new market-if-touch order.
func (t *tradingREST) SendMarketIfTouchOrder(accountId uint64, clOrdId string, symbol string, side Side, stopPx string, orderQty string) (*xmsg.ExecutionReport, error) {
	order := CreateMarketIfTouchOrder(clOrdId, symbol, side, orderQty, accountId, stopPx).Build()
	return t.NewOrder(order)
}

func (t *tradingREST) SendCancel(cmd *xmsg.OrderCancelRequest) (*xmsg.ExecutionReport, error) {
	query := NewQuery("order", "cancel")
	er := new(xmsg.ExecutionReport)
	err := t.sendPost(query, cmd, er)
	return er, err
}

func (t *tradingREST) SendMassCancelCmd(cmd *xmsg.OrderMassCancelRequest) (*xmsg.ExecutionReport, error) {
	query := NewQuery("order", "mass-cancel")
	er := new(xmsg.ExecutionReport)
	err := t.sendPost(query, cmd, er)
	return er, err
}

func (t *tradingREST) GetBalance(accountId uint64) (*xmsg.BalanceSnapshotRefresh, error) {
	query := NewQuery("accounts", strconv.FormatUint(accountId, 10), "balance")
	resp := new(xmsg.BalanceSnapshotRefresh)
	err := t.sendGet(query, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *tradingREST) GetMarginRequirements(accountId uint64) (*xmsg.MarginRequirementReport, error) {
	query := NewQuery("accounts", strconv.FormatUint(accountId, 10), "margin-requirements")
	resp := new(xmsg.MarginRequirementReport)
	err := t.sendGet(query, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *tradingREST) GetTradeHistory(accountId uint64, request TradeHistoryRequest) ([]*xmsg.ExecutionReport, error) {
	query := NewQuery("accounts", strconv.FormatUint(accountId, 10), "trade-history")
	query.AddQueryf("trade_id", request.TradeId).
		AddQueryf("client_order_id", request.ClOrdId).
		AddQueryf("symbol", request.Symbol).
		AddQueryf("from", request.From).
		AddQueryf("to", request.To).
		AddQueryf("page", request.PageNumber).
		AddQueryf("limit", request.Limit)
	resp := make([]*xmsg.ExecutionReport, 0)
	err := t.sendGet(query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//SendCancelByClOrdId cancels an existing order by original client order id.
func (t *tradingREST) SendCancelByClOrdId(account uint64, clOrdId, origClOrdId, symbol string, side Side) (*xmsg.ExecutionReport, error) {
	var request = xmsg.OrderCancelRequest{
		MsgType:      xmsg.MsgType_OrderCancelRequestMsgType,
		ClOrdId:      clOrdId,
		OrigClOrdId:  origClOrdId,
		Symbol:       symbol,
		Side:         string(side),
		Account:      account,
		TransactTime: time.Now().UnixNano(),
	}
	return t.SendCancel(&request)
}

//SendCancelByOrderId cancel an existing order by order id.
func (t *tradingREST) SendCancelByOrderId(account uint64, clOrdId, orderId, symbol string, side Side) (*xmsg.ExecutionReport, error) {
	var request = xmsg.OrderCancelRequest{
		MsgType:      xmsg.MsgType_OrderCancelRequestMsgType,
		ClOrdId:      clOrdId,
		OrderId:      orderId,
		Symbol:       symbol,
		Side:         string(side),
		Account:      account,
		TransactTime: time.Now().UnixNano(),
	}
	return t.SendCancel(&request)
}

// SendMassCancel send OrderMassCancelRequest request.
func (t *tradingREST) SendMassCancel(account uint64, clOrdId string) (*xmsg.ExecutionReport, error) {
	return t.SendMassCancelCmd(newOrderMassCancel(account, clOrdId).Build())
}

// GetPositions request all positions for account.
func (t *tradingREST) GetPositions(account uint64, requestId string) ([]*xmsg.PositionReport, error) {
	query := NewQuery("accounts", strconv.FormatUint(account, 10), "positions")
	//request := xmsg.PositionsRequest{
	//	MsgType:  xmsg.MsgType_RequestForPositions,
	//	Account:  account,
	//	PosReqId: requestId,
	//}
	resp := make([]*xmsg.PositionReport, 0)
	err := t.sendGet(query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *tradingREST) sendPost(query query, cmd, ret interface{}) error {
	data, err := json.Marshal(cmd)
	if err != nil {
		return err
	}
	query, err = query.AddSecret(t.apiSecret)
	if err != nil {
		t.config.logger.Errorf("%s on query.AddSecret()", err)
		return nil
	}
	resp, body, err := t.post(query, data)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s, body %s", resp.Status, string(body))
	}

	err = json.Unmarshal(body, ret)
	if err != nil {
		t.config.logger.Errorf("%s. on json.Unmarshal(%s) code: %d", err, string(body), resp.StatusCode)
	}
	return err
}

func (t *tradingREST) sendGet(query query, ret interface{}) error {
	query, err := query.AddSecret(t.apiSecret)
	if err != nil {
		t.config.logger.Errorf("%s on query.AddSecret()", err)
		return nil
	}
	resp, body, err := t.get(query)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp == nil && err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s, body %s", resp.Status, string(body))
	}

	err = json.Unmarshal(body, ret)
	if err != nil {
		t.config.logger.Errorf("%s. on json.Unmarshal(%s)", err, string(body))
	}
	return err
}

func (t *tradingREST) GetOrders(account uint64, requestId string) ([]*xmsg.ExecutionReport, error) {
	query := NewQuery("accounts", strconv.FormatUint(account, 10), "orders")
	resp := make([]*xmsg.ExecutionReport, 0)
	err := t.sendGet(query, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SendReplace cancel an existing order and replace.
func (t *tradingREST) SendReplace(request xmsg.OrderCancelReplaceRequest) (*xmsg.ExecutionReport, error) {
	query := NewQuery("order", "replace")
	resp := new(xmsg.ExecutionReport)
	err := t.sendPost(query, request, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// SendCollapsePositions collapse all existing positions for margin account and symbol.
func (t *tradingREST) SendCollapsePositions(account uint64, symbol string, requestId string) (*xmsg.PositionMaintenanceReport, error) {
	request := xmsg.PositionMaintenanceRequest{
		MsgType:        xmsg.MsgType_PositionMaintenanceRequest,
		Account:        account,
		Symbol:         symbol,
		PosReqId:       requestId,
		PosTransType:   PosTransTypeCollapse,
		PosMaintAction: PosMaintActionReplace,
	}
	query := NewQuery("position", "maintenance")
	resp := new(xmsg.PositionMaintenanceReport)
	err := t.sendPost(query, request, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (t *tradingREST) GetAccounts() ([]*xmsg.AccountInfo, error) {
	query := NewQuery("accounts")
	resp := struct {
		Accounts []*xmsg.AccountInfo `json:"accounts"`
	}{
		Accounts: make([]*xmsg.AccountInfo, 0),
	}
	err := t.sendGet(query, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Accounts, nil
}
