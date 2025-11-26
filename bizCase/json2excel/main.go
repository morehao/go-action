package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 获取当前目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取当前目录失败: %v\n", err)
		os.Exit(1)
	}

	// 从命令行参数获取 JSON 文件路径，如果没有则使用默认值
	jsonPath := "es_data_v1.json"
	if len(os.Args) > 1 {
		jsonPath = os.Args[1]
	}
	
	// 如果不是绝对路径，则相对于当前目录
	if !filepath.IsAbs(jsonPath) {
		jsonPath = filepath.Join(dir, jsonPath)
	}

	// 读取 JSON 文件
	jsonData, err := os.ReadFile(jsonPath)
	if err != nil {
		fmt.Printf("读取 JSON 文件失败: %v\n", err)
		os.Exit(1)
	}

	// 根据文件名获取对应的处理器
	processor, err := GetProcessor(jsonPath)
	if err != nil {
		fmt.Printf("获取处理器失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("使用 %s 版本处理器处理文件: %s\n", processor.Version(), jsonPath)

	// 解析 JSON 数据
	rows, err := processor.Parse(jsonData)
	if err != nil {
		fmt.Printf("解析 JSON 失败: %v\n", err)
		os.Exit(1)
	}

	// 生成 Excel 文件
	excelPath := filepath.Join(dir, "output.xlsx")
	if err := GenerateExcel(rows, excelPath); err != nil {
		fmt.Printf("生成 Excel 文件失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("成功生成 Excel 文件: %s\n", excelPath)
	fmt.Printf("共处理 %d 条记录\n", len(rows))
}
