package service

import (
	"testing"

	"github.com/morehao/go-action/bizCase/messageCenter/model"
	"github.com/stretchr/testify/assert"
)

func TestTemplateService_ValidateTemplate(t *testing.T) {
	service := NewTemplateService(nil)

	t.Run("有效模板", func(t *testing.T) {
		template := &model.MessageTemplate{
			TemplateCode:    "TEST_CODE",
			TemplateName:    "测试模板",
			TemplateContent: "这是 {{name}} 的测试消息",
			MsgType:         "test",
			JumpUrl:         "/test?id={{id}}",
			Priority:        1,
			Status:          model.TemplateStatusEnabled,
		}

		err := service.ValidateTemplate(template)
		assert.NoError(t, err)
	})

	t.Run("模板对象为空", func(t *testing.T) {
		err := service.ValidateTemplate(nil)
		assert.Error(t, err)
		assert.Equal(t, "模板对象不能为空", err.Error())
	})

	t.Run("模板编码为空", func(t *testing.T) {
		template := &model.MessageTemplate{
			TemplateName:    "测试模板",
			TemplateContent: "测试内容",
			MsgType:         "test",
		}

		err := service.ValidateTemplate(template)
		assert.Error(t, err)
		assert.Equal(t, "模板编码不能为空", err.Error())
	})

	t.Run("模板内容格式错误", func(t *testing.T) {
		template := &model.MessageTemplate{
			TemplateCode:    "TEST_CODE",
			TemplateName:    "测试模板",
			TemplateContent: "占位符未闭合 {{name",
			MsgType:         "test",
		}

		err := service.ValidateTemplate(template)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "占位符未正确闭合")
	})
}
