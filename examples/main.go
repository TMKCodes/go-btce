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
	fee, err := btce.Fee("btc_usd");
	if err != nil {
		fmt.Printf("Fee: %v\n", err);
	} else {
		fmt.Printf("Fee: %#v\n", fee);
	}
	ticker, err := btce.Ticker("btc_usd");
	if err != nil {
		fmt.Printf("Ticker: %v\n", err);
	} else {
		fmt.Printf("Ticker: %#v\n", ticker);
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
	Info, err := btce.GetInfo();
	if err != nil {
		fmt.Printf("Info: %v\n", err);
	} else {
		fmt.Printf("Info: %#v\n", Info);
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
