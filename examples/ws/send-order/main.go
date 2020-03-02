package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/xenaex/client-go/xena"
	"github.com/xenaex/client-go/xena/xmsg"
)

// set api ket in environment.
// export XENA_API_KEY=""
// export XENA_API_SECRET=""

var accountId = uint64(1000000000)

func main() {
	log.Printf("Start")

	apiKey := os.Getenv("XENA_API_KEY")
	apiSecret := os.Getenv("XENA_API_SECRET")

	if len(apiKey) == 0 || len(apiSecret) == 0 {
		log.Println("api key or api secret not found.")
		return
	}

	// var client = xena.NewTradingClient(apiKey, apiSecret, []uint64{}, xena.WithTradingURL(), xena.WithDebug(), xena.WithURL("ws://api.xena.rc/ws/trading"))
	var client = xena.NewTradingClient(apiKey, apiSecret, []uint64{accountId}, xena.WithTradingURL(), xena.WithDebug(), xena.WithURL("ws://localhost/api/ws/trading"))
	client.SetDisconnectHandler(xena.DefaultTradingDisconnectHandler)

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

	var err error
	var logonResponse *xmsg.Logon
	for {
		logonResponse, err = client.ConnectAndLogon()
		if err != nil {
			log.Printf("logon err: %s\n", err)
			time.Sleep(time.Second)
			continue
		}
		log.Printf("resp: %s\n", logonResponse)
		if len(logonResponse.RejectText) > 0 {
			return
		}
		break
	}
	log.Println("GOT logonResponse ", logonResponse)
	log.Println("logon completed")

	err = client.GetAccountStatusReport(accountId, "")
	if err != nil {
		log.Println(err)
	}

	order := xena.CreateLimitOrder(fmt.Sprint(rand.Int()), xena.XBTUSD.String(), xena.SideSell, "1", accountId, "7523.4").Build()
	err = client.Send(order)
	err = xena.CreateLimitOrder(fmt.Sprint(rand.Int()), xena.XBTUSD.String(), xena.SideSell, "1", accountId, "7523.4").Send(client)
	if err != nil {
		log.Println(err)
	}

	limitOrder := xena.CreateLimitOrder(fmt.Sprint(rand.Int()), xena.XBTUSD.String(), xena.SideSell, "1", accountId, "7523.4").SetTimeInForce("").SetPositionId(0).SetTakeProfitPrice("8000").Build()
	err = client.Send(limitOrder)
	if err != nil {
		log.Println(err)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt

	log.Printf("End")
}

func onListStatusHandler(_ xena.TradingClient, m *xmsg.ListStatus) {
	log.Println("GOT onListStatus ", m)
}

func onRejectHandler(_ xena.TradingClient, m *xmsg.Reject) {
	log.Println("GOT onReject ", m)
}

func onPositionMaintenanceReportHandler(_ xena.TradingClient, m *xmsg.PositionMaintenanceReport) {
	log.Println("GOT onPositionMaintenanceReport ", m)
}

func onMassPositionReportHandler(_ xena.TradingClient, m *xmsg.MassPositionReport) {
	log.Println("GOT onMassPositionReport ", m)
}

func onPositionReportHandler(_ xena.TradingClient, m *xmsg.PositionReport) {
	log.Println("GOT onPositionReport ", m)
}

func onOrderMassStatusResponseHandler(_ xena.TradingClient, m *xmsg.OrderMassStatusResponse) {
	log.Println("GOT onOrderMassStatusResponse ", m)
}

func onOrderCancelRejectHandler(_ xena.TradingClient, m *xmsg.OrderCancelReject) {
	log.Println("GOT onOrderCancelReject ", m)
}

func onExecutionReportHandler(_ xena.TradingClient, m *xmsg.ExecutionReport) {
	log.Println("GOT onExecutionReport ", m)
}

func onMarginRequirementReportHandler(_ xena.TradingClient, m *xmsg.MarginRequirementReport) {
	log.Println("GOT onMarginRequirementReport ", m)
}

func onBalanceIncrementalRefreshHandler(_ xena.TradingClient, m *xmsg.BalanceIncrementalRefresh) {
	log.Println("GOT onBalanceIncrementalRefresh ", m)
}

func onBalanceSnapshotRefreshHandler(_ xena.TradingClient, m *xmsg.BalanceSnapshotRefresh) {
	log.Println("GOT onBalanceSnapshotRefresh ", m)
}

func logonHandler(_ xena.TradingClient, m *xmsg.Logon) {
	log.Println("GOT onLogon ", m)
}
