package btce;

type Btce struct {
	Public *btcePublic
	Private *btcePrivate
}

func New(public string, secret string) *Btce {
	pub := NewPublic();
	priv := NewPrivate(public, secret);
	return &Btce{pub, priv};
}

func (this *Btce) GetInfo() (*Info, error) {
	return this.Private.GetInfo();
}

func (this *Btce) TransHistory(From int, Count int, FromID int, EndID int, Order string, Since string, End string) (*TransHistory, error) {
	return this.Private.TransHistory(From, Count, FromID, EndID, Order, Since, End);
}

func (this *Btce) TradeHistory(From int, Count int, FromID int, EndID int, Order string, Since string, End string, Pair string) (*TradeHistory, error) {
	return this.Private.TradeHistory(From, Count, FromID, EndID, Order, Since, End, Pair);
}

func (this *Btce) TradeHistoryDefault(Pair string) (*TradeHistory, error) {
	return this.Private.TradeHistoryDefault(Pair);
}

func (this *Btce) ActiveOrders(Pair string) (*ActiveOrders, error) {
	return this.Private.ActiveOrders(Pair);
}

func (this *Btce) Trade(Pair string, Type string, Rate float64, Amount float64) (*Trade, error) {
	return this.Private.Trade(Pair, Type, Rate, Amount);
}

func (this *Btce) CancelOrder(OrderID int) (*CancelOrder, error) {
	return this.Private.CancelOrder(OrderID);
}

func (this *Btce) Fee(Pair string) (*Fee, error) {
	return this.Public.Fee(Pair);
}

func (this *Btce) Ticker(Pair string) (*Ticker, error) {
	return this.Public.Ticker(Pair);
}

func (this *Btce) Trades(Pair string) (*Trades, error) {
	return this.Public.Trades(Pair);
}

func (this *Btce) Depth(Pair string) (*Depth, error) {
	return this.Public.Depth(Pair);
}
