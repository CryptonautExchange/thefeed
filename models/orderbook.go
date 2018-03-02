package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Orderbook struct {
	ID             bson.ObjectId `storm:"id"`
	Bids           []Bid         // Buy Orders
	Asks           []Ask         // Sell Orders
	LastActivityAt time.Time
	// Orderbook Stats
	HighestBid string
	LowestAsk  string
	Spread     string
	BidDepth   string
	AskDepth   string
}
