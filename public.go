package btce

import (
	"net/url"
	"io/ioutil"
	"encoding/json"
)

const (
	PUBLIC_API_ENDPOINT = "https://btc-e.com/api/3/";
)

type btcePublic struct {
	Client *clientTLS
}

type PublicInfo struct {
	ServerTime float64 `json:"server_time"`
	Pairs map[string] InfoPair `json:"pairs"`
}

type InfoPair struct {
	DecimalPlaces int `json:"decimal_places"`
	MinPrice float64 `json:"min_price"`
	MaxPrice float64 `json:"max_price"`
	MinAmount float64 `json:"min_amount"`
	Hidden int `json:"hidden"`
	Fee int `json:"fee"`
}

type Ticker map[string] TickerPair;

type TickerPair struct {
	High float64 `json:"high"`
	Low float64 `json:"low"`
	AVG float64 `json:"avg"`
	Vol float64 `json:"vol"`
	VolCur float64 `json:"vol_cur"`
	Last float64 `json:"last"`
	Buy float64 `json:"buy"`
	Sell float64 `json:"sell"`
	Updated int `json:"updated"`
}

type Depth map[string] DepthPair;

type DepthPair struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
}

type Trades map[string] TradesPair;

type TradesPair []TradesPairTrade;

type TradesPairTrade struct {
	Type string `json:"type"`
	Price float64 `json:"price"`
	Amount float64 `json:"amount"`
	TID int `json:"tid"`
	Timestamp int `json:"timestamp"` 
}

func NewPublic() *btcePublic {
	client := NewClient("", "", false);
	return &btcePublic{client};
}

func (this *btcePublic) Info() (*PublicInfo, error) {
	data := url.Values{};
	location := PUBLIC_API_ENDPOINT + "/info";
	response, err := this.Client.Request(data, location);
	if err != nil {
		return nil, err;
	}
	defer response.Body.Close();
	body, err := ioutil.ReadAll(response.Body);
	if err != nil {
		return nil, err;
	}
	Info := new(PublicInfo);
	err = json.Unmarshal([]byte(body), &Info);
	if err != nil {
		return nil, err;
	}
	return Info, nil;
}


func (this *btcePublic) Ticker(pair string) (*Ticker, error) {
	data := url.Values{};
	location := PUBLIC_API_ENDPOINT + "ticker/" + pair;
	response, err := this.Client.Request(data, location);
	if err != nil {
		return nil, err;
	}
	defer response.Body.Close();
	body, err := ioutil.ReadAll(response.Body);
	if err != nil {
		return nil, err;
	}
	Ticker := new(Ticker);
	err = json.Unmarshal([]byte(body), &Ticker);
	if err != nil {
		return nil, err;
	}
	return Ticker, nil;
}

func (this *btcePublic) Depth(pair string) (*Depth, error) {
	data := url.Values{};
	location := PUBLIC_API_ENDPOINT + "depth/" + pair;
	response, err := this.Client.Request(data, location);
	if err != nil {
		return nil, err;
	}
	defer response.Body.Close();
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

func (this *btcePublic) Trades(pair string) (*Trades, error) {
	data := url.Values{};
	location := PUBLIC_API_ENDPOINT + "trades/" + pair;
	response, err := this.Client.Request(data, location);
	if err != nil {
		return nil, err;
	}
	defer response.Body.Close();
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

