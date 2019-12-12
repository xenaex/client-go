package main

import (
	"github.com/xenaex/client-go/xena"
	"github.com/xenaex/client-go/xena/xmsg"
	"log"
	"os"
	"os/signal"
)

func main() {
	log.Printf("Start")

	//md := xena.NewMarketData(xena.WithDebug(), xena.WithURL("ws://trading.xena.rc/api/ws/market-data"))
	md := xena.NewMarketData(xena.WithDebug())

	//id, err := md.SubscribeOnDOM(xena.BTCUSDT, domHandler, xena.AggregateBook10, xena.ThrottleDOM5s)
	id, err := md.SubscribeOnDOM(xena.BTCUSDT, domHandler, xena.ThrottleDOM500ms)
	//id, err := md.SubscribeOnDOM(xena.BTCUSDT, domHandler, xena.AggregateBook10)
	//id, err := md.SubscribeOnDOM(xena.XBTUSD, domHandler)
	log.Println(id, err)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt

	log.Printf("End")
}

func domHandler(md xena.MarketDataClient, m *xmsg.MarketDataRefresh) {
	log.Println("GOT", m)
}