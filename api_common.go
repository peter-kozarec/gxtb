package gxtb

type LoginRequest struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
	AppId    string `json:"appId"`
	AppName  string `json:"appName"`
}

type LoginResponse struct {
	Status          bool   `json:"status"`
	StreamSessionId string `json:"streamSessionId"`
}

type SymbolRecord struct {
	Ask                float64  `json:"ask"`
	Bid                float64  `json:"bid"`
	CategoryName       string   `json:"categoryName"`
	ContractSize       int      `json:"contractSize"`
	Currency           string   `json:"currency"`
	CurrencyPair       bool     `json:"currencyPair"`
	CurrencyProfit     string   `json:"currencyProfit"`
	Description        string   `json:"description"`
	Expiration         *string  `json:"expiration"` // Nullable field
	GroupName          string   `json:"groupName"`
	High               float64  `json:"high"`
	InitialMargin      float64  `json:"initialMargin"`
	InstantMaxVolume   float64  `json:"instantMaxVolume"`
	Leverage           float64  `json:"leverage"`
	LongOnly           bool     `json:"longOnly"`
	LotMax             float64  `json:"lotMax"`
	LotMin             float64  `json:"lotMin"`
	LotStep            float64  `json:"lotStep"`
	Low                float64  `json:"low"`
	MarginHedged       float64  `json:"marginHedged"`
	MarginHedgedStrong bool     `json:"marginHedgedStrong"`
	MarginMaintenance  *float64 `json:"marginMaintenance"` // Nullable field
	MarginMode         int      `json:"marginMode"`
	Percentage         float64  `json:"percentage"`
	Precision          int      `json:"precision"`
	ProfitMode         int      `json:"profitMode"`
	QuoteId            int      `json:"quoteId"`
	ShortSelling       bool     `json:"shortSelling"`
	SpreadRaw          float64  `json:"spreadRaw"`
	SpreadTable        float64  `json:"spreadTable"`
	Starting           *string  `json:"starting"` // Nullable field
	StepRuleId         int      `json:"stepRuleId"`
	StopsLevel         float64  `json:"stopsLevel"`
	SwapRollover3Days  float64  `json:"swap_rollover3days"`
	SwapEnable         bool     `json:"swapEnable"`
	SwapLong           float64  `json:"swapLong"`
	SwapShort          float64  `json:"swapShort"`
	SwapType           int      `json:"swapType"`
	Symbol             string   `json:"symbol"`
	TickSize           float64  `json:"tickSize"`
	TickValue          float64  `json:"tickValue"`
	Time               int64    `json:"time"`
	TimeString         string   `json:"timeString"`
	TrailingEnabled    bool     `json:"trailingEnabled"`
	Type               int      `json:"type"`
}

type CalendarRecord struct {
	Country  string `json:"country"`
	Current  string `json:"current"`
	Forecast string `json:"forecast"`
	Impact   string `json:"impact"`
	Period   string `json:"period"`
	Previous string `json:"previous"`
	Time     int64  `json:"time"`
	Title    string `json:"title"`
}

type ChartLastRequest struct {
	Period int    `json:"period"`
	Start  int64  `json:"start"`
	Symbol string `json:"symbol"`
}

type ChartRangeRequest struct {
	End    int64  `json:"end"`
	Period int    `json:"period"`
	Start  int64  `json:"start"`
	Symbol string `json:"symbol"`
	Ticks  int    `json:"ticks"`
}

type RateInfoRecord struct {
	Close     float64 `json:"close"`
	Ctm       int64   `json:"ctm"`
	CtmString string  `json:"ctmString"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Open      float64 `json:"open"`
	Vol       float64 `json:"vol"`
}

type CommissionRequest struct {
	Symbol string  `json:"symbol"`
	Volume float32 `json:"volume"`
}

type CommissionData struct {
	Commission     float32 `json:"commission"`
	RateOfExchange float32 `json:"rateOfExchange"`
}

type UserData struct {
	CompanyUnit        int     `json:"companyUnit"`
	Currency           string  `json:"currency"`
	Group              string  `json:"group"`
	IbAccount          bool    `json:"ibAccount"`
	Leverage           int     `json:"leverage"`
	LeverageMultiplier float64 `json:"leverageMultiplier"`
	SpreadType         string  `json:"spreadType"`
	TrailingStop       bool    `json:"trailingStop"`
}

type MarginData struct {
	Balance     float64 `json:"balance"`
	Credit      float64 `json:"credit"`
	Currency    string  `json:"currency"`
	Equity      float64 `json:"equity"`
	Margin      float64 `json:"margin"`
	MarginFree  float64 `json:"margin_free"`
	MarginLevel float64 `json:"margin_level"`
}
