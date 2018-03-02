package exchange

import (
	"fmt"

	"github.com/CryptonautExchange/thefeed/exchanges/poloniex"
	. "github.com/hackwave/color"
)

type Poloniex struct {
	MarketUpdates    []poloniex.MarketUpdate
	Websocket        *poloniex.WSClient
	Orderbook        poloniex.OrderBook
	SubscribedMarket string
	Enabled          bool
}

func (self Poloniex) Initialize(subscribedMarket string) (err error) {
	// TODO: Move to a propper logging system
	fmt.Println(Gray("Initializing Poloniex..."))
	self.Websocket, err = poloniex.NewWSClient(true)
	if err != nil {
		return err
	}
	err = self.Websocket.SubscribeMarket(subscribedMarket)
	if err != nil {
		return err
	}
	self.SubscribedMarket = subscribedMarket
	fmt.Println(Green("Successfully subscribed to"), ":", self.SubscribedMarket)
	return nil
}

func (self Poloniex) Start() {
	fmt.Println("Starting Websockets Connection To Poloniex...")
	for {
		receive := <-self.Websocket.Subs[self.SubscribedMarket]
		self.MarketUpdates = receive.([]poloniex.MarketUpdate)
	}

}

func (self Poloniex) Stop() {
	fmt.Println("Stopping Websockets Connection To Poloniex...")
}

func (self Poloniex) Status() string {
	fmt.Println("Status Of Websockets Connection Is...")
	return ""
}

func (self Poloniex) PopulateOrderbook() {
	fmt.Println("Populating Poloniex Orderbook...")
	// Return asks and bids to fill an orderbook in the model calling this
}
