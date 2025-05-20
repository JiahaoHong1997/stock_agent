package young_wife_tactics

import (
	"agent/app/common/enums/stock_enums"
	"agent/app/common/params/stock_params"
	"agent/app/infra/stock/stock_indicators"
	"agent/app/infra/stock/stock_info"
	"context"
	"fmt"
	"time"
)

func (y *YoungWifeTactics) GetName() string {
	return "young_wife_tactics"
}

func (y *YoungWifeTactics) GetDescription() string {
	return "少妇战法"
}

func (y *YoungWifeTactics) GetPrompt() string {
	return "你是一个少妇战法的专家，擅长使用少妇战法来解决问题。请根据以下信息提供建议和指导："
}

func (y *YoungWifeTactics) SetPrompt(string) error {
	return nil
}

func (y *YoungWifeTactics) SetStockIndicators() {
	y.KdjN = 9
	y.KdjM1 = 3
	y.KdjM2 = 3
}

func (y *YoungWifeTactics) SelectMarketTiming(ctx context.Context) error {
	return nil
}

func (y *YoungWifeTactics) SelectStock(ctx context.Context) (string, error) {
	return "", nil
}

func (y *YoungWifeTactics) GetStockTradingDataForPastXDays(ctx context.Context, stockCode string, x int) error {
	// 1.查询公司基本信息
	_, market, err := stock_info.GetCompanyBasicInfo(ctx, stockCode)
	if err != nil {
		return err
	}
	if market == "" {
		return fmt.Errorf("market is empty")
	}

	endDate := time.Now().Format("20060102")
	startDate := time.Now().AddDate(0, 0, -x).Format("20060102")
	// 2.获取历史股票交易信息
	info, err := stock_info.GetHistoryStockTradeInfo(ctx, stockCode, market, stock_enums.TimeLevel_1Day, startDate, endDate)
	if err != nil {
		return err
	}
	STCtx := new(stock_params.StockTradeContext)
	STCtx.StockTradeInfo = info
	STCtx.StockCode = stockCode
	STCtx.MarketCode = market
	STCtx.ClosePrices = make([]float64, 0)
	STCtx.HighPrices = make([]float64, 0)
	STCtx.LowPrices = make([]float64, 0)
	for _, item := range info {
		STCtx.ClosePrices = append(STCtx.ClosePrices, item.ClosePrice)
		STCtx.HighPrices = append(STCtx.HighPrices, item.HighestPrice)
		STCtx.LowPrices = append(STCtx.LowPrices, item.LowestPrice)
	}
	y.Ctx = STCtx
	return nil
}

func (y *YoungWifeTactics) B1Buy(ctx context.Context) error {
	// 1.计算kdj指标
	_, _, y.Ctx.JValues = stock_indicators.KDJ(y.Ctx.ClosePrices, y.Ctx.HighPrices, y.Ctx.LowPrices, y.KdjN, y.KdjM1, y.KdjM2)
	return nil
}

func (y *YoungWifeTactics) RiskControl(ctx context.Context) error {
	return nil
}

func (y *YoungWifeTactics) TakeProfit(ctx context.Context, percent int64) error {
	return nil
}
