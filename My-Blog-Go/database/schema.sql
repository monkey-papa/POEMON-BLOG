-- ============================================================
-- POEMON-BLOG 数据库初始化脚本
-- 数据库: myblog | 编码: utf8mb4
-- ============================================================
-- 使用方法:
--   mysql -u root -p -e "CREATE DATABASE myblog CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"
--   mysql -u root -p myblog < schema.sql
--
-- 初始管理员账号: admin / admin123（登录后请及时修改密码）
-- ============================================================

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- 用户认证表
-- ----------------------------
DROP TABLE IF EXISTS `auth_user`;
CREATE TABLE `auth_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `password` varchar(128) NOT NULL,
  `last_login` datetime(3) DEFAULT NULL,
  `is_superuser` tinyint(1) DEFAULT '0',
  `username` varchar(150) NOT NULL,
  `first_name` varchar(150) NOT NULL DEFAULT '',
  `last_name` varchar(150) NOT NULL DEFAULT '',
  `email` varchar(254) NOT NULL DEFAULT '',
  `is_staff` tinyint(1) DEFAULT '0',
  `is_active` tinyint(1) DEFAULT '1',
  `date_joined` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户认证表';

-- ----------------------------
-- Token 表
-- ----------------------------
DROP TABLE IF EXISTS `authtoken_token`;
CREATE TABLE `authtoken_token` (
  `key` varchar(40) NOT NULL,
  `created` datetime(3) NOT NULL,
  `user_id` int NOT NULL,
  PRIMARY KEY (`key`),
  UNIQUE KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Token表';

-- ----------------------------
-- 用户信息表
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `user_id` int NOT NULL,
  `username` varchar(32) NOT NULL,
  `phone_number` varchar(16) DEFAULT NULL,
  `email` varchar(32) DEFAULT NULL,
  `user_status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否启用',
  `gender` bigint DEFAULT '0' COMMENT '性别: 0-保密, 1-男, 2-女',
  `open_id` varchar(128) DEFAULT NULL,
  `avatar` varchar(256) DEFAULT NULL COMMENT '头像',
  `admire` varchar(32) DEFAULT NULL COMMENT '赞赏',
  `introduction` varchar(4096) DEFAULT NULL COMMENT '简介',
  `user_type` bigint DEFAULT '2' COMMENT '用户类型: 0-Boss, 1-管理员, 2-普通用户, 3-访客',
  `create_time` datetime(3) DEFAULT NULL,
  `update_time` datetime(3) DEFAULT NULL,
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
  `id` int NOT NULL AUTO_INCREMENT,
  `sort_name` varchar(32) NOT NULL COMMENT '分类名称',
  `sort_description` varchar(256) NOT NULL COMMENT '分类描述',
  `sort_type` bigint DEFAULT '1' COMMENT '分类类型: 0-导航栏分类, 1-普通分类',
  `priority` bigint DEFAULT NULL COMMENT '优先级',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='分类表';

-- ----------------------------
-- 标签表
-- ----------------------------
DROP TABLE IF EXISTS `label`;
CREATE TABLE `label` (
  `id` int NOT NULL AUTO_INCREMENT,
  `sort_id` bigint unsigned DEFAULT NULL COMMENT '分类ID',
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
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户ID',
  `sort_id` bigint unsigned DEFAULT NULL COMMENT '分类ID',
  `label_id` bigint unsigned DEFAULT NULL COMMENT '标签ID',
  `article_cover` varchar(256) DEFAULT NULL COMMENT '封面',
  `article_title` varchar(32) NOT NULL COMMENT '文章标题',
  `article_content` longtext NOT NULL COMMENT '文章内容',
  `summary` text COMMENT '文章摘要',
  `view_count` bigint DEFAULT '0' COMMENT '浏览量',
  `like_count` bigint DEFAULT '0' COMMENT '点赞数',
  `view_status` tinyint(1) DEFAULT '1' COMMENT '是否可见',
  `password` varchar(128) DEFAULT NULL COMMENT '密码',
  `recommend_status` tinyint(1) DEFAULT '0' COMMENT '是否推荐',
  `comment_status` tinyint(1) DEFAULT '1' COMMENT '是否启用评论',
  `create_time` datetime(3) DEFAULT NULL,
  `update_time` datetime(3) DEFAULT NULL,
  `update_by` varchar(32) DEFAULT NULL COMMENT '最终修改人',
  `deleted` tinyint(1) DEFAULT '0' COMMENT '是否删除',
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
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id_id` bigint unsigned DEFAULT NULL COMMENT '用户ID',
  `article_id_id` bigint NOT NULL COMMENT '文章ID',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id_id`),
  KEY `idx_article_id` (`article_id_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='文章点赞表';

-- ----------------------------
-- 评论表
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` int NOT NULL AUTO_INCREMENT,
  `source` bigint DEFAULT NULL COMMENT '评论来源标识',
  `type` varchar(32) NOT NULL COMMENT '评论来源类型',
  `parent_comment_id` bigint unsigned DEFAULT NULL COMMENT '父评论ID',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '发表用户ID',
  `floor_comment_id` bigint unsigned DEFAULT NULL COMMENT '楼层评论ID',
  `parent_user_id` bigint unsigned DEFAULT NULL COMMENT '父发表用户ID',
  `like_count` bigint DEFAULT '0' COMMENT '点赞数',
  `comment_content` varchar(1024) NOT NULL COMMENT '评论内容',
  `comment_info` varchar(256) DEFAULT NULL COMMENT '评论额外信息',
  `create_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_comment_source` (`source`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='评论表';

-- ----------------------------
-- 家庭/表白墙表
-- ----------------------------
DROP TABLE IF EXISTS `family`;
CREATE TABLE `family` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id_id` bigint unsigned DEFAULT NULL COMMENT '用户ID',
  `bg_cover` varchar(256) NOT NULL COMMENT '背景封面',
  `man_cover` varchar(256) NOT NULL COMMENT '男生头像',
  `woman_cover` varchar(256) NOT NULL COMMENT '女生头像',
  `man_name` varchar(32) NOT NULL COMMENT '男生昵称',
  `woman_name` varchar(32) NOT NULL COMMENT '女生昵称',
  `timing` varchar(32) NOT NULL COMMENT '计时',
  `countdown_title` varchar(32) DEFAULT NULL COMMENT '倒计时标题',
  `countdown_time` varchar(32) DEFAULT NULL COMMENT '倒计时时间',
  `status` tinyint(1) DEFAULT '1' COMMENT '是否启用',
  `family_info` varchar(1024) DEFAULT NULL COMMENT '额外信息',
  `like_count` bigint DEFAULT '0' COMMENT '点赞数',
  `create_time` datetime(3) DEFAULT NULL,
  `update_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='家庭/表白墙表';

-- ----------------------------
-- 网站信息表
-- ----------------------------
DROP TABLE IF EXISTS `web_info`;
CREATE TABLE `web_info` (
  `id` int NOT NULL AUTO_INCREMENT,
  `web_name` varchar(16) NOT NULL COMMENT '网站名称',
  `web_title` varchar(512) NOT NULL COMMENT '网站信息',
  `notices` varchar(512) DEFAULT NULL COMMENT '公告',
  `footer` varchar(256) NOT NULL COMMENT '页脚',
  `background_image` varchar(256) DEFAULT NULL COMMENT '背景',
  `avatar` varchar(256) NOT NULL COMMENT '头像',
  `random_avatar` text DEFAULT NULL COMMENT '随机头像',
  `random_name` varchar(4096) DEFAULT NULL COMMENT '随机名称',
  `random_cover` text DEFAULT NULL COMMENT '随机封面',
  `waifu_json` text DEFAULT NULL COMMENT '看板娘消息',
  `status` tinyint(1) DEFAULT '1' COMMENT '是否启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='网站信息表';

-- ----------------------------
-- 资源表
-- ----------------------------
DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id_id` bigint unsigned DEFAULT NULL COMMENT '用户ID',
  `type` varchar(32) NOT NULL COMMENT '资源类型',
  `path` varchar(256) NOT NULL COMMENT '资源路径',
  `size` bigint DEFAULT NULL COMMENT '资源大小(字节)',
  `mime_type` varchar(256) DEFAULT NULL COMMENT 'MIME类型',
  `status` tinyint(1) DEFAULT '1' COMMENT '是否启用',
  `create_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='资源表';

-- ----------------------------
-- 资源路径表（友链等）
-- ----------------------------
DROP TABLE IF EXISTS `resource_path`;
CREATE TABLE `resource_path` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(64) NOT NULL COMMENT '标题',
  `classify` varchar(32) DEFAULT NULL COMMENT '分类',
  `cover` varchar(256) DEFAULT NULL COMMENT '封面',
  `url` varchar(256) DEFAULT NULL COMMENT '链接',
  `introduction` varchar(1024) DEFAULT NULL COMMENT '简介',
  `type` varchar(32) NOT NULL COMMENT '资源类型',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否启用',
  `remark` text DEFAULT NULL COMMENT '备注',
  `create_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='资源路径表';

-- ----------------------------
-- 树洞表
-- ----------------------------
DROP TABLE IF EXISTS `tree_hole`;
CREATE TABLE `tree_hole` (
  `id` int NOT NULL AUTO_INCREMENT,
  `avatar` varchar(256) DEFAULT NULL COMMENT '头像',
  `user_id` bigint unsigned DEFAULT '0' COMMENT '用户ID',
  `username` varchar(32) NOT NULL COMMENT '用户名',
  `message` varchar(64) NOT NULL COMMENT '留言',
  `create_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_tree_hole_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='树洞表';

-- ----------------------------
-- 微言表
-- ----------------------------
DROP TABLE IF EXISTS `wei_yan`;
CREATE TABLE `wei_yan` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户ID',
  `like_count` bigint DEFAULT '0' COMMENT '点赞数',
  `content` varchar(1024) NOT NULL COMMENT '内容',
  `type` varchar(32) NOT NULL COMMENT '类型',
  `source` bigint DEFAULT NULL COMMENT '来源标识',
  `is_public` tinyint(1) DEFAULT '0' COMMENT '是否公开',
  `create_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='微言表';

-- ----------------------------
-- 禁用词表
-- ----------------------------
DROP TABLE IF EXISTS `word`;
CREATE TABLE `word` (
  `id` int NOT NULL AUTO_INCREMENT,
  `avatar` varchar(256) DEFAULT NULL COMMENT '头像',
  `username` varchar(32) NOT NULL COMMENT '用户名',
  `message` varchar(64) NOT NULL COMMENT '留言',
  `create_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='禁用词表';

-- ----------------------------
-- 验证码表
-- ----------------------------
DROP TABLE IF EXISTS `appone_code`;
CREATE TABLE `appone_code` (
  `id` int NOT NULL AUTO_INCREMENT,
  `email` longtext DEFAULT NULL COMMENT '邮箱',
  `code` varchar(10) NOT NULL COMMENT '验证码',
  `create_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_email_create_time` (`email`(254), `create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='验证码表';

-- ----------------------------
-- IP 地址表
-- ----------------------------
DROP TABLE IF EXISTS `ip`;
CREATE TABLE `ip` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` bigint DEFAULT '0' COMMENT '用户ID',
  `ip` varchar(20) DEFAULT NULL COMMENT 'IP地址',
  `province` varchar(20) DEFAULT NULL COMMENT '省份',
  `city` varchar(20) DEFAULT NULL COMMENT '城市',
  `create_time` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `k_time` datetime(3) DEFAULT NULL COMMENT '修改时间',
  `mcount` bigint DEFAULT '0' COMMENT '时间段内登录次数',
  `dcount` bigint DEFAULT '0' COMMENT '登录天数',
  PRIMARY KEY (`id`),
  KEY `idx_ip_create_time` (`ip`, `create_time`),
  KEY `idx_user_id_create_time` (`user_id`, `create_time`),
  KEY `idx_create_time` (`create_time`),
  KEY `idx_province` (`province`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='IP地址表';

SET FOREIGN_KEY_CHECKS = 1;

-- ============================================================
-- 初始数据
-- ============================================================

-- 管理员账号（密码: admin123）
INSERT INTO `auth_user` (`id`, `password`, `is_superuser`, `username`, `first_name`, `last_name`, `email`, `is_staff`, `is_active`, `date_joined`)
VALUES (1, 'pbkdf2_sha256$720000$admin123defaultsalt1234$Y2V3iODhAyc1jUEJwRL4XqS5YigRzdGsYzHuX8Nbbn4=', 1, 'admin', '', '', 'admin@example.com', 1, 1, NOW());

INSERT INTO `user` (`user_id`, `username`, `email`, `user_status`, `gender`, `user_type`, `create_time`, `update_time`, `deleted`)
VALUES (1, 'admin', 'admin@example.com', 1, 0, 0, NOW(), NOW(), 0);

-- 默认分类
INSERT INTO `sort` (`id`, `sort_name`, `sort_description`, `sort_type`, `priority`) VALUES
(1, '技术', '技术相关文章', 1, 1),
(2, '生活', '生活随笔', 1, 2),
(3, '导航', '工具导航', 0, 3);

-- 默认标签
INSERT INTO `label` (`id`, `sort_id`, `label_name`, `label_description`) VALUES
(1, 1, '前端', '前端开发相关'),
(2, 1, '后端', '后端开发相关'),
(3, 1, 'Vue', 'Vue.js 框架'),
(4, 1, 'Go', 'Go 语言'),
(5, 2, '随笔', '日常随笔'),
(6, 3, '工具', '实用工具');

-- 网站信息
INSERT INTO `web_info` (`web_name`, `web_title`, `notices`, `footer`, `avatar`, `status`)
VALUES ('My Blog', '欢迎来到我的博客', '欢迎访问！这是一个基于 Vue3 + Go 的个人博客系统。', 'Copyright © 2024 My Blog. All rights reserved.', '', 1);

-- 示例文章
INSERT INTO `article` (`user_id`, `sort_id`, `label_id`, `article_title`, `article_content`, `summary`, `view_count`, `like_count`, `view_status`, `recommend_status`, `comment_status`, `create_time`, `update_time`, `deleted`)
VALUES (1, 1, 1, '你好，世界！', '# 欢迎使用 POEMON-BLOG\n\n这是你的第一篇文章，你可以在后台管理面板中编辑或删除它。\n\n## 功能特性\n\n- Markdown 编辑器\n- 代码高亮\n- 评论系统\n- 分类标签\n- 主题切换\n\n开始写作吧！', '这是博客的第一篇示例文章，展示基本的 Markdown 功能。', 0, 0, 1, 1, 1, NOW(), NOW(), 0);
