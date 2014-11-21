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

type Info struct {
	Success int `json:"success"`
	Return InfoReturn `json:"return"`
	Error string `json:"error"`
}

type InfoReturn struct {
	Funds InfoFunds `json:"funds"`
	Rights InfoRights `json:"rights"`
	TransactionsCount int `json:"Transaction_count"`
	OpenOrders int `json:"open_orders"`
	Time int `json:"server_time"`
}

type InfoRights struct {
	Info int `json:"Info"`
	Trade int `json:"Trade"`
	Withdraw int `json:"withdraw"`
}

type InfoFunds struct {
	USD float64 `json:"usd"`
	BTC float64 `json:"btc"`
	LTC float64 `json:"ltc"`
	NMC float64 `json:"nmc"`
	RUR float64 `json:"rur"`
	EUR float64 `json:"eur"`
	NVC float64 `json:"nvc"`
	TRC float64 `json:"trc"`
	PPC float64 `json:"ppc"`
	FTC float64 `json:"ftc"`
	XPM float64 `json:"xpm"`
}

type TransHistory struct {
	Success int `json:"success"`
	Return TransHistoryReturn `json:"return"`
	Error string `json:"error"`
}

type TransHistoryReturn map[string]TransHistoryTransaction;

type TransHistoryTransaction struct {
	Type int `json:"type"`
	Amount float64 `json:"amount"`
	Currency string `json:"currency"`
	Desc string `json:"desc"`
	Status int `json:"status"`
	Timestamp int `json:"timestamp"`
}

type TradeHistory struct {
	Success int `json:"success"`
	Return TradeHistoryReturn `json:"return"`
	Error string `json:"error"`
}

type TradeHistoryReturn map[string]TradeHistoryTrade;

type TradeHistoryTrade struct {
	Pair string `json:"pair"`
	Type string `json:"type"`
	Amount float64 `json:"amount"`
	Rate float64 `json:"rate"`
	OrderID int `json:"order_id"`
	IsYourOrder int `json:"is_your_order"`
	Timestamp int `json:"timestamp"`
}

type Trade struct {
	Success int `json:"success"`
	Return TradeReturn `json:"return"`
	Error string `json:"error"`
}

type TradeReturn struct {
	Received float64 `json:"received"`
	Remains float64 `json:"remains"`
	Order int `json:"order_id"`
	Funds InfoFunds `json:"funds"`
}

type ActiveOrders struct {
	Success int `json:"success"`
	Return ActiveOrdersReturn `json:"return"`
	Error string `json:"error"`
}

type ActiveOrdersReturn map[string]ActiveOrdersOrder;

type ActiveOrdersOrder struct {
	Pair string `json:"pair"`
	Type string `json:"type"`
	Amount float64 `json:"amount"`
	Rate float64 `json:"rate"`
	OrderID int `json:"order_id"`
	TimestampCreated int `json:"timestamp_created"`
	Status int `json:"status"`
}

type CancelOrder struct {
	Success int `json:"success"`
	Return CancelOrderReturn `json:"return"`
	Error string `json:"error"`
}

type CancelOrderReturn struct {
	OrderID int `json:"order_id"`
	Funds InfoFunds `json:"funds"`
}

func NewPrivate(public string, private string) *btcePrivate {
	client := NewClient(public, private, true);
	return &btcePrivate{client};
}


func (this *btcePrivate) GetInfo() (*Info, error) {
	data := url.Values{};
	data.Add("method", "getInfo");
	data.Add("nonce", strconv.Itoa(int(time.Now().Unix())));
	response, err := this.Client.Request(data, "");
	if err != nil {
		return nil, err;
	}
	defer response.Body.Close();
	body, err := ioutil.ReadAll(response.Body);
	if err != nil {
		return nil, err;
	}
	Info := new(Info);
	err = json.Unmarshal([]byte(body), &Info);
	if err != nil {
		return nil, err;
	}
	return Info, nil;
}

func (this *btcePrivate) TransHistory(From int, Count int, FromID int, EndID int, Order string, Since string, End string) (*TransHistory, error) {
	data := url.Values{};
	data.Add("method", "TransHistory");
	data.Add("nonce", strconv.Itoa(int(time.Now().Unix())));
	data.Add("from", strconv.Itoa(From));
	data.Add("count", strconv.Itoa(Count));
	data.Add("from_id", strconv.Itoa(FromID));
	data.Add("end_id", strconv.Itoa(EndID));
	data.Add("order", Order);
	data.Add("since", Since);
	data.Add("end", End);
	response, err := this.Client.Request(data, "");
	if err != nil {
		return nil, err;
	}
	defer response.Body.Close();
	body, err := ioutil.ReadAll(response.Body);
	if err != nil {
		return nil, err;
	}
	TransHistory := new(TransHistory);
	err = json.Unmarshal([]byte(body), &TransHistory);
	if err != nil {
		return nil, err;
	}
	return TransHistory, nil;
}

func (this *btcePrivate) TradeHistory(From int, Count int, FromID int, EndID int, Order string, Since string, End string, Pair string) (*TradeHistory, error) {
	data := url.Values{};
	data.Add("method", "TradeHistory");
	data.Add("nonce", strconv.Itoa(int(time.Now().Unix())));
	data.Add("from", strconv.Itoa(From));
	data.Add("count", strconv.Itoa(Count));
	data.Add("from_id", strconv.Itoa(FromID));
	data.Add("end_id", strconv.Itoa(EndID));
	data.Add("order", Order);
	data.Add("since", Since);
	data.Add("end", End);
	data.Add("pair", Pair);
	response, err := this.Client.Request(data, "");
	if err != nil {
		return nil, err;
	}
	defer response.Body.Close();
	body, err := ioutil.ReadAll(response.Body);
	if err != nil {
		return nil, err;
	}
	TradeHistory := new(TradeHistory);
	err = json.Unmarshal([]byte(body), &TradeHistory);
	if err != nil {
		return nil, err;
	}
	return TradeHistory, nil;
}

func (this *btcePrivate) TradeHistoryDefault(Pair string) (*TradeHistory, error) {
	data := url.Values{};
	data.Add("method", "TradeHistory");
	data.Add("nonce", strconv.Itoa(int(time.Now().Unix())));
	data.Add("pair", Pair);
	response, err := this.Client.Request(data, "");
	if err != nil {
		return nil, err;
	}
	defer response.Body.Close();
	body, err := ioutil.ReadAll(response.Body);
	if err != nil {
		return nil, err;
	}
	TradeHistory := new(TradeHistory);
	err = json.Unmarshal([]byte(body), &TradeHistory);
	if err != nil {
		return nil, err;
	}
	return TradeHistory, nil;
}

func (this *btcePrivate) ActiveOrders(Pair string) (*ActiveOrders, error) {
	data := url.Values{};
	data.Add("method", "Trade");
	data.Add("nonce", strconv.Itoa(int(time.Now().Unix())));
	data.Add("pair", Pair);
	response, err := this.Client.Request(data, "");
	if err != nil {
		return nil, err;
	}
	defer response.Body.Close();
	body, err := ioutil.ReadAll(response.Body);
	if err != nil {
		return nil, err;
	}
	ActiveOrders := new(ActiveOrders);
	err = json.Unmarshal([]byte(body), &ActiveOrders);
	if err != nil {
		return nil, err;
	}
	return ActiveOrders, nil;

}

func (this *btcePrivate) Trade(Pair string, Type string, Rate float64, Amount float64) (*Trade, error) {
	data := url.Values{};
	data.Add("method", "Trade");
	data.Add("nonce", strconv.Itoa(int(time.Now().Unix())));
	data.Add("pair", Pair);
	data.Add("type", Type);
	data.Add("rate", strconv.FormatFloat(Rate, 'f', 5, 64));
	data.Add("amount", strconv.FormatFloat(Amount, 'f', 8, 64));
	response, err := this.Client.Request(data, "");
	if err != nil {
		return nil, err;
	}
	defer response.Body.Close();
	body, err := ioutil.ReadAll(response.Body);
	if err != nil {
		return nil, err;
	}
	Trade := new(Trade);
	err = json.Unmarshal([]byte(body), &Trade);
	if err != nil {
		return nil, err;
	}
	return Trade, nil;
}

func (this *btcePrivate) CancelOrder(OrderID int) (*CancelOrder, error) {
	data := url.Values{};
	data.Add("method", "CancelOrder");
	data.Add("nonce", strconv.Itoa(int(time.Now().Unix())));
	data.Add("order_id", strconv.Itoa(OrderID));
	response, err := this.Client.Request(data, "");
	if err != nil {
		return nil, err;
	}
	defer response.Body.Close();
	body, err := ioutil.ReadAll(response.Body);
	if err != nil {
		return nil, err;
	}
	CancelOrder := new(CancelOrder);
	err = json.Unmarshal([]byte(body), &CancelOrder);
	if err != nil {
		return nil, err;
	}
	return CancelOrder, nil;

}
