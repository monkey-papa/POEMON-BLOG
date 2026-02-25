<div align="center">

# POEMON-BLOG

**Vue 3 + Go 全栈个人博客系统**

一个功能丰富、界面美观的现代化博客，前后端分离，开箱即用。

[![Stars](https://img.shields.io/github/stars/monkey-papa/POEMON-BLOG?style=flat-square&logo=github)](https://github.com/monkey-papa/POEMON-BLOG/stargazers)
[![Forks](https://img.shields.io/github/forks/monkey-papa/POEMON-BLOG?style=flat-square&logo=github)](https://github.com/monkey-papa/POEMON-BLOG/network/members)
[![License](https://img.shields.io/github/license/monkey-papa/POEMON-BLOG?style=flat-square)](LICENSE)
[![Vue](https://img.shields.io/badge/Vue-3.x-4FC08D?style=flat-square&logo=vue.js)](https://vuejs.org/)
[![Go](https://img.shields.io/badge/Go-1.24-00ADD8?style=flat-square&logo=go)](https://go.dev/)
[![TypeScript](https://img.shields.io/badge/TypeScript-5.x-3178C6?style=flat-square&logo=typescript)](https://www.typescriptlang.org/)

[在线演示](https://blog.zjh2002.icu) · [功能介绍](#功能特性) · [快速开始](#快速开始) · [部署指南](#生产环境部署) · [API 文档](My-Blog-Go/API.md)

</div>

---

## 预览

<!-- 请替换为你自己的截图 -->
<!-- 建议截图尺寸：1280x720 或更大 -->

| 首页 | 文章详情 |
|:---:|:---:|
| ![首页](<img width="1470" height="799" alt="image" src="https://github.com/user-attachments/assets/1a6c378d-5e20-42e4-aa84-6ff23e718c4a" />
) | ![文章详情](<img width="1470" height="799" alt="image" src="https://github.com/user-attachments/assets/08cba9b5-9b74-4acf-9bed-72a889fea44b" />
) |

| 后台管理 | 移动端适配 |
|:---:|:---:|
| ![后台管理](<img width="1470" height="799" alt="image" src="https://github.com/user-attachments/assets/df163a90-5533-4749-a528-ac62d6f1a7bb" />
) | ![移动端](<img width="383" height="671" alt="image" src="https://github.com/user-attachments/assets/66412777-a8b1-4441-8a62-b458488e2172" />
) |

## 版本历史

| 版本 | 技术栈 | 分支/链接 |
|------|--------|----------|
| **v3.0** (当前) | Vue 3 + Go (Gin) | `Vue3-Go` 分支 |
| v2.0 | Vue 3 + Python (Django) | [POEMON-BLOG-v2.0](https://github.com/monkey-papa/POEMON-BLOG-v2.0) |
| v1.0 | Vue 2 + Python (Django) | `main` 分支 |

## 技术栈

| 层级   | 技术                                               |
| ------ | -------------------------------------------------- |
| 前端   | Vue 3 + TypeScript + Vite + Element Plus + Pinia   |
| 后端   | Go 1.24 + Gin + GORM                               |
| 数据库 | MySQL 5.7+                                         |
| 其他   | APlayer 音乐播放器、Live2D 看板娘、Markdown 编辑器 |

## 功能特性

<table>
<tr>
<td>

**内容管理**
- 文章发布 / 编辑 / 分类 / 标签
- Markdown 编辑器 + 代码高亮
- AI 文章摘要（DeepSeek / OpenAI）
- 旅行日记

</td>
<td>

**社交互动**
- 评论系统（表情、多级回复）
- 树洞 / 微言（匿名留言）
- 表白墙
- 友链管理

</td>
</tr>
<tr>
<td>

**个性化**
- 背景主题切换（图片/渐变/纯色）
- APlayer 音乐播放器
- Live2D 看板娘
- 天气信息展示

</td>
<td>

**系统能力**
- JWT 认证 + 三级权限体系
- 后台管理面板
- 响应式设计，适配移动端
- Docker 一键部署

</td>
</tr>
</table>

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
git clone https://github.com/monkey-papa/POEMON-BLOG.git
cd POEMON-BLOG
git checkout Vue3-Go
```

### 2. 初始化数据库

```bash
mysql -u root -p -e "CREATE DATABASE myblog CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"
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
go mod download
go run main.go
```

后端启动后监听 `http://localhost:8000`

<details>
<summary>开发模式推荐使用热重载</summary>

```bash
go install github.com/air-verse/air@latest
air
```

</details>

### 5. 启动前端

```bash
cd My-Blog-Vue3
npm install
npm run dev
```

前端启动后访问 `http://localhost:81`

### 6. 开始使用

| 地址 | 说明 |
|------|------|
| `http://localhost:81` | 博客前台首页 |
| `http://localhost:81/admin` | 后台管理面板 |
| `http://localhost:8000/api/...` | API 接口 |

## 个性化配置

部署前，修改以下文件定制你自己的博客：

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

**`My-Blog-Vue3/index.html`** — 修改 `<title>` 和加载动画文字

**`My-Blog-Vue3/src/views/about.vue`** — 替换关于页背景图

**`My-Blog-Vue3/src/views/common/footer.vue`** / **`myAside.vue`** — 修改 GitHub 链接

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
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o server .
```

#### 3. 上传到服务器

```bash
scp -r My-Blog-Vue3/dist/* user@your-server:/var/www/blog/
scp My-Blog-Go/server user@your-server:/opt/blog/
scp .env.example user@your-server:/opt/blog/.env
# 登录服务器后编辑 /opt/blog/.env 填入生产环境配置
```

#### 4. 服务器配置

<details>
<summary>systemd 服务配置</summary>

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

</details>

<details>
<summary>Nginx 配置</summary>

```nginx
server {
    listen 80;
    server_name your-domain.com;

    root /var/www/blog;
    index index.html;

    location / {
        try_files $uri $uri/ /index.html;
    }

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

</details>

### 方式二：Docker 部署

```bash
cd My-Blog-Go

docker build -t poemon-blog-backend .

docker run -d \
  --name blog-backend \
  --env-file .env \
  -p 8000:8000 \
  --restart unless-stopped \
  poemon-blog-backend
```

前端构建后作为静态文件由 Nginx 提供服务。

### HTTPS 配置

```bash
# Let's Encrypt 自动申请
sudo apt install certbot python3-certbot-nginx
sudo certbot --nginx -d your-domain.com
```

或通过**宝塔面板**一键申请 SSL 证书。

## 常用命令

```bash
# 后端服务管理
sudo systemctl status my-blog      # 查看状态
sudo systemctl restart my-blog     # 重启
sudo journalctl -u my-blog -f      # 实时日志

# 前端开发
cd My-Blog-Vue3 && npm run dev     # 开发服务器
cd My-Blog-Vue3 && npm run build   # 生产构建

# 后端开发
cd My-Blog-Go && go run main.go    # 直接启动
cd My-Blog-Go && air               # 热重载
```

## API 文档

详见 [My-Blog-Go/API.md](My-Blog-Go/API.md) 和 [My-Blog-Go/README.md](My-Blog-Go/README.md)。

## Star 趋势

[![Star History Chart](https://api.star-history.com/svg?repos=monkey-papa/POEMON-BLOG&type=Date)](https://star-history.com/#monkey-papa/POEMON-BLOG&Date)

## 许可证

[MIT License](LICENSE) - 免费使用，欢迎贡献代码。

## 致谢

如果这个项目对你有帮助，请点一个 **Star** 支持一下！

<div align="center">
<a href="https://github.com/monkey-papa/POEMON-BLOG/stargazers">
<img src="https://img.shields.io/github/stars/monkey-papa/POEMON-BLOG?style=social" alt="Star this repo">
</a>
</div>
