package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
	"github.com/joho/godotenv"
	"github.com/kfcampbell/bazaar/bazaar"
)

// make sure there's only one instance of the trading client
var alpacaClient bazaar.ClientContainer

func main() {
	if err := realMain(); err != nil {
		log.Fatalf(err.Error())
	}
}

func realMain() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	apiKeyID := os.Getenv("API_KEY_ID")
	apiKeySecret := os.Getenv("API_KEY_SECRET")

	baseURL := "https://paper-api.alpaca.markets"
	alpaca.SetBaseUrl(baseURL)

	// obnoxiously copy over environment variables for the alpaca package
	if common.Credentials().ID == "" {
		os.Setenv(common.EnvApiKeyID, apiKeyID)
	}
	if common.Credentials().Secret == "" {
		os.Setenv(common.EnvApiSecretKey, apiKeySecret)
	}

	allStocks := []bazaar.StockField{}
	stockList := []string{"DOMO", "TLRY", "SQ", "MRO", "AAPL", "GM", "SNAP",
		"SHOP", "SPLK", "BA", "AMZN", "SUI", "SUN", "TSLA", "CGC", "SPWR", "NIO",
		"CAT", "MSFT", "PANW", "OKTA", "TWTR", "TM",
		"RTN", "ATVI", "GS", "BAC", "MS", "TWLO", "QCOM"}
	for _, stock := range stockList {
		allStocks = append(allStocks, bazaar.NewStockField(stock, 0))
	}

	alpacaClient := bazaar.NewClientContainer(
		alpaca.NewClient(common.Credentials()),
		bazaar.NewBucket([]string{}, -1, -1, 0),
		bazaar.NewBucket([]string{}, -1, -1, 0),
		make([]bazaar.StockField, len(allStocks)),
		[]string{},
	)

	copy(alpacaClient.AllStocks, allStocks)

	asset, err := alpacaClient.Client.GetAsset("MSFT")
	if err != nil {
		return err
	}
	fmt.Printf(asset.Symbol)

	/*assetKey := ""
	ordert := alpaca.OrderType.Market
	orderReq := alpaca.PlaceOrderRequest{
		AccountID: "",
		AssetKey:  &assetKey,
		Qty:       decimal.New(1, 1),
		Side:      "buy",
		Type:      "market",
		//TimeInForce: ,
	}
	alpacaClient.client.PlaceOrder()*/

	// list orders
	status, until, limit := "open", time.Now(), 100
	orders, _ := alpacaClient.Client.ListOrders(&status, &until, &limit, nil)
	for _, order := range orders {
		fmt.Printf("order: %v\n", order)
	}

	return nil
}
