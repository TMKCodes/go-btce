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
	TransHistory, err := btce.TransHistory(0, 1000, 0, 0, "DESC", "", "");
	if err != nil {
		fmt.Printf("TransHistory: %v\n", err);
	} else {
		fmt.Printf("TransHistory: %#v\n", TransHistory);
	}
}
