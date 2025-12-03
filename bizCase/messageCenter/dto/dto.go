package dto

import "time"

// SendMessageRequest 发送消息请求
type SendMessageRequest struct {
	UserID       uint64            `json:"user_id" binding:"required"`       // 用户ID
	TemplateCode string            `json:"template_code" binding:"required"` // 模板编码
	Title        string            `json:"title" binding:"required"`         // 消息标题
	Params       map[string]string `json:"params"`                           // 模板参数
	BizID        string            `json:"biz_id"`                           // 业务关联ID
	BizType      string            `json:"biz_type"`                         // 业务类型
}

// SendMessageResponse 发送消息响应
type SendMessageResponse struct {
	MessageID uint   `json:"message_id"` // 消息ID
	Success   bool   `json:"success"`    // 是否成功
	Message   string `json:"message"`    // 提示信息
}

// GetUserMessagesRequest 获取用户消息列表请求
type GetUserMessagesRequest struct {
	UserID   uint64  `json:"user_id" binding:"required"` // 用户ID
	IsRead   *int8   `json:"is_read"`                    // 已读状态：0-未读，1-已读，nil-全部
	MsgType  string  `json:"msg_type"`                   // 消息类型
	Page     int     `json:"page"`                       // 页码，从1开始
	PageSize int     `json:"page_size"`                  // 每页数量
	BizType  string  `json:"biz_type"`                   // 业务类型
}

// GetUserMessagesResponse 获取用户消息列表响应
type GetUserMessagesResponse struct {
	Total    int64                `json:"total"`    // 总数
	Page     int                  `json:"page"`     // 当前页
	PageSize int                  `json:"page_size"` // 每页数量
	List     []UserMessageVO      `json:"list"`     // 消息列表
}

// UserMessageVO 用户消息视图对象
type UserMessageVO struct {
	ID         uint       `json:"id"`          // 消息ID
	UserID     uint64     `json:"user_id"`     // 用户ID
	TemplateID uint       `json:"template_id"` // 模板ID
	Title      string     `json:"title"`       // 消息标题
	Content    string     `json:"content"`     // 消息内容
	MsgType    string     `json:"msg_type"`    // 消息类型
	JumpUrl    string     `json:"jump_url"`    // 跳转链接
	IsRead     int8       `json:"is_read"`     // 已读状态
	ReadTime   *time.Time `json:"read_time"`   // 阅读时间
	BizID      string     `json:"biz_id"`      // 业务关联ID
	BizType    string     `json:"biz_type"`    // 业务类型
	CreatedAt  time.Time  `json:"created_at"`  // 创建时间
}

// MarkAsReadRequest 标记已读请求
type MarkAsReadRequest struct {
	UserID     uint64 `json:"user_id" binding:"required"`     // 用户ID
	MessageID  uint   `json:"message_id" binding:"required"`  // 消息ID
}

// GetUnreadCountRequest 获取未读数量请求
type GetUnreadCountRequest struct {
	UserID  uint64 `json:"user_id" binding:"required"` // 用户ID
	MsgType string `json:"msg_type"`                   // 消息类型，为空则查全部
}

// GetUnreadCountResponse 获取未读数量响应
type GetUnreadCountResponse struct {
	UnreadCount int64 `json:"unread_count"` // 未读数量
}
