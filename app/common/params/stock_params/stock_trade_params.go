package stock_params

// StockRealTimeInfoResponse 股票实时信息响应结构体
type StockRealTimeInfoResponse struct {
	NewestPrice              float64 `json:"p"`   // 最新价格
	OpenPrice                float64 `json:"o"`   // 开盘价
	HighestPrice             float64 `json:"h"`   // 最高价
	LowestPrice              float64 `json:"l"`   // 最低价
	YesterdayPrice           float64 `json:"yc"`  // 昨收价
	TotalTransactionAmount   float64 `json:"cje"` // 成交总额
	TotalTradingVolume       float64 `json:"v"`   // 成交总量（成交量单位：元）
	OriginTotalTradingVolume float64 `json:"pv"`  // 原始成交总量（成交量单位：分）
	UpdateTime               string  `json:"t"`   // 更新时间
}

type HistoryStockTradeInfoResponse struct {
	TradeTime    string  `json:"t"`  // 交易时间
	OpenPrice    float64 `json:"o"`  // 开盘价
	HighestPrice float64 `json:"h"`  // 最高价
	LowestPrice  float64 `json:"l"`  // 最低价
	ClosePrice   float64 `json:"c"`  // 收盘价
	Volume       float64 `json:"v"`  // 成交量
	Amount       float64 `json:"a"`  // 成交额
	PrePrice     float64 `json:"yc"` // 前收盘价
	Stop         int8    `json:"sf"` // 停牌 1停牌，0 不停牌
}
