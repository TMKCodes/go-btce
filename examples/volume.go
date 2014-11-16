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

	ticker, err := btce.Depth("ltc_btc");
	if err != nil {
		fmt.Printf("Ticker: %v\n", err);
	} else {
		askVolume := 0.0;
		for i := range ticker.Asks {
			askVolume += ticker.Asks[i][1];
		}
		fmt.Printf("Ask volume: %v LTC\n", askVolume);
		bidVolume := 0.0;
		for i := range ticker.Bids {
			bidVolume += ticker.Bids[i][1];
		}
		fmt.Printf("Bid volume: %v LTC\n", bidVolume);
	}
}
