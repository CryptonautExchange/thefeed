package exchange

import (
	"fmt"

	"github.com/CryptonautExchange/thefeed/exchanges/bittrex"
	"github.com/CryptonautExchange/thefeed/models"
	. "github.com/hackwave/color"
)

type Bittrex struct {
	SubscribedMarket string
	Enabled          bool
}

func (self Bittrex) Initialize(subscribedMarket string) (err error) {
	// TODO: Move to a propper logging system
	fmt.Println(Gray("Initializing Bittrex..."))
	return nil
}

func (self Bittrex) Start() {
	fmt.Println("Starting Websockets Connection To Bittrex...")
}

func (self Bittrex) Stop() {
	fmt.Println("Stopping Websockets Connection To Bittrex...")
}

func (self Bittrex) Status() string {
	fmt.Println("Status Of Websockets Connection Is...")
}

func (self Bittrex) PopulateOrderbook() {
	fmt.Println("Populating Bittrex Orderbook...")
	// Return asks and bids to fill an orderbook in the model calling this
}
