package main

import (
	"os"
	"testing"
)

// TestIntegration 集成测试：完整的 JSON 转 Excel 流程
func TestIntegration(t *testing.T) {
	// 读取真实的测试文件
	jsonPath := "es_data_v1.json"
	jsonData, err := os.ReadFile(jsonPath)
	if err != nil {
		t.Skipf("跳过测试：无法读取测试文件 %s: %v", jsonPath, err)
	}

	// 获取处理器
	processor, err := GetProcessor(jsonPath)
	if err != nil {
		t.Fatalf("获取处理器失败: %v", err)
	}

	// 解析数据
	rows, err := processor.Parse(jsonData)
	if err != nil {
		t.Fatalf("解析数据失败: %v", err)
	}

	// 验证解析结果
	if len(rows) == 0 {
		t.Error("期望至少有一条记录，但实际为空")
	}

	// 生成 Excel 到当前目录
	outputPath := "integration_test.xlsx"
	err = GenerateExcel(rows, outputPath)
	if err != nil {
		t.Fatalf("生成 Excel 失败: %v", err)
	}

	// 验证文件存在
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Fatalf("Excel 文件未生成: %s", outputPath)
	}

	t.Logf("成功处理 %d 条记录，生成文件: %s", len(rows), outputPath)
}
