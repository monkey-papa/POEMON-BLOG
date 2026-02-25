-- My Blog Go 数据库结构
-- 数据库: myblog
-- 编码: utf8mb4

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- 创建数据库
-- ----------------------------
CREATE DATABASE IF NOT EXISTS `myblog` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE `myblog`;

-- ----------------------------
-- 用户认证表 (Django auth_user 兼容)
-- ----------------------------
DROP TABLE IF EXISTS `auth_user`;
CREATE TABLE `auth_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `password` varchar(128) NOT NULL,
  `last_login` datetime(6) DEFAULT NULL,
  `is_superuser` tinyint(1) NOT NULL DEFAULT 0,
  `username` varchar(150) NOT NULL,
  `first_name` varchar(150) NOT NULL DEFAULT '',
  `last_name` varchar(150) NOT NULL DEFAULT '',
  `email` varchar(254) NOT NULL DEFAULT '',
  `is_staff` tinyint(1) NOT NULL DEFAULT 0,
  `is_active` tinyint(1) NOT NULL DEFAULT 1,
  `date_joined` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户认证表';

-- ----------------------------
-- Token表 (Django REST framework 兼容)
-- ----------------------------
DROP TABLE IF EXISTS `authtoken_token`;
CREATE TABLE `authtoken_token` (
  `key` varchar(40) NOT NULL,
  `created` datetime(6) NOT NULL,
  `user_id` int(11) NOT NULL,
  PRIMARY KEY (`key`),
  UNIQUE KEY `user_id` (`user_id`),
  CONSTRAINT `authtoken_token_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `auth_user` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Token表';

-- ----------------------------
-- 用户信息表
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `user_id` int(11) NOT NULL,
  `username` varchar(32) NOT NULL,
  `phone_number` varchar(16) DEFAULT NULL,
  `email` varchar(32) DEFAULT NULL,
  `user_status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否启用',
  `gender` smallint(6) DEFAULT 0 COMMENT '性别: 0-保密, 1-男, 2-女',
  `open_id` varchar(128) DEFAULT NULL,
  `avatar` varchar(256) DEFAULT NULL COMMENT '头像',
  `admire` varchar(32) DEFAULT NULL COMMENT '赞赏',
  `introduction` varchar(4096) DEFAULT NULL COMMENT '简介',
  `user_type` smallint(6) NOT NULL DEFAULT 2 COMMENT '用户类型: 0-Boss, 1-管理员, 2-普通用户, 3-访客',
  `create_time` datetime(6) NOT NULL,
  `update_time` datetime(6) NOT NULL,
  `update_by` varchar(32) DEFAULT NULL,
  `deleted` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除',
  `province` varchar(20) DEFAULT NULL COMMENT '省份',
  `qiniu_domain` varchar(128) DEFAULT NULL COMMENT '七牛云域名',
  `qiniu_bucket_name` varchar(128) DEFAULT NULL COMMENT '七牛云bucket',
  `qiniu_secret_key` varchar(128) DEFAULT NULL COMMENT '七牛云secret_key',
  `qiniu_access_key` varchar(128) DEFAULT NULL COMMENT '七牛云access_key',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `username` (`username`),
  CONSTRAINT `user_auth_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `auth_user` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户信息表';

-- ----------------------------
-- 分类表
-- ----------------------------
DROP TABLE IF EXISTS `sort`;
CREATE TABLE `sort` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sort_name` varchar(32) NOT NULL COMMENT '分类名称',
  `sort_description` varchar(256) NOT NULL COMMENT '分类描述',
  `sort_type` int(11) NOT NULL DEFAULT 1 COMMENT '分类类型: 0-导航栏分类, 1-普通分类',
  `priority` int(11) DEFAULT NULL COMMENT '优先级',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='分类表';

-- ----------------------------
-- 标签表
-- ----------------------------
DROP TABLE IF EXISTS `label`;
CREATE TABLE `label` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sort_id` int(11) NOT NULL COMMENT '分类ID',
  `label_name` varchar(32) NOT NULL COMMENT '标签名称',
  `label_description` varchar(256) NOT NULL COMMENT '标签描述',
  PRIMARY KEY (`id`),
  KEY `idx_sort_id` (`sort_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='标签表';

-- ----------------------------
-- 文章表
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL COMMENT '用户ID',
  `sort_id` int(11) NOT NULL COMMENT '分类ID',
  `label_id` int(11) NOT NULL COMMENT '标签ID',
  `article_cover` varchar(256) DEFAULT NULL COMMENT '封面',
  `article_title` varchar(32) NOT NULL COMMENT '文章标题',
  `article_content` longtext NOT NULL COMMENT '文章内容',
  `summary` longtext COMMENT '文章摘要',
  `view_count` int(11) NOT NULL DEFAULT 0 COMMENT '浏览量',
  `like_count` int(11) NOT NULL DEFAULT 0 COMMENT '点赞数',
  `view_status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否可见',
  `password` varchar(128) DEFAULT NULL COMMENT '密码',
  `recommend_status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否推荐',
  `comment_status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否启用评论',
  `create_time` datetime(6) NOT NULL,
  `update_time` datetime(6) NOT NULL,
  `update_by` varchar(32) DEFAULT NULL COMMENT '最终修改人',
  `deleted` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_sort_id` (`sort_id`),
  KEY `idx_label_id` (`label_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='文章表';

-- ----------------------------
-- 文章点赞表
-- ----------------------------
DROP TABLE IF EXISTS `like`;
CREATE TABLE `like` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id_id` int(11) NOT NULL COMMENT '用户ID',
  `article_id_id` int(11) NOT NULL COMMENT '文章ID',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id_id`),
  KEY `idx_article_id` (`article_id_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='文章点赞表';

-- ----------------------------
-- 评论表
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `source` int(11) NOT NULL COMMENT '评论来源标识',
  `type` varchar(32) NOT NULL COMMENT '评论来源类型',
  `parent_comment_id` int(11) DEFAULT NULL COMMENT '父评论ID',
  `user_id` int(11) NOT NULL COMMENT '发表用户ID',
  `floor_comment_id` int(11) DEFAULT NULL COMMENT '楼层评论ID',
  `parent_user_id` int(11) DEFAULT NULL COMMENT '父发表用户ID',
  `like_count` int(11) NOT NULL DEFAULT 0 COMMENT '点赞数',
  `comment_content` varchar(1024) NOT NULL COMMENT '评论内容',
  `comment_info` varchar(256) DEFAULT NULL COMMENT '评论额外信息',
  `create_time` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_comment_source` (`source`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='评论表';

-- ----------------------------
-- 家庭/友链表
-- ----------------------------
DROP TABLE IF EXISTS `family`;
CREATE TABLE `family` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id_id` int(11) NOT NULL COMMENT '用户ID',
  `bg_cover` varchar(256) NOT NULL COMMENT '背景封面',
  `man_cover` varchar(256) NOT NULL COMMENT '男生头像',
  `woman_cover` varchar(256) NOT NULL COMMENT '女生头像',
  `man_name` varchar(32) NOT NULL COMMENT '男生昵称',
  `woman_name` varchar(32) NOT NULL COMMENT '女生昵称',
  `timing` varchar(32) NOT NULL COMMENT '计时',
  `countdown_title` varchar(32) DEFAULT NULL COMMENT '倒计时标题',
  `countdown_time` varchar(32) DEFAULT NULL COMMENT '倒计时时间',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否启用',
  `family_info` varchar(1024) DEFAULT NULL COMMENT '额外信息',
  `like_count` int(11) NOT NULL DEFAULT 0 COMMENT '点赞数',
  `create_time` datetime(6) NOT NULL,
  `update_time` datetime(6) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='家庭/友链表';

-- ----------------------------
-- 网站信息表
-- ----------------------------
DROP TABLE IF EXISTS `web_info`;
CREATE TABLE `web_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `web_name` varchar(16) NOT NULL COMMENT '网站名称',
  `web_title` varchar(512) NOT NULL COMMENT '网站信息',
  `notices` varchar(512) DEFAULT NULL COMMENT '公告',
  `footer` varchar(256) NOT NULL COMMENT '页脚',
  `background_image` varchar(256) DEFAULT NULL COMMENT '背景',
  `avatar` varchar(256) NOT NULL COMMENT '头像',
  `random_avatar` longtext DEFAULT NULL COMMENT '随机头像',
  `random_name` varchar(4096) DEFAULT NULL COMMENT '随机名称',
  `random_cover` longtext DEFAULT NULL COMMENT '随机封面',
  `waifu_json` longtext DEFAULT NULL COMMENT '看板娘消息',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='网站信息表';

-- ----------------------------
-- 资源表
-- ----------------------------
DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id_id` int(11) NOT NULL COMMENT '用户ID',
  `type` varchar(32) NOT NULL COMMENT '资源类型',
  `path` varchar(256) NOT NULL COMMENT '资源路径',
  `size` int(11) DEFAULT NULL COMMENT '资源大小(字节)',
  `mime_type` varchar(256) DEFAULT NULL COMMENT 'MIME类型',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否启用',
  `create_time` datetime(6) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='资源表';

-- ----------------------------
-- 资源路径表
-- ----------------------------
DROP TABLE IF EXISTS `resource_path`;
CREATE TABLE `resource_path` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(64) NOT NULL COMMENT '标题',
  `classify` varchar(32) DEFAULT NULL COMMENT '分类',
  `cover` varchar(256) DEFAULT NULL COMMENT '封面',
  `url` varchar(256) DEFAULT NULL COMMENT '链接',
  `introduction` varchar(1024) DEFAULT NULL COMMENT '简介',
  `type` varchar(32) NOT NULL COMMENT '资源类型',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否启用',
  `remark` longtext DEFAULT NULL COMMENT '备注',
  `create_time` datetime(6) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='资源路径表';

-- ----------------------------
-- 树洞表
-- ----------------------------
DROP TABLE IF EXISTS `tree_hole`;
CREATE TABLE `tree_hole` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `avatar` varchar(256) DEFAULT NULL COMMENT '头像',
  `username` varchar(32) NOT NULL COMMENT '用户名',
  `message` varchar(64) NOT NULL COMMENT '留言',
  `create_time` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='树洞表';

-- ----------------------------
-- 微言表
-- ----------------------------
DROP TABLE IF EXISTS `wei_yan`;
CREATE TABLE `wei_yan` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL COMMENT '用户ID',
  `like_count` int(11) NOT NULL DEFAULT 0 COMMENT '点赞数',
  `content` varchar(1024) NOT NULL COMMENT '内容',
  `type` varchar(32) NOT NULL COMMENT '类型',
  `source` int(11) DEFAULT NULL COMMENT '来源标识',
  `is_public` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否公开',
  `create_time` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='微言表';

-- ----------------------------
-- 禁用词表
-- ----------------------------
DROP TABLE IF EXISTS `word`;
CREATE TABLE `word` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `avatar` varchar(256) DEFAULT NULL COMMENT '头像',
  `username` varchar(32) NOT NULL COMMENT '用户名',
  `message` varchar(64) NOT NULL COMMENT '留言',
  `create_time` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='禁用词表';

-- ----------------------------
-- 验证码表
-- ----------------------------
DROP TABLE IF EXISTS `appone_code`;
CREATE TABLE `appone_code` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(254) NOT NULL COMMENT '邮箱',
  `code` varchar(10) NOT NULL COMMENT '验证码',
  `create_time` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_email_create_time` (`email`, `create_time` DESC)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='验证码表';

-- ----------------------------
-- IP地址表
-- ----------------------------
DROP TABLE IF EXISTS `ip`;
CREATE TABLE `ip` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `ip` varchar(20) DEFAULT NULL COMMENT 'IP地址',
  `province` varchar(20) DEFAULT NULL COMMENT '省份',
  `city` varchar(20) DEFAULT NULL COMMENT '城市',
  `create_time` datetime(6) NOT NULL COMMENT '创建时间',
  `k_time` datetime(6) DEFAULT NULL COMMENT '修改时间',
  `mcount` int(11) NOT NULL DEFAULT 0 COMMENT '时间段内登录次数',
  `dcount` int(11) NOT NULL DEFAULT 0 COMMENT '登录天数',
  PRIMARY KEY (`id`),
  KEY `idx_ip_create_time` (`ip`, `create_time`),
  KEY `idx_user_id_create_time` (`user_id`, `create_time`),
  KEY `idx_create_time` (`create_time`),
  KEY `idx_province` (`province`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='IP地址表';

SET FOREIGN_KEY_CHECKS = 1;

-- ----------------------------
-- 初始化管理员账号
-- 用户名: admin
-- 密码: admin123
-- ----------------------------
INSERT INTO `auth_user` (`id`, `password`, `is_superuser`, `username`, `first_name`, `last_name`, `email`, `is_staff`, `is_active`, `date_joined`)
VALUES (1, 'pbkdf2_sha256$720000$admin123salt$hash', 1, 'admin', '', '', 'admin@example.com', 1, 1, NOW());

INSERT INTO `user` (`user_id`, `username`, `email`, `user_status`, `gender`, `user_type`, `create_time`, `update_time`, `deleted`)
VALUES (1, 'admin', 'admin@example.com', 1, 0, 0, NOW(), NOW(), 0);

-- 初始化网站信息
INSERT INTO `web_info` (`web_name`, `web_title`, `notices`, `footer`, `avatar`, `status`)
VALUES ('My Blog', '欢迎来到我的博客', '欢迎访问！', 'Copyright © 2024 My Blog', 'https://example.com/avatar.jpg', 1);
