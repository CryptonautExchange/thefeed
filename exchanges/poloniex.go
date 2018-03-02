package exchange

import (
	"fmt"

	"github.com/CryptonautExchange/thefeed/exchanges/poloniex"
	"github.com/CryptonautExchange/thefeed/models"
	. "github.com/hackwave/color"
)

type Poloniex struct {
	Incoming         chan interface{} // Seems like it should be the message we are expecting like JSON types
	MarketUpdates    []poloniex.MarketUpdates
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
		self.Incoming = <-ws.Subs[market]
		self.MarketUpdates = self.Incoming.([]poloniex.MarketUpdate)
	}

}

func (self Poloniex) Stop() {
	fmt.Println("Stopping Websockets Connection To Poloniex...")
}

func (self Poloniex) Status() string {
	fmt.Println("Status Of Websockets Connection Is...")
}

func (self Poloniex) PopulateOrderbook() {
	fmt.Println("Populating Poloniex Orderbook...")
	// Return asks and bids to fill an orderbook in the model calling this
}
