package utils

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"text/template"
)

// TemplateRenderer 模板渲染器
type TemplateRenderer struct{}

// NewTemplateRenderer 创建模板渲染器
func NewTemplateRenderer() *TemplateRenderer {
	return &TemplateRenderer{}
}

// Render 渲染模板
// template: 模板内容，如 "您的订单 {{orderNo}} 已支付成功"
// params: 参数映射，如 {"orderNo": "123456"}
// validateRequired: 是否验证必需参数
func (r *TemplateRenderer) Render(templateStr string, params map[string]string, validateRequired bool) (string, error) {
	if templateStr == "" {
		return "", errors.New("模板内容不能为空")
	}

	// 查找所有占位符
	placeholders := r.findPlaceholders(templateStr)

	// 验证必需参数
	if validateRequired && len(placeholders) > 0 {
		if err := r.validateParams(placeholders, params); err != nil {
			return "", err
		}
	}

	// 将 {{key}} 格式转换为 text/template 的 {{.key}} 格式
	convertedTemplate := r.convertTemplateSyntax(templateStr)

	// 创建模板并解析
	tmpl, err := template.New("render").Parse(convertedTemplate)
	if err != nil {
		return "", fmt.Errorf("模板解析失败: %w", err)
	}

	// 执行模板渲染
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, params)
	if err != nil {
		return "", fmt.Errorf("模板渲染失败: %w", err)
	}

	return buf.String(), nil
}

// convertTemplateSyntax 将 {{key}} 格式转换为 {{.key}} 格式
func (r *TemplateRenderer) convertTemplateSyntax(templateStr string) string {
	result := templateStr
	placeholders := r.findPlaceholders(templateStr)

	// 从后往前替换，避免索引问题
	for i := len(placeholders) - 1; i >= 0; i-- {
		key := placeholders[i]
		oldPattern := fmt.Sprintf("{{%s}}", key)
		newPattern := fmt.Sprintf("{{.%s}}", key)
		result = strings.ReplaceAll(result, oldPattern, newPattern)
	}

	return result
}

// findPlaceholders 查找模板中的所有占位符
// 使用简单的字符串遍历，查找 {{xxx}} 格式的占位符
func (r *TemplateRenderer) findPlaceholders(template string) []string {
	var placeholders []string
	seen := make(map[string]bool)

	i := 0
	for i < len(template) {
		// 查找 {{ 的位置
		start := strings.Index(template[i:], "{{")
		if start == -1 {
			break
		}
		start += i

		// 查找对应的 }} 的位置
		end := strings.Index(template[start+2:], "}}")
		if end == -1 {
			break
		}
		end += start + 2

		// 提取占位符名称
		placeholderName := template[start+2 : end]
		placeholderName = strings.TrimSpace(placeholderName)

		// 验证占位符名称是否有效（只包含字母、数字、下划线）
		if placeholderName != "" && isValidPlaceholderName(placeholderName) && !seen[placeholderName] {
			placeholders = append(placeholders, placeholderName)
			seen[placeholderName] = true
		}

		i = end + 2
	}

	return placeholders
}

// isValidPlaceholderName 验证占位符名称是否有效
// 只允许字母、数字、下划线
func isValidPlaceholderName(name string) bool {
	if name == "" {
		return false
	}
	for _, ch := range name {
		if !((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || ch == '_') {
			return false
		}
	}
	return true
}

// validateParams 验证参数是否完整
func (r *TemplateRenderer) validateParams(placeholders []string, params map[string]string) error {
	var missingParams []string

	for _, placeholder := range placeholders {
		if value, exists := params[placeholder]; !exists || value == "" {
			missingParams = append(missingParams, placeholder)
		}
	}

	if len(missingParams) > 0 {
		return fmt.Errorf("缺少必需参数: %s", strings.Join(missingParams, ", "))
	}

	return nil
}

// ExtractPlaceholders 提取模板中的占位符列表
func (r *TemplateRenderer) ExtractPlaceholders(template string) []string {
	return r.findPlaceholders(template)
}

// ValidateTemplate 验证模板格式是否正确
func (r *TemplateRenderer) ValidateTemplate(templateStr string) error {
	if templateStr == "" {
		return errors.New("模板内容不能为空")
	}

	// 检查占位符格式是否正确（简单检查是否成对出现）
	openCount := strings.Count(templateStr, "{{")
	closeCount := strings.Count(templateStr, "}}")

	if openCount != closeCount {
		return errors.New("模板格式错误：占位符未正确闭合")
	}

	// 使用 text/template 验证模板语法
	convertedTemplate := r.convertTemplateSyntax(templateStr)
	_, err := template.New("validate").Parse(convertedTemplate)
	if err != nil {
		return fmt.Errorf("模板格式错误: %w", err)
	}

	return nil
}

// 全局默认渲染器实例
var defaultRenderer = NewTemplateRenderer()

// Render 使用默认渲染器渲染模板
func Render(template string, params map[string]string) (string, error) {
	return defaultRenderer.Render(template, params, true)
}

// RenderWithoutValidation 使用默认渲染器渲染模板（不验证必需参数）
func RenderWithoutValidation(template string, params map[string]string) (string, error) {
	return defaultRenderer.Render(template, params, false)
}

// ExtractPlaceholders 使用默认渲染器提取占位符
func ExtractPlaceholders(template string) []string {
	return defaultRenderer.ExtractPlaceholders(template)
}

// ValidateTemplate 使用默认渲染器验证模板
func ValidateTemplate(template string) error {
	return defaultRenderer.ValidateTemplate(template)
}
