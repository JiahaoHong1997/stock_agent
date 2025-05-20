package stock_info

import (
	"agent/app/common/constants/stock_constants"
	"agent/app/common/params/stock_params"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetCompanyBasicInfo(ctx context.Context, stockCode string) (*stock_params.CompanyBasicInfo, string, error) {
	// 创建 HTTP 客户端
	client := &http.Client{}

	// 构建请求
	req, err := http.NewRequest("GET", fmt.Sprintf(stock_constants.GetCompanyBasicInfoURL, stockCode, stock_constants.License), nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return nil, "", err
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求失败:", err)
		return nil, "", err
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return nil, "", err
	}

	// 输出响应
	fmt.Println("响应状态码:", resp.StatusCode)
	fmt.Println("响应内容:", string(body))

	companyBasicInfo := new(stock_params.CompanyBasicInfo)
	err = json.Unmarshal(body, &companyBasicInfo)
	if err != nil {
		fmt.Printf("Unmarshal info err: %v\n", err)
		return nil, "", err
	}

	var market string
	if companyBasicInfo.Market == "深圳证券交易所" {
		market = "SZ"
	} else if companyBasicInfo.Market == "上海证券交易所" {
		market = "SH"
	}
	return companyBasicInfo, market, nil
}

// GetCompanyIndex 获取公司所属指数
func GetCompanyIndex(ctx context.Context, stockCode string) ([]*stock_params.CompanyIndex, error) {
	// 创建 HTTP 客户端
	client := &http.Client{}

	// 构建请求
	req, err := http.NewRequest("GET", fmt.Sprintf(stock_constants.GetCompanyIndexURL, stockCode, stock_constants.License), nil)
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
	companyIndex := make([]*stock_params.CompanyIndex, 0)
	err = json.Unmarshal(body, &companyIndex)
	if err != nil {
		fmt.Printf("Unmarshal info err: %v\n", err)
		return nil, err
	}
	return companyIndex, nil
}
