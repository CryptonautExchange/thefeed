package main

import (
	"fmt"

	exchange "github.com/CryptonautExchange/thefeed/exchanges"

	"github.com/asdine/storm"
	. "github.com/hackwave/color"

	"libs/ui/spinner"
)

func main() {
	// Initialize Orderbook Database
	fmt.Println(Magenta("The Feed"), Gray(":"), Blue("Market Data Mocking Library"))
	fmt.Println(Gray("====================================="))
	db, err := storm.Open("orderbook.db")
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
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
