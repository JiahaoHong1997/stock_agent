package stock_info

import (
	"agent/app/common/constants/stock_constants"
	"agent/app/common/enums/stock_enums"
	"agent/app/common/params/stock_params"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// https://api.mairuiapi.com/hsstock/real/time/股票代码/证书您的licence
// GetStockTradeInfo 获取（沪深）股票实时交易数据
func GetStockTradeInfo(ctx context.Context, stockCode string) (*stock_params.StockRealTimeInfoResponse, error) {
	// 创建 HTTP 客户端
	client := &http.Client{}

	// 构建请求
	req, err := http.NewRequest("GET", fmt.Sprintf(stock_constants.StockRealTimeInfoURL, stockCode, stock_constants.License), nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return nil, err
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求失败:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return nil, err
	}

	// 输出响应
	fmt.Println("响应状态码:", resp.StatusCode)
	fmt.Println("响应内容:", string(body))
	realTimeInfo := new(stock_params.StockRealTimeInfoResponse)
	err = json.Unmarshal(body, realTimeInfo)
	if err != nil {
		fmt.Printf("Unmarshal info err: %v\n", err)
		return nil, err
	}
	return realTimeInfo, nil
}

// BatchGetStockTradeInfo 批量获取（沪深）股票实时交易数据，每批不超过20
func BatchGetStockTradeInfo(ctx context.Context, stockCodes []string) ([]*stock_params.StockRealTimeInfoResponse, error) {
	// 创建 HTTP 客户端
	client := &http.Client{}

	// 构建请求
	req, err := http.NewRequest("GET", fmt.Sprintf(stock_constants.StockRealTimeBatchInfoURL, stock_constants.License, strings.Join(stockCodes, ",")), nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return nil, err
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求失败:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return nil, err
	}

	// 输出响应
	fmt.Println("响应状态码:", resp.StatusCode)
	fmt.Println("响应内容:", string(body))
	realTimeInfos := make([]*stock_params.StockRealTimeInfoResponse, 0)
	err = json.Unmarshal(body, &realTimeInfos)
	if err != nil {
		fmt.Printf("Unmarshal info err: %v\n", err)
		return nil, err
	}
	return realTimeInfos, nil
}

func GetHistoryStockTradeInfo(ctx context.Context, stockCode string, timeLevel stock_enums.TimeLevel, startTime, endTime string) ([]*stock_params.HistoryStockTradeInfoResponse, error) {
	// 创建 HTTP 客户端
	client := &http.Client{}

	// 构建请求
	req, err := http.NewRequest("GET", fmt.Sprintf(stock_constants.HistoryStockTradeInfoURL, stockCode, timeLevel, "f", stock_constants.License, startTime, endTime), nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return nil, err
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求失败:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return nil, err
	}

	// 输出响应
	fmt.Println("响应状态码:", resp.StatusCode)
	fmt.Println("响应内容:", string(body))
	historyInfo := make([]*stock_params.HistoryStockTradeInfoResponse, 0)
	err = json.Unmarshal(body, &historyInfo)
	if err != nil {
		fmt.Printf("Unmarshal info err: %v\n", err)
		return nil, err
	}
	return historyInfo, nil
}
