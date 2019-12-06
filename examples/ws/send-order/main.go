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

	apiKey := "mMrGotJjPzvB2AIHo50PmkYnIV77VzVr1MdMBR7azqQ="
	apiSecret := "307702010104202b9897a48cd4ff109adc8857f9d634bcc9e0c915b64ac881c8dd68bd62d7c8cca00a06082a8648ce3d030107a14403420004f7a99c11874ac6004bdac9b390ef85f037e9f2892d2d66507fb5679f5d220544fe92f5d53a516052996698db246604beccc29bd0166c63f49dc965e36ab7d782"

	client := xena.NewTradingClient(apiKey, apiSecret, xena.WithDebug())
	client.ListenLogon(logonHandler)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt

	log.Printf("End")
}

func logonHandler(t xena.TradingClient, m *messages.Logon) {
	log.Println("GOT onLogon", m)

	err := t.SendLimitOrder(1012833471, xena.ID("clOrdID_"), xena.XBTUSD, xena.SideBuy, "8815", "1")
	log.Println("LimitOrder: ",err)
}