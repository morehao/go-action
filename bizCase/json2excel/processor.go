package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/xuri/excelize/v2"
)

// DataRow 表示一行数据
type DataRow struct {
	CreatedAt string
	Question  string
	Answer    string
}

// Processor 定义版本处理器的接口
type Processor interface {
	// Parse 解析 JSON 数据并返回数据行
	Parse(jsonData []byte) ([]DataRow, error)
	// Version 返回版本号
	Version() string
}

// GetProcessor 根据文件名获取对应的处理器
func GetProcessor(filename string) (Processor, error) {
	baseName := filepath.Base(filename)

	// 从文件名中提取版本号，例如 es_data_v1.json -> v1
	if strings.Contains(baseName, "_v1") {
		return &ProcessorV1{}, nil
	}

	// 默认使用 v1 版本
	return &ProcessorV1{}, nil
}

// GenerateExcel 生成 Excel 文件
func GenerateExcel(rows []DataRow, outputPath string) error {
	// 创建 Excel 文件
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("关闭 Excel 文件失败: %v\n", err)
		}
	}()

	// 获取默认工作表
	sheet := f.GetSheetName(0)

	// 设置表头
	headers := []string{"创建时间", "问题", "答案"}
	for i, header := range headers {
		cell := fmt.Sprintf("%c1", 'A'+i)
		if err := f.SetCellValue(sheet, cell, header); err != nil {
			return fmt.Errorf("设置表头失败: %w", err)
		}
	}

	// 写入数据
	for i, row := range rows {
		rowNum := i + 2 // 从第 2 行开始（第 1 行是表头）

		// 创建时间
		if err := f.SetCellValue(sheet, fmt.Sprintf("A%d", rowNum), row.CreatedAt); err != nil {
			return fmt.Errorf("写入创建时间失败: %w", err)
		}

		// 问题
		if err := f.SetCellValue(sheet, fmt.Sprintf("B%d", rowNum), row.Question); err != nil {
			return fmt.Errorf("写入问题失败: %w", err)
		}

		// 答案
		if err := f.SetCellValue(sheet, fmt.Sprintf("C%d", rowNum), row.Answer); err != nil {
			return fmt.Errorf("写入答案失败: %w", err)
		}
	}

	// 设置列宽（可选，让内容更易读）
	if err := f.SetColWidth(sheet, "A", "A", 25); err != nil {
		return fmt.Errorf("设置列宽失败: %w", err)
	}
	if err := f.SetColWidth(sheet, "B", "B", 30); err != nil {
		return fmt.Errorf("设置列宽失败: %w", err)
	}
	if err := f.SetColWidth(sheet, "C", "C", 80); err != nil {
		return fmt.Errorf("设置列宽失败: %w", err)
	}

	// 保存 Excel 文件
	if err := f.SaveAs(outputPath); err != nil {
		return fmt.Errorf("保存 Excel 文件失败: %w", err)
	}

	return nil
}
