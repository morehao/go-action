package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {

	config := Config{
		FilePath:  "./test.md",
		OldPrefix: "http://12.13.14.15:4530",
		NewPrefix: "https://test.morehao.com:4530",
	}

	if err := replaceURLsInFile(config); err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}
}

// Config 配置结构
type Config struct {
	FilePath  string
	OldPrefix string
	NewPrefix string
}

func replaceURLsInFile(config Config) error {
	// 读取文件
	content, err := os.ReadFile(config.FilePath)
	if err != nil {
		return fmt.Errorf("读取文件失败: %w", err)
	}

	// 转义特殊字符并构建正则表达式
	pattern := regexp.MustCompile(regexp.QuoteMeta(config.OldPrefix) + "/test-bucket/servicea/" + `[^\s\)]*`)

	// 查找所有匹配
	matches := pattern.FindAllString(string(content), -1)
	if len(matches) == 0 {
		fmt.Println("✓ 未找到需要替换的 URL")
		return nil
	}

	// 显示找到的 URL
	fmt.Printf("找到 %d 个需要替换的 URL:\n", len(matches))
	uniqueMatches := uniqueStrings(matches)
	for i, match := range uniqueMatches {
		fmt.Printf("  %d. %s\n", i+1, match)
	}

	// 执行替换：直接替换前缀
	replacedContent := []byte(pattern.ReplaceAllStringFunc(string(content), func(match string) string {
		return config.NewPrefix + match[len(config.OldPrefix):]
	}))

	// 写入文件
	if err := os.WriteFile(config.FilePath, replacedContent, 0644); err != nil {
		return fmt.Errorf("写入文件失败: %w", err)
	}

	fmt.Printf("\n✓ 成功替换 %d 个 URL\n", len(matches))
	fmt.Printf("✓ 文件已更新: %s\n", config.FilePath)
	return nil
}

// 去重辅助函数
func uniqueStrings(slice []string) []string {
	seen := make(map[string]bool)
	result := []string{}
	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}
