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

// NewClientContainer constructs and returns a pointer to a ClientContainer
func NewClientContainer(alpacaClient *alpaca.Client, long Bucket, short Bucket,
	allStocks []StockField, blacklist []string) *ClientContainer {
	return &ClientContainer{
		Client:    alpacaClient,
		Long:      long,
		Short:     short,
		AllStocks: allStocks,
		Blacklist: blacklist,
	}
}

// Bucket holds a bucket...of something. I don't know, I don't know money things
type Bucket struct {
	List        []string
	Qty         int
	AdjustedQty int
	EquityAmt   float64
}

// NewBucket creates and returns a bucket
// (not a pointer because the alpaca interface doesn't like them for some reason)
func NewBucket(list []string, qty int, adjustedQty int, equityAmt float64) Bucket {
	return Bucket{
		List:        list,
		Qty:         qty,
		AdjustedQty: adjustedQty,
		EquityAmt:   equityAmt,
	}
}

// StockField does something
type StockField struct {
	Name string
	PC   float64
}

// NewStockField creates and returns a StockField
// (not a pointer because the alpaca interface doesn't like them for some reason)
func NewStockField(name string, pc float64) StockField {
	return StockField{
		Name: name,
		PC:   pc,
	}
}
