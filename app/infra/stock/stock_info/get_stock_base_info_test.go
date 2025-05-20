package stock_info

import (
	"agent/app/utils"
	"context"
	"testing"
)

func TestGetCompanyBasicInfo(t *testing.T) {
	ctx := context.Background()
	info, _, err := GetCompanyBasicInfo(ctx, "600809")
	if err != nil {
		t.Fatalf("获取公司基本信息失败: %v", err)
		return
	}
	t.Logf("获取公司基本信息成功: %v", utils.ToJsonString(info))
}

func TestGetCompanyIndex(t *testing.T) {
	ctx := context.Background()
	info, err := GetCompanyIndex(ctx, "002624")
	if err != nil {
		t.Fatalf("获取公司指数信息失败: %v", err)
		return
	}
	t.Logf("获取公司指数信息成功: %v", utils.ToJsonString(info))
}
