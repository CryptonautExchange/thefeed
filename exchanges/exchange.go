package exchange

type Exchange interface {
	Initialize(string)
	Start()
	Stop()
	Status() string
	PopulateOrderbook() int
}
