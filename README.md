# POEMON-BLOG

一个基于 **Vue 3 + Go (Gin)** 的全栈个人博客系统，前后端分离，功能丰富，界面美观。

## 历史

> 1.0版本：main分支

> 2.0版本：https://github.com/monkey-papa/POEMON-BLOG-v2.0

## 预览

> 在线演示站点：https://blog.zjh2002.icu

## 技术栈

| 层级   | 技术                                               |
| ------ | -------------------------------------------------- |
| 前端   | Vue 3 + TypeScript + Vite + Element Plus + Pinia   |
| 后端   | Go 1.24 + Gin + GORM                               |
| 数据库 | MySQL 5.7+                                         |
| 其他   | APlayer 音乐播放器、Live2D 看板娘、Markdown 编辑器 |

## 功能特性

- 文章管理（发布、编辑、分类、标签、搜索）
- Markdown 编辑器，支持代码高亮
- 评论系统（支持表情、回复）
- 用户系统（注册、登录、个人主页）
- 树洞 / 微言（匿名留言板）
- 表白墙
- 友链管理
- 旅行日记
- 背景主题切换（图片、渐变色、纯色）
- APlayer 音乐播放器
- Live2D 看板娘
- AI 文章摘要生成（支持 DeepSeek / OpenAI 兼容接口）
- 天气信息展示
- 后台管理面板
- 响应式设计，适配移动端
- JWT 认证 + 三级权限体系

## 项目结构

```
POEMON-BLOG/
├── My-Blog-Vue3/          # 前端项目
│   ├── src/
│   │   ├── views/         # 页面组件
│   │   ├── utils/         # 工具函数 & 常量配置
│   │   ├── stores/        # Pinia 状态管理
│   │   ├── router/        # 路由配置
│   │   ├── assets/        # 静态资源
│   │   └── types/         # TypeScript 类型
│   ├── .env.production    # 生产环境变量
│   └── vite.config.ts     # Vite 配置
│
├── My-Blog-Go/            # 后端项目
│   ├── config/            # 配置（数据库、应用配置）
│   ├── handlers/          # API 接口处理
│   ├── middleware/        # 中间件（认证、CORS、限流）
│   ├── models/            # 数据模型
│   ├── routes/            # 路由注册
│   ├── utils/             # 工具函数
│   ├── database/          # 数据库初始化脚本
│   ├── static/            # 上传文件存储
│   ├── API.md             # API 接口文档
│   └── main.go            # 入口文件
│
├── .env.example           # 环境变量模板
└── README.md
```

## 快速开始

### 环境要求

- **Node.js** >= 20
- **Go** >= 1.24
- **MySQL** >= 5.7

### 1. 克隆项目

```bash
git clone https://github.com/your-username/POEMON-BLOG.git
cd POEMON-BLOG
```

### 2. 初始化数据库

创建 MySQL 数据库并导入表结构：

```bash
mysql -u root -p -e "CREATE DATABASE myblog CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;"
mysql -u root -p myblog < My-Blog-Go/database/schema.sql
```

> 初始管理员账号：`admin` / `admin123`（登录后请及时修改密码）

### 3. 配置后端

```bash
cd My-Blog-Go

# 复制环境变量模板
cp ../.env.example .env

# 编辑 .env，至少修改以下配置：
# - DB_PASSWORD: 你的数据库密码
# - JWT_SECRET: 改为一个随机字符串（openssl rand -base64 32）
```

<details>
<summary>环境变量说明</summary>

| 变量                   | 必填 | 说明                               |
| ---------------------- | ---- | ---------------------------------- |
| `DB_HOST`              | 是   | 数据库地址，默认 `127.0.0.1`       |
| `DB_PORT`              | 是   | 数据库端口，默认 `3306`            |
| `DB_USER`              | 是   | 数据库用户名                       |
| `DB_PASSWORD`          | 是   | 数据库密码                         |
| `DB_NAME`              | 是   | 数据库名称                         |
| `JWT_SECRET`           | 是   | JWT 密钥，**生产环境务必修改**     |
| `EMAIL_FROM`           | 否   | 发件邮箱（用于验证码、通知）       |
| `EMAIL_PASSWORD`       | 否   | 邮箱 SMTP 授权码                   |
| `AI_API_KEY`           | 否   | AI 摘要服务 API Key（DeepSeek 等） |
| `AI_API_URL`           | 否   | AI 接口地址，默认 DeepSeek         |
| `SITE_NAME`            | 否   | 站点名称                           |
| `SITE_URL`             | 否   | 站点地址                           |
| `WEATHER_API_ID`       | 否   | 天气 API ID（apihz.cn）            |
| `WEATHER_API_KEY`      | 否   | 天气 API Key                       |
| `CORS_ALLOWED_ORIGINS` | 否   | CORS 白名单，多个用逗号分隔        |
| `SERVER_PORT`          | 否   | 后端端口，默认 `8000`              |
| `ENVIRONMENT`          | 否   | `development` 或 `production`      |

</details>

### 4. 启动后端

```bash
cd My-Blog-Go

# 安装依赖
go mod download

# 启动
go run main.go
```

后端启动后监听 `http://localhost:8000`

**开发模式推荐使用热重载：**

```bash
# 安装 air
go install github.com/air-verse/air@latest

# 启动（修改代码自动重编译重启）
air
```

### 5. 启动前端

```bash
cd My-Blog-Vue3

# 安装依赖
npm install

# 启动开发服务器
npm run dev
```

前端启动后访问 `http://localhost:81`

### 6. 访问博客

- 前台首页：`http://localhost:81`
- 后台管理：`http://localhost:81/admin`（使用 admin 账号登录）
- API 接口：`http://localhost:8000/api/...`

## 个性化配置

部署前，你需要修改以下文件来定制你自己的博客：

### 前端配置

**`My-Blog-Vue3/src/utils/constant.ts`**：

| 配置项                | 说明                         |
| --------------------- | ---------------------------- |
| `siteName`            | 站点名称                     |
| `bossEmail`           | 博主邮箱                     |
| `qqNumber`            | 博主QQ                       |
| `wechatId`            | 博主微信                     |
| `userId`              | 博主用户ID（注册后的用户ID） |
| `defaultAvatar`       | 默认头像图片地址             |
| `defaultBackground`   | 默认背景图地址               |
| `qiniuUploadEntrance` | 图床地址（七牛云等）         |
| `favoriteVideo`       | 收藏视频地址                 |
| `about`               | 关于页展示图片               |

**`My-Blog-Vue3/index.html`**：

- 修改 `<title>` 标签为你的站点名称
- 修改加载动画文字

**`My-Blog-Vue3/src/views/about.vue`**：

- 替换关于页面的背景图片

**`My-Blog-Vue3/src/views/common/footer.vue`** 和 **`myAside.vue`**：

- 修改 GitHub 链接为你的仓库地址

### 后端配置

通过 `.env` 文件配置，参见上方环境变量说明。

## 生产环境部署

### 方式一：直接部署

#### 1. 构建前端

```bash
cd My-Blog-Vue3

# 先修改 .env.production 中的域名
# VITE_API_BASE_URL=https://your-domain.com/api
# VITE_WEB_URL=https://your-domain.com
# VITE_SITE_URL=https://your-domain.com

npm run build
```

构建产物在 `dist/` 目录。

#### 2. 编译后端

```bash
cd My-Blog-Go

# 交叉编译 Linux 二进制（如果服务器是 Linux）
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o server .
```

#### 3. 上传到服务器

```bash
# 上传前端静态文件到 Web 服务器目录
scp -r My-Blog-Vue3/dist/* user@your-server:/var/www/blog/

# 上传后端二进制和配置
scp My-Blog-Go/server user@your-server:/opt/blog/
scp .env.example user@your-server:/opt/blog/.env
# 登录服务器后编辑 /opt/blog/.env 填入生产环境配置
```

#### 4. 服务器配置

**创建 systemd 服务管理后端：**

```bash
sudo cat > /etc/systemd/system/my-blog.service << 'EOF'
[Unit]
Description=My Blog Go Backend
After=network.target

[Service]
Type=simple
WorkingDirectory=/opt/blog
ExecStart=/opt/blog/server
Restart=on-failure
RestartSec=5

[Install]
WantedBy=multi-user.target
EOF

sudo systemctl daemon-reload
sudo systemctl enable my-blog
sudo systemctl start my-blog
```

**Nginx 配置（反向代理 + 前端静态文件）：**

```nginx
server {
    listen 80;
    server_name your-domain.com;

    # 前端静态文件
    root /var/www/blog;
    index index.html;

    location / {
        try_files $uri $uri/ /index.html;
    }

    # 后端 API 反向代理
    location /api {
        proxy_pass http://127.0.0.1:8000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_read_timeout 300s;
    }
}
```

### 方式二：Docker 部署

后端提供了 Dockerfile，可以使用 Docker 部署：

```bash
cd My-Blog-Go

# 构建镜像
docker build -t poemon-blog-backend .

# 运行容器
docker run -d \
  --name blog-backend \
  --env-file .env \
  -p 8000:8000 \
  --restart unless-stopped \
  poemon-blog-backend
```

前端构建后作为静态文件由 Nginx 提供服务。

### HTTPS 配置

推荐使用 **Let's Encrypt** 免费证书：

```bash
# 使用 certbot 自动申请
sudo apt install certbot python3-certbot-nginx
sudo certbot --nginx -d your-domain.com
```

或者通过宝塔面板一键申请 SSL 证书。

## 常用命令

```bash
# 后端服务管理
sudo systemctl status my-blog      # 查看状态
sudo systemctl restart my-blog     # 重启
sudo systemctl stop my-blog        # 停止
sudo journalctl -u my-blog -f      # 实时日志

# 前端本地开发
cd My-Blog-Vue3 && npm run dev     # 启动开发服务器
cd My-Blog-Vue3 && npm run build   # 生产构建

# 后端本地开发
cd My-Blog-Go && go run main.go    # 直接启动
cd My-Blog-Go && air               # 热重载启动
```

## API 文档

详见 [My-Blog-Go/API.md](My-Blog-Go/API.md) 和 [My-Blog-Go/README.md](My-Blog-Go/README.md)。

## 许可证

MIT License

## 致谢

如果这个项目对你有帮助，欢迎 Star 支持！
