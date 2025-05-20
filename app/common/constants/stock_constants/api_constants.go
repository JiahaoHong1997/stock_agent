package stock_constants

// 账户 license
const License = "BF9D6459-8F2C-4417-9FCC-0347F55C47C9"

const (
	StockRealTimeInfoURL      = "https://api.mairuiapi.com/hsstock/real/time/%s/%s"                 // 获取股票实时交易数据接口
	StockRealTimeBatchInfoURL = "http://api.mairuiapi.com/hsrl/ssjy_more/%s?stock_codes=%s"         // 批量获取股票实时交易数据接口
	HistoryStockTradeInfoURL  = "https://api.mairuiapi.com/hsstock/history/%s/%s/%s/%s?st=%s&et=%s" // 获取历史交易数据接口
	GetCompanyBasicInfoURL    = "http://api.mairuiapi.com/hscp/gsjj/%s/%s"                          // 获取公司基本信息接口
	GetCompanyIndexURL        = "http://api.mairuiapi.com/hscp/sszs/%s/%s"                          // 获取公司所属指数接口
)
