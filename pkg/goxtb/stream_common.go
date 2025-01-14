package goxtb

type RecordType string

const (
	InvalidRecordType     RecordType = "InvalidRecord"
	BalanceRecordType     RecordType = "BalanceRecord"
	CandlesRecordType     RecordType = "CandlesRecord"
	KeepAliveRecordType   RecordType = "KeepAliveRecord"
	NewsRecordType        RecordType = "NewsRecord"
	ProfitsRecordType     RecordType = "ProfitsRecord"
	TickPricesRecordType  RecordType = "TickPricesRecord"
	TradesRecordType      RecordType = "TradesRecord"
	TradeStatusRecordType RecordType = "TradeStatusRecord"
)

type Record interface {
	Type() RecordType
}

type RecordMessage struct {
	Value Record
	Err   error
}

type BalanceRecord struct {
	Balance     float32 `json:"balance"`
	Credit      float32 `json:"credit"`
	Equity      float32 `json:"equity"`
	Margin      float32 `json:"margin"`
	MarginFree  float32 `json:"marginFree"`
	MarginLevel float32 `json:"marginLevel"`
}

func (r BalanceRecord) Type() RecordType {
	return BalanceRecordType
}

type CandleRecord struct {
	Close     float32 `json:"close"`
	Ctm       int64   `json:"ctm"`
	CtmString string  `json:"ctmString"`
	High      float32 `json:"high"`
	Low       float32 `json:"low"`
	Open      float32 `json:"open"`
	QuoteId   int     `json:"quoteId"`
	Symbol    string  `json:"symbol"`
	Volume    float32 `json:"vol"`
}

func (r CandleRecord) Type() RecordType {
	return CandlesRecordType
}

type KeepAliveRecord struct {
	Timestamp int64 `json:"timestamp"`
}

func (r KeepAliveRecord) Type() RecordType {
	return KeepAliveRecordType
}

type NewsRecord struct {
	Body  string `json:"body"`
	Key   string `json:"key"`
	Time  int64  `json:"time"`
	Title string `json:"title"`
}

func (r NewsRecord) Type() RecordType {
	return NewsRecordType
}

type ProfitsRecord struct {
	Order    int     `json:"order"`
	Order2   int     `json:"order2"`
	Position int     `json:"position"`
	Profit   float32 `json:"profit"`
}

func (r ProfitsRecord) Type() RecordType {
	return ProfitsRecordType
}

type TickPricesRecord struct {
	Ask         float32 `json:"ask"`
	AskVolume   int     `json:"askVolume"`
	Bid         float32 `json:"bid"`
	BidVolume   int     `json:"bidVolume"`
	High        float32 `json:"high"`
	Level       int     `json:"level"`
	Low         float32 `json:"low"`
	QuoteId     int     `json:"quoteId"`
	SpreadRaw   float32 `json:"spreadRaw"`
	SpreadTable float32 `json:"spreadTable"`
	Symbol      string  `json:"symbol"`
	Timestamp   int64   `json:"timestamp"`
}

func (r TickPricesRecord) Type() RecordType {
	return TickPricesRecordType
}

type TradesRecord struct {
	ClosePrice    float32  `json:"close_price"`
	CloseTime     *int64   `json:"close_time"`
	Closed        bool     `json:"closed"`
	Cmd           int      `json:"cmd"`
	Comment       string   `json:"comment"`
	Commission    float32  `json:"commission"`
	CustomComment string   `json:"customComment"`
	Digits        int      `json:"digits"`
	Expiration    *int64   `json:"expiration"`
	MarginRate    float32  `json:"margin_rate"`
	Offset        int      `json:"offset"`
	OpenPrice     float32  `json:"open_price"`
	OpenTime      int64    `json:"open_time"`
	Order         int      `json:"order"`
	Order2        int      `json:"order2"`
	Position      int      `json:"position"`
	Profit        *float32 `json:"profit"`
	StopLoss      float32  `json:"sl"`
	State         string   `json:"state"`
	Storage       float32  `json:"storage"`
	Symbol        string   `json:"symbol"`
	TakeProfit    float32  `json:"tp"`
	TradeType     int      `json:"type"`
	Volume        float32  `json:"volume"`
}

func (r TradesRecord) Type() RecordType {
	return TradesRecordType
}

type TradeStatusRecord struct {
	CustomComment string  `json:"customComment"`
	Message       *string `json:"message"`
	Order         int     `json:"order"`
	Price         float32 `json:"price"`
	RequestStatus int     `json:"requestStatus"`
}

func (r TradeStatusRecord) Type() RecordType {
	return TradeStatusRecordType
}
