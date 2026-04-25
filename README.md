# Elian Blog: 基于go-zero与gorm的个人博客系统
<div align="center">

[![简体中文](https://img.shields.io/badge/简体中文-blue?style=for-the-badge&logo=book&logoColor=white)](./README.md)
[![English](https://img.shields.io/badge/English-orange?style=for-the-badge&logo=language&logoColor=white)](./README.md)
</div>

<p align="center">
  <img src="./draw/framework.drawio.svg" alt="https://elian.net.cn" width="800"/>
</p>

---

## 项目结构

```
elian-blog/
├── cmd/
│   └── server/main.go          # 服务入口
│   └── test/main.go            # 初始化管理员/修复密码
├── configs/config.yaml         # 配置文件
├── internal/
│   ├── config/                 # 配置结构体
│   ├── model/                  # GORM 模型 + 自动迁移
│   ├── dao/                    # 数据访问层
│   ├── svc/                    # ServiceContext 依赖注入
│   ├── handler/                # HTTP 处理器
│   │   ├── admin_handler.go    # 旧版 Admin CRUD
│   │   ├── blog_handler.go     # 前台 Blog API
│   │   ├── vecrud_handler.go   # ve-admin CRUD
│   │   └── response.go         # 响应工具函数
│   ├── logic/
│   │   ├── blog/               # 前台业务逻辑
│   │   └── admin/              # 后台管理逻辑
│   ├── types/                  # 请求/响应类型
│   ├── middleware/             # JWT / RBAC / CORS / RateLimit
│   ├── routes/                 # 路由注册
│   └── utils/                  # JWT / 密码加密 (bcrypt)
├── pkg/
│   ├── logger/                 # Zap 日志
│   └── response/               # Blog API 响应格式
├── web-blog/                   # 前台 (Vue3 + Naive UI + Pinia)
└── web-admin/                  # 后台 (Vue3 + Element Plus + ve-admin-element)
```

## 架构

```
HTTP Request → Router → Middleware (JWT/RBAC/CORS) → Handler → Logic → DAO → MySQL
```

| 层级 | 目录 | 职责 |
|------|------|------|
| Handler | `internal/handler/` | 解析请求，调用 Logic，返回响应 |
| Logic | `internal/logic/` | 业务逻辑，调用 DAO |
| DAO | `internal/dao/` | 数据库 CRUD (GORM) |
| Model | `internal/model/` | GORM 模型定义 |
| Types | `internal/types/` | 请求/响应结构体 |
| Middleware | `internal/middleware/` | JWT 认证、RBAC 鉴权 |

| 前端 | 目录 | 技术栈 |
|------|------|--------|
| 博客前台 | `web-blog/` | Vue3 + Naive UI + Pinia + TypeScript + Vite |
| 管理后台 | `web-admin/` | Vue3 + Element Plus + UnoCSS + Pinia + TypeScript + Vite |

---

## 快速开始

### 环境要求

- Go 1.22+
- Node.js 18+ / pnpm
- MySQL 8.0+
- Redis 7.0+

### 1. 克隆项目

```bash
git clone https://github.com/2Elian/elian-blog.git
cd elian-blog
```

### 2. 准备数据库

```sql
CREATE DATABASE elian_blog CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
```

启动时 GORM `AutoMigrate` 自动建表，`SeedMenus` 自动插入默认菜单和角色。

### 3. 修改配置

编辑 `configs/config.yaml`：

```yaml
MySQL:
  DataSource: "root:your_password@tcp(127.0.0.1:3306)/elian_blog?charset=utf8mb4&parseTime=True&loc=Local"

Auth:
  AccessSecret: "your-jwt-secret-key"
```

### 开发模式

```bash
# Backend
go mod tidy
go run cmd/server/main.go          # 启动后端 (localhost:8080)
go build -o server.exe cmd/server/main.go  # 编译

# 初始化管理员账号 / 修复明文密码
go run cmd/test/main.go

# Frontend - 博客前台 (localhost:3000)
cd web-blog && pnpm install && pnpm dev

# Frontend - 管理后台 (localhost:3001)
cd web-admin && pnpm install && pnpm dev
```

### 生产模式
```bash
bash build.sh
./server.exe -f configs/config.yaml
```

### 6. 创建管理员

```bash
go run cmd/test/main.go
```

默认管理员：`admin / admin123`

---

## 站点配置说明

启动后在管理后台 **站点设置** 页面可配置：

| 配置项 | 说明 |
|--------|------|
| 网站名称/头像/简介 | 前台 Header、Footer、Login 页动态展示 |
| 网站公告 | 前台侧边栏通知区域 |
| 备案号 | Footer 显示 |
| Hero 名称/描述/打字文本 | 首页 Hero 区域动态内容 |
| 社交链接 | 前台侧边栏展示 |
| 功能开关 | 评论审核、留言审核、聊天室、音乐播放器、AI 助手、打赏 |

### 关于页面

后台 **关于页面** 编辑器支持结构化编辑：
- 作者信息（姓名、角色、简介、链接、信息卡片）
- 技术栈（分组标签）
- 代表项目（标题、描述、技术标签、图片、链接、竞赛标签）
- 科研与竞赛

数据以 JSON 存储在 `site_config` 表，前台渲染为展示卡片。

---

## 关键设计

### 认证

- JWT (HS256)，72小时过期
- Claims: UserID, Username, Role
- 密码 bcrypt 加密

### RBAC 权限

- 角色: admin, editor, user
- 菜单权限绑定角色，用户关联角色
- 后台 API 通过中间件校验角色

### API 格式

**Admin API** (`/admin-api/v1/`)：
```json
{"flag": 0, "code": 200, "data": {}, "msg": "success"}
```

**Blog API** (`/blog-api/v1/`)：
```json
{"code": 0, "message": "success", "data": {}}
```

### 数据库模型

| 模型 | 说明 |
|------|------|
| User | 用户 (username, password, nickname, avatar, roles) |
| Role | 角色 (menus 多对多) |
| Menu | 菜单 (树形结构) |
| Article | 文章 (summary, content markdown) |
| Category / Tag | 分类/标签 |
| Comment | 评论 (关联 User, Article) |
| Product | 产品 (content markdown, type varchar) |
| FriendLink | 友链 |
| Message | 留言 |
| Page | 页面 |
| SiteConfig | 站点配置 (key-value, JSON) |
| OperationLog | 操作日志 |

---

## API 接口

### 博客前台 `/blog-api/v1/`

| 方法 | 路径 | 说明 | 认证 |
|------|------|------|------|
| POST | /login | 登录 | 否 |
| POST | /register | 注册 | 否 |
| GET | /articles | 文章列表 | 否 |
| GET | /articles/:id | 文章详情 | 否 |
| GET | /articles/:id/comments | 文章评论 | 否 |
| GET | /categories | 分类列表 | 否 |
| GET | /tags | 标签列表 | 否 |
| GET | /products | 产品列表 | 否 |
| GET | /products/:id | 产品详情 | 否 |
| GET | /friend-links | 友链列表 | 否 |
| GET | /site/config | 站点配置 | 否 |
| GET | /site/about | 关于页面 | 否 |
| GET | /user/info | 用户信息 | JWT |
| POST | /comments | 发表评论 | JWT |

### 管理后台 `/admin-api/v1/`

完整的 CRUD 接口覆盖：文章、分类、标签、评论、友链、留言、页面、角色、菜单、用户、产品、站点配置。

---

## 技术栈

| 组件 | 技术 |
|------|------|
| 后端框架 | Go-Zero REST |
| ORM | GORM |
| 数据库 | MySQL 8.0 |
| 缓存 | Redis 7.0 |
| 认证 | JWT (HS256) + bcrypt |
| 日志 | Zap |
| 博客前台 | Vue3 + Naive UI + Pinia + TypeScript + Vite |
| 管理后台 | Vue3 + Element Plus + UnoCSS + Pinia + TypeScript + Vite |
| Markdown | md-editor-v3 (后台) / marked (前台) |

---

## 参考项目

- [ve-admin-element](https://github.com/ve-weiyi/ve-admin-element) — 后台管理前端框架
- [md-editor-v3](https://github.com/imzbf/md-editor-v3) — markdown编辑
- [ve-blog-golang](https://github.com/ve-weiyi/ve-blog-golang) — Go-Zero 后端架构参考
- [ve-blog-naive](https://github.com/ve-weiyi/ve-blog-naive) — 博客前台设计参考
