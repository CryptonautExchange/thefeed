package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Market struct {
	ID        bson.ObjectId `storm:"id"`
	Currency  string
	Commodity string
	Orderbook
}
