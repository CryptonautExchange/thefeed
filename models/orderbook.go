package models

import (
	"time"

	"labix.org/v2/mgo/bson"
)

type Orderbook struct {
	ID             bson.ObjectId `storm:"id"`
	Bids           []bid         // Buy Orders
	Asks           []ask         // Sell Orders
	LastActivityAt time.Time
	// Orderbook Stats
	HighestBid bid
	LowestAsk  ask
	Spread     string
	BidDepth   string
	AskDepth   string
}
