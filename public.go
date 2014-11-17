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

type Fee struct {
	Trade float64 `json:"trade"`
}

type tmpTicker struct {
	Ticker Ticker `json:"ticker"`
}

type Ticker struct {
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

type Trades []TradesTrade;

type TradesTrade struct {
	Date int `json:"data"`
	Price float64 `json:"price"`
	Amount float64 `json:"amount"`
	TID int `json:"tid"`
	PriceCurrency string `json:"price_currency"`
	Item string `json:"item"`
	TradeType string `json:"trade_type"`
}

type Depth struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
}

func NewPublic() *btcePublic {
	client := NewClient("", "", false);
	return &btcePublic{client};
}


func (this *btcePublic) Fee(pair string) (*Fee, error) {
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
	Fee := new(Fee);
	err = json.Unmarshal([]byte(body), &Fee);
	if err != nil {
		return nil, err;
	}
	return Fee, nil;
}

func (this *btcePublic) Ticker(pair string) (*Ticker, error) {
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


func (this *btcePublic) Trades(pair string) (*Trades, error) {
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
	Trades := new(Trades);
	err = json.Unmarshal([]byte(body), &Trades);
	if err != nil {
		return nil, err;
	}
	return Trades, nil;
}

func (this *btcePublic) Depth(pair string) (*Depth, error) {
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
	Depth := new(Depth);
	err = json.Unmarshal([]byte(body), &Depth);
	if err != nil {
		return nil, err;
	}
	return Depth, nil;
}


