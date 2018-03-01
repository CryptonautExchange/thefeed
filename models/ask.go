package orderbook

import (
	//"fmt"
	"time"

	"labix.org/v2/mgo/bson"
)

type Ask struct {
	ID        bson.ObjectId `storm:"id"`
	CreatedAt time.Time     `storm:"index"`
	Currency  string        `storm:"index"`
	Commodity string        `storm:"index"`
	Price     string        `storm:"index"`
}
