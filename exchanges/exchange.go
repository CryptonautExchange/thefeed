package exchange

import (
//"fmt"
)

type Exchange interface {
	Initialize(string, string)
	Start()
	Stop()

	Orderbook()

	IsEnabled() bool
	SetEnabled(bool)
}
