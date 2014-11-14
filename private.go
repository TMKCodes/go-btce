package btce

import (
	"time"
	"strconv"
	"net/url"
	"io/ioutil"
	"encoding/json"
)

type btcePrivate struct {
	Client *clientTLS
}

type info struct {
	Success int `json:"success"`
	Return infoReturn `json:"return"`
}

type infoReturn struct {
	Funds infoFunds `json:"funds"`
	Rights infoRights `json:"rights"`
	TransactionsCount int `json:"transaction_count"`
	OpenOrders int `json:"open_orders"`
	Time int `json:"server_time"`
}

type infoRights struct {
	Info int `json:"info"`
	Trade int `json:"trade"`
	Withdraw int `json:"withdraw"`
}

type infoFunds struct {
	USD int `json:"usd"`
	BTC int `json:"btc"`
	LTC int `json:"ltc"`
	NMC int `json:"nmc"`
	RUR int `json:"rur"`
	EUR int `json:"eur"`
	NVC int `json:"nvc"`
	TRC int `json:"trc"`
	PPC int `json:"ppc"`
	FTC int `json:"ftc"`
	XPM int `json:"xpm"`
}

type transHistory struct {
	Success int `json:"success"`
	Return transHistoryReturn `json:"return"`
}

type transHistoryReturn map[string]transHistoryTransaction;

type transHistoryTransaction struct {
	Type int `json:"type"`
	Amount float32 `json:"amount"`
	Currency string `json:"currency"`
	Desc string `json:"desc"`
	Status int `json:"status"`
	Timestamp int `json:"timestamp"`
}

type tradeHistory struct {
	Success int `json:"success"`
	Return tradeHistoryReturn `json:"return"`
}

type tradeHistoryReturn map[string]tradeHistoryTrade;

type tradeHistoryTrade struct {
	Pair string `json:"pair"`
	Type string `json:"type"`
	Amount float32 `json:"amount"`
	Rate float32 `json:"rate"`
	OrderID int `json:"order_id"`
	IsYourOrder int `json:"is_your_order"`
	Timestamp int `json:"timestamp"`
}

type trade struct {
	Success int `json:"success"`
	Return tradeReturn `json:"return"`
}

type tradeReturn struct {
	Received float32 `json:"received"`
	Remains float32 `json:"remains"`
	Order int `json:"order_id"`
	Funds infoFunds `json:"funds"`
}

type cancelOrder struct {
	Success int `json:"success"`
	Return cancelOrderReturn `json:"return"`
}

type cancelOrderReturn struct {
	OrderID int `json:"order_id"`
	Funds infoFunds `json:"funds"`
}

func New(public string, private string) *btcePrivate {
	client := NewClient(public, private);
	return &btcePrivate{client};
}


func (this *btcePrivate) getInfo() (*info, error) {
	data := url.Values{};
	data.Add("method", "getInfo");
	data.Add("nonce", strconv.Itoa(int(time.Now().Unix())));
	response, err := this.Client.Request(data);
	defer response.Body.Close();
	if err != nil {
		return nil, err;
	}
	body, err := ioutil.ReadAll(response.Body);
	if err != nil {
		return nil, err;
	}
	info := new(info);
	err = json.Unmarshal([]byte(body), &info);
	if err != nil {
		return nil, err;
	}
	return info, nil;
}

func (this *btcePrivate) TransHistory(From int, Count int, FromID int, EndID int, Order string, Since string, End string) (*transHistory, error) {
	data := url.Values{};
	data.Add("method", "TransHistory");
	data.Add("nonce", strconv.Itoa(int(time.Now().Unix())));
	if From != nil {
		data.Add("from", strconv.Itoa(From));
	}
	if Count != nil {
		data.Add("count", strconv.Itoa(Count));
	}
	if FromID != nil {
		data.Add("from_id", strconv.Itoa(FromID));
	}
	if EndID != nil {
		data.Add("end_id", strconv.Itoa(EndID));
	}
	if Order != nil {
		data.Add("order", Order);
	}
	if Since != nil {
		data.Add("since", Since);
	}
	if End != nil {
		data.Add("end", End);
	}
	response, err := this.Client.Request(data);
	defer response.Body.Close();
	if err != nil {
		return nil, err;
	}
	body, err := ioutil.ReadAll(response.Body);
	if err != nil {
		return nil, err;
	}
	transHistory := new(transHistory);
	err = json.Unmarshal([]byte(body), &transHistory);
	if err != nil {
		return nil, err;
	}
	return transHistory, nil;
}

func (this *btcePrivate) TradeHistory() (*tradeHistory, error) {
	data := url.Values{};
	data.Add("method", "TradeHistory");
	data.Add("nonce", strconv.Itoa(int(time.Now().Unix())));
	response, err := this.Client.Request(data);
	defer response.Body.Close();
	if err != nil {
		return nil, err;
	}
	body, err := ioutil.ReadAll(response.Body);
	if err != nil {
		return nil, err;
	}
	tradeHistory := new(tradeHistory);
	err = json.Unmarshal([]byte(body), &tradeHistory);
	if err != nil {
		return nil, err;
	}
	return tradeHistory, nil;
}

func (this *btcePrivate) Trade(Pair string, Type string, Rate float64, Amount float64) (*trade, error) {
	data := url.Values{};
	data.Add("method", "Trade");
	data.Add("nonce", strconv.Itoa(int(time.Now().Unix())));
	data.Add("pair", Pair);
	data.Add("type", Type);
	data.Add("rate", strconv.FormatFloat(Rate, 'f', 6, 32));
	data.Add("amount", strconv.FormatFloat(Amount, 'f', 6, 32));
	response, err := this.Client.Request(data);
	defer response.Body.Close();
	if err != nil {
		return nil, err;
	}
	body, err := ioutil.ReadAll(response.Body);
	if err != nil {
		return nil, err;
	}
	trade := new(trade);
	err = json.Unmarshal([]byte(body), &trade);
	if err != nil {
		return nil, err;
	}
	return trade, nil;
}

func (this *btcePrivate) CancelOrder(Order int) (*cancelOrder, error) {
	data := url.Values{};
	data.Add("method", "CancelOrder");
	data.Add("nonce", strconv.Itoa(int(time.Now().Unix())));
	data.Add("order_id", strconv.Itoa(Order));
	response, err := this.Client.Request(data);
	defer response.Body.Close();
	if err != nil {
		return nil, err;
	}
	body, err := ioutil.ReadAll(response.Body);
	if err != nil {
		return nil, err;
	}
	cancelOrder := new(cancelOrder);
	err = json.Unmarshal([]byte(body), &cancelOrder);
	if err != nil {
		return nil, err;
	}
	return cancelOrder, nil;

}
