package young_wife_tactics

import "agent/app/common/params/stock_params"

type YoungWifeTactics struct {
	Ctx   *stock_params.StockTradeContext // 股票交易上下文
	KdjN  int                             // kdj N值
	KdjM1 int                             // kdj M1值
	KdjM2 int                             // kdj M2值
}

func NewYoungWifeTactics() *YoungWifeTactics {
	return &YoungWifeTactics{}
}
