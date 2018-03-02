package main

import (
	"fmt"

	exchange "github.com/CryptonautExchange/thefeed/exchanges"

	"github.com/asdine/storm"
	. "github.com/hackwave/color"
)

func main() {
	// Initialize Orderbook Database
	fmt.Println(Magenta("The Feed"), Gray(":"), Blue("Market Data Mocking Library"))
	fmt.Println(Gray("====================================="))
	db, err := storm.Open("orderbook.db")
	if err != nil {
		fmt.Println(Red("Fatal Error:"), Gray("Failed to load orderbook database, exiting..."))
		fmt.Println(Red("["), Gray(err), Red("]"))
	}
	defer db.Close()

	poloniex := exchange.Poloniex{}

	poloniex.Initialize("BTC_ETH")

}
