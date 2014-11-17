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

	Trades, err := btce.Trade("ltc_btc", "buy", 0.00991, (0.00269559 / 0.00991) - (0.00269559 / 0.00991 * 0.002));
	if err != nil {
		fmt.Printf("Trades: %v\n", err);
	} else {
		fmt.Printf("Trades: %#v\n", Trades);
	}
}
