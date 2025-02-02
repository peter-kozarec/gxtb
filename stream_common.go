package gxtb

type Balance struct {
	Balance     float64 `json:"balance"`
	Credit      float64 `json:"credit"`
	Equity      float64 `json:"equity"`
	Margin      float64 `json:"margin"`
	MarginFree  float64 `json:"marginFree"`
	MarginLevel float64 `json:"marginLevel"`
}

type Candle struct {
	Close     float64 `json:"close"`
	Ctm       int64   `json:"ctm"`
	CtmString string  `json:"ctmString"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Open      float64 `json:"open"`
	QuoteId   int     `json:"quoteId"`
	Symbol    string  `json:"symbol"`
	Volume    float64 `json:"vol"`
}

type KeepAlive struct {
	Timestamp int64 `json:"timestamp"`
}

type News struct {
	Body  string `json:"body"`
	Key   string `json:"key"`
	Time  int64  `json:"time"`
	Title string `json:"title"`
}

type Profit struct {
	Order    int     `json:"order"`
	Order2   int     `json:"order2"`
	Position int     `json:"position"`
	Profit   float64 `json:"profit"`
}

type TickPrice struct {
	Ask         float64 `json:"ask"`
	AskVolume   int     `json:"askVolume"`
	Bid         float64 `json:"bid"`
	BidVolume   int     `json:"bidVolume"`
	High        float64 `json:"high"`
	Level       int     `json:"level"`
	Low         float64 `json:"low"`
	QuoteId     int     `json:"quoteId"`
	SpreadRaw   float64 `json:"spreadRaw"`
	SpreadTable float64 `json:"spreadTable"`
	Symbol      string  `json:"symbol"`
	Timestamp   int64   `json:"timestamp"`
}

type Trade struct {
	ClosePrice    float64  `json:"close_price"`
	CloseTime     *int64   `json:"close_time"`
	Closed        bool     `json:"closed"`
	Cmd           int      `json:"cmd"`
	Comment       string   `json:"comment"`
	Commission    float64  `json:"commission"`
	CustomComment string   `json:"customComment"`
	Digits        int      `json:"digits"`
	Expiration    *int64   `json:"expiration"`
	MarginRate    float64  `json:"margin_rate"`
	Offset        int      `json:"offset"`
	OpenPrice     float64  `json:"open_price"`
	OpenTime      int64    `json:"open_time"`
	Order         int      `json:"order"`
	Order2        int      `json:"order2"`
	Position      int      `json:"position"`
	Profit        *float64 `json:"profit"`
	StopLoss      float64  `json:"sl"`
	State         string   `json:"state"`
	Storage       float64  `json:"storage"`
	Symbol        string   `json:"symbol"`
	TakeProfit    float64  `json:"tp"`
	TradeType     int      `json:"type"`
	Volume        float64  `json:"volume"`
}

type TradeStatus struct {
	CustomComment string  `json:"customComment"`
	Message       *string `json:"message"`
	Order         int     `json:"order"`
	Price         float64 `json:"price"`
	RequestStatus int     `json:"requestStatus"`
}
