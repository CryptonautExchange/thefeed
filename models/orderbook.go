package models

import (
//"fmt"
)

type Orderbook struct {
	Market    string
	Currency  string
	Commodity string
	Bids      []bid // Buy Orders
	Asks      []ask // Sell Orders
	// Orderbook Stats
	HighestBid bid
	LowestAsk  ask
	Spread     string
}
