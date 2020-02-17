package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/xenaex/client-go/xena"
	"github.com/xenaex/client-go/xena/xmsg"
)

func main() {
	log.Printf("Start")

	md := xena.NewMarketData(
		xena.WithDebug(),
		xena.WithMarketDataURL(),
	)
	md.SetDisconnectHandler(xena.DefaultMarketDisconnectHandler)

	var err error
	var resp *xmsg.Logon
	for {
		resp, err = md.Connect()
		if err != nil {
			log.Printf("error %s on md.Connect()", err)
			time.Sleep(time.Second)
			continue
		}
		log.Printf("logon message %s", resp)
		if len(resp.RejectText) > 0 {
			return
		}
		break
	}
	id := ""
	id, err = md.SubscribeOnCandles(xena.XBTUSD.String(), "1m", handler, xena.ThrottleCandles1s, xena.AggregateBook25)
	log.Println(id, err)
	id, err = md.SubscribeOnDom(xena.XBTUSD.String(), handler, xena.ThrottleDOM5s, xena.MarketDepth10)
	log.Println(id, err)
	id, err = md.SubscribeOnTrades(xena.XBTUSD.String(), handler, xena.ThrottleTrades5s)
	log.Println(id, err)
	id, err = md.SubscribeOnMarketWatch(xena.XBTUSD.String(), handler)
	log.Println(id, err)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt

	log.Printf("End")
}

func handler(md xena.MarketDataClient, r *xmsg.MarketDataRequestReject, m *xmsg.MarketDataRefresh) {
	log.Println("GOT", r, m)
	time.Sleep(20 * time.Millisecond)
}

func domHandler(_ xena.MarketDataClient, m *xmsg.MarketDataRefresh) {
	log.Println("GOT", m)
}
