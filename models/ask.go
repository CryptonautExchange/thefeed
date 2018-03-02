package models

import (
	//"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Ask struct {
	ID        bson.ObjectId `storm:"id"`
	CreatedAt time.Time     `storm:"index"`
	Currency  string        `storm:"index"`
	Commodity string        `storm:"index"`
	Price     string        `storm:"index"`
}
