package strategy

import (
	"context"
)

type IStrategy interface {
	// GetName 获取策略名称
	GetName() string
	// GetDescription 获取策略描述
	GetDescription() string
	// GetPrompt 获取策略提示词
	GetPrompt() string
	// SetPrompt 设置策略提示词
	SetPrompt(string) error
	// SetStockIndicators 设置股票指标
	SetStockIndicators()

	//  ------ 主要策略（择时、选股、B1买入、设置止损、止盈放飞一半（BBI以上两根中长阳线）、止盈完全放飞（跌破BBI连续两根k线））----------
	// SelectMarketTiming 择时
	SelectMarketTiming(ctx context.Context) error
	// SelectStock 选股
	SelectStock(ctx context.Context) (string, error)
	// GetStockTradingDataForPastXDays 获取过去X天的股票交易数据
	GetStockTradingDataForPastXDays(ctx context.Context, stockCode string, x int) error
	// B1Buy B1买入点
	B1Buy(ctx context.Context) error
	// RiskControl 风控（止损）
	RiskControl(ctx context.Context) error
	// TakeProfit 止盈点
	TakeProfit(ctx context.Context, percent int64) error
}
