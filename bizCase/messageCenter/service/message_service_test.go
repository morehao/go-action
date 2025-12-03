package service

import (
	"testing"

	"github.com/morehao/go-action/bizCase/messageCenter/dto"
	"github.com/stretchr/testify/assert"
)

func TestMessageService_SendMessage(t *testing.T) {
	service := NewMessageService(nil)

	t.Run("请求参数为空", func(t *testing.T) {
		resp, err := service.SendMessage(nil)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, "请求参数不能为空", err.Error())
	})

	t.Run("用户ID为空", func(t *testing.T) {
		req := &dto.SendMessageRequest{
			TemplateCode: "ORDER_PAID",
			Title:        "测试",
		}

		resp, err := service.SendMessage(req)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, "用户ID不能为空", err.Error())
	})
}

func TestMessageService_BatchMarkAsRead(t *testing.T) {
	service := NewMessageService(nil)

	t.Run("消息ID列表为空", func(t *testing.T) {
		err := service.BatchMarkAsRead(123, []uint{})

		assert.Error(t, err)
		assert.Equal(t, "消息ID列表不能为空", err.Error())
	})
}

func TestMessageService_GetUserMessages(t *testing.T) {
	service := NewMessageService(nil)

	t.Run("用户ID为空", func(t *testing.T) {
		req := &dto.GetUserMessagesRequest{
			Page:     1,
			PageSize: 10,
		}

		resp, err := service.GetUserMessages(req)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, "用户ID不能为空", err.Error())
	})
}
