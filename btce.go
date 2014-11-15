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

func (this *btce) GetInfo() (*info, error) {
	return this.Private.GetInfo();
}

func (this *btce) TransHistory(From int, Count int, FromID int, EndID int, Order string, Since string, End string) (*transHistory, error) {
	return this.Private.TransHistory(From, Count, FromID, EndID, Order, Since, End);
}

func (this *btce) TradeHistory(From int, Count int, FromID int, EndID int, Order string, Since string, End string, Pair string) (*tradeHistory, error) {
	return this.Private.TradeHistory(From, Count, FromID, EndID, Order, Since, End, Pair);
}

func (this *btce) ActiveOrders(Pair string) (*activeOrders, error) {
	return this.Private.ActiveOrders(Pair);
}

func (this *btce) Trade(Pair string, Type string, Rate float64, Amount float64) (*trade, error) {
	return this.Private.Trade(Pair, Type, Rate, Amount);
}

func (this *btce) CancelOrder(OrderID int) (*cancelOrder, error) {
	return this.Private.CancelOrder(OrderID);
}

func (this *btce) Fee(Pair string) (*fee, error) {
	return this.Public.Fee(Pair);
}

func (this *btce) Ticker(Pair string) (*ticker, error) {
	return this.Public.Ticker(Pair);
}

func (this *btce) Trades(Pair string) (*trades, error) {
	return this.Public.Trades(Pair);
}

func (this *btce) Depth(Pair string) (*depth, error) {
	return this.Public.Depth(Pair);
}
