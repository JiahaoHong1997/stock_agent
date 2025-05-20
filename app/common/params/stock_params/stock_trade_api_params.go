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

// CompanyBasicInfo 公司基本信息
type CompanyBasicInfo struct {
	Name      string `json:"name"`      // 公司名称
	EName     string `json:"ename"`     // 公司英文名称
	Market    string `json:"market"`    // 上市市场
	Idea      string `json:"idea"`      // 概念及板块，多个概念由英文逗号分隔
	LDate     string `json:"ldate"`     // 上市日期，格式yyyy-MM-dd
	SPrice    string `json:"sprice"`    // 发行价格（元）
	Principal string `json:"principal"` // 主承销商
	RDate     string `json:"rdate"`     // 成立日期
	RPrice    string `json:"rprice"`    // 注册资本
	InstType  string `json:"instype"`   // 机构类型
	Organ     string `json:"organ"`     // 组织形式
	Secre     string `json:"secre"`     // 董事会秘书
	Phone     string `json:"phone"`     // 公司电话
	SPhone    string `json:"sphone"`    // 董秘电话
	Fax       string `json:"fax"`       // 公司传真
	SFax      string `json:"sfax"`      // 董秘传真
	Email     string `json:"email"`     // 公司电子邮箱
	SEmail    string `json:"semail"`    // 董秘电子邮箱
	Site      string `json:"site"`      // 公司网站
	Post      string `json:"post"`      // 邮政编码
	InfoSite  string `json:"infosite"`  // 信息披露网址
	OName     string `json:"oname"`     // 证券简称更名历史
	Addr      string `json:"addr"`      // 注册地址
	OAddr     string `json:"oaddr"`     // 办公地址
	Desc      string `json:"desc"`      // 公司简介
	BScope    string `json:"bscope"`    // 经营范围
	PrintType string `json:"printype"`  // 承销方式
	Referrer  string `json:"referrer"`  // 上市推荐人
	PuType    string `json:"putype"`    // 发行方式
	PE        string `json:"pe"`        // 发行市盈率（按发行后总股本）
	FirGu     string `json:"firgu"`     // 首发前总股本（万股）
	LastGu    string `json:"lastgu"`    // 首发后总股本（万股）
	RealGu    string `json:"realgu"`    // 实际发行量（万股）
	PlanM     string `json:"planm"`     // 预计募集资金（万元）
	RealM     string `json:"realm"`     // 实际募集资金合计（万元）
	PubFee    string `json:"pubfee"`    // 发行费用总额（万元）
	Collect   string `json:"collect"`   // 募集资金净额（万元）
	SignFee   string `json:"signfee"`   // 承销费用（万元）
	PDate     string `json:"pdate"`     // 招股公告日
}

type CompanyIndex struct {
	Mc   string `json:"mc"`   // 指数名称
	Dm   string `json:"dm"`   // 指数代码
	Ind  string `json:"ind"`  // 进入日期
	Outd string `json:"outd"` // 退出日期
}
