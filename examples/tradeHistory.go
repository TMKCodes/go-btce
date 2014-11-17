package main

import (
	"fmt"
	"sort"
	"github.com/TMKCodes/go-btce"
)

const (
	KEY = "M7Z2I8HR-07RJRR6L-PO7DDGGX-LCJ7VW3B-NX7PCSET"
	SECRET = "4a720e1de1ba08bec388209996df8141e48813c57dbc12425985eebf3e8e3b87"
)

type ByTimestamp []btce.TradeHistoryTrade;

func (this ByTimestamp) Len() int { return len(this); }
func (this ByTimestamp) Swap(i, j int) { this[i], this[j] = this[j], this[i]; }
func (this ByTimestamp) Less(i, j int) bool { return this[i].Timestamp > this[j].Timestamp; }

func main() {
	btc := btce.New(KEY, SECRET);
	TradeHistory, err := btc.TradeHistory(0, 1000, 0, 0, "ASC", "", "", "ltc_eur");
	if err != nil {
		fmt.Printf("TradeHistory: %v\n", err);
	} else {
		temp := make([]btce.TradeHistoryTrade, 0);
		for key, value := range TradeHistory.Return {
			fmt.Printf("TradeHistory.Return[%#v] := %#v\n", key, value);
			temp = append(temp, value);
		}
		sort.Sort(ByTimestamp(temp));
		for i := range temp {
			fmt.Printf("TradeHistory.Return[%v] := %#v\n", i, temp[i]);
		}
	}

}
