package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/xenaex/client-go/xena"
)

const Day = 24 * time.Hour

func main() {
	var err error
	var resp interface{}

	if strings.Contains(os.Args[0], "/") && strings.Contains(os.Args[0], "main") {
		os.Args = os.Args[1:]
	}
	client := xena.NewMarketDataREST(
		xena.WithRestMarketDataHost,
		//		xena.WithRestHost("http://api.xena.test/"),
	)

	examples := make(map[string]func())
	examples["server-time"] = func() {
		resp, err = client.GetServerTime()
		fmt.Printf("resp: %s, \nerror: %#v\n", resp, err)
	}
	examples["instruments"] = func() {
		resp, err = client.GetInstruments()
		fmt.Printf("resp: %s, \nerror: %#v\n", resp, err)
	}

	examples["trades"] = func() {
		resp, err = client.GetTrades(xena.XBTUSD.String(), time.Now().Add(-10*Day), time.Now(), 1, 10)
		fmt.Printf("resp: %s, \nerror: %#v\n", resp, err)
	}

	examples["dom"] = func() {
		resp, err = client.GetDom(xena.XBTUSD.String())
		fmt.Printf("resp: %s, \nerror: %#v\n", resp, err)
	}

	examples["candles"] = func() {
		resp, err = client.GetCandles(xena.BTCUSDT.String(), "1m", time.Now().Add(-5*time.Minute), time.Now())
		fmt.Printf("resp: %s, \nerror: %#v\n", resp, err)
	}
	for i, a := range os.Args {
		fmt.Printf("%d - %s\n", i, a)
		if strings.EqualFold(a, "-h") ||
			strings.EqualFold(a, "h") ||
			strings.EqualFold(a, "help") ||
			strings.EqualFold(a, "--help") ||
			strings.EqualFold(a, "-help") {
			fmt.Println("list of available examples")
			for k := range examples {
				fmt.Printf("\t%s\n", k)
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
		f, ok := examples[key]
		if !ok {
			fmt.Printf("key not found %s\n", key)
			continue
		}
		f()
	}
}
