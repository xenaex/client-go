package xena

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/xenaex/client-go/xena/xmsg"
)

// NewTradingREST creates rest client of Xena trading api.
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

// TradingClient is rest client interface of the Xena trading api.
type TradingREST interface {
	// NewOrder places new order.
	NewOrder(cmd *xmsg.NewOrderSingle) (*xmsg.ExecutionReport, error)

	// SendMarketOrder places new market order.
	SendMarketOrder(accountId uint64, clOrdId string, symbol string, side Side, orderQty string) (*xmsg.ExecutionReport, error)

	// SendLimitOrder places new limit order.
	SendLimitOrder(accountId uint64, clOrdId string, symbol string, side Side, price string, orderQty string) (*xmsg.ExecutionReport, error)

	// SendStopOrder places new stop order.
	SendStopOrder(accountId uint64, clOrdId string, symbol string, side Side, stopPx string, orderQty string) (*xmsg.ExecutionReport, error)

	// SendMarketIfTouchOrder place new market-if-touch order.
	SendMarketIfTouchOrder(accountId uint64, clOrdId string, symbol string, side Side, stopPx string, orderQty string) (*xmsg.ExecutionReport, error)

	// SendCancel function cancels existing order.
	SendCancel(cmd *xmsg.OrderCancelRequest) (*xmsg.ExecutionReport, error)

	// SendMassCancelCmd cancels all existing orders.
	SendMassCancelCmd(cmd *xmsg.OrderMassCancelRequest) (*xmsg.ExecutionReport, error)

	// SendCancelByClOrdId cancels existing order by original client order id.
	SendCancelByClOrdId(account uint64, clOrdId, origClOrdId, symbol string, side Side) (*xmsg.ExecutionReport, error)

	// SendCancelByOrderId cancels existing order by order id.
	SendCancelByOrderId(account uint64, clOrdId, orderId, symbol string, side Side) (*xmsg.ExecutionReport, error)

	// SendMassCancel cancels all existing orders.
	SendMassCancel(account uint64, clOrdId string) (*xmsg.ExecutionReport, error)

	// SendReplace cancels existing order and replaces by new order.
	SendReplace(request xmsg.OrderCancelReplaceRequest) (*xmsg.ExecutionReport, error)

	// SendCollapsePositions collapses all existing positions for margin account and symbol.
	SendCollapsePositions(account uint64, symbol string, requestId string) (*xmsg.PositionMaintenanceReport, error)

	// GetPositions requests all positions.
	GetPositions(account uint64, requestId string) ([]*xmsg.PositionReport, error)

	// GetPositionsHistory requests all positions history.
	GetPositionsHistory(accountId uint64, request PositionsHistoryRequest) ([]*xmsg.PositionReport, error)

	// GetOrder returns last execution report for order or cancel/replace request.
	GetOrder(account uint64, request OrderRequest) (*xmsg.ExecutionReport, error)

	// GetActiveOrders return list of active orders
	GetActiveOrders(account uint64, requeset ActiveOrdersRequest) ([]*xmsg.ExecutionReport, error)

	// GetLastOrderStatuses return list of last execution reports for non-active orders
	GetLastOrderStatuses(account uint64, requeset LastOrderStatusesRequest) ([]*xmsg.ExecutionReport, error)

	// GetOrderHistory return list of historical execution reports
	GetOrderHistory(account uint64, requeset OrderHistoryRequest) ([]*xmsg.ExecutionReport, error)

	// GetTradeHistory returns trade history.
	GetTradeHistory(accountId uint64, request TradeHistoryRequest) ([]*xmsg.ExecutionReport, error)

	// GetTradeHistory returns balance.
	GetBalance(accountId uint64) (*xmsg.BalanceSnapshotRefresh, error)

	// GetTradeHistory returns MarginRequirements.
	GetMarginRequirements(accountId uint64) (*xmsg.MarginRequirementReport, error)

	// GetTradeHistory returns accounts.
	GetAccounts() ([]*xmsg.AccountInfo, error)

	// SendApplicationHeartbeat sends application heartbeat.
	SendApplicationHeartbeat(groupId string, heartbeatInSec int32) error
}

type tradingREST struct {
	baseREST
	apiSecret string
}

func (t *tradingREST) sendPost(query query, cmd, ret interface{}) error {
	data, err := json.Marshal(cmd)
	if err != nil {
		return err
	}
	query, err = query.addSecret(t.apiSecret)
	if err != nil {
		t.config.logger.Errorf("%s on query.addSecret()", err)
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

func (t *tradingREST) sendPostWithoutResp(query query, cmd interface{}) error {
	data, err := json.Marshal(cmd)
	if err != nil {
		return err
	}
	query, err = query.addSecret(t.apiSecret)
	if err != nil {
		t.config.logger.Errorf("%s on query.addSecret()", err)
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

	return err
}

func (t *tradingREST) sendGet(query query, ret interface{}) error {
	query, err := query.addSecret(t.apiSecret)
	if err != nil {
		t.config.logger.Errorf("%s on query.addSecret()", err)
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

func (t *tradingREST) NewOrder(cmd *xmsg.NewOrderSingle) (*xmsg.ExecutionReport, error) {
	query := newQuery("order", "new")
	er := new(xmsg.ExecutionReport)
	return er, t.sendPost(query, cmd, er)
}

func (t *tradingREST) SendMarketOrder(accountId uint64, clOrdId string, symbol string, side Side, orderQty string) (*xmsg.ExecutionReport, error) {
	order := CreateMarketOrder(clOrdId, symbol, side, orderQty, accountId).Build()
	return t.NewOrder(order)
}

func (t *tradingREST) SendLimitOrder(accountId uint64, clOrdId string, symbol string, side Side, price string, orderQty string) (*xmsg.ExecutionReport, error) {
	order := CreateLimitOrder(clOrdId, symbol, side, orderQty, accountId, price).Build()
	return t.NewOrder(order)
}

func (t *tradingREST) SendStopOrder(accountId uint64, clOrdId string, symbol string, side Side, stopPx string, orderQty string) (*xmsg.ExecutionReport, error) {
	order := CreateStopOrder(clOrdId, symbol, side, orderQty, accountId, stopPx).Build()
	return t.NewOrder(order)
}

func (t *tradingREST) SendMarketIfTouchOrder(accountId uint64, clOrdId string, symbol string, side Side, stopPx string, orderQty string) (*xmsg.ExecutionReport, error) {
	order := CreateMarketIfTouchOrder(clOrdId, symbol, side, orderQty, accountId, stopPx).Build()
	return t.NewOrder(order)
}

func (t *tradingREST) SendCancel(cmd *xmsg.OrderCancelRequest) (*xmsg.ExecutionReport, error) {
	query := newQuery("order", "cancel")
	er := new(xmsg.ExecutionReport)
	return er, t.sendPost(query, cmd, er)
}

func (t *tradingREST) SendMassCancelCmd(cmd *xmsg.OrderMassCancelRequest) (*xmsg.ExecutionReport, error) {
	query := newQuery("order", "mass-cancel")
	er := new(xmsg.ExecutionReport)
	return er, t.sendPost(query, cmd, er)
}

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

func (t *tradingREST) SendMassCancel(account uint64, clOrdId string) (*xmsg.ExecutionReport, error) {
	return t.SendMassCancelCmd(newOrderMassCancel(account, clOrdId).Build())
}

func (t *tradingREST) SendReplace(request xmsg.OrderCancelReplaceRequest) (*xmsg.ExecutionReport, error) {
	query := newQuery("order", "replace")
	resp := new(xmsg.ExecutionReport)
	return resp, t.sendPost(query, request, resp)
}

func (t *tradingREST) GetAccounts() ([]*xmsg.AccountInfo, error) {
	query := newQuery("accounts")
	resp := struct {
		Accounts []*xmsg.AccountInfo `json:"accounts"`
	}{
		Accounts: make([]*xmsg.AccountInfo, 0),
	}
	if err := t.sendGet(query, &resp); err != nil {
		return nil, err
	}
	return resp.Accounts, nil
}

func (t *tradingREST) GetBalance(accountId uint64) (*xmsg.BalanceSnapshotRefresh, error) {
	query := newQuery("accounts", strconv.FormatUint(accountId, 10), "balance")
	resp := new(xmsg.BalanceSnapshotRefresh)
	return resp, t.sendGet(query, resp)
}

func (t *tradingREST) GetMarginRequirements(accountId uint64) (*xmsg.MarginRequirementReport, error) {
	query := newQuery("accounts", strconv.FormatUint(accountId, 10), "margin-requirements")
	resp := new(xmsg.MarginRequirementReport)
	return resp, t.sendGet(query, resp)
}

func (t *tradingREST) GetOrder(account uint64, opt OrderRequest) (*xmsg.ExecutionReport, error) {
	query := newQuery("accounts", strconv.FormatUint(account, 10), "order")
	query.addQueryf("client_order_id", opt.ClOrdId).
		addQueryf("order_id", opt.OrderId)

	resp := new(xmsg.ExecutionReport)
	return resp, t.sendGet(query, resp)
}

func (t *tradingREST) GetActiveOrders(account uint64, opt ActiveOrdersRequest) ([]*xmsg.ExecutionReport, error) {
	query := newQuery("accounts", strconv.FormatUint(account, 10), "active-orders")
	query.addQueryf("symbol", opt.Symbol)
	resp := make([]*xmsg.ExecutionReport, 0)
	return resp, t.sendGet(query, &resp)
}

func (t *tradingREST) GetLastOrderStatuses(accountId uint64, opt LastOrderStatusesRequest) ([]*xmsg.ExecutionReport, error) {
	query := newQuery("accounts", strconv.FormatUint(accountId, 10), "last-order-statuses")
	query.addQueryf("symbol", opt.Symbol).
		addQueryf("from", opt.From).
		addQueryf("to", opt.To).
		addQueryf("page", opt.PageNumber).
		addQueryf("limit", opt.Limit)

	resp := make([]*xmsg.ExecutionReport, 0)
	return resp, t.sendGet(query, &resp)
}

func (t *tradingREST) GetOrderHistory(accountId uint64, opt OrderHistoryRequest) ([]*xmsg.ExecutionReport, error) {
	query := newQuery("accounts", strconv.FormatUint(accountId, 10), "order-history")
	query.addQueryf("symbol", opt.Symbol).
		addQueryf("client_order_id", opt.ClOrdId).
		addQueryf("order_id", opt.OrderId).
		addQueryf("from", opt.From).
		addQueryf("to", opt.To).
		addQueryf("page", opt.PageNumber).
		addQueryf("limit", opt.Limit)

	resp := make([]*xmsg.ExecutionReport, 0)
	return resp, t.sendGet(query, &resp)
}

func (t *tradingREST) GetTradeHistory(accountId uint64, opt TradeHistoryRequest) ([]*xmsg.ExecutionReport, error) {
	query := newQuery("accounts", strconv.FormatUint(accountId, 10), "trade-history")
	query.addQueryf("trade_id", opt.TradeId).
		addQueryf("client_order_id", opt.ClOrdId).
		addQueryf("symbol", opt.Symbol).
		addQueryf("from", opt.From).
		addQueryf("to", opt.To).
		addQueryf("page", opt.PageNumber).
		addQueryf("limit", opt.Limit)

	resp := make([]*xmsg.ExecutionReport, 0)
	return resp, t.sendGet(query, &resp)
}

func (t *tradingREST) GetPositions(account uint64, opt string) ([]*xmsg.PositionReport, error) {
	query := newQuery("accounts", strconv.FormatUint(account, 10), "positions")
	resp := make([]*xmsg.PositionReport, 0)
	return resp, t.sendGet(query, &resp)
}
func (t *tradingREST) GetPositionsHistory(accountId uint64, opt PositionsHistoryRequest) ([]*xmsg.PositionReport, error) {
	query := newQuery("accounts", strconv.FormatUint(accountId, 10), "positions-history")
	query.addQueryf("id", opt.PositionId)
	query.addQueryf("parentid", opt.ParentPositionId)
	query.addQueryf("symbol", opt.Symbol)
	query.addQueryf("openfrom", opt.OpenFrom)
	query.addQueryf("opento", opt.OpenTo)
	query.addQueryf("closefrom", opt.CloseFrom)
	query.addQueryf("closeto", opt.CloseTo)
	query.addQueryf("page", opt.PageNumber)
	query.addQueryf("limit", opt.Limit)

	resp := make([]*xmsg.PositionReport, 0)
	return resp, t.sendGet(query, &resp)
}

func (t *tradingREST) SendCollapsePositions(account uint64, symbol string, requestId string) (*xmsg.PositionMaintenanceReport, error) {
	request := xmsg.PositionMaintenanceRequest{
		MsgType:        xmsg.MsgType_PositionMaintenanceRequest,
		Account:        account,
		Symbol:         symbol,
		PosReqId:       requestId,
		PosTransType:   PosTransTypeCollapse,
		PosMaintAction: PosMaintActionReplace,
	}
	query := newQuery("position", "maintenance")
	resp := new(xmsg.PositionMaintenanceReport)
	return resp, t.sendPost(query, request, resp)
}

func (t *tradingREST) SendApplicationHeartbeat(groupId string, heartbeatInSec int32) error {
	cmd := xmsg.ApplicationHeartbeat{
		MsgType:    xmsg.MsgType_ApplicationHeartbeat,
		GrpID:      groupId,
		HeartBtInt: heartbeatInSec,
	}
	query := newQuery("order", "heartbeat")
	return t.sendPostWithoutResp(query, &cmd)
}
