/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80035 (8.0.35)
 Source Host           : localhost:3306
 Source Schema         : myblog-open

 Target Server Type    : MySQL
 Target Server Version : 80035 (8.0.35)
 File Encoding         : 65001

 Date: 02/09/2024 10:36:16
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for appone_code
-- ----------------------------
DROP TABLE IF EXISTS `appone_code`;
CREATE TABLE `appone_code`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `email` varchar(254) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `code` varchar(10) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `create_time` datetime(6) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 159 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of appone_code
-- ----------------------------

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `article_cover` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `article_title` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `article_content` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `view_count` int NOT NULL,
  `like_count` int NOT NULL,
  `view_status` tinyint(1) NOT NULL,
  `password` varchar(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `recommend_status` tinyint(1) NOT NULL,
  `comment_status` tinyint(1) NOT NULL,
  `create_time` datetime(6) NOT NULL,
  `update_time` datetime(6) NOT NULL,
  `update_by` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `deleted` tinyint(1) NOT NULL,
  `label_id` int NOT NULL,
  `sort_id` int NOT NULL,
  `user_id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `article_label_id_id_2af2b7f7_fk_label_id`(`label_id` ASC) USING BTREE,
  INDEX `article_sort_id_id_5870156c_fk_sort_id`(`sort_id` ASC) USING BTREE,
  INDEX `article_user_id_id_19d49033_fk_auth_user_id`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 64 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of article
-- ----------------------------

-- ----------------------------
-- Table structure for auth_group
-- ----------------------------
DROP TABLE IF EXISTS `auth_group`;
CREATE TABLE `auth_group`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(150) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_group
-- ----------------------------

-- ----------------------------
-- Table structure for auth_group_permissions
-- ----------------------------
DROP TABLE IF EXISTS `auth_group_permissions`;
CREATE TABLE `auth_group_permissions`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `group_id` int NOT NULL,
  `permission_id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `auth_group_permissions_group_id_permission_id_0cd325b0_uniq`(`group_id` ASC, `permission_id` ASC) USING BTREE,
  INDEX `auth_group_permissio_permission_id_84c5c92e_fk_auth_perm`(`permission_id` ASC) USING BTREE,
  CONSTRAINT `auth_group_permissio_permission_id_84c5c92e_fk_auth_perm` FOREIGN KEY (`permission_id`) REFERENCES `auth_permission` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `auth_group_permissions_group_id_b120cbf9_fk_auth_group_id` FOREIGN KEY (`group_id`) REFERENCES `auth_group` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_group_permissions
-- ----------------------------

-- ----------------------------
-- Table structure for auth_permission
-- ----------------------------
DROP TABLE IF EXISTS `auth_permission`;
CREATE TABLE `auth_permission`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `content_type_id` int NOT NULL,
  `codename` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `auth_permission_content_type_id_codename_01ab375a_uniq`(`content_type_id` ASC, `codename` ASC) USING BTREE,
  CONSTRAINT `auth_permission_content_type_id_2f476e4b_fk_django_co` FOREIGN KEY (`content_type_id`) REFERENCES `django_content_type` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 93 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_permission
-- ----------------------------
INSERT INTO `auth_permission` VALUES (1, 'Can add log entry', 1, 'add_logentry');
INSERT INTO `auth_permission` VALUES (2, 'Can change log entry', 1, 'change_logentry');
INSERT INTO `auth_permission` VALUES (3, 'Can delete log entry', 1, 'delete_logentry');
INSERT INTO `auth_permission` VALUES (4, 'Can view log entry', 1, 'view_logentry');
INSERT INTO `auth_permission` VALUES (5, 'Can add permission', 2, 'add_permission');
INSERT INTO `auth_permission` VALUES (6, 'Can change permission', 2, 'change_permission');
INSERT INTO `auth_permission` VALUES (7, 'Can delete permission', 2, 'delete_permission');
INSERT INTO `auth_permission` VALUES (8, 'Can view permission', 2, 'view_permission');
INSERT INTO `auth_permission` VALUES (9, 'Can add group', 3, 'add_group');
INSERT INTO `auth_permission` VALUES (10, 'Can change group', 3, 'change_group');
INSERT INTO `auth_permission` VALUES (11, 'Can delete group', 3, 'delete_group');
INSERT INTO `auth_permission` VALUES (12, 'Can view group', 3, 'view_group');
INSERT INTO `auth_permission` VALUES (13, 'Can add user', 4, 'add_user');
INSERT INTO `auth_permission` VALUES (14, 'Can change user', 4, 'change_user');
INSERT INTO `auth_permission` VALUES (15, 'Can delete user', 4, 'delete_user');
INSERT INTO `auth_permission` VALUES (16, 'Can view user', 4, 'view_user');
INSERT INTO `auth_permission` VALUES (17, 'Can add content type', 5, 'add_contenttype');
INSERT INTO `auth_permission` VALUES (18, 'Can change content type', 5, 'change_contenttype');
INSERT INTO `auth_permission` VALUES (19, 'Can delete content type', 5, 'delete_contenttype');
INSERT INTO `auth_permission` VALUES (20, 'Can view content type', 5, 'view_contenttype');
INSERT INTO `auth_permission` VALUES (21, 'Can add session', 6, 'add_session');
INSERT INTO `auth_permission` VALUES (22, 'Can change session', 6, 'change_session');
INSERT INTO `auth_permission` VALUES (23, 'Can delete session', 6, 'delete_session');
INSERT INTO `auth_permission` VALUES (24, 'Can view session', 6, 'view_session');
INSERT INTO `auth_permission` VALUES (25, 'Can add 网站信息表', 7, 'add_webinfo');
INSERT INTO `auth_permission` VALUES (26, 'Can change 网站信息表', 7, 'change_webinfo');
INSERT INTO `auth_permission` VALUES (27, 'Can delete 网站信息表', 7, 'delete_webinfo');
INSERT INTO `auth_permission` VALUES (28, 'Can view 网站信息表', 7, 'view_webinfo');
INSERT INTO `auth_permission` VALUES (29, 'Can add 用户信息表', 8, 'add_client');
INSERT INTO `auth_permission` VALUES (30, 'Can change 用户信息表', 8, 'change_client');
INSERT INTO `auth_permission` VALUES (31, 'Can delete 用户信息表', 8, 'delete_client');
INSERT INTO `auth_permission` VALUES (32, 'Can view 用户信息表', 8, 'view_client');
INSERT INTO `auth_permission` VALUES (33, 'Can add Token', 9, 'add_token');
INSERT INTO `auth_permission` VALUES (34, 'Can change Token', 9, 'change_token');
INSERT INTO `auth_permission` VALUES (35, 'Can delete Token', 9, 'delete_token');
INSERT INTO `auth_permission` VALUES (36, 'Can view Token', 9, 'view_token');
INSERT INTO `auth_permission` VALUES (37, 'Can add token', 10, 'add_tokenproxy');
INSERT INTO `auth_permission` VALUES (38, 'Can change token', 10, 'change_tokenproxy');
INSERT INTO `auth_permission` VALUES (39, 'Can delete token', 10, 'delete_tokenproxy');
INSERT INTO `auth_permission` VALUES (40, 'Can view token', 10, 'view_tokenproxy');
INSERT INTO `auth_permission` VALUES (41, 'Can add 分类', 11, 'add_sort');
INSERT INTO `auth_permission` VALUES (42, 'Can change 分类', 11, 'change_sort');
INSERT INTO `auth_permission` VALUES (43, 'Can delete 分类', 11, 'delete_sort');
INSERT INTO `auth_permission` VALUES (44, 'Can view 分类', 11, 'view_sort');
INSERT INTO `auth_permission` VALUES (45, 'Can add 标签', 12, 'add_label');
INSERT INTO `auth_permission` VALUES (46, 'Can change 标签', 12, 'change_label');
INSERT INTO `auth_permission` VALUES (47, 'Can delete 标签', 12, 'delete_label');
INSERT INTO `auth_permission` VALUES (48, 'Can view 标签', 12, 'view_label');
INSERT INTO `auth_permission` VALUES (49, 'Can add 文章表', 13, 'add_article');
INSERT INTO `auth_permission` VALUES (50, 'Can change 文章表', 13, 'change_article');
INSERT INTO `auth_permission` VALUES (51, 'Can delete 文章表', 13, 'delete_article');
INSERT INTO `auth_permission` VALUES (52, 'Can view 文章表', 13, 'view_article');
INSERT INTO `auth_permission` VALUES (53, 'Can add 文章评论表', 14, 'add_comment');
INSERT INTO `auth_permission` VALUES (54, 'Can change 文章评论表', 14, 'change_comment');
INSERT INTO `auth_permission` VALUES (55, 'Can delete 文章评论表', 14, 'delete_comment');
INSERT INTO `auth_permission` VALUES (56, 'Can view 文章评论表', 14, 'view_comment');
INSERT INTO `auth_permission` VALUES (57, 'Can add 资源信息', 15, 'add_resource');
INSERT INTO `auth_permission` VALUES (58, 'Can change 资源信息', 15, 'change_resource');
INSERT INTO `auth_permission` VALUES (59, 'Can delete 资源信息', 15, 'delete_resource');
INSERT INTO `auth_permission` VALUES (60, 'Can view 资源信息', 15, 'view_resource');
INSERT INTO `auth_permission` VALUES (61, 'Can add 树洞', 16, 'add_treehole');
INSERT INTO `auth_permission` VALUES (62, 'Can change 树洞', 16, 'change_treehole');
INSERT INTO `auth_permission` VALUES (63, 'Can delete 树洞', 16, 'delete_treehole');
INSERT INTO `auth_permission` VALUES (64, 'Can view 树洞', 16, 'view_treehole');
INSERT INTO `auth_permission` VALUES (65, 'Can add 禁用词', 17, 'add_words');
INSERT INTO `auth_permission` VALUES (66, 'Can change 禁用词', 17, 'change_words');
INSERT INTO `auth_permission` VALUES (67, 'Can delete 禁用词', 17, 'delete_words');
INSERT INTO `auth_permission` VALUES (68, 'Can view 禁用词', 17, 'view_words');
INSERT INTO `auth_permission` VALUES (69, 'Can add 家庭信息', 18, 'add_family');
INSERT INTO `auth_permission` VALUES (70, 'Can change 家庭信息', 18, 'change_family');
INSERT INTO `auth_permission` VALUES (71, 'Can delete 家庭信息', 18, 'delete_family');
INSERT INTO `auth_permission` VALUES (72, 'Can view 家庭信息', 18, 'view_family');
INSERT INTO `auth_permission` VALUES (73, 'Can add 微言表', 19, 'add_weiyan');
INSERT INTO `auth_permission` VALUES (74, 'Can change 微言表', 19, 'change_weiyan');
INSERT INTO `auth_permission` VALUES (75, 'Can delete 微言表', 19, 'delete_weiyan');
INSERT INTO `auth_permission` VALUES (76, 'Can view 微言表', 19, 'view_weiyan');
INSERT INTO `auth_permission` VALUES (77, 'Can add 资源路径', 20, 'add_resourcepath');
INSERT INTO `auth_permission` VALUES (78, 'Can change 资源路径', 20, 'change_resourcepath');
INSERT INTO `auth_permission` VALUES (79, 'Can delete 资源路径', 20, 'delete_resourcepath');
INSERT INTO `auth_permission` VALUES (80, 'Can view 资源路径', 20, 'view_resourcepath');
INSERT INTO `auth_permission` VALUES (81, 'Can add code', 21, 'add_code');
INSERT INTO `auth_permission` VALUES (82, 'Can change code', 21, 'change_code');
INSERT INTO `auth_permission` VALUES (83, 'Can delete code', 21, 'delete_code');
INSERT INTO `auth_permission` VALUES (84, 'Can view code', 21, 'view_code');
INSERT INTO `auth_permission` VALUES (85, 'Can add IP地址表', 22, 'add_ip');
INSERT INTO `auth_permission` VALUES (86, 'Can change IP地址表', 22, 'change_ip');
INSERT INTO `auth_permission` VALUES (87, 'Can delete IP地址表', 22, 'delete_ip');
INSERT INTO `auth_permission` VALUES (88, 'Can view IP地址表', 22, 'view_ip');
INSERT INTO `auth_permission` VALUES (89, 'Can add 点赞', 23, 'add_article_like');
INSERT INTO `auth_permission` VALUES (90, 'Can change 点赞', 23, 'change_article_like');
INSERT INTO `auth_permission` VALUES (91, 'Can delete 点赞', 23, 'delete_article_like');
INSERT INTO `auth_permission` VALUES (92, 'Can view 点赞', 23, 'view_article_like');

-- ----------------------------
-- Table structure for auth_user
-- ----------------------------
DROP TABLE IF EXISTS `auth_user`;
CREATE TABLE `auth_user`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `password` varchar(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `last_login` datetime(6) NULL DEFAULT NULL,
  `is_superuser` tinyint(1) NOT NULL,
  `username` varchar(150) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `first_name` varchar(150) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `last_name` varchar(150) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `email` varchar(254) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `is_staff` tinyint(1) NOT NULL,
  `is_active` tinyint(1) NOT NULL,
  `date_joined` datetime(6) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 84 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_user
-- ----------------------------
INSERT INTO `auth_user` VALUES (9, 'pbkdf2_sha256$390000$8jFLb9l829VHXbcs5BI2RS$pCuxMB6o+lIW28ZJF10PCCwx5YuF8oMukCjY2pT3Shs=', NULL, 0, 'monkey-papa', '', '', '', 0, 1, '2023-07-17 02:19:01.386899');

-- ----------------------------
-- Table structure for auth_user_groups
-- ----------------------------
DROP TABLE IF EXISTS `auth_user_groups`;
CREATE TABLE `auth_user_groups`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `group_id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `auth_user_groups_user_id_group_id_94350c0c_uniq`(`user_id` ASC, `group_id` ASC) USING BTREE,
  INDEX `auth_user_groups_group_id_97559544_fk_auth_group_id`(`group_id` ASC) USING BTREE,
  CONSTRAINT `auth_user_groups_group_id_97559544_fk_auth_group_id` FOREIGN KEY (`group_id`) REFERENCES `auth_group` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `auth_user_groups_user_id_6a12ed8b_fk_auth_user_id` FOREIGN KEY (`user_id`) REFERENCES `auth_user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_user_groups
-- ----------------------------

-- ----------------------------
-- Table structure for auth_user_user_permissions
-- ----------------------------
DROP TABLE IF EXISTS `auth_user_user_permissions`;
CREATE TABLE `auth_user_user_permissions`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `permission_id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `auth_user_user_permissions_user_id_permission_id_14a6b632_uniq`(`user_id` ASC, `permission_id` ASC) USING BTREE,
  INDEX `auth_user_user_permi_permission_id_1fbb5f2c_fk_auth_perm`(`permission_id` ASC) USING BTREE,
  CONSTRAINT `auth_user_user_permi_permission_id_1fbb5f2c_fk_auth_perm` FOREIGN KEY (`permission_id`) REFERENCES `auth_permission` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `auth_user_user_permissions_user_id_a95ead1b_fk_auth_user_id` FOREIGN KEY (`user_id`) REFERENCES `auth_user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_user_user_permissions
-- ----------------------------

-- ----------------------------
-- Table structure for authtoken_token
-- ----------------------------
DROP TABLE IF EXISTS `authtoken_token`;
CREATE TABLE `authtoken_token`  (
  `key` varchar(40) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `created` datetime(6) NOT NULL,
  `user_id` int NOT NULL,
  PRIMARY KEY (`key`) USING BTREE,
  UNIQUE INDEX `user_id`(`user_id` ASC) USING BTREE,
  CONSTRAINT `authtoken_token_user_id_35299eff_fk_auth_user_id` FOREIGN KEY (`user_id`) REFERENCES `auth_user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of authtoken_token
-- ----------------------------
INSERT INTO `authtoken_token` VALUES ('1d1f93ff371fd596ba8ae00fee892795cb03ec41', '2023-07-17 02:19:01.815348', 9);

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `source` int NOT NULL,
  `type` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `parent_comment_id` int NULL DEFAULT NULL,
  `user_id` int NOT NULL,
  `floor_comment_id` int NULL DEFAULT NULL,
  `parent_user_id` int NULL DEFAULT NULL,
  `like_count` int NOT NULL,
  `comment_content` varchar(1024) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `comment_info` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `create_time` datetime(6) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_comment_source`(`source` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 203 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of comment
-- ----------------------------

-- ----------------------------
-- Table structure for django_admin_log
-- ----------------------------
DROP TABLE IF EXISTS `django_admin_log`;
CREATE TABLE `django_admin_log`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `action_time` datetime(6) NOT NULL,
  `object_id` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `object_repr` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `action_flag` smallint UNSIGNED NOT NULL,
  `change_message` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `content_type_id` int NULL DEFAULT NULL,
  `user_id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `django_admin_log_content_type_id_c4bce8eb_fk_django_co`(`content_type_id` ASC) USING BTREE,
  INDEX `django_admin_log_user_id_c564eba6_fk_auth_user_id`(`user_id` ASC) USING BTREE,
  CONSTRAINT `django_admin_log_content_type_id_c4bce8eb_fk_django_co` FOREIGN KEY (`content_type_id`) REFERENCES `django_content_type` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `django_admin_log_user_id_c564eba6_fk_auth_user_id` FOREIGN KEY (`user_id`) REFERENCES `auth_user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 251 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of django_admin_log
-- ----------------------------

-- ----------------------------
-- Table structure for django_content_type
-- ----------------------------
DROP TABLE IF EXISTS `django_content_type`;
CREATE TABLE `django_content_type`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `app_label` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `model` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `django_content_type_app_label_model_76bd3d3b_uniq`(`app_label` ASC, `model` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of django_content_type
-- ----------------------------
INSERT INTO `django_content_type` VALUES (1, 'admin', 'logentry');
INSERT INTO `django_content_type` VALUES (13, 'appone', 'article');
INSERT INTO `django_content_type` VALUES (23, 'appone', 'article_like');
INSERT INTO `django_content_type` VALUES (8, 'appone', 'client');
INSERT INTO `django_content_type` VALUES (21, 'appone', 'code');
INSERT INTO `django_content_type` VALUES (14, 'appone', 'comment');
INSERT INTO `django_content_type` VALUES (18, 'appone', 'family');
INSERT INTO `django_content_type` VALUES (22, 'appone', 'ip');
INSERT INTO `django_content_type` VALUES (12, 'appone', 'label');
INSERT INTO `django_content_type` VALUES (15, 'appone', 'resource');
INSERT INTO `django_content_type` VALUES (20, 'appone', 'resourcepath');
INSERT INTO `django_content_type` VALUES (11, 'appone', 'sort');
INSERT INTO `django_content_type` VALUES (16, 'appone', 'treehole');
INSERT INTO `django_content_type` VALUES (7, 'appone', 'webinfo');
INSERT INTO `django_content_type` VALUES (19, 'appone', 'weiyan');
INSERT INTO `django_content_type` VALUES (17, 'appone', 'words');
INSERT INTO `django_content_type` VALUES (3, 'auth', 'group');
INSERT INTO `django_content_type` VALUES (2, 'auth', 'permission');
INSERT INTO `django_content_type` VALUES (4, 'auth', 'user');
INSERT INTO `django_content_type` VALUES (9, 'authtoken', 'token');
INSERT INTO `django_content_type` VALUES (10, 'authtoken', 'tokenproxy');
INSERT INTO `django_content_type` VALUES (5, 'contenttypes', 'contenttype');
INSERT INTO `django_content_type` VALUES (6, 'sessions', 'session');

-- ----------------------------
-- Table structure for django_migrations
-- ----------------------------
DROP TABLE IF EXISTS `django_migrations`;
CREATE TABLE `django_migrations`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `app` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `applied` datetime(6) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 38 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of django_migrations
-- ----------------------------
INSERT INTO `django_migrations` VALUES (1, 'contenttypes', '0001_initial', '2023-06-06 15:49:28.650150');
INSERT INTO `django_migrations` VALUES (2, 'auth', '0001_initial', '2023-06-06 15:49:29.467032');
INSERT INTO `django_migrations` VALUES (3, 'admin', '0001_initial', '2023-06-06 15:49:29.605645');
INSERT INTO `django_migrations` VALUES (4, 'admin', '0002_logentry_remove_auto_add', '2023-06-06 15:49:29.613195');
INSERT INTO `django_migrations` VALUES (5, 'admin', '0003_logentry_add_action_flag_choices', '2023-06-06 15:49:29.620019');
INSERT INTO `django_migrations` VALUES (6, 'appone', '0001_initial', '2023-06-06 15:49:29.746964');
INSERT INTO `django_migrations` VALUES (7, 'contenttypes', '0002_remove_content_type_name', '2023-06-06 15:49:29.869233');
INSERT INTO `django_migrations` VALUES (8, 'auth', '0002_alter_permission_name_max_length', '2023-06-06 15:49:29.931507');
INSERT INTO `django_migrations` VALUES (9, 'auth', '0003_alter_user_email_max_length', '2023-06-06 15:49:30.003809');
INSERT INTO `django_migrations` VALUES (10, 'auth', '0004_alter_user_username_opts', '2023-06-06 15:49:30.013780');
INSERT INTO `django_migrations` VALUES (11, 'auth', '0005_alter_user_last_login_null', '2023-06-06 15:49:30.079605');
INSERT INTO `django_migrations` VALUES (12, 'auth', '0006_require_contenttypes_0002', '2023-06-06 15:49:30.083606');
INSERT INTO `django_migrations` VALUES (13, 'auth', '0007_alter_validators_add_error_messages', '2023-06-06 15:49:30.092571');
INSERT INTO `django_migrations` VALUES (14, 'auth', '0008_alter_user_username_max_length', '2023-06-06 15:49:30.164983');
INSERT INTO `django_migrations` VALUES (15, 'auth', '0009_alter_user_last_name_max_length', '2023-06-06 15:49:30.233095');
INSERT INTO `django_migrations` VALUES (16, 'auth', '0010_alter_group_name_max_length', '2023-06-06 15:49:30.300073');
INSERT INTO `django_migrations` VALUES (17, 'auth', '0011_update_proxy_permissions', '2023-06-06 15:49:30.309527');
INSERT INTO `django_migrations` VALUES (18, 'auth', '0012_alter_user_first_name_max_length', '2023-06-06 15:49:30.386015');
INSERT INTO `django_migrations` VALUES (19, 'authtoken', '0001_initial', '2023-06-06 15:49:30.481610');
INSERT INTO `django_migrations` VALUES (20, 'authtoken', '0002_auto_20160226_1747', '2023-06-06 15:49:30.505546');
INSERT INTO `django_migrations` VALUES (21, 'authtoken', '0003_tokenproxy', '2023-06-06 15:49:30.510555');
INSERT INTO `django_migrations` VALUES (22, 'sessions', '0001_initial', '2023-06-06 15:49:30.558728');
INSERT INTO `django_migrations` VALUES (23, 'appone', '0002_sort', '2023-07-18 06:53:26.792773');
INSERT INTO `django_migrations` VALUES (24, 'appone', '0003_label', '2023-07-18 07:54:00.042160');
INSERT INTO `django_migrations` VALUES (25, 'appone', '0004_article', '2023-07-18 08:21:15.118957');
INSERT INTO `django_migrations` VALUES (26, 'appone', '0002_comment_comment_idx_comment_source', '2023-07-20 08:24:02.800646');
INSERT INTO `django_migrations` VALUES (27, 'appone', '0003_resource', '2023-07-21 10:57:36.185119');
INSERT INTO `django_migrations` VALUES (28, 'appone', '0004_treehole', '2023-07-22 13:52:52.175300');
INSERT INTO `django_migrations` VALUES (29, 'appone', '0002_words', '2023-07-23 05:31:15.574012');
INSERT INTO `django_migrations` VALUES (30, 'appone', '0002_alter_comment_parent_comment_id', '2023-08-03 09:46:02.305827');
INSERT INTO `django_migrations` VALUES (31, 'appone', '0003_family', '2023-08-03 09:46:02.393524');
INSERT INTO `django_migrations` VALUES (32, 'appone', '0004_weiyan_weiyan_wei_yan_user_id_32c565_idx', '2023-08-04 05:57:53.058930');
INSERT INTO `django_migrations` VALUES (33, 'appone', '0005_resourcepath', '2023-08-06 10:21:04.990636');
INSERT INTO `django_migrations` VALUES (34, 'appone', '0002_code', '2023-08-11 09:47:57.743533');
INSERT INTO `django_migrations` VALUES (35, 'appone', '0002_ip', '2023-08-14 07:31:49.093882');
INSERT INTO `django_migrations` VALUES (36, 'appone', '0002_article_like', '2023-08-14 08:30:39.605842');
INSERT INTO `django_migrations` VALUES (37, 'appone', '0002_ip_k_time', '2023-08-16 02:36:22.230090');

-- ----------------------------
-- Table structure for django_session
-- ----------------------------
DROP TABLE IF EXISTS `django_session`;
CREATE TABLE `django_session`  (
  `session_key` varchar(40) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `session_data` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `expire_date` datetime(6) NOT NULL,
  PRIMARY KEY (`session_key`) USING BTREE,
  INDEX `django_session_expire_date_a5c62663`(`expire_date` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of django_session
-- ----------------------------
INSERT INTO `django_session` VALUES ('2zszf1cbmt8k5fk55oquvkge7szoua3o', '.eJxVjEEOwiAQRe_C2hCHToG6dN8zEJgZpGpoUtqV8e7apAvd_vfef6kQt7WErckSJlYXBer0u6VID6k74Hust1nTXNdlSnpX9EGbHmeW5_Vw_w5KbOVbi0NLmKijjokgWUkIYnkAjFYcoI9IvYOzY98bk6VzWZKhgQ16yF69P_p0ODY:1qM3di:X6g3vpUpInxrMRbQnv8rQR_CdlNHqAuJRafJi3NfVm4', '2023-08-02 09:38:34.033098');
INSERT INTO `django_session` VALUES ('6zoohxkry2yus7k4785488rhte4aan0z', '.eJxVjEEOwiAQRe_C2hCHToG6dN8zEJgZpGpoUtqV8e7apAvd_vfef6kQt7WErckSJlYXBer0u6VID6k74Hust1nTXNdlSnpX9EGbHmeW5_Vw_w5KbOVbi0NLmKijjokgWUkIYnkAjFYcoI9IvYOzY98bk6VzWZKhgQ16yF69P_p0ODY:1qLk2P:uBcuTQxBvbOrJkHYZsmOXTMGv7Z5Q5dxFyFrlf09vxU', '2023-08-01 12:42:45.083258');
INSERT INTO `django_session` VALUES ('83cf9l7wq3wt38jlzxheq3e3v9q4qzej', '.eJxVjEEOwiAQRe_C2hCHToG6dN8zEJgZpGpoUtqV8e7apAvd_vfef6kQt7WErckSJlYXBer0u6VID6k74Hust1nTXNdlSnpX9EGbHmeW5_Vw_w5KbOVbi0NLmKijjokgWUkIYnkAjFYcoI9IvYOzY98bk6VzWZKhgQ16yF69P_p0ODY:1rRCtC:usfq1u4Ui9VT6qAPRtt2LTnoJzcRBW7cshfNNd2DlQY', '2024-02-03 23:04:06.953425');
INSERT INTO `django_session` VALUES ('9usdfswmh5afyybeovfx1wq8wkxzki9o', '.eJxVjEEOwiAQRe_C2hCHToG6dN8zEJgZpGpoUtqV8e7apAvd_vfef6kQt7WErckSJlYXBer0u6VID6k74Hust1nTXNdlSnpX9EGbHmeW5_Vw_w5KbOVbi0NLmKijjokgWUkIYnkAjFYcoI9IvYOzY98bk6VzWZKhgQ16yF69P_p0ODY:1qRobu:funOvcGzJ4r2WN0Nv3sZdz33NVbhPBskodPJXlF04cA', '2023-08-18 06:48:30.650223');
INSERT INTO `django_session` VALUES ('gchc1lt5i4grm67nq4d5qif4f0dd7wxp', '.eJxVjEEOwiAQRe_C2hCHToG6dN8zEJgZpGpoUtqV8e7apAvd_vfef6kQt7WErckSJlYXBer0u6VID6k74Hust1nTXNdlSnpX9EGbHmeW5_Vw_w5KbOVbi0NLmKijjokgWUkIYnkAjFYcoI9IvYOzY98bk6VzWZKhgQ16yF69P_p0ODY:1qRoMs:FMAWyIuuJEEE2N5WoHCfrPdfAbZCvxps_8IMO9Mx7vo', '2023-08-18 06:32:58.997447');
INSERT INTO `django_session` VALUES ('i1wr4qpsknv8ad58846u4xbfnza7zide', '.eJxVjEEOwiAQRe_C2hCHToG6dN8zEJgZpGpoUtqV8e7apAvd_vfef6kQt7WErckSJlYXBer0u6VID6k74Hust1nTXNdlSnpX9EGbHmeW5_Vw_w5KbOVbi0NLmKijjokgWUkIYnkAjFYcoI9IvYOzY98bk6VzWZKhgQ16yF69P_p0ODY:1qYRVr:Ig4tdB9ZENIaZiy6CZD_Ml7J9phZeyau9BiCXbICz6c', '2023-09-05 21:33:39.827606');
INSERT INTO `django_session` VALUES ('j8jnhkeji33byw7rcdorshnacer496cq', '.eJxVjEEOwiAQRe_C2hCHToG6dN8zEJgZpGpoUtqV8e7apAvd_vfef6kQt7WErckSJlYXBer0u6VID6k74Hust1nTXNdlSnpX9EGbHmeW5_Vw_w5KbOVbi0NLmKijjokgWUkIYnkAjFYcoI9IvYOzY98bk6VzWZKhgQ16yF69P_p0ODY:1qWZat:PMFyd4nO4kQBykB6z6avYrPXZZ2Y40zqy93y90kB67I', '2023-08-31 09:47:07.171390');
INSERT INTO `django_session` VALUES ('ndxbwnsxmyomcqzxzemvbtuu4flbu0k8', '.eJxVjEEOwiAQRe_C2hCHToG6dN8zEJgZpGpoUtqV8e7apAvd_vfef6kQt7WErckSJlYXBer0u6VID6k74Hust1nTXNdlSnpX9EGbHmeW5_Vw_w5KbOVbi0NLmKijjokgWUkIYnkAjFYcoI9IvYOzY98bk6VzWZKhgQ16yF69P_p0ODY:1qXfU3:HvdtuvrnhbX6CA0rqUpcfeLJNEA_teVEftMhc25GykY', '2023-09-03 18:16:35.494251');
INSERT INTO `django_session` VALUES ('qle497s4mc4djo7kcttsyo3bnaif0vf1', '.eJxVjEEOwiAQRe_C2hCHToG6dN8zEJgZpGpoUtqV8e7apAvd_vfef6kQt7WErckSJlYXBer0u6VID6k74Hust1nTXNdlSnpX9EGbHmeW5_Vw_w5KbOVbi0NLmKijjokgWUkIYnkAjFYcoI9IvYOzY98bk6VzWZKhgQ16yF69P_p0ODY:1rBpT7:DwirEII92jctpP8MGjv1Hq3sVxwhw6tp0jIMduoWHHc', '2023-12-23 13:01:37.291866');
INSERT INTO `django_session` VALUES ('ula2jxcexe6tagrkvr1fagxx8r6gp511', '.eJxVjEEOwiAQRe_C2hCHToG6dN8zEJgZpGpoUtqV8e7apAvd_vfef6kQt7WErckSJlYXBer0u6VID6k74Hust1nTXNdlSnpX9EGbHmeW5_Vw_w5KbOVbi0NLmKijjokgWUkIYnkAjFYcoI9IvYOzY98bk6VzWZKhgQ16yF69P_p0ODY:1qYcyY:R2KKO0DsFFEzlxe1L3gv_13iWCH-fGmK5a6SYvCL7co', '2023-09-06 09:48:02.316478');
INSERT INTO `django_session` VALUES ('xhwvxs57qddehzuorfpo15axuqleznvd', '.eJxVjEEOwiAQRe_C2hCHToG6dN8zEJgZpGpoUtqV8e7apAvd_vfef6kQt7WErckSJlYXBer0u6VID6k74Hust1nTXNdlSnpX9EGbHmeW5_Vw_w5KbOVbi0NLmKijjokgWUkIYnkAjFYcoI9IvYOzY98bk6VzWZKhgQ16yF69P_p0ODY:1qRUuV:IiqxclTAylQgbkz3BWgyn0MKDq0zmB3M00WqJ91niYs', '2023-08-17 09:46:23.219775');
INSERT INTO `django_session` VALUES ('y8pmiijq0l0ia2bnzoipbkudytvxoqym', '.eJxVjDEOwjAMRe-SGUUhLpXDyM4ZIjt2SAElUtNOiLtDpQ6w_vfef5lI61Li2nWOk5iz8ebwuzGlh9YNyJ3qrdnU6jJPbDfF7rTbaxN9Xnb376BQL996TFnCAHzCkbPzgJxz1mNAJD8wJ2JAUCEInkEUCBUzsaJT8QjOvD8PGzko:1q6Yx8:NTxf4MOCPbXlYgLq0NPsNljlBhhHpLBFiyCNuGAg-To', '2023-06-20 15:50:34.333675');
INSERT INTO `django_session` VALUES ('ybrb0hsi7bqdaw2o0hqq9gsx85xyvkdp', '.eJxVjEEOwiAQRe_C2hCHToG6dN8zEJgZpGpoUtqV8e7apAvd_vfef6kQt7WErckSJlYXBer0u6VID6k74Hust1nTXNdlSnpX9EGbHmeW5_Vw_w5KbOVbi0NLmKijjokgWUkIYnkAjFYcoI9IvYOzY98bk6VzWZKhgQ16yF69P_p0ODY:1qLDp4:5yxQOxZD14TaNaA8RjEpHX1KYRzvReUE6wwN3cs-tEY', '2023-07-31 02:18:50.199201');
INSERT INTO `django_session` VALUES ('ymbrf6xri357mwg2114cg7rxh6df4t34', '.eJxVjEEOwiAQRe_C2hCHToG6dN8zEJgZpGpoUtqV8e7apAvd_vfef6kQt7WErckSJlYXBer0u6VID6k74Hust1nTXNdlSnpX9EGbHmeW5_Vw_w5KbOVbi0NLmKijjokgWUkIYnkAjFYcoI9IvYOzY98bk6VzWZKhgQ16yF69P_p0ODY:1qXgWd:zdbGuPy7Bbfc2Vus744dfP7yx7666LO9XRIj2L3HcuI', '2023-09-03 19:23:19.916391');
INSERT INTO `django_session` VALUES ('yz8wcklon1bluj3v3nck2r78ylptcnra', '.eJxVjEEOwiAQRe_C2hCHToG6dN8zEJgZpGpoUtqV8e7apAvd_vfef6kQt7WErckSJlYXBer0u6VID6k74Hust1nTXNdlSnpX9EGbHmeW5_Vw_w5KbOVbi0NLmKijjokgWUkIYnkAjFYcoI9IvYOzY98bk6VzWZKhgQ16yF69P_p0ODY:1qLbTj:U7Wd8JfRWr8ZlMp5nx4pM1Cdma5AIe49eqijl3BzDRI', '2023-08-01 03:34:23.490109');

-- ----------------------------
-- Table structure for family
-- ----------------------------
DROP TABLE IF EXISTS `family`;
CREATE TABLE `family`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `bg_cover` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `man_cover` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `woman_cover` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `man_name` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `woman_name` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `timing` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `countdown_title` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `countdown_time` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `status` tinyint(1) NOT NULL,
  `family_info` varchar(1024) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `like_count` int NOT NULL,
  `create_time` datetime(6) NOT NULL,
  `update_time` datetime(6) NOT NULL,
  `user_id_id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `family_user_id_id_f0564d8e_fk_auth_user_id`(`user_id_id` ASC) USING BTREE,
  CONSTRAINT `family_user_id_id_f0564d8e_fk_auth_user_id` FOREIGN KEY (`user_id_id`) REFERENCES `auth_user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of family
-- ----------------------------

-- ----------------------------
-- Table structure for ip
-- ----------------------------
DROP TABLE IF EXISTS `ip`;
CREATE TABLE `ip`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `ip` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `province` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `city` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `create_time` datetime(6) NOT NULL,
  `k_time` datetime(6) NULL DEFAULT NULL,
  `mcount` int NOT NULL,
  `dcount` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5457 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ip
-- ----------------------------
INSERT INTO `ip` VALUES (5456, 0, '127.0.0.1', NULL, NULL, '2024-08-02 21:44:43.767080', NULL, 0, 1);

-- ----------------------------
-- Table structure for label
-- ----------------------------
DROP TABLE IF EXISTS `label`;
CREATE TABLE `label`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `label_name` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `label_description` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `sort_id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `label_sort_id_id_439e7832_fk_sort_id`(`sort_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 39 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of label
-- ----------------------------
INSERT INTO `label` VALUES (15, 'JS', 'JS高级（coderwhy）', 1);
INSERT INTO `label` VALUES (16, 'CSS', 'CSS（了解+真难）', 1);
INSERT INTO `label` VALUES (17, 'Vue', 'Vue2/3', 1);
INSERT INTO `label` VALUES (18, '高仿网易云音乐', 'https://www.zjh2002.icu/', 6);
INSERT INTO `label` VALUES (19, '王者农药', '王者农药', 3);
INSERT INTO `label` VALUES (20, 'My-Blog', 'My-Blog', 6);
INSERT INTO `label` VALUES (21, '不良人-画江湖', '不良人-画江湖', 4);
INSERT INTO `label` VALUES (22, '总结', '总结', 7);
INSERT INTO `label` VALUES (23, '未定义', '未定义', 11);
INSERT INTO `label` VALUES (24, 'NodeJS', 'NodeJS', 1);
INSERT INTO `label` VALUES (25, '博文', '无聊写写', 1);
INSERT INTO `label` VALUES (26, '浏览器', '跨域、请求、渲染等', 1);
INSERT INTO `label` VALUES (27, '优化', '前端项目优化', 1);
INSERT INTO `label` VALUES (28, '前端', '前端', 1);
INSERT INTO `label` VALUES (29, '旅行', '到处“乱跑”', 2);
INSERT INTO `label` VALUES (30, 'axios', '神器axios', 1);
INSERT INTO `label` VALUES (31, 'IPhone', 'IPhone', 13);
INSERT INTO `label` VALUES (32, 'MacBook', 'MacBook', 13);
INSERT INTO `label` VALUES (33, '一些好吃的', '就是炫', 14);
INSERT INTO `label` VALUES (34, '音乐', 'rap', 2);
INSERT INTO `label` VALUES (35, '渡一面试题（JS）', '渡一面试题（JS）', 1);
INSERT INTO `label` VALUES (36, '渡一面试题（Promise）', '渡一面试题（Promise）', 1);
INSERT INTO `label` VALUES (37, '网络', '网络工程', 1);
INSERT INTO `label` VALUES (38, '渡一面试题（工程化）', '渡一面试题（工程化）', 1);

-- ----------------------------
-- Table structure for like
-- ----------------------------
DROP TABLE IF EXISTS `like`;
CREATE TABLE `like`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `article_id_id` bigint NOT NULL,
  `user_id_id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `like_article_id_id_63839db3_fk_article_id`(`article_id_id` ASC) USING BTREE,
  INDEX `like_user_id_id_d204a180_fk_auth_user_id`(`user_id_id` ASC) USING BTREE,
  CONSTRAINT `like_article_id_id_63839db3_fk_article_id` FOREIGN KEY (`article_id_id`) REFERENCES `article` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `like_user_id_id_d204a180_fk_auth_user_id` FOREIGN KEY (`user_id_id`) REFERENCES `auth_user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of like
-- ----------------------------

-- ----------------------------
-- Table structure for resource
-- ----------------------------
DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `type` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `path` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `size` int NULL DEFAULT NULL,
  `mime_type` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` datetime(6) NOT NULL,
  `user_id_id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `resource_user_id_id_f8636fa7_fk_auth_user_id`(`user_id_id` ASC) USING BTREE,
  CONSTRAINT `resource_user_id_id_f8636fa7_fk_auth_user_id` FOREIGN KEY (`user_id_id`) REFERENCES `auth_user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 165 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of resource
-- ----------------------------

-- ----------------------------
-- Table structure for resource_path
-- ----------------------------
DROP TABLE IF EXISTS `resource_path`;
CREATE TABLE `resource_path`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `classify` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `cover` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `url` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `introduction` varchar(1024) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `type` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `status` tinyint(1) NOT NULL,
  `remark` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `create_time` datetime(6) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 118 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of resource_path
-- ----------------------------

-- ----------------------------
-- Table structure for sort
-- ----------------------------
DROP TABLE IF EXISTS `sort`;
CREATE TABLE `sort`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `sort_name` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `sort_description` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `sort_type` int NOT NULL,
  `priority` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sort
-- ----------------------------
INSERT INTO `sort` VALUES (1, '学习笔记', '个人成长史上自学的笔记', 0, 1);
INSERT INTO `sort` VALUES (2, '生活、娱乐、趣闻', '生活中的有趣奇葩事', 0, 3);
INSERT INTO `sort` VALUES (3, '游戏', '劳逸结合', 0, 4);
INSERT INTO `sort` VALUES (4, '动漫', '就看过这一部动漫', 0, 5);
INSERT INTO `sort` VALUES (5, '教程', '教程', 0, 6);
INSERT INTO `sort` VALUES (6, '我的项目', '我做的0-1的项目', 0, 2);
INSERT INTO `sort` VALUES (7, '随想', 'Ding~~', 0, 7);
INSERT INTO `sort` VALUES (11, '未定义', '未定义', 1, 999999);
INSERT INTO `sort` VALUES (13, '苹果全家桶', '苹果全家桶', 1, 8);
INSERT INTO `sort` VALUES (14, '美食', '饿的睡不着', 1, 9);

-- ----------------------------
-- Table structure for tree_hole
-- ----------------------------
DROP TABLE IF EXISTS `tree_hole`;
CREATE TABLE `tree_hole`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `avatar` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `message` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `username` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `create_time` datetime(6) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 53 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of tree_hole
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `phone_number` varchar(16) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `email` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `user_status` tinyint(1) NOT NULL,
  `gender` smallint NULL DEFAULT NULL,
  `open_id` varchar(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `avatar` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `admire` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `introduction` varchar(4096) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `user_type` smallint NOT NULL,
  `create_time` datetime(6) NOT NULL,
  `update_time` datetime(6) NOT NULL,
  `update_by` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `deleted` tinyint(1) NOT NULL,
  `user_id` int NOT NULL,
  `province` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `qiniu_domain` varchar(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `qiniu_bucket_name` varchar(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `qiniu_secret_key` varchar(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `qiniu_access_key` varchar(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username` ASC) USING BTREE,
  UNIQUE INDEX `user_id`(`user_id` ASC) USING BTREE,
  CONSTRAINT `user_user_id_3ad375f5_fk_auth_user_id` FOREIGN KEY (`user_id`) REFERENCES `auth_user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 81 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (7, 'monkey-papa', '15523692545', '1816298537@qq.com', 1, 0, NULL, 'https://www.qiniuyun.zjh2002.icu/images/e61a0eb0423611ee8b6500163e10317e', '0', '这里介绍不了我', 0, '2023-07-17 02:19:01.790727', '2024-07-08 00:25:17.978806', NULL, 0, 9, '湖北', NULL, NULL, NULL, NULL);

-- ----------------------------
-- Table structure for web_info
-- ----------------------------
DROP TABLE IF EXISTS `web_info`;
CREATE TABLE `web_info`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `web_name` varchar(16) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `web_title` varchar(512) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `notices` varchar(512) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `footer` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `background_image` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `avatar` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `random_avatar` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `random_name` varchar(4096) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `random_cover` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `waifu_json` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `status` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of web_info
-- ----------------------------
INSERT INTO `web_info` VALUES (1, 'Monkey-PaPa', 'Blog | Monkey-PaPa', '[\"欢迎来到『Monkey-PaPa』的小站 (￣▽￣)～■干杯□～(￣▽￣) 如果有什么加载不出来或者其它滴小bug,「刷新网页/Ctrl+F5」应该、或许大概能解决\"]', '- 为了寻找你 我搬进鸟的眼睛 经常盯着路过的风 -', 'https://www.qiniuyun.zjh2002.icu/images/changeBg2', 'https://www.qiniuyun.zjh2002.icu/images/avatar', '[]', '[\"Monkey-PaPa\"]', '[\"https://www.qiniuyun.zjh2002.icu/avatar/errorBG.png \",\"https://www.qiniuyun.zjh2002.icu/gif/lazy.gif \",\"https://www.qiniuyun.zjh2002.icu/otherPhoto/switch-feiji.png\",\"https://www.qiniuyun.zjh2002.icu/otherPhoto/switch-guaishou.png\",\"https://www.qiniuyun.zjh2002.icu/avatar/love%E5%9B%BE1.jpg\",\"https://www.qiniuyun.zjh2002.icu/avatar/love%E5%9B%BE2.jpg\",\"https://www.qiniuyun.zjh2002.icu/avatar/%E5%85%89%E9%81%87.webp\",\"https://www.qiniuyun.zjh2002.icu/love/%E7%88%B1%E5%BF%83.png\",\"https://www.qiniuyun.zjh2002.icu/letter/letter-4.jpg\",\"https://www.qiniuyun.zjh2002.icu/letter/letter-2.jpg\",\"https://www.qiniuyun.zjh2002.icu/letter/letter-3.jpg\",\"https://www.qiniuyun.zjh2002.icu/images/%E5%8F%8B%E9%93%BEbg.jpg\",\"https://www.qiniuyun.zjh2002.icu/images/e61a0eb0423611ee8b6500163e10317e\"]', '你好啊', 1);

-- ----------------------------
-- Table structure for wei_yan
-- ----------------------------
DROP TABLE IF EXISTS `wei_yan`;
CREATE TABLE `wei_yan`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `like_count` int NOT NULL,
  `content` varchar(1024) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `type` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `source` int NULL DEFAULT NULL,
  `is_public` tinyint(1) NOT NULL,
  `create_time` datetime(6) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `wei_yan_user_id_32c565_idx`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 44 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of wei_yan
-- ----------------------------

-- ----------------------------
-- Table structure for word
-- ----------------------------
DROP TABLE IF EXISTS `word`;
CREATE TABLE `word`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `message` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `avatar` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `username` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `create_time` datetime(6) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 38 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of word
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
