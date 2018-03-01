package exchange

import (
//"fmt"
)

type Exchange interface {
	Initialize([]string)
	Start()
	Stop()
	Status() string
	PopulateOrderbook() int
}
