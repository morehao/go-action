package model

import (
	"fmt"
	"time"
)

// UserMessage 用户消息模型
type UserMessage struct {
	ID         uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     uint64     `gorm:"type:bigint unsigned;index:idx_user_id;index:idx_user_id_is_read;not null" json:"user_id"`
	TemplateID uint       `gorm:"type:bigint unsigned;index:idx_template_id;not null" json:"template_id"`
	Title      string     `gorm:"type:varchar(256);not null" json:"title"`
	Content    string     `gorm:"type:text;not null" json:"content"`
	MsgType    string     `gorm:"type:varchar(32);not null" json:"msg_type"`
	JumpUrl    string     `gorm:"type:varchar(512)" json:"jump_url"`
	IsRead     int8       `gorm:"type:tinyint;index:idx_user_id_is_read;not null;default:0" json:"is_read"`
	ReadTime   *time.Time `gorm:"type:datetime" json:"read_time"`
	BizID      string     `gorm:"type:varchar(128);index:idx_biz" json:"biz_id"`
	BizType    string     `gorm:"type:varchar(64);index:idx_biz" json:"biz_type"`
	CreatedAt  time.Time  `gorm:"type:datetime;index:idx_created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName 指定表名（支持分表）
func (m *UserMessage) TableName() string {
	// 默认使用主表，实际项目中可根据 user_id 哈希值返回不同的表名
	return "user_message"
}

// GetShardTableName 获取分表名称
// 示例：按 user_id 哈希到 10 张表
func GetShardTableName(userID uint64, shardCount int) string {
	if shardCount <= 1 {
		return "user_message"
	}
	shardIndex := userID % uint64(shardCount)
	return fmt.Sprintf("user_message_%d", shardIndex)
}

// 消息已读状态常量
const (
	MessageUnread = 0 // 未读
	MessageRead   = 1 // 已读
)

// IsRead 判断消息是否已读
func (m *UserMessage) IsUnread() bool {
	return m.IsRead == MessageUnread
}

// MarkAsRead 标记消息为已读
func (m *UserMessage) MarkAsRead() {
	m.IsRead = MessageRead
	now := time.Now()
	m.ReadTime = &now
}
