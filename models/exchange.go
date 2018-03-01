package models

import (
	"labix.org/v2/mgo/bson"
)

type Exchange struct {
	ID bson.ObjectId `storm:"id"`
}
