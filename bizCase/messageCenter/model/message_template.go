package model

import (
	"time"
)

// MessageTemplate 消息模板模型
type MessageTemplate struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	TemplateCode    string    `gorm:"type:varchar(64);uniqueIndex:uk_template_code;not null" json:"template_code"`
	TemplateName    string    `gorm:"type:varchar(128);not null" json:"template_name"`
	TemplateContent string    `gorm:"type:text;not null" json:"template_content"`
	MsgType         string    `gorm:"type:varchar(32);index:idx_msg_type;not null" json:"msg_type"`
	JumpUrl         string    `gorm:"type:varchar(512)" json:"jump_url"`
	Priority        int       `gorm:"type:int;not null;default:0" json:"priority"`
	Status          int8      `gorm:"type:tinyint;index:idx_status;not null;default:1" json:"status"`
	CreatedAt       time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName 指定表名
func (MessageTemplate) TableName() string {
	return "message_template"
}

// 模板状态常量
const (
	TemplateStatusDisabled = 0 // 禁用
	TemplateStatusEnabled  = 1 // 启用
)

// IsEnabled 判断模板是否启用
func (t *MessageTemplate) IsEnabled() bool {
	return t.Status == TemplateStatusEnabled
}
