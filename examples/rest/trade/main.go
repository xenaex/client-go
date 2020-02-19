package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/xenaex/client-go/xena"
	"github.com/xenaex/client-go/xena/xmsg"
)

func main() {
	apiKey := os.Getenv("XENA_API_KEY")
	apiSecret := os.Getenv("XENA_API_SECRET")

	if len(apiKey) == 0 || len(apiSecret) == 0 {
		log.Println("api key or api secret not found.")
		return
	}
	client := xena.NewTradingREST(apiKey, apiSecret, xena.WithRestTradingHost)
	accounts, err := client.GetAccounts()
	if err != nil {
		log.Println(err)
	}

	symbol := xena.XBTUSD.String()

	if len(accounts) == 0 || accounts == nil {
		log.Println("account not founds")
		return
	}
	indexAccountId := rand.Intn(len(accounts) - 1)
	log.Println(indexAccountId)
	var accountId = accounts[indexAccountId].Id
	bestAsk, bestBid := GetBests(symbol)
	log.Printf("bestAsk: %f, bestBid: %f\n", bestAsk, bestBid)

	if strings.Contains(os.Args[0], "/") && strings.Contains(os.Args[0], "main") {
		os.Args = os.Args[1:]
	}
	examples := make(map[string]func())
	examples["market-order"] = func() {
		resp, err := client.SendMarketOrder(accountId, xena.ID("mo-"), symbol, xena.SideBuy, "1")
		log.Printf("resp: %s\n", resp)
		if err != nil {
			log.Printf("error: %v\n", err)
		}

		// or using helpers method
		mOrder := xena.CreateMarketOrder(xena.ID("mo-"), symbol, xena.SideSell, "1", accountId).SetText("my comment").Build()
		resp, err = client.NewOrder(mOrder)
		log.Printf("resp: %s\n", resp)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
	}
	examples["limit-order"] = func() {
		resp, err := client.SendLimitOrder(accountId, xena.ID("lo-"), symbol, xena.SideBuy, fmt.Sprintf("%.1f", bestAsk+10), "1")
		log.Printf("resp: %s\n", resp)
		if err != nil {
			log.Printf("error: %v\n", err)
		}

		// or using helpers method
		mOrder := xena.CreateLimitOrder(xena.ID("lo-"), symbol, xena.SideSell, "1", accountId, fmt.Sprintf("%.1f", bestBid-10)).SetText("my comment").Build()
		resp, err = client.NewOrder(mOrder)
		log.Printf("resp: %s\n", resp)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
	}
	examples["stop-order"] = func() {
		resp, err := client.SendStopOrder(accountId, xena.ID("so-"), symbol, xena.SideBuy, fmt.Sprintf("%.1f", bestAsk+10), "1")
		log.Printf("resp: %s\n", resp)
		if err != nil {
			log.Printf("error: %v\n", err)
		}

		// or using helpers method
		mOrder := xena.CreateStopOrder(xena.ID("so-"), symbol, xena.SideSell, "1", accountId, fmt.Sprintf("%.1f", bestBid-10)).SetText("stop order").Build()
		resp, err = client.NewOrder(mOrder)
		log.Printf("resp: %s\n", resp)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
	}
	examples["sltp-group"] = func() {
		// example_of_sltp_group
		mOrder := xena.CreateLimitOrder(xena.ID("sltp-"), symbol, xena.SideBuy, "1", accountId, fmt.Sprintf("%.1f", bestAsk+10)).
			SetText("my comment").
			SetTakeProfitPrice(fmt.Sprintf("%.1f0000572", bestBid+500)).
			SetStopLossPrice("500.00000572").Build()
		resp, err := client.NewOrder(mOrder)
		log.Printf("resp: %s\n", resp)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
		// or using helpers method
		mOrder = xena.CreateLimitOrder(xena.ID("sltp-"), symbol, xena.SideSell, "1", accountId, fmt.Sprintf("%.1f", bestBid-10)).
			SetText("my comment").
			SetTakeProfitPrice(fmt.Sprintf("%.1f0000572", bestBid+500)).
			SetStopLossPrice("500.00000572").Build()
		resp, err = client.NewOrder(mOrder)
		log.Printf("resp: %s\n", resp)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
	}
	examples["stop-loss-for-existing-position"] = func() {
		// example_of_stop_loss_for_existing_position
		positionId := uint64(130723016)
		mOrder := xena.CreateLimitOrder(xena.ID("stop-order"), symbol, xena.SideBuy, "1", accountId, fmt.Sprintf("%.1f", bestAsk+10)).SetText("my comment").SetPositionId(positionId).Build()
		resp, err := client.NewOrder(mOrder)
		log.Printf("resp: %s\n", resp)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
		// or using helpers method
		mOrder = xena.CreateStopOrder(xena.ID("stop-order"), symbol, xena.SideSell, "1", accountId, fmt.Sprintf("%.1f", bestBid-10)).SetText("my comment").SetPositionId(positionId).Build()
		resp, err = client.NewOrder(mOrder)
		log.Printf("resp: %s\n", resp)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
	}
	examples["take-profit-for-existing-position"] = func() {
		// example_of_take_profit_for_existing_position
		positionId := uint64(130723016)
		mOrder := xena.CreateLimitOrder(xena.ID("limit-order"), symbol, xena.SideSell, "1", accountId, fmt.Sprintf("%.1f", bestBid-10)).SetText("my comment").SetPositionId(positionId).Build()
		resp, err := client.NewOrder(mOrder)
		log.Printf("resp: %s\n", resp)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
	}
	examples["cancel-by-client-order-id"] = func() {
		// example_of_cancel
		ids := []string{xena.ID("limit-order-1"), xena.ID("limit-order-2"), xena.ID("limit-order-3")}
		resps := make([]*xmsg.ExecutionReport, 0)
		for _, id := range ids {
			resp, err := client.SendLimitOrder(accountId, id, symbol, xena.SideBuy, fmt.Sprintf("%.1f", bestAsk-1), "1")
			log.Printf("resp: %s\n", resp)
			if err != nil {
				log.Printf("error: %v\n", err)
			}
			if err == nil {
				resps = append(resps, resp)
			}
		}
		for i, resp := range resps {
			if resp.ClOrdId == ids[i] {
				cancelResp, err := client.SendCancelByClOrdId(accountId, xena.ID("cancel-1"), resp.ClOrdId, symbol, xena.SideBuy)
				log.Printf("resp: %s\n", cancelResp)
				if err != nil {
					log.Printf("error: %v\n", err)
				}
			}
		}
	}
	examples["cancel-by-order-id"] = func() {
		ids := []string{xena.ID("limit-order-1"), xena.ID("limit-order-2"), xena.ID("limit-order-3")}
		resps := make([]*xmsg.ExecutionReport, 0)
		for _, id := range ids {
			resp, err := client.SendLimitOrder(accountId, id, symbol, xena.SideBuy, fmt.Sprintf("%.1f", bestAsk-1), "1")
			log.Printf("resp: %s\n", resp)
			if err != nil {
				log.Printf("error: %v\n", err)
			}
			if err == nil {
				resps = append(resps, resp)
			}
		}
		for i, resp := range resps {
			if resp.ClOrdId == ids[i] {
				cancelResp, err := client.SendCancelByOrderId(accountId, xena.ID("cancel-1"), resp.OrderId, symbol, xena.SideBuy)
				log.Printf("resp: %s\n", cancelResp)
				if err != nil {
					log.Printf("error: %v\n", err)
				}
			}
		}
	}
	examples["mass-cancel-1"] = func() {
		// example_of_mass_cancel
		massCancel := xena.CreateOrderMassCancel(accountId, xena.ID("mass-cancel-1-"))
		resp, err := client.SendMassCancelCmd(massCancel.Build())
		log.Printf("resp: %s\n", resp)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
	}
	examples["mass-cancel-2"] = func() {
		massCancel := xena.CreateOrderMassCancel(accountId, xena.ID("mass-cancel-2-")).SetSymbol(symbol)
		resp, err := client.SendMassCancelCmd(massCancel.Build())
		log.Printf("resp: %s\n", resp)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
	}

	examples["mass-cancel-3"] = func() {
		massCancel := xena.CreateOrderMassCancel(accountId, xena.ID("mass-cancel-3-")).SetSymbol(symbol).SetSide(xena.SideBuy)
		resp, err := client.SendMassCancelCmd(massCancel.Build())
		log.Printf("resp: %s\n", resp)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
	}
	examples["mass-cancel-4"] = func() {
		massCancel := xena.CreateOrderMassCancel(accountId, xena.ID("mass-cancel-4-")).SetSymbol(symbol).SetPositionEffect(xena.PositionEffectOpen)
		resp, err := client.SendMassCancelCmd(massCancel.Build())
		log.Printf("resp: %s\n", resp)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
	}
	examples["replace"] = func() {
		resp, err := client.SendLimitOrder(accountId, xena.ID("limit-order"), symbol, xena.SideBuy, fmt.Sprintf("%.1f", bestAsk-1), "1")
		log.Printf("resp: %s\n", resp)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
		replace := xena.CreateReplace(xena.ID(""), resp)
		replace.OrderQty = "10"
		resp, err = client.SendReplace(replace)
		log.Printf("resp: %s\n", resp)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
	}
	examples["positions-collapse"] = func() {
		for _, acc := range accounts {
			if xena.IsMargin(acc.Id) {
				respColl, err := client.SendCollapsePositions(acc.Id, symbol, "")
				log.Printf("resp: %s\n", respColl)
				if err != nil {
					log.Printf("error: %v, account: %d, symbol: %s\n", err, acc.Id, acc.Currency)
				}
			}
		}
	}
	examples["positions"] = func() {
		for _, acc := range accounts {
			if xena.IsMargin(acc.Id) {
				respColl, err := client.GetPositions(acc.Id, "")
				log.Printf("resp: %s\n", respColl)
				if err != nil {
					log.Printf("error: %v, account: %d, symbol: %s\n", err, acc.Id, acc.Currency)
				}
			}
		}
	}
	examples["positions-history"] = func() {
		for _, acc := range accounts {
			if xena.IsMargin(acc.Id) {
				respColl, err := client.GetPositionsHistory(acc.Id, xena.PositionsHistoryRequest{})
				log.Printf("resp: %s\n", respColl)
				if err != nil {
					log.Printf("error: %v, account: %d, symbol: %s\n", err, acc.Id, acc.Currency)
				}
			}
		}
	}
	examples["orders"] = func() {
		for _, acc := range accounts {
			if xena.IsMargin(acc.Id) {
				respColl, err := client.GetOrders(acc.Id, "")
				log.Printf("resp: %s\n", respColl)
				if err != nil {
					log.Printf("error: %v, account: %d, symbol: %s\n", err, acc.Id, acc.Currency)
				}
			}
		}
	}
	examples["trade-history"] = func() {
		for _, acc := range accounts {
			if xena.IsMargin(acc.Id) {
				respColl, err := client.GetTradeHistory(acc.Id, xena.TradeHistoryRequest{}.SetPage(1, 100))
				log.Printf("resp: %s\n", respColl)
				if err != nil {
					log.Printf("error: %v, account: %d, symbol: %s\n", err, acc.Id, acc.Currency)
				}
			}
		}
	}
	examples["balances"] = func() {
		for _, acc := range accounts {
			respColl, err := client.GetBalance(acc.Id)
			log.Printf("resp: %s\n", respColl)
			if err != nil {
				log.Printf("error: %v, account: %d, symbol: %s\n", err, acc.Id, acc.Currency)
			}
		}
	}
	examples["margin-requirements"] = func() {
		for _, acc := range accounts {
			if xena.IsMargin(acc.Id) {
				respColl, err := client.GetMarginRequirements(acc.Id)
				log.Printf("resp: %s\n", respColl)
				if err != nil {
					log.Printf("error: %v, account: %d, symbol: %s\n", err, acc.Id, acc.Currency)
				}
			}
		}
	}
	examples["heartbeat"] = func() {
		gid := "42"
		hb := int32(42)
		err := client.SendApplicationHeartbeat(gid, hb)
		if err != nil {
			log.Printf("error: %v on SendApplicationHeartbeat(%s, %d)", err, gid, hb)
		}
	}

	for i, a := range os.Args {
		log.Printf("%d - %s\n", i, a)
		if strings.EqualFold(a, "-h") ||
			strings.EqualFold(a, "h") ||
			strings.EqualFold(a, "help") ||
			strings.EqualFold(a, "--help") ||
			strings.EqualFold(a, "-help") {
			log.Println("list of available examples")
			for k := range examples {
				log.Printf("\t%s\n", k)
			}
			return
		}
	}

	keyExamples := os.Args
	if len(keyExamples) == 0 {
		for k := range examples {
			keyExamples = append(keyExamples, k)
		}
	}

	sort.Strings(keyExamples)
	for _, key := range keyExamples {
		log.Printf("--- run key %s\n", key)
		f, ok := examples[key]
		if !ok {
			log.Printf("key not found %s\n", key)
			time.Sleep(time.Millisecond)
			continue
		}
		f()
	}

}

func GetBests(symbol string) (bestAsk, bestBid float64) {
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
