package stock_params

type CarveResult struct {
	Value    float64 `json:"value"`     // 值
	DateTime string  `json:"date_time"` // 日期时间
}

type BBICarveParams struct {
	StockCode       string         `json:"stock_code"`        // 股票代码
	BBICarveResults []*CarveResult `json:"bbi_carve_results"` // BBI指标
}

type KDJCarveParams struct {
	StockCode     string         `json:"stock_code"`      // 股票代码
	KCarveResults []*CarveResult `json:"k_carve_results"` // K值
	DCarveResults []*CarveResult `json:"d_carve_results"` // D值
	JCarveResults []*CarveResult `json:"j_carve_results"` // J值
}
