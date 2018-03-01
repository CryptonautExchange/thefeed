package main

import (
	"fmt"

	"github.com/CryptonautExchange/thefeed/exchanges/poloniex"
)

func main() {

	ws, err := poloniex.NewWSClient(true)
	if err != nil {
		return
	}

	market := "btc_eth"

	err = ws.SubscribeMarket(market)
	if err != nil {
		return
	}

	var m poloniex.OrderBook

	for {
		receive := <-ws.Subs[market]
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
