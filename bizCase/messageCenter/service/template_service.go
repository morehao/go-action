package service

import (
	"errors"

	"github.com/morehao/go-action/bizCase/messageCenter/model"
	"github.com/morehao/go-action/bizCase/messageCenter/utils"
	"gorm.io/gorm"
)

// TemplateService 模板服务
type TemplateService struct {
	db *gorm.DB
}

// NewTemplateService 创建模板服务
func NewTemplateService(db *gorm.DB) *TemplateService {
	return &TemplateService{
		db: db,
	}
}

// GetTemplateByCode 根据模板编码获取模板
func (s *TemplateService) GetTemplateByCode(templateCode string) (*model.MessageTemplate, error) {
	if templateCode == "" {
		return nil, errors.New("模板编码不能为空")
	}

	var template model.MessageTemplate
	err := s.db.Where("template_code = ?", templateCode).First(&template).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("模板不存在")
		}
		return nil, err
	}

	return &template, nil
}

// GetTemplateByID 根据模板ID获取模板
func (s *TemplateService) GetTemplateByID(templateID uint) (*model.MessageTemplate, error) {
	if templateID == 0 {
		return nil, errors.New("模板ID不能为空")
	}

	var template model.MessageTemplate
	err := s.db.First(&template, templateID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("模板不存在")
		}
		return nil, err
	}

	return &template, nil
}

// ValidateTemplate 验证模板配置
func (s *TemplateService) ValidateTemplate(template *model.MessageTemplate) error {
	if template == nil {
		return errors.New("模板对象不能为空")
	}

	if template.TemplateCode == "" {
		return errors.New("模板编码不能为空")
	}

	if template.TemplateName == "" {
		return errors.New("模板名称不能为空")
	}

	if template.TemplateContent == "" {
		return errors.New("模板内容不能为空")
	}

	if template.MsgType == "" {
		return errors.New("消息类型不能为空")
	}

	// 验证模板格式
	if err := utils.ValidateTemplate(template.TemplateContent); err != nil {
		return err
	}

	// 如果有跳转链接模板，也需要验证
	if template.JumpUrl != "" {
		if err := utils.ValidateTemplate(template.JumpUrl); err != nil {
			return err
		}
	}

	return nil
}

// GetEnabledTemplateByCode 获取已启用的模板
func (s *TemplateService) GetEnabledTemplateByCode(templateCode string) (*model.MessageTemplate, error) {
	template, err := s.GetTemplateByCode(templateCode)
	if err != nil {
		return nil, err
	}

	if !template.IsEnabled() {
		return nil, errors.New("模板未启用")
	}

	return template, nil
}

// CreateTemplate 创建模板（用于管理功能，虽然不在核心需求中，但作为完整性补充）
func (s *TemplateService) CreateTemplate(template *model.MessageTemplate) error {
	if err := s.ValidateTemplate(template); err != nil {
		return err
	}

	// 检查模板编码是否已存在
	var count int64
	err := s.db.Model(&model.MessageTemplate{}).Where("template_code = ?", template.TemplateCode).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("模板编码已存在")
	}

	return s.db.Create(template).Error
}

// UpdateTemplate 更新模板
func (s *TemplateService) UpdateTemplate(template *model.MessageTemplate) error {
	if template.ID == 0 {
		return errors.New("模板ID不能为空")
	}

	if err := s.ValidateTemplate(template); err != nil {
		return err
	}

	return s.db.Save(template).Error
}

// ListTemplates 获取模板列表
func (s *TemplateService) ListTemplates(msgType string, status *int8, page, pageSize int) ([]model.MessageTemplate, int64, error) {
	query := s.db.Model(&model.MessageTemplate{})

	// 按消息类型筛选
	if msgType != "" {
		query = query.Where("msg_type = ?", msgType)
	}

	// 按状态筛选
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	// 获取总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	var templates []model.MessageTemplate
	if page > 0 && pageSize > 0 {
		offset := (page - 1) * pageSize
		query = query.Offset(offset).Limit(pageSize)
	}

	err := query.Order("priority DESC, id DESC").Find(&templates).Error
	if err != nil {
		return nil, 0, err
	}

	return templates, total, nil
}
