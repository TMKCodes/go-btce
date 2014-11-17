package btce;

type btce struct {
	Public *btcePublic
	Private *btcePrivate
}

func New(public string, secret string) *btce {
	pub := NewPublic();
	priv := NewPrivate(public, secret);
	return &btce{pub, priv};
}

func (this *btce) GetInfo() (*Info, error) {
	return this.Private.GetInfo();
}

func (this *btce) TransHistory(From int, Count int, FromID int, EndID int, Order string, Since string, End string) (*TransHistory, error) {
	return this.Private.TransHistory(From, Count, FromID, EndID, Order, Since, End);
}

func (this *btce) TradeHistory(From int, Count int, FromID int, EndID int, Order string, Since string, End string, Pair string) (*TradeHistory, error) {
	return this.Private.TradeHistory(From, Count, FromID, EndID, Order, Since, End, Pair);
}

func (this *btce) TradeHistoryDefault(Pair string) (*TradeHistory, error) {
	return this.Private.TradeHistoryDefault(Pair);
}

func (this *btce) ActiveOrders(Pair string) (*ActiveOrders, error) {
	return this.Private.ActiveOrders(Pair);
}

func (this *btce) Trade(Pair string, Type string, Rate float64, Amount float64) (*Trade, error) {
	return this.Private.Trade(Pair, Type, Rate, Amount);
}

func (this *btce) CancelOrder(OrderID int) (*CancelOrder, error) {
	return this.Private.CancelOrder(OrderID);
}

func (this *btce) Fee(Pair string) (*Fee, error) {
	return this.Public.Fee(Pair);
}

func (this *btce) Ticker(Pair string) (*Ticker, error) {
	return this.Public.Ticker(Pair);
}

func (this *btce) Trades(Pair string) (*Trades, error) {
	return this.Public.Trades(Pair);
}

func (this *btce) Depth(Pair string) (*Depth, error) {
	return this.Public.Depth(Pair);
}
