package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/xenaex/client-go/xena"
	"github.com/xenaex/client-go/xena/xmsg"
)

var client xena.TradingClient
var accounts []*xmsg.AccountInfo
var accountId uint64
var symbol = xena.XBTUSD.String()
var bestAsk, bestBid = 10000.0, 10000.0

func main() {
	log.Printf("Start")

	apiKey := os.Getenv("XENA_API_KEY")
	apiSecret := os.Getenv("XENA_API_SECRET")

	if len(apiKey) == 0 || len(apiSecret) == 0 {
		log.Println("api key or api secret not found.")
		return
	}

	var err error
	restClient := xena.NewTradingREST(apiKey, apiSecret, xena.WithRestTradingHost)
	accounts, err = restClient.GetAccounts()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(accounts)
	if len(accounts) == 0 || accounts == nil {
		fmt.Println("account not founds")
		return
	}
	indexAccountId := rand.Intn(len(accounts) - 1)
	fmt.Println(indexAccountId)
	accountId = accounts[indexAccountId].Id

	client = xena.NewTradingClient(
		apiKey,
		apiSecret,
		xena.WithTradingURL(),
		xena.WithDebug(),
	)
	resp, err := client.ConnectAndLogon()
	if err != nil {
		log.Printf("loggon err: %s\n", err)
		return
	}
	log.Printf("resp: %s\n", resp)
	if len(resp.RejectText) > 0 {
		return
	}

	bestAsk, bestBid = GetBests(symbol)

	examples := make(map[string]func())
	examples["market_order"] = exampleMarketOrder
	examples["limit-order"] = exampleLimitOrder
	examples["limit_order_post_only"] = exampleLimitOrderPostOnly
	examples["stop-order"] = exampleStopOrder
	examples["sltp-group"] = exampleSLTPGroup
	examples["stop-loss-for-existing-position"] = exampleStopLossForExistingPosition
	examples["take-profit-for-existing-position"] = exampleTakeProfitForExistingPosition
	examples["cancel-by-client-order-id"] = exampleCancelByClientOrderId
	examples["cancel-by-order-id"] = exampleCancelByClientClOrdId
	examples["receiving_all_order_and_canceling"] = exampleReceivingAllOrderAndCanceling
	examples["mass-cancel-1"] = exampleMassCancel1
	examples["mass-cancel-2"] = exampleMassCancel2
	examples["mass-cancel-3"] = exampleMassCancel3
	examples["mass-cancel-4"] = exampleMassCancel4
	examples["replace"] = exampleReplace
	examples["positions-collapse"] = examplePositionsCollapse
	examples["positions"] = examplePositions
	examples["orders"] = exampleOrders
	examples["balances"] = exampleBalances
	examples["margin-requirements"] = exampleMarginRequirements

	log.Println("market_order---------------------------------------------")

	for k, f := range examples {
		for i := 0; i < 4; i++ {
			fmt.Println()
		}
		log.Printf("%s ---------------------------------------------", k)
		f()
	}
}

func GetBests(symbol string) (bestAsk, bestBid float64) {
	log.Println("get dom")
	client := xena.NewMarketDataREST(
		xena.WithRestMarketDataHost,
	)
	dom, err := client.GetDom(symbol)
	if err != nil {
		bestBid = 10000.0
		bestAsk = 10000.0
		return
	}
	bestBid, err = strconv.ParseFloat(dom.BestBid, 10)
	if err != nil {
		bestBid = 10000.0
	}
	bestAsk, err = strconv.ParseFloat(dom.BestAsk, 10)
	if err != nil {
		bestAsk = 10000.0
	}
	return
}

func exampleMarketOrder() {
	defer client.ListenExecutionReport(nil)
	defer client.ListenReject(nil)
	done := make(chan bool)
	var resp *xmsg.ExecutionReport
	var reject *xmsg.Reject
	client.ListenExecutionReport(func(t xena.TradingClient, m *xmsg.ExecutionReport) {
		resp = m
		log.Printf("resp: %s\n", m)
		if m.ExecType == xmsg.ExecType_NewExec {
			done <- true
		}
		if m.ExecType == xmsg.ExecType_RejectedExec || m.ExecType == xmsg.ExecType_CanceledExec {
			done <- false
		}
	})
	client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
		reject = m
		log.Printf("reject: %s\n", m)
		done <- false
	})

	err := client.MarketOrder(accountId, xena.ID("mo-"), symbol, xena.SideBuy, "1")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	log.Printf("is ok %t\n", <-done)

	client.ListenExecutionReport(nil)
	client.ListenReject(nil)

	done = make(chan bool)
	resp = nil
	reject = nil
	client.ListenExecutionReport(func(t xena.TradingClient, m *xmsg.ExecutionReport) {
		resp = m
		log.Printf("resp: %s\n", m)
		if m.ExecType == xmsg.ExecType_NewExec {
			done <- true
		}
		if m.ExecType == xmsg.ExecType_RejectedExec || m.ExecType == xmsg.ExecType_CanceledExec {
			done <- false
		}
	})
	client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
		reject = m
		log.Printf("reject: %s\n", m)
		done <- false
	})

	// or using helpers method
	mOrder := xena.CreateMarketOrder(xena.ID("mo-"), symbol, xena.SideSell, "1", accountId).SetText("my comment").Build()
	err = client.Send(mOrder)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	log.Printf("is ok %t\n", <-done)
}

func exampleLimitOrder() {
	defer client.ListenExecutionReport(nil)
	defer client.ListenReject(nil)
	done := make(chan bool)
	var resp *xmsg.ExecutionReport
	var reject *xmsg.Reject
	client.ListenExecutionReport(func(t xena.TradingClient, m *xmsg.ExecutionReport) {
		resp = m
		log.Printf("resp: %s\n", m)
		if resp.ExecType == xmsg.ExecType_NewExec {
			done <- true
		}
		if resp.ExecType == xmsg.ExecType_RejectedExec || resp.ExecType == xmsg.ExecType_CanceledExec {
			done <- false
		}
	})
	client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
		reject = m
		log.Printf("reject: %s\n", m)
		done <- false
	})

	err := client.LimitOrder(accountId, xena.ID("lo-"), symbol, xena.SideBuy, fmt.Sprintf("%.1f", bestAsk+10), "1")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	log.Printf("is ok %t\n", <-done)

	client.ListenExecutionReport(nil)
	client.ListenReject(nil)

	done = make(chan bool)
	resp = nil
	reject = nil
	client.ListenExecutionReport(func(t xena.TradingClient, m *xmsg.ExecutionReport) {
		resp = m
		log.Printf("resp: %s\n", m)
		if resp.ExecType == xmsg.ExecType_NewExec {
			done <- true
		}
		if resp.ExecType == xmsg.ExecType_RejectedExec || resp.ExecType == xmsg.ExecType_CanceledExec {
			done <- false
		}
	})
	client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
		reject = m
		log.Printf("reject: %s\n", m)
		done <- false
	})

	// or using helpers method
	mOrder := xena.CreateLimitOrder(xena.ID("lo-"), symbol, xena.SideSell, "1", accountId, fmt.Sprintf("%.1f", bestBid-10)).SetText("my comment").Build()
	err = client.Send(mOrder)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	log.Printf("is ok %t\n", <-done)
}

func exampleLimitOrderPostOnly() {
	defer client.ListenExecutionReport(nil)
	defer client.ListenReject(nil)
	done := make(chan bool)
	var resp *xmsg.ExecutionReport
	var reject *xmsg.Reject
	id := xena.ID("limit-order-")
	client.ListenExecutionReport(func(t xena.TradingClient, m *xmsg.ExecutionReport) {
		resp = m
		log.Printf("resp: %s\n", m)
		if resp.ClOrdId != id && resp.OrigClOrdId != id {
			return
		}
		if resp.ExecType == xmsg.ExecType_NewExec {
			done <- true
		}
		if resp.ExecType == xmsg.ExecType_RejectedExec || resp.ExecType == xmsg.ExecType_CanceledExec {
			done <- false
		}
	})
	client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
		log.Printf("resp: %s\n", m)
		done <- false
	})

	mOrder := xena.CreateLimitOrder(id, symbol, xena.SideSell, "1", accountId, fmt.Sprintf("%.1f", bestBid-10)).SetText("my comment").AddExecInst(xmsg.ExecInst_StayOnOfferSide).Build()
	err := client.Send(mOrder)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	log.Printf("is ok %t\n", <-done)

	client.ListenExecutionReport(nil)
	client.ListenReject(nil)

	id = xena.ID("limit-order-")
	done = make(chan bool)
	resp = nil
	reject = nil
	client.ListenExecutionReport(func(t xena.TradingClient, m *xmsg.ExecutionReport) {
		resp = m
		log.Printf("resp1: %s\n", m)
		if resp.ClOrdId != id && resp.OrigClOrdId != id {
			return
		}
		if resp.ExecType == xmsg.ExecType_NewExec {
			done <- true
		}
		if resp.ExecType == xmsg.ExecType_RejectedExec || resp.ExecType == xmsg.ExecType_CanceledExec {
			done <- false
		}
	})
	client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
		reject = m
		log.Printf("reject: %s\n", m)
		done <- false
	})

	// or using helpers method
	mOrder = xena.CreateLimitOrder(id, symbol, xena.SideSell, "1", accountId, fmt.Sprintf("%.1f", bestBid-10)).SetText("my comment").AddExecInst(xmsg.ExecInst_PegToOfferSide).Build()
	err = client.Send(mOrder)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	log.Printf("is ok %t\n", <-done)
}

func exampleStopOrder() {
	defer client.ListenExecutionReport(nil)
	defer client.ListenReject(nil)
	done := make(chan bool)
	var resp *xmsg.ExecutionReport
	var reject *xmsg.Reject
	id := xena.ID("limit-order-")
	client.ListenExecutionReport(func(t xena.TradingClient, m *xmsg.ExecutionReport) {
		resp = m
		log.Printf("resp1: %s\n", m)
		if resp.ClOrdId != id && resp.OrigClOrdId != id {
			return
		}
		if resp.ExecType == xmsg.ExecType_NewExec {
			done <- true
		}
		if resp.ExecType == xmsg.ExecType_RejectedExec || resp.ExecType == xmsg.ExecType_CanceledExec {
			done <- false
		}
	})
	client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
		reject = m
		log.Printf("reject: %s\n", m)
		done <- false
	})

	err := client.StopOrder(accountId, id, symbol, xena.SideBuy, fmt.Sprintf("%.1f", bestAsk+10), "1")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	log.Printf("is ok %t\n", <-done)

	client.ListenExecutionReport(nil)
	client.ListenReject(nil)

	done = make(chan bool)
	id = xena.ID("limit-order-")
	resp = nil
	reject = nil
	client.ListenExecutionReport(func(t xena.TradingClient, m *xmsg.ExecutionReport) {
		resp = m
		log.Printf("resp1: %s\n", m)
		if resp.ClOrdId != id && resp.OrigClOrdId != id {
			return
		}
		if resp.ExecType == xmsg.ExecType_NewExec {
			done <- true
		}
		if resp.ExecType == xmsg.ExecType_RejectedExec || resp.ExecType == xmsg.ExecType_CanceledExec {
			done <- false
		}
	})
	client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
		reject = m
		log.Printf("reject: %s\n", m)
		done <- false
	})

	// or using helpers method
	mOrder := xena.CreateStopOrder(id, symbol, xena.SideSell, "1", accountId, fmt.Sprintf("%.1f", bestBid-10)).SetText("stop order").Build()
	err = client.Send(mOrder)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	log.Printf("is ok %t\n", <-done)
}

func exampleSLTPGroup() {
	defer client.ListenExecutionReport(nil)
	defer client.ListenReject(nil)
	done := make(chan bool)
	var resp *xmsg.ExecutionReport
	var reject *xmsg.Reject
	client.ListenExecutionReport(func(t xena.TradingClient, m *xmsg.ExecutionReport) {
		resp = m
		log.Printf("resp: %s\n", m)
		if resp.ExecType == xmsg.ExecType_NewExec {
			done <- true
		}
		if resp.ExecType == xmsg.ExecType_RejectedExec || resp.ExecType == xmsg.ExecType_CanceledExec {
			done <- false
		}
	})
	client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
		reject = m
		log.Printf("reject: %s\n", m)
		done <- false
	})

	mOrder := xena.CreateLimitOrder(xena.ID("sltp-"), symbol, xena.SideBuy, "1", accountId, fmt.Sprintf("%.1f", bestAsk+10)).
		SetText("my comment").
		SetTakeProfitPrice(fmt.Sprintf("%.1f0000572", bestBid+500)).
		SetStopLossPrice("500.00000572").Build()
	err := client.Send(mOrder)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	log.Printf("is ok %t\n", <-done)

	client.ListenExecutionReport(nil)
	client.ListenReject(nil)

	done = make(chan bool)
	resp = nil
	reject = nil
	client.ListenExecutionReport(func(t xena.TradingClient, m *xmsg.ExecutionReport) {
		resp = m
		log.Printf("resp: %s\n", m)
		if resp.ExecType == xmsg.ExecType_NewExec {
			done <- true
		}
		if resp.ExecType == xmsg.ExecType_RejectedExec || resp.ExecType == xmsg.ExecType_CanceledExec {
			done <- false
		}
	})
	client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
		reject = m
		log.Printf("reject: %s\n", m)
		done <- false
	})

	// or using helpers method
	mOrder = xena.CreateLimitOrder(xena.ID("sltp-"), symbol, xena.SideSell, "1", accountId, fmt.Sprintf("%.1f", bestBid-10)).
		SetText("my comment").
		SetTakeProfitPrice(fmt.Sprintf("%.1f0000572", bestBid+500)).
		SetStopLossPrice("500.00000572").Build()
	err = client.Send(mOrder)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	log.Printf("is ok %t\n", <-done)
}

func exampleStopLossForExistingPosition() {
	positionId := uint64(130723016)

	defer client.ListenExecutionReport(nil)
	defer client.ListenReject(nil)
	done := make(chan bool)
	var resp *xmsg.ExecutionReport
	var reject *xmsg.Reject
	client.ListenExecutionReport(func(t xena.TradingClient, m *xmsg.ExecutionReport) {
		resp = m
		log.Printf("resp: %s\n", m)
		if resp.ExecType == xmsg.ExecType_NewExec {
			done <- true
		}
		if resp.ExecType == xmsg.ExecType_RejectedExec || resp.ExecType == xmsg.ExecType_CanceledExec {
			done <- false
		}
	})
	client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
		reject = m
		log.Printf("reject: %s\n", m)
		done <- false
	})

	mOrder := xena.CreateLimitOrder(xena.ID("stop-order"), symbol, xena.SideBuy, "1", accountId, fmt.Sprintf("%.1f", bestAsk+10)).SetText("my comment").SetPositionId(positionId).Build()
	err := client.Send(mOrder)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	log.Printf("is ok %t\n", <-done)

	client.ListenExecutionReport(nil)
	client.ListenReject(nil)

	done = make(chan bool)
	resp = nil
	reject = nil
	client.ListenExecutionReport(func(t xena.TradingClient, m *xmsg.ExecutionReport) {
		resp = m
		log.Printf("resp: %s\n", m)
		if resp.ExecType == xmsg.ExecType_NewExec {
			done <- true
		}
		if resp.ExecType == xmsg.ExecType_RejectedExec || resp.ExecType == xmsg.ExecType_CanceledExec {
			done <- false
		}
	})
	client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
		reject = m
		log.Printf("reject: %s\n", m)
		done <- false
	})

	mOrder = xena.CreateStopOrder(xena.ID("stop-order"), symbol, xena.SideSell, "1", accountId, fmt.Sprintf("%.1f", bestBid-10)).SetText("my comment").SetPositionId(positionId).Build()
	err = client.Send(mOrder)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	log.Printf("is ok %t\n", <-done)
}

func exampleTakeProfitForExistingPosition() {
	positionId := uint64(130723016)

	defer client.ListenExecutionReport(nil)
	defer client.ListenReject(nil)
	done := make(chan bool)
	var resp *xmsg.ExecutionReport
	var reject *xmsg.Reject
	client.ListenExecutionReport(func(t xena.TradingClient, m *xmsg.ExecutionReport) {
		resp = m
		log.Printf("resp: %s\n", m)
		if resp.ExecType == xmsg.ExecType_NewExec {
			done <- true
		}
		if resp.ExecType == xmsg.ExecType_RejectedExec || resp.ExecType == xmsg.ExecType_CanceledExec {
			done <- false
		}
	})
	client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
		reject = m
		log.Printf("reject: %s\n", m)
		done <- false
	})

	mOrder := xena.CreateLimitOrder(xena.ID("limit-order"), symbol, xena.SideSell, "1", accountId, fmt.Sprintf("%.1f", bestBid-10)).SetText("my comment").SetPositionId(positionId).Build()
	err := client.Send(mOrder)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	log.Printf("is ok %t\n", <-done)
}

func exampleCancelByClientOrderId() {
	ids := []string{xena.ID("limit-order-1"), xena.ID("limit-order-2"), xena.ID("limit-order-3")}

	defer client.ListenExecutionReport(nil)
	defer client.ListenReject(nil)
	done := make(chan bool, 1)
	var reject *xmsg.Reject
	wg := &sync.WaitGroup{}
	client.ListenExecutionReport(func(t xena.TradingClient, m *xmsg.ExecutionReport) {
		if m.ExecType == xmsg.ExecType_NewExec {
			for _, id := range ids {
				if m.ClOrdId == id {
					err := client.CancelByOrderId(accountId, xena.ID("cancel-1"), m.OrderId, symbol, xena.SideBuy)
					// fmt.Printf("resp: %s\n", cancelResp)
					if err != nil {
						fmt.Printf("error: %v\n", err)
					}
					log.Printf("exec resp: %s\n", m)
				}
			}
		}
		if m.ExecType == xmsg.ExecType_CanceledExec {
			for _, id := range ids {
				if m.OrigClOrdId == id {
					log.Printf("cancel resp: %s  %v\n", m, *wg)
					wg.Done()
				}
			}
		}
		if m.ExecType == xmsg.ExecType_RejectedExec {
			for _, id := range ids {
				if m.OrigClOrdId == id {
					log.Printf("reject resp: %s\n", m)
					wg.Done()
				}
			}
		}
	})
	client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
		reject = m
		log.Printf("reject: %s\n", m)
		done <- false
	})
	for _, id := range ids {
		err := client.LimitOrder(accountId, id, symbol, xena.SideBuy, fmt.Sprintf("%.1f", bestAsk-1), "1")
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}
		wg.Add(1)
	}
	wg.Wait()
	done <- true
	log.Printf("is ok %t", <-done)
}

func exampleCancelByClientClOrdId() {
	ids := []string{xena.ID("limit-order-1"), xena.ID("limit-order-2"), xena.ID("limit-order-3")}

	defer client.ListenExecutionReport(nil)
	defer client.ListenReject(nil)
	done := make(chan bool, 1)
	var reject *xmsg.Reject
	wg := &sync.WaitGroup{}
	client.ListenExecutionReport(func(t xena.TradingClient, m *xmsg.ExecutionReport) {
		if m.ExecType == xmsg.ExecType_NewExec {
			for _, id := range ids {
				if m.ClOrdId == id {
					err := client.CancelByClOrdId(accountId, xena.ID("cancel-1"), m.ClOrdId, symbol, xena.SideBuy)
					// fmt.Printf("resp: %s\n", cancelResp)
					if err != nil {
						fmt.Printf("error: %v\n", err)
					}
					log.Printf("exec resp: %s\n", m)
				}
			}
		}
		if m.ExecType == xmsg.ExecType_CanceledExec {
			for _, id := range ids {
				if m.OrigClOrdId == id {
					log.Printf("cancel resp: %s  %v\n", m, *wg)
					wg.Done()
				}
			}
		}
		if m.ExecType == xmsg.ExecType_RejectedExec {
			for _, id := range ids {
				if m.OrigClOrdId == id {
					log.Printf("reject resp: %s\n", m)
					wg.Done()
				}
			}
		}
	})
	client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
		reject = m
		log.Printf("reject: %s\n", m)
		done <- false
	})
	for _, id := range ids {
		err := client.LimitOrder(accountId, id, symbol, xena.SideBuy, fmt.Sprintf("%.1f", bestAsk-1), "1")
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}
		wg.Add(1)
	}
	wg.Wait()
	done <- true
	log.Printf("is ok %t", <-done)
}

func exampleReceivingAllOrderAndCanceling() {
	defer client.ListenOrderMassStatusResponse(nil)
	defer client.ListenExecutionReport(nil)
	defer client.ListenReject(nil)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	resps := make([]*xmsg.ExecutionReport, 0)

	client.ListenExecutionReport(func(t xena.TradingClient, m *xmsg.ExecutionReport) {
		if m.ExecType == xmsg.ExecType_CanceledExec || m.ExecType == xmsg.ExecType_RejectedExec {
			wg.Done()
		}
		if m.ExecType == xmsg.ExecType_CanceledExec {
			resps = append(resps, m)
		}
		fmt.Printf("resp: %v\n", m)
	})
	client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
		fmt.Printf("reject resp: %v\n", m)
		done <- false
	})
	client.ListenOrderMassStatusResponse(func(t xena.TradingClient, m *xmsg.OrderMassStatusResponse) {
		wg.Add(len(m.Orders))
		for _, er := range m.Orders {
			cancelCmd := xena.CreateCancelRequestFromExecutionReport(xena.ID("cancel-"), er)
			err := client.Cancel(cancelCmd)
			if err != nil {
				fmt.Printf("error: %v\n", err)
			}
		}
		done <- true
	})
	err := client.Orders(accountId, "")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	ok := <-done
	log.Printf("is ok %t\n", ok)
	if !ok {
		return
	}
	wg.Wait()
	log.Printf("%d orders were cancelled", len(resps))
}

func exampleMassCancel1() {
	defer client.ListenOrderMassCancelReport(nil)
	defer client.ListenExecutionReport(nil)
	defer client.ListenReject(nil)
	done := make(chan bool)
	resps := make([]*xmsg.ExecutionReport, 0)
	var respReport *xmsg.OrderMassCancelReport
	var reject *xmsg.Reject
	client.ListenOrderMassCancelReport(func(t xena.TradingClient, m *xmsg.OrderMassCancelReport) {
		respReport = m
		log.Printf("resp: %s\n", m)
		time.Sleep(time.Second)
		done <- true
	})
	client.ListenExecutionReport(func(t xena.TradingClient, m *xmsg.ExecutionReport) {
		resps = append(resps, m)
	})
	client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
		reject = m
		log.Printf("reject: %s\n", m)
		done <- false
	})

	err := client.MassCancel(accountId, xena.ID("mass-cancel-1-"))
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	log.Printf("is ok %t\n", <-done)
	for _, cancel := range resps {
		log.Printf("%s was cancel \n", cancel.ClOrdId)
	}
}

func exampleMassCancel2() {
	defer client.ListenOrderMassCancelReport(nil)
	defer client.ListenExecutionReport(nil)
	defer client.ListenReject(nil)
	done := make(chan bool)
	resps := make([]*xmsg.ExecutionReport, 0)
	var respReport *xmsg.OrderMassCancelReport
	var reject *xmsg.Reject
	client.ListenOrderMassCancelReport(func(t xena.TradingClient, m *xmsg.OrderMassCancelReport) {
		respReport = m
		log.Printf("resp: %s\n", m)
		time.Sleep(time.Second)
		done <- true
	})
	client.ListenExecutionReport(func(t xena.TradingClient, m *xmsg.ExecutionReport) {
		resps = append(resps, m)
	})
	client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
		reject = m
		log.Printf("reject: %s\n", m)
		done <- false
	})

	massCancel := xena.CreateOrderMassCancel(accountId, xena.ID("mass-cancel-2-")).SetSymbol(symbol)
	err := client.Send(massCancel.Build())
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	log.Printf("is ok %t\n", <-done)
	for _, cancel := range resps {
		log.Printf("%s was cancel \n", cancel.ClOrdId)
	}
}

func exampleMassCancel3() {
	defer client.ListenOrderMassCancelReport(nil)
	defer client.ListenExecutionReport(nil)
	defer client.ListenReject(nil)
	done := make(chan bool)
	resps := make([]*xmsg.ExecutionReport, 0)
	var respReport *xmsg.OrderMassCancelReport
	var reject *xmsg.Reject
	client.ListenOrderMassCancelReport(func(t xena.TradingClient, m *xmsg.OrderMassCancelReport) {
		respReport = m
		log.Printf("resp: %s\n", m)
		time.Sleep(time.Second)
		done <- true
	})
	client.ListenExecutionReport(func(t xena.TradingClient, m *xmsg.ExecutionReport) {
		resps = append(resps, m)
	})
	client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
		reject = m
		log.Printf("reject: %s\n", m)
		done <- false
	})

	massCancel := xena.CreateOrderMassCancel(accountId, xena.ID("mass-cancel-3-")).SetSymbol(symbol).SetSide(xena.SideBuy)
	err := client.Send(massCancel.Build())
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	log.Printf("is ok %t\n", <-done)
	for _, cancel := range resps {
		log.Printf("%s was cancel \n", cancel.ClOrdId)
	}
}

func exampleMassCancel4() {
	defer client.ListenOrderMassCancelReport(nil)
	defer client.ListenExecutionReport(nil)
	defer client.ListenReject(nil)
	done := make(chan bool)
	resps := make([]*xmsg.ExecutionReport, 0)
	var respReport *xmsg.OrderMassCancelReport
	var reject *xmsg.Reject
	client.ListenOrderMassCancelReport(func(t xena.TradingClient, m *xmsg.OrderMassCancelReport) {
		respReport = m
		log.Printf("resp: %s\n", m)
		time.Sleep(time.Second)
		done <- true
	})
	client.ListenExecutionReport(func(t xena.TradingClient, m *xmsg.ExecutionReport) {
		resps = append(resps, m)
	})
	client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
		reject = m
		log.Printf("reject: %s\n", m)
		done <- false
	})

	massCancel := xena.CreateOrderMassCancel(accountId, xena.ID("mass-cancel-4-")).SetSymbol(symbol).SetPositionEffect(xena.PositionEffectOpen)
	err := client.Send(massCancel.Build())
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	log.Printf("is ok %t\n", <-done)
	for _, cancel := range resps {
		log.Printf("%s was cancel \n", cancel.ClOrdId)
	}
}

func exampleReplace() {
	id := xena.ID("limit-order-")
	defer client.ListenExecutionReport(nil)
	defer client.ListenReject(nil)
	done := make(chan bool)
	var resp *xmsg.ExecutionReport
	var reject *xmsg.Reject
	client.ListenExecutionReport(func(t xena.TradingClient, m *xmsg.ExecutionReport) {
		resp = m
		log.Printf("resp: %s\n", m)
		if resp.ClOrdId == id {
			if resp.ExecType == xmsg.ExecType_PendingNewExec {
				log.Printf("order %s is panding new\n", resp.ClOrdId)
				return
			}
			if resp.ExecType == xmsg.ExecType_RejectedExec {
				log.Printf("order %s is panding new\n", resp.ClOrdId)
				done <- false
				return
			}
			replace := xena.CreateReplace(xena.ID("replace-"), resp)
			replace.OrderQty = "10"
			err := client.Replace(replace)
			log.Printf("resp: %s\n", resp)
			if err != nil {
				log.Printf("error: %v\n", err)
			}
			return
		}
		if strings.Contains(resp.ClOrdId, "replace-") {
			if resp.ExecType == xmsg.ExecType_ReplacedExec {
				done <- true
			}
		}
	})
	client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
		reject = m
		log.Printf("reject: %s\n", m)
		done <- false
	})

	err := client.LimitOrder(accountId, id, symbol, xena.SideBuy, fmt.Sprintf("%.1f", bestAsk-1), "1")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	log.Printf("is ok %t\n", <-done)

}

func examplePositionsCollapse() {
	for _, acc := range accounts {
		if xena.IsMargin(acc.Id) {
			func() {
				defer client.ListenMarginRequirementReport(nil)
				defer client.ListenReject(nil)
				done := make(chan bool)
				var resp *xmsg.MarginRequirementReport
				var reject *xmsg.Reject
				client.ListenMarginRequirementReport(func(t xena.TradingClient, m *xmsg.MarginRequirementReport) {
					resp = m
					log.Printf("resp: %s\n", m)
					done <- true
				})
				client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
					reject = m
					log.Printf("reject: %s\n", m)
					done <- false
				})

				err := client.CollapsePositions(acc.Id, symbol, "")
				if err != nil {
					fmt.Printf("error: %v, account: %d, symbol: %s\n", err, acc.Id, acc.Currency)
				}

				log.Printf("is ok %t\n", <-done)
			}()
		}
	}
}

func examplePositions() {
	for _, acc := range accounts {
		if xena.IsMargin(acc.Id) {
			func() {
				defer client.ListenMassPositionReport(nil)
				defer client.ListenReject(nil)
				done := make(chan bool)
				var resp *xmsg.MassPositionReport
				var reject *xmsg.Reject
				client.ListenMassPositionReport(func(t xena.TradingClient, m *xmsg.MassPositionReport) {
					resp = m
					log.Printf("resp: %s\n", m)
					done <- true
				})
				client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
					reject = m
					log.Printf("reject: %s\n", m)
					done <- false
				})

				err := client.GetPositions(acc.Id, "")
				if err != nil {
					fmt.Printf("error: %v, account: %d, symbol: %s\n", err, acc.Id, acc.Currency)
				}

				log.Printf("is ok %t\n", <-done)
			}()
		}
	}
}

func exampleOrders() {
	for _, acc := range accounts {
		if xena.IsMargin(acc.Id) {
			func() {
				defer client.ListenOrderMassStatusResponse(nil)
				defer client.ListenReject(nil)
				done := make(chan bool)
				var resp *xmsg.OrderMassStatusResponse
				var reject *xmsg.Reject
				client.ListenOrderMassStatusResponse(func(t xena.TradingClient, m *xmsg.OrderMassStatusResponse) {
					resp = m
					log.Printf("resp: %s\n", m)
					done <- true
				})
				client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
					reject = m
					log.Printf("reject: %s\n", m)
					done <- false
				})

				err := client.Orders(acc.Id, "")
				if err != nil {
					fmt.Printf("error: %v, account: %d, symbol: %s\n", err, acc.Id, acc.Currency)
				}

				log.Printf("is ok %t\n", <-done)
			}()
		}
	}
}

func exampleBalances() {
	for _, acc := range accounts {
		if xena.IsMargin(acc.Id) {
			func() {
				defer client.ListenMarginRequirementReport(nil)
				defer client.ListenBalanceSnapshotRefresh(nil)
				defer client.ListenReject(nil)
				done := make(chan bool)
				var marginRequirement *xmsg.MarginRequirementReport
				var balanceSnapshot *xmsg.BalanceSnapshotRefresh
				var reject *xmsg.Reject
				client.ListenMarginRequirementReport(func(t xena.TradingClient, m *xmsg.MarginRequirementReport) {
					marginRequirement = m
					log.Printf("resp: %s\n", m)
					done <- true
				})
				client.ListenBalanceSnapshotRefresh(func(t xena.TradingClient, m *xmsg.BalanceSnapshotRefresh) {
					balanceSnapshot = m
					log.Printf("resp: %s\n", m)
					done <- true
				})
				client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
					reject = m
					log.Printf("reject: %s\n", m)
					done <- false
				})

				err := client.AccountStatusReport(acc.Id, "")
				if err != nil {
					fmt.Printf("error: %v, account: %d, symbol: %s\n", err, acc.Id, acc.Currency)
				}

				log.Printf("is ok %t\n", <-done)
				log.Printf("is ok %t\n", <-done)
			}()
		}
	}
}

func exampleMarginRequirements() {
	for _, acc := range accounts {
		if xena.IsMargin(acc.Id) {
			func() {
				defer client.ListenOrderMassStatusResponse(nil)
				defer client.ListenReject(nil)
				done := make(chan bool)
				var reject *xmsg.Reject
				client.ListenOrderMassStatusResponse(func(t xena.TradingClient, m *xmsg.OrderMassStatusResponse) {
					log.Printf("resp: %s\n", m)
					done <- true
				})
				client.ListenReject(func(t xena.TradingClient, m *xmsg.Reject) {
					reject = m
					log.Printf("reject: %s\n", m)
					done <- false
				})

				err := client.Orders(acc.Id, "")
				if err != nil {
					fmt.Printf("error: %v, account: %d, symbol: %s\n", err, acc.Id, acc.Currency)
				}
				log.Printf("is ok %t\n", <-done)
			}()
		}
	}
}
