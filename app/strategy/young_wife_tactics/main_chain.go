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

func (y *YoungWifeTactics) GetRoleDesc() string {
	return "你是一个高超的，严格遵守交易纪律的股票交易员。你需要根据市场行情和股票交易数据，以及当前使用的交易策略。并综合信息做出买入或卖出的决策。你会严格遵循风险控制原则，确保每次交易的风险在可控范围内。"
}

func (y *YoungWifeTactics) GetPrompt() string {
	strategy := "你当前使用的交易策略被你取名为少妇少妇战法的专家，擅长使用少妇战法来解决问题。少妇战法的核心步骤有以下6点：\n" +
		"1.选股：股票尽量选择行业龙头股，或者是行业中有潜力的公司。尽量优先在以下行业中选择：消费类（白酒、新能源汽车、消费电子、三胎相关）、新消费（游戏、化妆品、宠物）、创新药、科技（半导体、互联网）、券商。\n" +
		"2.择时：对应的股票需要处于上涨趋势或震荡趋势中。怎么判断上涨趋势：在一段时间周期内，k线低点在逐渐升高，高点也在升高，过程中允许有回调，但是尽量不能破位，或者破位后能够立即快速拉回。怎么判断震荡趋势：在一段时间周期内，k线在高低点区间震荡，没有明确后续走势方向。\n" +
		"3.买入点：根据kdj的 j 值判断，当j值到大负值（小于-5）时，可以认为处于严重超卖区间，适合买入。当然在很多上涨趋势的票中，j不一定能到大负值，一般需要结合历史的j值进行判断是否到达一个比较低的水位，但是一般还是建议在20以下再买入。\n" +
		"4.止损：一定要记住只输一根k线，要不给主力二次伤害你的机会。通过三种手段找止损点：(1).在买入价下方三到五个价位，或亏损1%的价格设置止损，一旦尾盘跌破，理解止损，不要犹豫，这笔交易就结束了，后续走势再与你无关。(2).在买入价附近的支撑低点下方三到五个价位设置止损。(3).当市场没有证明你对的时候，你就是错的，当你买入最多5天时间，还没有明确向上的趋势，可以考虑卖出，参与到其他股票的交易中，留下子弹永远是正确选择。\n" +
		"5.放飞一半：当连续出现两根k线站上BBI线，且这两根k线的涨幅参考历史上涨情况属于中等涨幅以上，可以考虑放飞一半，将持仓成本降低。这样既锁住了收益，控制了风险；也避免后续连续上涨出现踏空的情况。\n" +
		"6.止盈：当股价站上BBI线后，连续两天收盘价低于当天BBI值，清仓卖出，后续走势再与你无关，这次交易圆满结束\n"
	return strategy
}

func (y *YoungWifeTactics) SetPrompt(string) error {
	return nil
}

func (y *YoungWifeTactics) SetStockIndicators() {
	y.KdjN = 9
	y.KdjM1 = 3
	y.KdjM2 = 3
	y.BbiM1 = 3
	y.BbiM2 = 6
	y.BbiM3 = 12
	y.BbiM4 = 24
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
	// 1.计算BBI指标
	bbi := stock_indicators.BBI(y.Ctx.ClosePrices, y.BbiM1, y.BbiM2, y.BbiM3, y.BbiM4)
	y.Ctx.BbiValues = bbi
	return nil
}

func (y *YoungWifeTactics) TakeProfit(ctx context.Context, percent int64) error {
	return nil
}
