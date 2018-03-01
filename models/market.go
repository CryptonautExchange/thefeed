package models

import (
	"labix.org/v2/mgo/bson"
)

type Market struct {
	ID        bson.ObjectId `storm:"id"`
	Currency  string
	Commodity string
	Orderbook
}
