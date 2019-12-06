package main

import (
	"github.com/xenaex/client-go/xena"
	"github.com/xenaex/client-go/xena/messages"
	"log"
	"os"
	"os/signal"
)

func main() {
	log.Printf("Start")

	client := xena.NewWsClient(xena.WithDebug(), xena.WithConnectHandler(onConnect))
	client.Connect()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt

	log.Printf("End")
}

func onConnect(c xena.WsClient)  {
	log.Print("Connection established")

	sm := messages.MarketDataRequest{
		MsgType:                 messages.MsgType_MarketDataRequest,
		MDStreamID:              "DOM:BTC/USDT:aggregated",
		SubscriptionRequestType: messages.SubscriptionRequestType_SnapshotAndUpdates ,
	}

	c.Write(sm)

}
