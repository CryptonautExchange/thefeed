package exchange

import (
	"fmt"

	"github.com/CryptonautExchange/thefeed/tree/master/models"
)

type Poloniex struct {
	//Connection
	Orderbook         poloniex.Orderbook
	SubscribedMarkets []string
	Enabled           bool
}

func (self Poloniex) Initialize(subscribedMarkets []string) (err error) {
	fmt.Println("Initializing Poloniex...")

	market := subscribedMarkets[0]

	err = ws.SubscribeMarket(market)
	if err != nil {
		return err
	}

}

func (self Poloniex) Start() {
	fmt.Println("Starting Websockets Connection To Poloniex...")

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
