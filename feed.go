package main

import (
	"fmt"

	exchange "github.com/CryptonautExchange/thefeed/exchanges"

	"github.com/asdine/storm"
	. "github.com/hackwave/textui/color"
	"github.com/hackwave/textui/loader"
)

func main() {
	// Initialize Orderbook Database
	fmt.Println(Magenta("The Feed"), Gray(":"), Blue("Market Data Mocking Library"))
	fmt.Println(Gray("====================================="))
	db, err := storm.Open("orderbook.db")

	s := ansi.Loader(loaders.Dots, 100)
	s.SetValue("Loading")
	s.Start()
	time.Sleep(4 * time.Second)

	fmt.Println(Gray("[Database] Initializing orderbook database..."))
	if err != nil {
		fmt.Println(Red("Fatal Error:"), Gray("Failed to load orderbook database, exiting..."))
		fmt.Println(Red("["), Gray(err), Red("]"))
	} else {
		s.Stop()
		fmt.Println(Green("[Database] Successfully opened the orderbook database."))
	}
	defer db.Close()

	poloniex := exchange.Poloniex{}

	poloniex.Initialize("BTC_ETH")

}
