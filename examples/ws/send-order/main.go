package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"

	"github.com/xenaex/client-go/xena"
	"github.com/xenaex/client-go/xena/api"
)

// set api ket in environment.
// export XENA_API_KEY=""
// export XENA_API_SECRET=""
// export XENA_HOST="ws://api.xena.rc/ws/trading/"

var accountId = uint64(1012000000)

func main() {
	log.Printf("Start")

	apiKey := os.Getenv("XENA_API_KEY")
	apiSecret := os.Getenv("XENA_API_SECRET")
	host := os.Getenv("XENA_HOST")

	if len(apiKey) == 0 || len(apiSecret) == 0 {
		fmt.Println("api key or api secret not found.")
		return
	}
	if len(host) == 0 {
		host = "wss://api.xena.exchange/ws/market-data/"
	}

	log.Println("will be connect to ", host)
	client := xena.NewTradingClient(apiKey, apiSecret, xena.WithURL(host), xena.WithDebug())
	client.ListenBalanceIncrementalRefresh(onBalanceIncrementalRefreshHandler)
	client.ListenBalanceSnapshotRefresh(onBalanceSnapshotRefreshHandler)
	client.ListenExecutionReport(onExecutionReportHandler)
	client.ListenListStatus(onListStatusHandler)
	client.ListenLogon(logonHandler)
	client.ListenMarginRequirementReport(onMarginRequirementReportHandler)
	client.ListenMassPositionReport(onMassPositionReportHandler)
	client.ListenOrderCancelReject(onOrderCancelRejectHandler)
	client.ListenOrderMassStatusResponse(onOrderMassStatusResponseHandler)
	client.ListenPositionMaintenanceReport(onPositionMaintenanceReportHandler)
	client.ListenPositionReport(onPositionReportHandler)
	client.ListenReject(onRejectHandler)

	logonResponse, err := client.ConnectAndLogon()
	if err != nil {
		fmt.Println(err)
	}
	log.Println("GOT logonResponse ", logonResponse)
	log.Println("logon completed")

	err = client.AccountStatusReport(accountId, "")
	if err != nil {
		fmt.Println(err)
	}

	order := xena.CreateLimitOrder(fmt.Sprint(rand.Int()), xena.XBTUSD, xena.SideSell, "1", accountId, "7523.4").Build()
	err = client.Send(order)
	err = xena.CreateLimitOrder(fmt.Sprint(rand.Int()), xena.XBTUSD, xena.SideSell, "1", accountId, "7523.4").Send(client)
	if err != nil {
		fmt.Println(err)
	}

	limitOrder := xena.CreateLimitOrder(fmt.Sprint(rand.Int()), xena.XBTUSD, xena.SideSell, "1", accountId, "7523.4").SetTimeInForce("").SetPositionId(0).SetTakeProfitPrice("8000").Build()
	err = client.Send(limitOrder)
	if err != nil {
		fmt.Println(err)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt

	log.Printf("End")
}

func onListStatusHandler(t xena.TradingClient, m *api.ListStatus) {
	log.Println("GOT onListStatus ", m)
}

func onRejectHandler(t xena.TradingClient, m *api.Reject) {
	log.Println("GOT onReject ", m)
}

func onPositionMaintenanceReportHandler(t xena.TradingClient, m *api.PositionMaintenanceReport) {
	log.Println("GOT onPositionMaintenanceReport ", m)
}

func onMassPositionReportHandler(t xena.TradingClient, m *api.MassPositionReport) {
	log.Println("GOT onMassPositionReport ", m)
}

func onPositionReportHandler(t xena.TradingClient, m *api.PositionReport) {
	log.Println("GOT onPositionReport ", m)
}

func onOrderMassStatusResponseHandler(t xena.TradingClient, m *api.OrderMassStatusResponse) {
	log.Println("GOT onOrderMassStatusResponse ", m)
}

func onOrderCancelRejectHandler(t xena.TradingClient, m *api.OrderCancelReject) {
	log.Println("GOT onOrderCancelReject ", m)
}

func onExecutionReportHandler(t xena.TradingClient, m *api.ExecutionReport) {
	log.Println("GOT onExecutionReport ", m)
}

func onMarginRequirementReportHandler(t xena.TradingClient, m *api.MarginRequirementReport) {
	log.Println("GOT onMarginRequirementReport ", m)
}

func onBalanceIncrementalRefreshHandler(t xena.TradingClient, m *api.BalanceIncrementalRefresh) {
	log.Println("GOT onBalanceIncrementalRefresh ", m)
}

func onBalanceSnapshotRefreshHandler(t xena.TradingClient, m *api.BalanceSnapshotRefresh) {
	log.Println("GOT onBalanceSnapshotRefresh ", m)
}

func logonHandler(t xena.TradingClient, m *api.Logon) {
	log.Println("GOT onLogon ", m)
}
