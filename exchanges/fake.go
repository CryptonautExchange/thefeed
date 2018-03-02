package exchange

import (
	"fmt"

	"github.com/CryptonautExchange/thefeed/models"

	"github.com/hackwave/color"
)

type FakeExchange struct {
	//Incoming         chan interface{} // Seems like it should be the message we are expecting like JSON types
	//MarketUpdates
	SubscribedMarket string
	Enabled          bool
}

func (self FakeExchange) Initialize(subscribedMarket string) (err error) {
	// TODO:
	// TODO: Move to a propper logging system
	fmt.Println(Gray("Initializing Poloniex..."))
	return nil
}

func (self FakeExchange) Start() {
	fmt.Println("Starting Websockets Connection To Poloniex...")

}

func (self FakeExchange) Stop() {
	fmt.Println("Stopping Websockets Connection To Poloniex...")
}

func (self FakeExchange) Status() string {
	fmt.Println("Status Of Websockets Connection Is...")
}

func (self FakeExchange) PopulateOrderbook() {
	fmt.Println("Populating Poloniex Orderbook...")
	// Return asks and bids to fill an orderbook in the model calling this
}
