package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplateRenderer_Render(t *testing.T) {
	renderer := NewTemplateRenderer()

	t.Run("成功渲染模板", func(t *testing.T) {
		template := "您的订单 {{orderNo}} 已支付成功，金额 {{amount}} 元"
		params := map[string]string{
			"orderNo": "20231201001",
			"amount":  "99.00",
		}

		result, err := renderer.Render(template, params, true)

		assert.NoError(t, err)
		assert.Equal(t, "您的订单 20231201001 已支付成功，金额 99.00 元", result)
	})

	t.Run("不验证必需参数", func(t *testing.T) {
		template := "您的订单 {{orderNo}} 已支付成功"
		params := map[string]string{} // 空参数

		result, err := renderer.Render(template, params, false)

		assert.NoError(t, err)
		assert.Equal(t, "您的订单 {{orderNo}} 已支付成功", result) // 占位符未替换
	})

	t.Run("模板为空", func(t *testing.T) {
		result, err := renderer.Render("", map[string]string{}, true)

		assert.Error(t, err)
		assert.Empty(t, result)
		assert.Equal(t, "模板内容不能为空", err.Error())
	})

	t.Run("缺少必需参数", func(t *testing.T) {
		template := "订单 {{orderNo}} 金额 {{amount}}"
		params := map[string]string{
			"orderNo": "123",
		}

		result, err := renderer.Render(template, params, true)

		assert.Error(t, err)
		assert.Empty(t, result)
		assert.Contains(t, err.Error(), "缺少必需参数")
		assert.Contains(t, err.Error(), "amount")
	})

	t.Run("多个占位符", func(t *testing.T) {
		template := "用户 {{userName}} 在 {{time}} 购买了 {{productName}}，价格 {{price}}"
		params := map[string]string{
			"userName":    "张三",
			"time":        "2023-12-01",
			"productName": "商品A",
			"price":       "100",
		}

		result, err := renderer.Render(template, params, true)

		assert.NoError(t, err)
		assert.Equal(t, "用户 张三 在 2023-12-01 购买了 商品A，价格 100", result)
	})

	t.Run("重复占位符", func(t *testing.T) {
		template := "{{name}} 您好，欢迎 {{name}}"
		params := map[string]string{
			"name": "李四",
		}

		result, err := renderer.Render(template, params, true)

		assert.NoError(t, err)
		assert.Equal(t, "李四 您好，欢迎 李四", result)
	})
}

func TestTemplateRenderer_ExtractPlaceholders(t *testing.T) {
	renderer := NewTemplateRenderer()

	t.Run("提取单个占位符", func(t *testing.T) {
		template := "您的订单 {{orderNo}} 已支付成功"
		placeholders := renderer.ExtractPlaceholders(template)

		assert.Len(t, placeholders, 1)
		assert.Contains(t, placeholders, "orderNo")
	})

	t.Run("提取多个占位符", func(t *testing.T) {
		template := "订单 {{orderNo}} 金额 {{amount}} 时间 {{time}}"
		placeholders := renderer.ExtractPlaceholders(template)

		assert.Len(t, placeholders, 3)
		assert.Contains(t, placeholders, "orderNo")
		assert.Contains(t, placeholders, "amount")
		assert.Contains(t, placeholders, "time")
	})

	t.Run("提取重复占位符（去重）", func(t *testing.T) {
		template := "{{name}} 您好，欢迎 {{name}}"
		placeholders := renderer.ExtractPlaceholders(template)

		assert.Len(t, placeholders, 1)
		assert.Contains(t, placeholders, "name")
	})

	t.Run("没有占位符", func(t *testing.T) {
		template := "这是一个没有占位符的模板"
		placeholders := renderer.ExtractPlaceholders(template)

		assert.Len(t, placeholders, 0)
	})

	t.Run("空模板", func(t *testing.T) {
		placeholders := renderer.ExtractPlaceholders("")

		assert.Len(t, placeholders, 0)
	})
}

func TestTemplateRenderer_ValidateTemplate(t *testing.T) {
	renderer := NewTemplateRenderer()

	t.Run("有效模板", func(t *testing.T) {
		template := "订单 {{orderNo}} 金额 {{amount}}"
		err := renderer.ValidateTemplate(template)

		assert.NoError(t, err)
	})

	t.Run("模板为空", func(t *testing.T) {
		err := renderer.ValidateTemplate("")

		assert.Error(t, err)
		assert.Equal(t, "模板内容不能为空", err.Error())
	})

	t.Run("占位符未正确闭合-缺少右括号", func(t *testing.T) {
		template := "订单 {{orderNo 金额 {{amount}}"
		err := renderer.ValidateTemplate(template)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "占位符未正确闭合")
	})

	t.Run("占位符未正确闭合-缺少左括号", func(t *testing.T) {
		template := "订单 orderNo}} 金额 {{amount}}"
		err := renderer.ValidateTemplate(template)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "占位符未正确闭合")
	})
}

func TestRender(t *testing.T) {
	t.Run("使用全局渲染器", func(t *testing.T) {
		template := "订单 {{orderNo}} 已完成"
		params := map[string]string{
			"orderNo": "123456",
		}

		result, err := Render(template, params)

		assert.NoError(t, err)
		assert.Equal(t, "订单 123456 已完成", result)
	})
}

func TestRenderWithoutValidation(t *testing.T) {
	t.Run("不验证参数", func(t *testing.T) {
		template := "订单 {{orderNo}} 金额 {{amount}}"
		params := map[string]string{
			"orderNo": "123456",
			// amount 参数缺失
		}

		result, err := RenderWithoutValidation(template, params)

		assert.NoError(t, err)
		assert.Equal(t, "订单 123456 金额 {{amount}}", result) // amount 未替换
	})
}

func TestExtractPlaceholders_Global(t *testing.T) {
	template := "{{name}} 您好，您的 {{product}} 已发货"
	placeholders := ExtractPlaceholders(template)

	assert.Len(t, placeholders, 2)
	assert.Contains(t, placeholders, "name")
	assert.Contains(t, placeholders, "product")
}

func TestValidateTemplate_Global(t *testing.T) {
	t.Run("有效模板", func(t *testing.T) {
		err := ValidateTemplate("订单 {{orderNo}}")
		assert.NoError(t, err)
	})

	t.Run("无效模板", func(t *testing.T) {
		err := ValidateTemplate("订单 {{orderNo")
		assert.Error(t, err)
	})
}

func TestTemplateRenderer_ComplexScenarios(t *testing.T) {
	renderer := NewTemplateRenderer()

	t.Run("中英文混合", func(t *testing.T) {
		template := "Hello {{userName}}, 您的订单 {{orderNo}} 已完成"
		params := map[string]string{
			"userName": "张三",
			"orderNo":  "ORDER123",
		}

		result, err := renderer.Render(template, params, true)

		assert.NoError(t, err)
		assert.Equal(t, "Hello 张三, 您的订单 ORDER123 已完成", result)
	})

	t.Run("数字占位符", func(t *testing.T) {
		template := "订单号 {{order123}} 金额 {{amount99}}"
		params := map[string]string{
			"order123": "ABC",
			"amount99": "100",
		}

		result, err := renderer.Render(template, params, true)

		assert.NoError(t, err)
		assert.Equal(t, "订单号 ABC 金额 100", result)
	})

	t.Run("下划线占位符", func(t *testing.T) {
		template := "用户 {{user_name}} 邮箱 {{user_email}}"
		params := map[string]string{
			"user_name":  "test",
			"user_email": "test@example.com",
		}

		result, err := renderer.Render(template, params, true)

		assert.NoError(t, err)
		assert.Equal(t, "用户 test 邮箱 test@example.com", result)
	})

	t.Run("空参数值", func(t *testing.T) {
		template := "订单 {{orderNo}} 备注 {{remark}}"
		params := map[string]string{
			"orderNo": "123",
			"remark":  "", // 空值
		}

		_, err := renderer.Render(template, params, true)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "remark")
	})

	t.Run("额外参数不影响渲染", func(t *testing.T) {
		template := "订单 {{orderNo}}"
		params := map[string]string{
			"orderNo": "123",
			"extra":   "不需要的参数",
		}

		result, err := renderer.Render(template, params, true)

		assert.NoError(t, err)
		assert.Equal(t, "订单 123", result)
	})
}
