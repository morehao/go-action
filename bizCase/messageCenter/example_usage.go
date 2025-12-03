package main

import (
	"fmt"
	"log"

	"github.com/morehao/go-action/bizCase/messageCenter/dto"
	"github.com/morehao/go-action/bizCase/messageCenter/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 这是一个使用示例文件，演示如何使用消息中心系统
// 注意：需要先配置好数据库连接和导入初始化脚本

func main() {
	// 1. 连接数据库
	dsn := "root:password@tcp(127.0.0.1:3306)/message_center?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 2. 创建消息服务
	messageService := service.NewMessageService(db)

	// ============ 示例1: 发送消息 ============
	fmt.Println("========== 示例1: 发送消息 ==========")
	sendReq := &dto.SendMessageRequest{
		UserID:       123456,
		TemplateCode: "ORDER_PAID",
		Title:        "订单支付成功",
		Params: map[string]string{
			"orderNo": "20231201001",
			"amount":  "99.00",
			"orderId": "1001",
		},
		BizID:   "1001",
		BizType: "order",
	}

	sendResp, err := messageService.SendMessage(sendReq)
	if err != nil {
		log.Printf("发送消息失败: %v\n", err)
	} else {
		fmt.Printf("✓ 消息发送成功，消息ID: %d\n", sendResp.MessageID)
	}

	// ============ 示例2: 查询用户消息列表（未读消息） ============
	fmt.Println("\n========== 示例2: 查询未读消息 ==========")
	isRead := int8(0) // 0-未读
	listReq := &dto.GetUserMessagesRequest{
		UserID:   123456,
		IsRead:   &isRead,
		Page:     1,
		PageSize: 10,
	}

	listResp, err := messageService.GetUserMessages(listReq)
	if err != nil {
		log.Printf("查询消息失败: %v\n", err)
	} else {
		fmt.Printf("✓ 未读消息总数: %d\n", listResp.Total)
		for i, msg := range listResp.List {
			fmt.Printf("  [%d] %s - %s (类型: %s)\n", i+1, msg.Title, msg.Content, msg.MsgType)
		}
	}

	// ============ 示例3: 获取未读消息数量 ============
	fmt.Println("\n========== 示例3: 获取未读数量 ==========")
	countReq := &dto.GetUnreadCountRequest{
		UserID: 123456,
	}

	countResp, err := messageService.GetUnreadCount(countReq)
	if err != nil {
		log.Printf("查询未读数量失败: %v\n", err)
	} else {
		fmt.Printf("✓ 未读消息数量: %d\n", countResp.UnreadCount)
	}

	// 按类型获取未读数量
	countByType, err := messageService.GetUnreadCountByType(123456)
	if err != nil {
		log.Printf("查询分类未读数量失败: %v\n", err)
	} else {
		fmt.Println("✓ 按类型统计未读数量:")
		for msgType, count := range countByType {
			fmt.Printf("  - %s: %d 条\n", msgType, count)
		}
	}

	// ============ 示例4: 标记消息为已读 ============
	fmt.Println("\n========== 示例4: 标记消息为已读 ==========")
	if sendResp != nil && sendResp.MessageID > 0 {
		readReq := &dto.MarkAsReadRequest{
			UserID:    123456,
			MessageID: sendResp.MessageID,
		}

		err = messageService.MarkAsRead(readReq)
		if err != nil {
			log.Printf("标记已读失败: %v\n", err)
		} else {
			fmt.Printf("✓ 消息 %d 已标记为已读\n", sendResp.MessageID)
		}
	}

	// ============ 示例5: 批量标记已读 ============
	fmt.Println("\n========== 示例5: 批量标记已读 ==========")
	messageIDs := []uint{1, 2, 3}
	err = messageService.BatchMarkAsRead(123456, messageIDs)
	if err != nil {
		log.Printf("批量标记失败: %v\n", err)
	} else {
		fmt.Printf("✓ 成功批量标记 %d 条消息为已读\n", len(messageIDs))
	}

	// ============ 示例6: 标记全部已读 ============
	fmt.Println("\n========== 示例6: 标记全部已读 ==========")
	err = messageService.MarkAllAsRead(123456, "order") // 只标记订单类消息
	if err != nil {
		log.Printf("标记全部已读失败: %v\n", err)
	} else {
		fmt.Println("✓ 所有订单类消息已标记为已读")
	}

	// ============ 示例7: 获取消息详情 ============
	fmt.Println("\n========== 示例7: 获取消息详情 ==========")
	if sendResp != nil && sendResp.MessageID > 0 {
		detail, err := messageService.GetMessageDetail(123456, sendResp.MessageID)
		if err != nil {
			log.Printf("获取消息详情失败: %v\n", err)
		} else {
			fmt.Printf("✓ 消息详情:\n")
			fmt.Printf("  标题: %s\n", detail.Title)
			fmt.Printf("  内容: %s\n", detail.Content)
			fmt.Printf("  类型: %s\n", detail.MsgType)
			fmt.Printf("  跳转链接: %s\n", detail.JumpUrl)
			fmt.Printf("  是否已读: %v\n", detail.IsRead == 1)
		}
	}

	// ============ 示例8: 查询所有消息（已读+未读） ============
	fmt.Println("\n========== 示例8: 查询所有消息 ==========")
	allReq := &dto.GetUserMessagesRequest{
		UserID:   123456,
		IsRead:   nil, // nil 表示查询全部
		Page:     1,
		PageSize: 20,
	}

	allResp, err := messageService.GetUserMessages(allReq)
	if err != nil {
		log.Printf("查询所有消息失败: %v\n", err)
	} else {
		fmt.Printf("✓ 消息总数: %d (第 %d 页，每页 %d 条)\n",
			allResp.Total, allResp.Page, allResp.PageSize)
		for i, msg := range allResp.List {
			status := "未读"
			if msg.IsRead == 1 {
				status = "已读"
			}
			fmt.Printf("  [%d] [%s] %s\n", i+1, status, msg.Title)
		}
	}

	// ============ 示例9: 按消息类型查询 ============
	fmt.Println("\n========== 示例9: 按消息类型查询 ==========")
	typeReq := &dto.GetUserMessagesRequest{
		UserID:   123456,
		MsgType:  "order",
		Page:     1,
		PageSize: 10,
	}

	typeResp, err := messageService.GetUserMessages(typeReq)
	if err != nil {
		log.Printf("查询订单消息失败: %v\n", err)
	} else {
		fmt.Printf("✓ 订单类消息总数: %d\n", typeResp.Total)
		for i, msg := range typeResp.List {
			fmt.Printf("  [%d] %s\n", i+1, msg.Title)
		}
	}

	// ============ 示例10: 删除消息 ============
	fmt.Println("\n========== 示例10: 删除消息 ==========")
	if sendResp != nil && sendResp.MessageID > 0 {
		err = messageService.DeleteMessage(123456, sendResp.MessageID)
		if err != nil {
			log.Printf("删除消息失败: %v\n", err)
		} else {
			fmt.Printf("✓ 消息 %d 已删除\n", sendResp.MessageID)
		}
	}

	fmt.Println("\n========== 示例演示完成 ==========")
}

