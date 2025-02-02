package gxtb

type Period int
type TradeCmd int
type TradeType int

const (
	PERIOD_M1  Period = 1
	PERIOD_M5  Period = 5
	PERIOD_M15 Period = 15
	PERIOD_M30 Period = 30
	PERIOD_H1  Period = 60
	PERIOD_H4  Period = 240
	PERIOD_D1  Period = 1440
	PERIOD_W1  Period = 10080
	PERIOD_MM1 Period = 43200
)

const (
	CMD_BUY TradeCmd = iota
	CMD_SELL
	CMD_BUY_LIMIT
	CMD_SELL_LIMIT
	CMD_BUY_STOP
	CMD_SELL_STOP
	CMD_BALANCE
	CMD_CREDIT
)

const (
	TYPE_OPEN TradeType = iota
	TYPE_PENDING
	TYPE_CLOSE
	TYPE_MODIFY
	TYPE_DELETE
)

type ChartLastInfo struct {
	Period Period `json:"period"`
	Start  int    `json:"start"`
	Symbol string `json:"symbol"`
}

type ChartRangeInfo struct {
	Period Period `json:"period"`
	Start  int    `json:"start"`
	End    int    `json:"end"`
	Symbol string `json:"symbol"`
	Ticks  int    `json:"ticks"`
}

type TransactionInfo struct {
	Cmd           TradeCmd  `json:"cmd"`
	CustomComment string    `json:"customComment"`
	Expiration    int64     `json:"expiration"`
	Offset        int       `json:"offset"`
	Order         int       `json:"order"`
	Price         float64   `json:"price"`
	Sl            float64   `json:"sl"`
	Symbol        string    `json:"symbol"`
	Tp            float64   `json:"tp"`
	Type          TradeType `json:"type"`
	Volume        float64   `json:"volume"`
}
