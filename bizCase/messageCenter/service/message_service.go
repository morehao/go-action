package service

import (
	"errors"
	"time"

	"github.com/morehao/go-action/bizCase/messageCenter/dto"
	"github.com/morehao/go-action/bizCase/messageCenter/model"
	"github.com/morehao/go-action/bizCase/messageCenter/utils"
	"gorm.io/gorm"
)

// MessageService 消息服务
type MessageService struct {
	db              *gorm.DB
	templateService *TemplateService
	renderer        *utils.TemplateRenderer
}

// NewMessageService 创建消息服务
func NewMessageService(db *gorm.DB) *MessageService {
	return &MessageService{
		db:              db,
		templateService: NewTemplateService(db),
		renderer:        utils.NewTemplateRenderer(),
	}
}

// SendMessage 根据模板发送消息
func (s *MessageService) SendMessage(req *dto.SendMessageRequest) (*dto.SendMessageResponse, error) {
	if req == nil {
		return nil, errors.New("请求参数不能为空")
	}

	if req.UserID == 0 {
		return nil, errors.New("用户ID不能为空")
	}

	if req.TemplateCode == "" {
		return nil, errors.New("模板编码不能为空")
	}

	if req.Title == "" {
		return nil, errors.New("消息标题不能为空")
	}

	// 1. 获取模板配置
	template, err := s.templateService.GetEnabledTemplateByCode(req.TemplateCode)
	if err != nil {
		return nil, err
	}

	// 2. 渲染消息内容
	content, err := s.renderer.Render(template.TemplateContent, req.Params, false)
	if err != nil {
		return nil, err
	}

	// 3. 渲染跳转链接
	var jumpUrl string
	if template.JumpUrl != "" {
		jumpUrl, err = s.renderer.Render(template.JumpUrl, req.Params, false)
		if err != nil {
			return nil, err
		}
	}

	// 4. 创建用户消息
	userMessage := &model.UserMessage{
		UserID:     req.UserID,
		TemplateID: template.ID,
		Title:      req.Title,
		Content:    content,
		MsgType:    template.MsgType,
		JumpUrl:    jumpUrl,
		IsRead:     model.MessageUnread,
		BizID:      req.BizID,
		BizType:    req.BizType,
	}

	// 5. 保存到数据库
	if err := s.db.Create(userMessage).Error; err != nil {
		return nil, err
	}

	return &dto.SendMessageResponse{
		MessageID: userMessage.ID,
		Success:   true,
		Message:   "消息发送成功",
	}, nil
}

// MarkAsRead 标记消息为已读
func (s *MessageService) MarkAsRead(req *dto.MarkAsReadRequest) error {
	if req == nil {
		return errors.New("请求参数不能为空")
	}

	if req.UserID == 0 {
		return errors.New("用户ID不能为空")
	}

	if req.MessageID == 0 {
		return errors.New("消息ID不能为空")
	}

	// 查询消息
	var message model.UserMessage
	err := s.db.Where("id = ? AND user_id = ?", req.MessageID, req.UserID).First(&message).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("消息不存在或无权限")
		}
		return err
	}

	// 如果已经是已读状态，直接返回
	if message.IsRead == model.MessageRead {
		return nil
	}

	// 标记为已读
	message.MarkAsRead()
	return s.db.Save(&message).Error
}

// BatchMarkAsRead 批量标记已读
func (s *MessageService) BatchMarkAsRead(userID uint64, messageIDs []uint) error {
	if userID == 0 {
		return errors.New("用户ID不能为空")
	}

	if len(messageIDs) == 0 {
		return errors.New("消息ID列表不能为空")
	}

	now := time.Now()
	return s.db.Model(&model.UserMessage{}).
		Where("user_id = ? AND id IN ? AND is_read = ?", userID, messageIDs, model.MessageUnread).
		Updates(map[string]interface{}{
			"is_read":   model.MessageRead,
			"read_time": now,
		}).Error
}

// MarkAllAsRead 标记所有消息为已读
func (s *MessageService) MarkAllAsRead(userID uint64, msgType string) error {
	if userID == 0 {
		return errors.New("用户ID不能为空")
	}

	query := s.db.Model(&model.UserMessage{}).
		Where("user_id = ? AND is_read = ?", userID, model.MessageUnread)

	// 按消息类型筛选
	if msgType != "" {
		query = query.Where("msg_type = ?", msgType)
	}

	now := time.Now()
	return query.Updates(map[string]interface{}{
		"is_read":   model.MessageRead,
		"read_time": now,
	}).Error
}

// GetUserMessages 查询用户消息列表
func (s *MessageService) GetUserMessages(req *dto.GetUserMessagesRequest) (*dto.GetUserMessagesResponse, error) {
	if req == nil {
		return nil, errors.New("请求参数不能为空")
	}

	if req.UserID == 0 {
		return nil, errors.New("用户ID不能为空")
	}

	// 设置默认分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}
	if req.PageSize > 100 {
		req.PageSize = 100 // 限制最大每页数量
	}

	// 构建查询
	query := s.db.Model(&model.UserMessage{}).Where("user_id = ?", req.UserID)

	// 按已读状态筛选
	if req.IsRead != nil {
		query = query.Where("is_read = ?", *req.IsRead)
	}

	// 按消息类型筛选
	if req.MsgType != "" {
		query = query.Where("msg_type = ?", req.MsgType)
	}

	// 按业务类型筛选
	if req.BizType != "" {
		query = query.Where("biz_type = ?", req.BizType)
	}

	// 获取总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// 分页查询
	var messages []model.UserMessage
	offset := (req.Page - 1) * req.PageSize
	err := query.Order("created_at DESC").
		Offset(offset).
		Limit(req.PageSize).
		Find(&messages).Error
	if err != nil {
		return nil, err
	}

	// 转换为VO
	list := make([]dto.UserMessageVO, 0, len(messages))
	for _, msg := range messages {
		list = append(list, dto.UserMessageVO{
			ID:         msg.ID,
			UserID:     msg.UserID,
			TemplateID: msg.TemplateID,
			Title:      msg.Title,
			Content:    msg.Content,
			MsgType:    msg.MsgType,
			JumpUrl:    msg.JumpUrl,
			IsRead:     msg.IsRead,
			ReadTime:   msg.ReadTime,
			BizID:      msg.BizID,
			BizType:    msg.BizType,
			CreatedAt:  msg.CreatedAt,
		})
	}

	return &dto.GetUserMessagesResponse{
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
		List:     list,
	}, nil
}

// GetUnreadCount 获取未读消息数量
func (s *MessageService) GetUnreadCount(req *dto.GetUnreadCountRequest) (*dto.GetUnreadCountResponse, error) {
	if req == nil {
		return nil, errors.New("请求参数不能为空")
	}

	if req.UserID == 0 {
		return nil, errors.New("用户ID不能为空")
	}

	query := s.db.Model(&model.UserMessage{}).
		Where("user_id = ? AND is_read = ?", req.UserID, model.MessageUnread)

	// 按消息类型筛选
	if req.MsgType != "" {
		query = query.Where("msg_type = ?", req.MsgType)
	}

	var count int64
	err := query.Count(&count).Error
	if err != nil {
		return nil, err
	}

	return &dto.GetUnreadCountResponse{
		UnreadCount: count,
	}, nil
}

// GetUnreadCountByType 按消息类型获取未读数量
func (s *MessageService) GetUnreadCountByType(userID uint64) (map[string]int64, error) {
	if userID == 0 {
		return nil, errors.New("用户ID不能为空")
	}

	type Result struct {
		MsgType string
		Count   int64
	}

	var results []Result
	err := s.db.Model(&model.UserMessage{}).
		Select("msg_type, COUNT(*) as count").
		Where("user_id = ? AND is_read = ?", userID, model.MessageUnread).
		Group("msg_type").
		Find(&results).Error
	if err != nil {
		return nil, err
	}

	countMap := make(map[string]int64)
	for _, result := range results {
		countMap[result.MsgType] = result.Count
	}

	return countMap, nil
}

// DeleteMessage 删除消息（软删除或真实删除）
func (s *MessageService) DeleteMessage(userID uint64, messageID uint) error {
	if userID == 0 {
		return errors.New("用户ID不能为空")
	}

	if messageID == 0 {
		return errors.New("消息ID不能为空")
	}

	// 验证消息归属
	result := s.db.Where("id = ? AND user_id = ?", messageID, userID).Delete(&model.UserMessage{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("消息不存在或无权限")
	}

	return nil
}

// GetMessageDetail 获取消息详情
func (s *MessageService) GetMessageDetail(userID uint64, messageID uint) (*dto.UserMessageVO, error) {
	if userID == 0 {
		return nil, errors.New("用户ID不能为空")
	}

	if messageID == 0 {
		return nil, errors.New("消息ID不能为空")
	}

	var message model.UserMessage
	err := s.db.Where("id = ? AND user_id = ?", messageID, userID).First(&message).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("消息不存在或无权限")
		}
		return nil, err
	}

	return &dto.UserMessageVO{
		ID:         message.ID,
		UserID:     message.UserID,
		TemplateID: message.TemplateID,
		Title:      message.Title,
		Content:    message.Content,
		MsgType:    message.MsgType,
		JumpUrl:    message.JumpUrl,
		IsRead:     message.IsRead,
		ReadTime:   message.ReadTime,
		BizID:      message.BizID,
		BizType:    message.BizType,
		CreatedAt:  message.CreatedAt,
	}, nil
}
