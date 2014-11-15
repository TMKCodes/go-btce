package btce

import (
	"net/url"
	"io/ioutil"
	"encoding/json"
)

const (
	PUBLIC_API_ENDPOINT = "https://btc-e.com/api/2/";
)

type btcePublic struct {
	Client *clientTLS
}

type fee struct {
	Trade float64 `json:"trade"`
}

type tmpTicker struct {
	Ticker ticker `json:"ticker"`
}

type ticker struct {
	High float64 `json:"high"`
	Low float64 `json:"low"`
	Avg float64 `json:"avg"`
	Vol float64 `json:"vol"`
	VolCur float64 `json:"vol_cur"`
	Last float64 `json:"last"`
	Buy float64 `json:"buy"`
	Sell float64 `json:"sell"`
	Updated int `json:"updated"`
	ServerTime int `json:"server_time"`
}

type trades []tradesTrade;

type tradesTrade struct {
	Date int `json:"data"`
	Price float64 `json:"price"`
	Amount float64 `json:"amount"`
	TID int `json:"tid"`
	PriceCurrency string `json:"price_currency"`
	Item string `json:"item"`
	TradeType string `json:"trade_type"`
}

type depth struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
}

func NewPublic() *btcePublic {
	client := NewClient("", "", false);
	return &btcePublic{client};
}


func (this *btcePublic) Fee(pair string) (*fee, error) {
	data := url.Values{};
	location := PUBLIC_API_ENDPOINT + pair + "/fee";
	response, err := this.Client.Request(data, location);
	defer response.Body.Close();
	if err != nil {
		return nil, err;
	}
	body, err := ioutil.ReadAll(response.Body);
	if err != nil {
		return nil, err;
	}
	fee := new(fee);
	err = json.Unmarshal([]byte(body), &fee);
	if err != nil {
		return nil, err;
	}
	return fee, nil;
}

func (this *btcePublic) Ticker(pair string) (*ticker, error) {
	data := url.Values{};
	location := PUBLIC_API_ENDPOINT + pair + "/ticker";
	response, err := this.Client.Request(data, location);
	defer response.Body.Close();
	if err != nil {
		return nil, err;
	}
	body, err := ioutil.ReadAll(response.Body);
	if err != nil {
		return nil, err;
	}
	tmpTicker := new(tmpTicker);
	err = json.Unmarshal([]byte(body), &tmpTicker);
	if err != nil {
		return nil, err;
	}
	return &tmpTicker.Ticker, nil;
}


func (this *btcePublic) Trades(pair string) (*trades, error) {
	data := url.Values{};
	location := PUBLIC_API_ENDPOINT + pair + "/trades";
	response, err := this.Client.Request(data, location);
	defer response.Body.Close();
	if err != nil {
		return nil, err;
	}
	body, err := ioutil.ReadAll(response.Body);
	if err != nil {
		return nil, err;
	}
	trades := new(trades);
	err = json.Unmarshal([]byte(body), &trades);
	if err != nil {
		return nil, err;
	}
	return trades, nil;
}

func (this *btcePublic) Depth(pair string) (*depth, error) {
	data := url.Values{};
	location := PUBLIC_API_ENDPOINT + pair + "/depth";
	response, err := this.Client.Request(data, location);
	defer response.Body.Close();
	if err != nil {
		return nil, err;
	}
	body, err := ioutil.ReadAll(response.Body);
	if err != nil {
		return nil, err;
	}
	depth := new(depth);
	err = json.Unmarshal([]byte(body), &depth);
	if err != nil {
		return nil, err;
	}
	return depth, nil;
}


