package stock_info

import (
	"agent/app/common/enums/stock_enums"
	"agent/app/common/params/stock_params"
	"agent/app/infra/stock/stock_indicators"
	"agent/app/utils"
	"context"
	"testing"
	"time"
)

func TestGetStockTradeInfo(t *testing.T) {
	ctx := context.Background()
	info, err := GetStockTradeInfo(ctx, "600600")
	if err != nil {
		t.Fatalf("获取股票交易信息失败: %v", err)
		return
	}
	t.Logf("获取股票交易信息成功: %v", utils.ToJsonString(info))
}

func TestBatchGetStockTradeInfo(t *testing.T) {
	ctx := context.Background()
	info, err := BatchGetStockTradeInfo(ctx, []string{"600600", "002594"})
	if err != nil {
		t.Fatalf("获取股票交易信息失败: %v", err)
		return
	}
	t.Logf("获取股票交易信息成功: %v", utils.ToJsonString(info))
}

func TestGetHistoryStockTradeInfo(t *testing.T) {
	stockCode := "600600"
	market := "SH"
	const (
		M1    = 3
		M2    = 6
		M3    = 12
		M4    = 24
		N     = 9
		KdjM1 = 3
		KdjM2 = 3
	)

	ctx := context.Background()

	// 获取历史股票交易信息
	info, err := GetHistoryStockTradeInfo(ctx, stockCode, market, stock_enums.TimeLevel_1Day, "20250101", "20250509")
	if err != nil {
		t.Fatalf("获取股票交易信息失败: %v", err)
		return
	}
	t.Logf("获取股票交易信息成功: %v", utils.ToJsonString(info))

	closePrices := make([]float64, 0)
	highPrices := make([]float64, 0)
	lowPrices := make([]float64, 0)
	dateTime := make([]string, 0)
	for _, item := range info {
		closePrices = append(closePrices, item.ClosePrice)
		highPrices = append(highPrices, item.HighestPrice)
		lowPrices = append(lowPrices, item.LowestPrice)
		// 解析时间字符串
		parsedTime, err := time.Parse("2006-01-02 15:04:05", item.TradeTime)
		if err != nil {
			t.Log("时间解析失败:", err)
			return
		}

		// 格式化为目标格式
		formattedTime := parsedTime.Format("2006-01-02")
		dateTime = append(dateTime, formattedTime)
	}

	// 计算BBI指标
	bbi := stock_indicators.BBI(closePrices, M1, M2, M3, M4)
	t.Logf("BBI: %v", bbi)
	bbiCarve := &stock_params.BBICarveParams{
		StockCode:       stockCode,
		BBICarveResults: make([]*stock_params.CarveResult, 0),
	}
	for i := M4 - 1; i < len(bbi); i++ {
		bbiCarve.BBICarveResults = append(bbiCarve.BBICarveResults, &stock_params.CarveResult{
			Value:    bbi[i],
			DateTime: dateTime[i],
		})
	}
	t.Logf("bbiCarve: %v", utils.ToJsonString(bbiCarve))

	// 计算KDJ指标
	k, d, j := stock_indicators.KDJ(closePrices, highPrices, lowPrices, N, KdjM1, KdjM2)
	t.Logf("K: %v", k)
	t.Logf("D: %v", d)
	t.Logf("J: %v", j)
	kdjCarve := &stock_params.KDJCarveParams{
		StockCode:     stockCode,
		KCarveResults: make([]*stock_params.CarveResult, 0),
		DCarveResults: make([]*stock_params.CarveResult, 0),
		JCarveResults: make([]*stock_params.CarveResult, 0),
	}
	for i := M4 - 1; i < len(k); i++ {
		kdjCarve.KCarveResults = append(kdjCarve.KCarveResults, &stock_params.CarveResult{
			Value:    k[i],
			DateTime: dateTime[i],
		})
		kdjCarve.DCarveResults = append(kdjCarve.DCarveResults, &stock_params.CarveResult{
			Value:    d[i],
			DateTime: dateTime[i],
		})
		kdjCarve.JCarveResults = append(kdjCarve.JCarveResults, &stock_params.CarveResult{
			Value:    j[i],
			DateTime: dateTime[i],
		})
	}

	t.Logf("kdjCarve: %v", utils.ToJsonString(kdjCarve))
}
