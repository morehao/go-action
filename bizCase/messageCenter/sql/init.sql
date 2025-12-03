-- 消息中心数据库初始化脚本
-- 创建数据库（如需要）
-- CREATE DATABASE IF NOT EXISTS message_center DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
-- USE message_center;

-- =====================================================
-- 消息模板表
-- =====================================================
CREATE TABLE IF NOT EXISTS `message_template` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `template_code` VARCHAR(64) NOT NULL COMMENT '模板编码，唯一标识',
    `template_name` VARCHAR(128) NOT NULL COMMENT '模板名称',
    `template_content` TEXT NOT NULL COMMENT '模板内容，支持 {{variable}} 占位符',
    `msg_type` VARCHAR(32) NOT NULL COMMENT '消息类型（用于区分跳转页面），如：system, order, promotion',
    `jump_url` VARCHAR(512) DEFAULT NULL COMMENT '跳转链接模板，支持占位符',
    `priority` INT NOT NULL DEFAULT 0 COMMENT '优先级，数字越大优先级越高',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '启用状态：0-禁用，1-启用',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_template_code` (`template_code`),
    KEY `idx_msg_type` (`msg_type`),
    KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='消息模板配置表';

-- =====================================================
-- 用户消息表（主表，实际使用时可按 user_id 哈希分表）
-- =====================================================
CREATE TABLE IF NOT EXISTS `user_message` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    `template_id` BIGINT UNSIGNED NOT NULL COMMENT '模板ID',
    `title` VARCHAR(256) NOT NULL COMMENT '消息标题',
    `content` TEXT NOT NULL COMMENT '渲染后的消息内容',
    `msg_type` VARCHAR(32) NOT NULL COMMENT '消息类型',
    `jump_url` VARCHAR(512) DEFAULT NULL COMMENT '实际跳转链接',
    `is_read` TINYINT NOT NULL DEFAULT 0 COMMENT '已读状态：0-未读，1-已读',
    `read_time` DATETIME DEFAULT NULL COMMENT '阅读时间',
    `biz_id` VARCHAR(128) DEFAULT NULL COMMENT '业务关联ID',
    `biz_type` VARCHAR(64) DEFAULT NULL COMMENT '业务类型',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`),
    KEY `idx_user_id_is_read` (`user_id`, `is_read`),
    KEY `idx_template_id` (`template_id`),
    KEY `idx_biz` (`biz_type`, `biz_id`),
    KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户消息表';

-- =====================================================
-- 分表示例（实际项目中可创建多个分表）
-- =====================================================
-- CREATE TABLE IF NOT EXISTS `user_message_0` LIKE `user_message`;
-- CREATE TABLE IF NOT EXISTS `user_message_1` LIKE `user_message`;
-- ...

-- =====================================================
-- 插入测试模板数据
-- =====================================================
INSERT INTO `message_template` (`template_code`, `template_name`, `template_content`, `msg_type`, `jump_url`, `priority`, `status`) VALUES
('ORDER_PAID', '订单支付成功通知', '您的订单 {{orderNo}} 已支付成功，支付金额 {{amount}} 元', 'order', '/order/detail?orderId={{orderId}}', 10, 1),
('ORDER_SHIPPED', '订单发货通知', '您的订单 {{orderNo}} 已发货，快递单号：{{expressNo}}', 'order', '/order/logistics?orderId={{orderId}}', 10, 1),
('SYSTEM_NOTICE', '系统通知', '{{content}}', 'system', '{{url}}', 5, 1),
('PROMOTION_COUPON', '优惠券到账通知', '恭喜您获得 {{couponName}}，有效期至 {{expireTime}}', 'promotion', '/coupon/list', 8, 1);
