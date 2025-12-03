# 快速入门指南

## 5分钟快速上手

### 第一步：初始化数据库

```bash
# 1. 创建数据库
mysql -u root -p -e "CREATE DATABASE message_center DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"

# 2. 导入初始化脚本（包含建表语句和示例数据）
mysql -u root -p message_center < sql/init.sql
```

### 第二步：安装依赖

> **注意**：本项目使用 monorepo 结构，依赖已在根目录的 `go.mod` 中管理，无需单独安装。

如需更新依赖，请在项目根目录执行：

```bash
cd /path/to/go-action
go mod tidy
```

### 第三步：运行测试

```bash
# 运行所有测试
go test ./... -v

# 查看测试覆盖率
go test ./... -cover
```

### 第四步：在你的项目中使用

创建一个 `main.go` 文件：

```go
package main

import (
    "fmt"
    "messageCenter/dto"
    "messageCenter/service"
    
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func main() {
    // 连接数据库
    dsn := "root:password@tcp(127.0.0.1:3306)/message_center?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }
    
    // 创建消息服务
    msgService := service.NewMessageService(db)
    
    // 发送消息
    resp, err := msgService.SendMessage(&dto.SendMessageRequest{
        UserID:       123456,
        TemplateCode: "ORDER_PAID",
        Title:        "订单支付成功",
        Params: map[string]string{
            "orderNo": "20231201001",
            "amount":  "99.00",
            "orderId": "1001",
        },
    })
    
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("消息发送成功，ID: %d\n", resp.MessageID)
}
```

## 核心功能示例

### 1. 发送消息

```go
msgService.SendMessage(&dto.SendMessageRequest{
    UserID:       123456,
    TemplateCode: "ORDER_PAID",
    Title:        "订单支付成功",
    Params: map[string]string{
        "orderNo": "20231201001",
        "amount":  "99.00",
    },
})
```

### 2. 查询消息列表

```go
// 查询未读消息
isRead := int8(0)
msgService.GetUserMessages(&dto.GetUserMessagesRequest{
    UserID:   123456,
    IsRead:   &isRead,
    Page:     1,
    PageSize: 20,
})
```

### 3. 标记已读

```go
// 单条标记
msgService.MarkAsRead(&dto.MarkAsReadRequest{
    UserID:    123456,
    MessageID: 1,
})

// 批量标记
msgService.BatchMarkAsRead(123456, []uint{1, 2, 3})

// 全部标记
msgService.MarkAllAsRead(123456, "order")
```

### 4. 获取未读数量

```go
msgService.GetUnreadCount(&dto.GetUnreadCountRequest{
    UserID: 123456,
})
```

## 预置模板

初始化脚本中包含了4个示例模板：

1. **ORDER_PAID** - 订单支付成功通知
   - 参数：`orderNo`, `amount`, `orderId`
   
2. **ORDER_SHIPPED** - 订单发货通知
   - 参数：`orderNo`, `expressNo`, `orderId`
   
3. **SYSTEM_NOTICE** - 系统通知
   - 参数：`content`, `url`
   
4. **PROMOTION_COUPON** - 优惠券到账通知
   - 参数：`couponName`, `expireTime`

## 自定义模板

### 创建新模板

```go
templateService := service.NewTemplateService(db)

template := &model.MessageTemplate{
    TemplateCode:    "CUSTOM_TEMPLATE",
    TemplateName:    "自定义模板",
    TemplateContent: "您好 {{userName}}，{{content}}",
    MsgType:         "custom",
    JumpUrl:         "/custom?id={{id}}",
    Priority:        5,
    Status:          model.TemplateStatusEnabled,
}

err := templateService.CreateTemplate(template)
```

### 模板占位符规则

- 使用 `{{variableName}}` 格式
- 变量名只能包含字母、数字和下划线
- 推荐使用驼峰命名法

## 常见问题

### Q: 如何支持分表？

A: 使用 `model.GetShardTableName()` 函数：

```go
tableName := model.GetShardTableName(userID, 10) // 分10张表
db.Table(tableName).Create(&message)
```

### Q: 如何自定义数据库连接？

A: 在创建 gorm.DB 时传入自定义配置：

```go
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
    // 自定义配置
})
```

### Q: 如何集成到 HTTP API？

A: 参考 `README.md` 中的 "HTTP API 集成" 章节。

## 下一步

- 查看 [README.md](README.md) 了解完整文档
- 查看 [example_usage.go](example_usage.go) 了解更多使用示例
- 运行测试了解各个功能的详细用法

## 技术支持

如有问题，请查看：
1. README.md 完整文档
2. 单元测试用例
3. example_usage.go 示例代码

