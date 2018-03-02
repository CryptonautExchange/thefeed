package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Exchange struct {
	ID      bson.ObjectId `storm:"id"`
	Markets []Market
}
