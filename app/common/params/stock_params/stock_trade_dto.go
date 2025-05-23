package stock_params

type StockTradeContext struct {
	StockCode      string                           // 股票代码
	MarketCode     string                           // 市场代码
	StockTradeInfo []*HistoryStockTradeInfoResponse // 历史股票交易信息
	ClosePrices    []float64                        // 收盘价
	HighPrices     []float64                        // 最高价
	LowPrices      []float64                        // 最低价
	JValues        []float64                        // J值
	BbiValues      []float64                        // BBI值
}
