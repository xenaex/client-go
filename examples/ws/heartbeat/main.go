package main

import (
	"log"
	"os"
	"time"

	"github.com/xenaex/client-go/xena"
	"github.com/xenaex/client-go/xena/xmsg"
)

var symbol = xena.XBTUSD.String()

func main() {
	log.Printf("Start")

	apiKey := os.Getenv("XENA_API_KEY")
	apiSecret := os.Getenv("XENA_API_SECRET")
	accountId := uint64(1000000000)

	if len(apiKey) == 0 || len(apiSecret) == 0 {
		log.Println("api key or api secret not found.")
		return
	}

	var err error
	client := xena.NewTradingClient(
		apiKey,
		apiSecret,
		xena.WithTradingURL(),
		xena.WithDebug(),
	)
	client.SetDisconnectHandler(xena.DefaultTradingDisconnectHandler)
	connected := make(chan struct{})
	client.ListenLogon(func(t xena.TradingClient, m *xmsg.Logon) {
		if len(m.RejectText) == 0 {
			connected <- struct{}{}
		}
	})
	resp, err := client.ConnectAndLogon()
	if err != nil {
		log.Printf("logon err: %s\n", err)
	} else {
		log.Printf("resp: %s\n", resp)
		if len(resp.RejectText) > 0 {
			return
		}
	}
	<-connected
	client.ListenLogon(nil)
	close(connected)
	client.ListenHeartbeat(func(t xena.TradingClient, m *xmsg.Heartbeat) {
		log.Printf("Heartbeat\n")
	})

	client.ListenExecutionReport(func(t xena.TradingClient, m *xmsg.ExecutionReport) {
		log.Printf("resp %v\n", m)
	})
	go func() {
		for {
			err := client.SendApplicationHeartbeat("42", 3)
			if err != nil {
				log.Printf("error %s", err)
			}
			time.Sleep(2 * time.Second)
		}
	}()
	started := time.Now()
	for time.Now().Sub(started) < time.Minute {
		go func() {
			ord := xena.CreateLimitOrder(xena.ID(""), xena.XBTUSD.String(), xena.SideSell, "1", accountId, "10000")
			ord = ord.SetGroupId("42")
			log.Printf("send order")

			err := client.Send(ord.Build())
			if err != nil {
				log.Printf("err %s", err)
			}
		}()
		time.Sleep(20 * time.Second)
	}
}
