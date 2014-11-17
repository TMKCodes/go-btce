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
	Depth, err := btce.Depth("btc_eur");
	if err != nil {
		fmt.Printf("Depth: %v\n", err);
	} else {
		fmt.Printf("Ask rate: %v\n", Depth.Asks[0][0]);
		fmt.Printf("Ask amount: %v\n", Depth.Asks[0][1]);
		fmt.Printf("Bid rate: %v\n", Depth.Bids[0][0]);
		fmt.Printf("Bid amount: %v\n", Depth.Bids[0][1]);
	}
}
