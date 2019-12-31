package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/xenaex/client-go/xena"
	"github.com/xenaex/client-go/xena/xmsg"
)

func main() {
	log.Printf("Start")

	host := "ws://trading.xena.rc/api/ws/market-data"
	// md := xena.NewMarketData(xena.WithDebug(), xena.WithURL("ws://trading.xena.rc/api/ws/market-data"))
	md := xena.NewMarketData(xena.DefaultMarketDisconnectHandler, xena.WithURL(host), xena.WithDebug())
	resp, err := md.Connect()
	if err != nil {
		log.Printf("error %s on md.Connect()", err)
	}
	log.Printf("loggon message %s", resp)

	// id, err := md.SubscribeOnDOM(xena.BTCUSDT, domHandler, xena.AggregateBook10, xena.ThrottleDOM5s)
	id, err := md.SubscribeOnDOM(xena.BTCUSDT, domHandler, xena.ThrottleDOM500ms)
	// id, err := md.SubscribeOnDOM(xena.BTCUSDT, domHandler, xena.AggregateBook10)
	// id, err := md.SubscribeOnDOM(xena.XBTUSD, domHandler)
	log.Println(id, err)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt

	log.Printf("End")
}

func domHandler(_ xena.MarketDataClient, m *xmsg.MarketDataRefresh) {
	log.Println("GOT", m)
}
