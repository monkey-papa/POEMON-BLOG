# My-Blog-Go 博客后端

基于 Go + Gin + GORM 构建的博客后端服务。

## 项目特点

- **统一 API 规范**：所有接口使用 POST 方法，JSON 格式传参
- **清晰的代码结构**：按功能模块划分 handlers，代码简洁易读
- **完善的权限控制**：三级权限体系（公开/登录/管理员）
- **JWT 认证**：支持 Token 自动刷新机制
- **性能优化**：数据库连接池、预编译语句等
- **参数验证**：完善的输入验证和 XSS 防护

## 项目结构

```
My-Blog-Go/
├── config/           # 配置文件
│   ├── config.go     # 应用配置
│   └── database.go   # 数据库配置
├── handlers/         # 接口处理函数
│   ├── common.go     # 公共结构体和辅助函数
│   ├── article.go    # 文章相关
│   ├── auth.go       # 认证相关
│   ├── user.go       # 用户管理
│   ├── comment.go    # 评论管理
│   ├── family.go     # 表白墙
│   ├── weiyan.go     # 微言
│   ├── resource.go   # 资源管理
│   ├── resourcepath.go # 资源路径
│   ├── sort.go       # 分类标签
│   ├── webinfo.go    # 网站信息
│   ├── treehole.go   # 树洞
│   ├── words.go      # 禁用词
│   ├── ip.go         # IP统计
│   └── image.go      # 图片上传
├── middleware/       # 中间件
│   ├── auth.go       # 认证中间件
│   └── cors.go       # CORS 中间件
├── models/           # 数据模型
├── response/         # 响应封装
├── routes/           # 路由配置
├── utils/            # 工具函数
│   ├── jwt.go        # JWT 工具
│   ├── password.go   # 密码工具
│   ├── email.go      # 邮件工具
│   ├── qiniu.go      # 七牛云工具
│   └── validator.go  # 参数验证工具
├── static/           # 静态文件
├── main.go           # 入口文件
├── API.md            # API 接口文档
└── README.md
```

## 环境要求

- Go 1.24+
- MySQL 5.7+

## 快速开始

### 1. 安装依赖

```bash
go mod download
```

### 2. 配置环境变量

创建 `.env` 文件：

```env
# 数据库配置
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=myblog

# JWT 密钥（生产环境务必修改）
JWT_SECRET=your-secret-key

# 邮件配置
EMAIL_FROM=your-email@qq.com
EMAIL_PASSWORD=your-email-password
EMAIL_SMTP_HOST=smtp.qq.com
EMAIL_SMTP_PORT=465

# AI 摘要服务（可选）
SUMMARY_URL=
SUMMARY_AUTH=

# 服务器配置
SERVER_PORT=8000
ENVIRONMENT=development

# 默认值配置（可选）
DEFAULT_SORT_ID=11
DEFAULT_LABEL_ID=23

# 验证配置（可选）
MAX_USERNAME_LENGTH=50
MAX_PASSWORD_LENGTH=128
MIN_PASSWORD_LENGTH=6
MAX_ARTICLE_TITLE_LENGTH=200
MAX_COMMENT_LENGTH=2000
MAX_TREEHOLE_LENGTH=500
```

### 3. 初始化数据库

```bash
mysql -u root -p < database/schema.sql
```

### 4. 启动服务

**普通启动：**

```bash
go run main.go
```

**热重载启动（推荐开发时使用）：**

```bash
# 首次使用需要安装 air
go install github.com/air-verse/air@latest

# 如果提示 command not found，需要添加 Go bin 到 PATH
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
source ~/.zshrc

# 使用 air 启动（修改代码后自动重新编译并重启）
air
```

服务启动后访问: `http://localhost:8000`

### 5. 热重载说明

项目已配置 `.air.toml`，使用 `air` 命令启动后：

- ✨ 修改任意 `.go` 文件后自动重新编译
- ✨ 编译成功后自动重启服务器
- ✨ 无需手动 Ctrl+C 再重新运行

| 启动方式         | 修改代码后              |
| ---------------- | ----------------------- |
| `go run main.go` | 手动停止 → 重新运行     |
| `air`            | **自动编译 → 自动重启** |

## API 文档

查看 `API.md` 文件获取完整的接口说明。

## API 规范

### 请求格式

- 所有接口统一使用 `POST` 方法
- 请求体使用 `JSON` 格式
- 路径不以斜杠 `/` 结尾

### 响应格式

```json
{
  "code": 0, // 状态码：0-成功，其他-失败
  "msg": "操作成功",
  "data": {}, // 响应数据
  "error": false,
  "success": true
}
```

### 分页响应

```json
{
  "code": 0,
  "msg": "操作成功",
  "data": {
    "totalCount": 100,
    "pageSize": 10,
    "totalPage": 10,
    "page": 1,
    "list": [],
    "empty": false
  },
  "error": false,
  "success": true
}
```

### 认证方式

在请求头中添加：

```
Authorization: Token <your-token>
```

或

```
Authorization: Bearer <your-token>
```

## API 列表

### 公开接口（无需登录）

| 接口                                   | 说明         |
| -------------------------------------- | ------------ |
| POST /api/user/login                   | 前台登录     |
| POST /api/user/registration            | 用户注册     |
| POST /api/user/updateForForgetPassword | 忘记密码     |
| POST /api/admin/user/login             | 后台登录     |
| POST /api/code                         | 发送验证码   |
| POST /api/article/list                 | 文章列表     |
| POST /api/article/detail               | 文章详情     |
| POST /api/summary                      | 获取文章摘要 |
| POST /api/webInfo/getSortInfo          | 分类信息     |
| POST /api/webInfo/getWebInfo           | 网站信息     |
| POST /api/comment/listComment          | 评论列表     |
| ...                                    |              |

### 需要登录的接口

| 接口                                    | 说明         |
| --------------------------------------- | ------------ |
| POST /api/user/updateUserInfo           | 更新用户信息 |
| POST /api/resource/saveResource         | 保存资源     |
| POST /api/resource/updateImage          | 上传图片     |
| POST /api/admin/comment/boss/addComment | 添加评论     |
| POST /api/article/articleLike           | 文章点赞     |
| ...                                     |              |

### 管理员接口

| 接口                               | 说明                     |
| ---------------------------------- | ------------------------ |
| POST /api/article/list             | 文章列表（viewAll=true） |
| POST /api/article/save             | 保存文章                 |
| POST /api/article/update           | 更新文章                 |
| POST /api/article/delete           | 删除文章                 |
| POST /api/user/list                | 用户列表（isAdmin=true） |
| POST /api/webInfo/saveOrUpdateSort | 保存或更新分类           |
| ...                                |                          |

## 权限说明

| 用户类型 | 值  | 说明       |
| -------- | --- | ---------- |
| Boss     | 0   | 超级管理员 |
| Admin    | 1   | 管理员     |
| Normal   | 2   | 普通用户   |
| Visitor  | 3   | 访客       |

## 参数验证

系统内置了完善的参数验证机制：

| 参数类型 | 验证规则                 | 默认限制       |
| -------- | ------------------------ | -------------- |
| 用户名   | 字母、数字、下划线、中文 | 最大 50 字符   |
| 密码     | 任意字符                 | 6-128 字符     |
| 邮箱     | 标准邮箱格式             | -              |
| 验证码   | 6 位字母数字             | -              |
| 文章标题 | 任意字符                 | 最大 200 字符  |
| 评论内容 | 任意字符                 | 最大 2000 字符 |
| 树洞留言 | 任意字符                 | 最大 500 字符  |

系统会自动进行 XSS 防护，清理危险的 HTML 标签和 JavaScript 代码。

## 安全建议

1. **生产环境务必修改 JWT_SECRET**
2. 使用 HTTPS
3. 配置防火墙规则
4. 定期备份数据库
5. 监控异常访问
6. 定期更新依赖包

## 许可证

MIT License
