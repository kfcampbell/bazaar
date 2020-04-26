package bazaar

import (
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
)

// ClientContainer holds a pointer to an alpaca.Client, and
// some convenience fields for stock access.
// todo(kfcampbell): is this really the construction we want to go with?
type ClientContainer struct {
	Client    *alpaca.Client
	Long      Bucket
	Short     Bucket
	AllStocks []StockField
	Blacklist []string
}

// Bucket holds a bucket...of something. I don't know, I don't know money things
type Bucket struct {
	List        []string
	Qty         int
	AdjustedQty int
	EquityAmt   float64
}

// StockField does something
type StockField struct {
	Name string
	PC   float64
}
