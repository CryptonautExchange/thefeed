package main

import (
	"fmt"

	exchange "github.com/CryptonautExchange/thefeed/exchanges"

	"github.com/asdine/storm"
)

func main() {
	// Initialize Orderbook Database
	db, err := storm.Open("orderbook.db")
	defer db.Close()

	poloniex := exchange.Poloniex{}

	poloniex.Initialize()
}
