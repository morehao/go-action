# 消息中心系统

## 项目简介

消息中心系统是一个基于模板的消息管理解决方案，支持动态模板管理、消息生成、状态管理和多类型消息跳转功能。

### 核心特性

- ✅ 消息模板管理：支持占位符配置和动态内容渲染
- ✅ 消息发送：根据模板自动生成用户消息
- ✅ 状态管理：支持已读/未读状态追踪
- ✅ 多类型消息：不同消息类型支持不同的跳转页面
- ✅ 分表支持：预留分表逻辑，支持大规模用户场景
- ✅ 完整测试：提供单元测试保障代码质量

## 项目结构

```
messageCenter/
├── README.md                       # 项目文档
├── sql/
│   └── init.sql                   # 数据库初始化脚本
├── model/
│   ├── message_template.go        # 消息模板数据模型
│   └── user_message.go            # 用户消息数据模型
├── dto/
│   └── dto.go                     # 数据传输对象（DTO）
├── service/
│   ├── template_service.go        # 模板管理服务
│   ├── template_service_test.go   # 模板服务测试
│   ├── message_service.go         # 消息服务
│   └── message_service_test.go    # 消息服务测试
└── utils/
    └── template_render.go         # 模板渲染工具
```

> **注意**：本项目使用 monorepo 结构，依赖管理使用根目录的 `go.mod`，无需单独的模块配置文件。

## 数据库设计

### message_template（消息模板表）

| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT | 主键ID |
| template_code | VARCHAR(64) | 模板编码（唯一） |
| template_name | VARCHAR(128) | 模板名称 |
| template_content | TEXT | 模板内容，支持 `{{variable}}` 占位符 |
| msg_type | VARCHAR(32) | 消息类型（如：system, order, promotion） |
| jump_url | VARCHAR(512) | 跳转链接模板 |
| priority | INT | 优先级（数字越大优先级越高） |
| status | TINYINT | 启用状态（0-禁用，1-启用） |
| created_at | DATETIME | 创建时间 |
| updated_at | DATETIME | 更新时间 |

### user_message（用户消息表）

| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT | 主键ID |
| user_id | BIGINT | 用户ID（分表字段） |
| template_id | BIGINT | 模板ID |
| title | VARCHAR(256) | 消息标题 |
| content | TEXT | 渲染后的消息内容 |
| msg_type | VARCHAR(32) | 消息类型 |
| jump_url | VARCHAR(512) | 实际跳转链接 |
| is_read | TINYINT | 已读状态（0-未读，1-已读） |
| read_time | DATETIME | 阅读时间 |
| biz_id | VARCHAR(128) | 业务关联ID |
| biz_type | VARCHAR(64) | 业务类型 |
| created_at | DATETIME | 创建时间 |
| updated_at | DATETIME | 更新时间 |

## 快速开始

### 1. 安装依赖

> **注意**：本项目使用 monorepo 结构，依赖已在根目录的 `go.mod` 中管理，无需单独安装。

如需更新依赖，请在项目根目录执行：

```bash
cd /path/to/go-action
go mod tidy
```

### 2. 初始化数据库

```bash
# 连接到 MySQL 数据库
mysql -u root -p

# 创建数据库
CREATE DATABASE message_center DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

# 导入初始化脚本
mysql -u root -p message_center < sql/init.sql
```

### 3. 配置数据库连接

在你的应用中配置数据库连接：

```go
import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

// 数据库连接配置
dsn := "root:password@tcp(127.0.0.1:3306)/message_center?charset=utf8mb4&parseTime=True&loc=Local"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
if err != nil {
    panic("failed to connect database")
}
```

### 4. 使用示例

#### 发送消息

```go
package main

import (
    "fmt"
    "messageCenter/dto"
    "messageCenter/service"
)

func main() {
    // 创建消息服务
    messageService := service.NewMessageService(db)
    
    // 发送消息请求
    req := &dto.SendMessageRequest{
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
    
    // 发送消息
    resp, err := messageService.SendMessage(req)
    if err != nil {
        fmt.Printf("发送消息失败: %v\n", err)
        return
    }
    
    fmt.Printf("消息发送成功，消息ID: %d\n", resp.MessageID)
}
```

#### 查询用户消息

```go
// 查询用户消息列表
isRead := int8(0) // 0-未读，1-已读
req := &dto.GetUserMessagesRequest{
    UserID:   123456,
    IsRead:   &isRead,
    Page:     1,
    PageSize: 20,
}

resp, err := messageService.GetUserMessages(req)
if err != nil {
    fmt.Printf("查询失败: %v\n", err)
    return
}

fmt.Printf("总数: %d, 当前页: %d\n", resp.Total, resp.Page)
for _, msg := range resp.List {
    fmt.Printf("消息: %s - %s\n", msg.Title, msg.Content)
}
```

#### 标记消息为已读

```go
// 标记单条消息为已读
req := &dto.MarkAsReadRequest{
    UserID:    123456,
    MessageID: 1,
}

err := messageService.MarkAsRead(req)
if err != nil {
    fmt.Printf("标记失败: %v\n", err)
    return
}

fmt.Println("标记成功")
```

#### 获取未读消息数量

```go
// 获取未读消息数量
req := &dto.GetUnreadCountRequest{
    UserID:  123456,
    MsgType: "order", // 可选，不传则查询所有类型
}

resp, err := messageService.GetUnreadCount(req)
if err != nil {
    fmt.Printf("查询失败: %v\n", err)
    return
}

fmt.Printf("未读消息数量: %d\n", resp.UnreadCount)
```

#### 批量标记已读

```go
// 批量标记消息为已读
userID := uint64(123456)
messageIDs := []uint{1, 2, 3, 4, 5}

err := messageService.BatchMarkAsRead(userID, messageIDs)
if err != nil {
    fmt.Printf("批量标记失败: %v\n", err)
    return
}

fmt.Println("批量标记成功")
```

## API 接口说明

### MessageService（消息服务）

#### SendMessage - 发送消息

根据模板发送消息给用户。

**参数：**
- `req *dto.SendMessageRequest`
  - `UserID`: 用户ID（必填）
  - `TemplateCode`: 模板编码（必填）
  - `Title`: 消息标题（必填）
  - `Params`: 模板参数（可选）
  - `BizID`: 业务关联ID（可选）
  - `BizType`: 业务类型（可选）

**返回：**
- `*dto.SendMessageResponse`: 包含消息ID和发送结果
- `error`: 错误信息

#### GetUserMessages - 获取用户消息列表

查询用户的消息列表，支持分页和筛选。

**参数：**
- `req *dto.GetUserMessagesRequest`
  - `UserID`: 用户ID（必填）
  - `IsRead`: 已读状态（可选，0-未读，1-已读，nil-全部）
  - `MsgType`: 消息类型（可选）
  - `BizType`: 业务类型（可选）
  - `Page`: 页码（默认1）
  - `PageSize`: 每页数量（默认20，最大100）

**返回：**
- `*dto.GetUserMessagesResponse`: 消息列表和分页信息
- `error`: 错误信息

#### MarkAsRead - 标记消息为已读

标记单条消息为已读状态。

**参数：**
- `req *dto.MarkAsReadRequest`
  - `UserID`: 用户ID（必填）
  - `MessageID`: 消息ID（必填）

**返回：**
- `error`: 错误信息

#### BatchMarkAsRead - 批量标记已读

批量标记多条消息为已读状态。

**参数：**
- `userID uint64`: 用户ID
- `messageIDs []uint`: 消息ID列表

**返回：**
- `error`: 错误信息

#### MarkAllAsRead - 标记全部已读

标记用户的所有未读消息为已读状态。

**参数：**
- `userID uint64`: 用户ID
- `msgType string`: 消息类型（可选，为空则标记全部）

**返回：**
- `error`: 错误信息

#### GetUnreadCount - 获取未读消息数量

获取用户的未读消息数量。

**参数：**
- `req *dto.GetUnreadCountRequest`
  - `UserID`: 用户ID（必填）
  - `MsgType`: 消息类型（可选）

**返回：**
- `*dto.GetUnreadCountResponse`: 未读消息数量
- `error`: 错误信息

#### GetUnreadCountByType - 按类型获取未读数量

获取用户各类型消息的未读数量。

**参数：**
- `userID uint64`: 用户ID

**返回：**
- `map[string]int64`: 消息类型与未读数量的映射
- `error`: 错误信息

#### GetMessageDetail - 获取消息详情

获取单条消息的详细信息。

**参数：**
- `userID uint64`: 用户ID
- `messageID uint`: 消息ID

**返回：**
- `*dto.UserMessageVO`: 消息详情
- `error`: 错误信息

#### DeleteMessage - 删除消息

删除指定的消息。

**参数：**
- `userID uint64`: 用户ID
- `messageID uint`: 消息ID

**返回：**
- `error`: 错误信息

### TemplateService（模板服务）

#### GetTemplateByCode - 根据模板编码获取模板

**参数：**
- `templateCode string`: 模板编码

**返回：**
- `*model.MessageTemplate`: 模板信息
- `error`: 错误信息

#### GetEnabledTemplateByCode - 获取已启用的模板

**参数：**
- `templateCode string`: 模板编码

**返回：**
- `*model.MessageTemplate`: 模板信息
- `error`: 错误信息

#### ValidateTemplate - 验证模板配置

**参数：**
- `template *model.MessageTemplate`: 模板对象

**返回：**
- `error`: 错误信息

#### CreateTemplate - 创建模板

**参数：**
- `template *model.MessageTemplate`: 模板对象

**返回：**
- `error`: 错误信息

#### UpdateTemplate - 更新模板

**参数：**
- `template *model.MessageTemplate`: 模板对象

**返回：**
- `error`: 错误信息

#### ListTemplates - 获取模板列表

**参数：**
- `msgType string`: 消息类型（可选）
- `status *int8`: 状态（可选）
- `page int`: 页码
- `pageSize int`: 每页数量

**返回：**
- `[]model.MessageTemplate`: 模板列表
- `int64`: 总数
- `error`: 错误信息

## 模板渲染

### 占位符语法

模板内容使用 `{{variable}}` 格式的占位符：

```
您的订单 {{orderNo}} 已支付成功，支付金额 {{amount}} 元
```

### 渲染示例

```go
import "messageCenter/utils"

// 使用默认渲染器
template := "您的订单 {{orderNo}} 已支付成功"
params := map[string]string{
    "orderNo": "20231201001",
}

result, err := utils.Render(template, params)
// 结果: "您的订单 20231201001 已支付成功"
```

### 高级用法

```go
// 创建自定义渲染器
renderer := utils.NewTemplateRenderer()

// 提取占位符
placeholders := renderer.ExtractPlaceholders("订单 {{orderNo}} 金额 {{amount}}")
// 结果: ["orderNo", "amount"]

// 验证模板格式
err := renderer.ValidateTemplate("订单 {{orderNo")
// 返回错误: "模板格式错误：占位符未正确闭合"
```

## 分表策略

系统预留了分表逻辑，用于支持大规模用户场景。

### 获取分表名称

```go
import "messageCenter/model"

// 按 user_id 哈希到 10 张表
tableName := model.GetShardTableName(userID, 10)
// 示例: user_message_3
```

### 使用分表

```go
// 在创建消息前指定表名
message := &model.UserMessage{
    UserID: 123456,
    // ... 其他字段
}

// 使用分表
shardCount := 10
tableName := model.GetShardTableName(message.UserID, shardCount)
db.Table(tableName).Create(message)
```

## 运行测试

```bash
# 运行所有测试
go test ./...

# 运行指定测试
go test ./service -v

# 查看测试覆盖率
go test ./... -cover
```

## 技术栈

- **Go**: 1.23.3
- **GORM**: ORM 框架
- **MySQL**: v8.0+（数据库）
- **testify**: 测试断言库

> **注意**：依赖版本由项目根目录的 `go.mod` 统一管理

## 注意事项

1. **模板编码唯一性**：每个模板的 `template_code` 必须唯一，用于快速查找模板。

2. **占位符命名规范**：占位符名称建议使用驼峰命名法，只能包含字母、数字和下划线。

3. **参数验证**：发送消息时建议提供所有必需参数，避免渲染失败。

4. **分表实施**：生产环境建议根据实际用户量规划分表数量（如 10、100、1000 张表）。

5. **索引优化**：根据实际查询场景优化索引，特别是 `user_id`、`is_read`、`created_at` 等常用字段。

6. **消息清理**：建议定期清理过期消息，避免数据库膨胀。

## 扩展建议

### HTTP API 集成

如需提供 HTTP API 接口，可参考以下示例：

```go
package main

import (
    "messageCenter/dto"
    "messageCenter/service"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    
    // 初始化服务
    messageService := service.NewMessageService(db)
    
    // 发送消息
    r.POST("/message/send", func(c *gin.Context) {
        var req dto.SendMessageRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        
        resp, err := messageService.SendMessage(&req)
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        
        c.JSON(200, resp)
    })
    
    // 获取消息列表
    r.POST("/message/list", func(c *gin.Context) {
        var req dto.GetUserMessagesRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        
        resp, err := messageService.GetUserMessages(&req)
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        
        c.JSON(200, resp)
    })
    
    // 标记已读
    r.POST("/message/read", func(c *gin.Context) {
        var req dto.MarkAsReadRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        
        err := messageService.MarkAsRead(&req)
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        
        c.JSON(200, gin.H{"message": "success"})
    })
    
    // 获取未读数量
    r.POST("/message/unread-count", func(c *gin.Context) {
        var req dto.GetUnreadCountRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        
        resp, err := messageService.GetUnreadCount(&req)
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        
        c.JSON(200, resp)
    })
    
    r.Run(":8080")
}
```

### 消息推送

可集成 WebSocket 或 Server-Sent Events（SSE）实现实时消息推送。

### 消息归档

可实现消息归档功能，将历史消息迁移到归档表。

## 许可证

MIT License

## 联系方式

如有问题或建议，欢迎提交 Issue 或 Pull Request。
