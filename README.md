# bazaar

This repo holds some rudimentary experiments with paper stock trading using alpaca.markets.

## Setup

- Get an account at alpaca.markets
- Generate a paper trading ID and secret
- Clone the repo
- Populate an `.env` file in the repo root with your `API_KEY_ID` and `API_KEY_SECRET` values
- Run `go run main.go` in the repo root
    - Alternately, you can `Run and Debug` in VSCode
- ???
- Lose money

## Useful Links

- General documentation: https://alpaca.markets/docs/
- Getting started: https://alpaca.markets/docs/get-started-with-alpaca/
- Go example: https://github.com/alpacahq/alpaca-trade-api-go/blob/master/examples/long-short/long-short.go
- Dashboard: https://app.alpaca.markets/paper/dashboard/overview
- Godoc: https://godoc.org/github.com/alpacahq/alpaca-trade-api-go/alpaca

## Some Other Cool Things

- Beta streaming API: https://alpaca.markets/docs/api-documentation/api-v2/market-data/streaming/