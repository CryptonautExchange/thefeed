package main

import (
	"fmt"

	"github.com/iowar/poloniex"
)

func main() {

	ws, err := poloniex.NewWSClient(true)
	if err != nil {
		return
	}

	err = ws.SubscribeMarket("usdt_btc")
	if err != nil {
		return
	}

	var m poloniex.OrderBook

	for {
		receive := <-ws.Subs["usdt_btc"]
		updates := receive.([]poloniex.MarketUpdate)
		for _, v := range updates {
			if v.TypeUpdate == "OrderBookRemove" || v.TypeUpdate == "OrderBookModify" {
				m = v.Data.(poloniex.OrderBook)

				fmt.Printf("Rate:%f, Type:%s, Amount:%f\n",
					m.Rate, m.TypeOrder, m.Amount)
			}
		}
	}
}
