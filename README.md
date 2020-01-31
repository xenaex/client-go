# Xena Exchange official websocket and rest clients for go lang

For API documentation check out [Help Center](https://support.xena.exchange/support/solutions/folders/44000161002)


#### Install

Add to project:
```
go get  github.com/xenaex/client-go
```


#### Market Data websocket example

```go
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
	md := xena.NewMarketData(xena.WithMarketDataURL())
	md.SetDisconnectHandler(xena.DefaultMarketDisconnectHandler)
	resp, err := md.Connect()
	if err != nil {
		log.Printf("error %s on md.Connect()", err)
	}
	log.Printf("loggon message %s", resp)

	id, err := md.SubscribeOnCandles(xena.XBTUSD.String(), "1m", handler, xena.ThrottleCandles1s, xena.AggregateBook25)
	log.Println(id, err)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
}

func handler(md xena.MarketDataClient, r *xmsg.MarketDataRequestReject, m *xmsg.MarketDataRefresh) {
	log.Println("GOT", r, m)
	time.Sleep(20 * time.Millisecond)
}
```

#### Trading websocket example

Register an account with [Xena](https://trading.xena.exchange/registration). Generate an API Key and assign relevant permissions.
	
```go
package main

import (
	"log"
	"time"

	"github.com/xenaex/client-go/xena"
	"github.com/xenaex/client-go/xena/xmsg"
)

func main() {
	apiKey := "your api key"
	apiSecret := "your api secret"
	accountId := uint64(10000000) // your account.
	client := xena.NewTradingClient(apiKey, apiSecret, xena.WithTradingURL())
	client.SetDisconnectHandler(xena.DefaultTradingDisconnectHandler)

	connected := make(chan struct{})
	client.ListenLogon(func(t xena.TradingClient, m *xmsg.Logon) {
		if len(m.RejectText) == 0 {
			connected <- struct{}{}
		}
	})
	client.ListenExecutionReport(func(t xena.TradingClient, m *xmsg.ExecutionReport) {
		log.Println("GOT", m)
		time.Sleep(20 * time.Millisecond)
	})

	resp, err := client.ConnectAndLogon()
	if err != nil {
		log.Printf("loggon err: %s\n", err)
		return
	} else {
		log.Printf("resp: %s\n", resp)
		if len(resp.RejectText) > 0 {
			return
		}
	}
	<-connected
	client.ListenLogon(nil)
	close(connected)

	err = client.MarketOrder(accountId, xena.ID(""), xena.XBTUSD.String(), xena.SideBuy, "1")
	if err != nil {
		log.Printf("err %s", err)
	}
	time.Sleep(15 * time.Second)
}
```

#### Trading rest example

Register an account with [Xena](https://trading.xena.exchange/registration). Generate an API Key and assign relevant permissions.

```go
package main

import (
	"fmt"
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
	apiKey := "your api key"
	apiSecret := "your api secret"
	accountId := uint64(10000000) // your account.

	client := xena.NewTradingREST(apiKey, apiSecret, xena.WithRestTradingHost)
	resp, err := client.SendMarketOrder(accountId, xena.ID(""), xena.XBTUSD.String(), xena.SideBuy, "1")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("resp: %s\n", resp)
}
```

#### Market data rest example

Register an account with [Xena](https://trading.xena.exchange/registration). Generate an API Key and assign relevant permissions.

```go
package main

import (
	"fmt"

	"github.com/xenaex/client-go/xena"
)

func main() {
	client := xena.NewMarketDataREST(xena.WithRestMarketDataHost)
	resp, err := client.GetInstruments()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("resp: %s\n", resp)
}
```


For more examples check out "examples" folder.