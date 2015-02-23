Go-btce
=======

Go-btce is an implementation of the [btc-e.com](https://btc-e.com) API (public and private) in Go.

This version implements the version 2 [btc-e.com](https://btc-e.com) API.

## Install

	go get github.com/TMKCodes/go-btce


## Import

	import "github.com/TMKCodes/go-btce"


## Example
```go
package main
import (
"fmt"
"github.com/TMKCodes/go-btce"
)
const (
KEY = "M7Z2I8HR-07RJRR6L-PO7DDGGX-LCJ7VW3B-NX7PCSET"
SECRET = "4a720e1de1ba08bec388209996df8141e48813c57dbc12425985eebf3e8e3b87"
)
func main() {
	btce := btce.New(KEY, SECRET);
	PublicInfo, err := btce.Info("btc_usd");
	if err != nil {
		fmt.Printf("Info: %v\n", err);
	} else {
		fmt.Printf("Info: %#v\n", Info);
	}
	Ticker, err := btce.Ticker("btc_usd");
	if err != nil {
		fmt.Printf("Ticker: %v\n", err);
	} else {
		fmt.Printf("Ticker: %#v\n", Ticker);
	}
	Trades, err := btce.Trades("btc_usd");
	if err != nil {
		fmt.Printf("Trades: %v\n", err);
	} else {
		fmt.Printf("Trades: %#v\n", Trades);
	}
	Depth, err := btce.Depth("btc_usd");
	if err != nil {
		fmt.Printf("Depth: %v\n", err);
	} else {
		fmt.Printf("Depth: %#v\n", Depth);
	}
	PrivateInfo, err := btce.GetInfo();
	if err != nil {
		fmt.Printf("Info: %v\n", err);
	} else {
		fmt.Printf("Info: %#v\n", PrivateInfo);
	}
	TransHistory, err := btce.TransHistory(0, 1000, 0, 0, "ASC", "", "");
	if err != nil {
		fmt.Printf("TransHistory: %v\n", err);
	} else {
		fmt.Printf("TransHistory: %#v\n", TransHistory);
	}
	TradeHistory, err := btce.TradeHistory(0, 1000, 0, 0, "DESC", "", "", "");
	if err != nil {
		fmt.Printf("TradeHistory: %v\n", err);
	} else {
		fmt.Printf("TradeHistory: %#v\n", TradeHistory);
	}
}
```
