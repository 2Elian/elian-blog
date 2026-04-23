# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Elian Blog 是基于 Go-Zero + Vue3 的全栈个人博客系统，采用前后端分离架构。后端使用 Go-Zero REST 框架，数据库为 MySQL (GORM)，缓存为 Redis，认证使用 JWT。

## Essential Commands

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

## Service Ports

- **Backend API**: 8080
- **MySQL**: 3306
- **Redis**: 6379
- **Blog Frontend**: 3000
- **Admin Frontend**: 3001

## Project Structure

```
elian-blog/
├── cmd/
│   └── server/main.go          # 服务入口
│   └── test/main.go            # 管理员初始化/密码修复工具
├── configs/config.yaml         # 配置文件
├── internal/
│   ├── config/                 # 配置结构体
│   ├── model/                  # 数据模型 (GORM)
│   ├── dao/                    # 数据访问层
│   ├── svc/                    # ServiceContext 依赖注入
│   ├── handler/                # HTTP 处理器
│   │   ├── admin_handler.go    # 旧版 Admin CRUD handler
│   │   ├── blog_handler.go     # 旧版 Blog handler
│   │   ├── veauth_handler.go   # ve-admin 认证 handler
│   │   ├── vecrud_handler.go   # ve-admin CRUD handler
│   │   └── response.go         # 响应工具函数
│   ├── logic/
│   │   ├── blog/               # 前台业务逻辑 (auth, article, category, tag, comment, friend, message, page, site)
│   │   └── admin/              # 后台管理逻辑 (article, category, tag, comment, friend, message, page, role, menu, site, user, dashboard, veauth)
│   ├── types/                  # 请求/响应类型
│   ├── middleware/             # JWT / RBAC / CORS / RateLimit / OperationLog
│   ├── routes/                 # 路由注册
│   └── utils/                  # JWT / 密码加密 (bcrypt)
├── pkg/
│   ├── logger/                 # Zap 日志
│   └── response/               # 旧版响应格式
├── web-blog/                   # 博客前台 (Vue3 + Naive UI + Pinia + TypeScript)
└── web-admin/                  # 后台管理 (Vue3 + Element Plus + Pinia + TypeScript, 基于 ve-admin-element)
```

## Architecture

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
| Middleware | `internal/middleware/` | JWT 认证、RBAC 鉴权、CORS、限流 |

## API Routes

### Blog API `/blog-api/v1/`

公开接口：login, register, articles, categories, tags, comments, friend-links, messages, pages, site/config

JWT 保护接口：user/info, comments (POST), messages (POST)

### Admin API `/admin-api/v1/`

公开接口：login

JWT 保护接口：logout, user/*, article/*, category/*, tag/*, comment/*, friend/*, message/*, page/*, role/*, menu/*, account/*, admin/*, upload/*

## API Response Formats

**Admin API** (ve-admin-element 格式):
```json
{"flag": 0, "code": 200, "data": {}, "msg": "success", "trace_id": ""}
```

**Blog API** (旧格式):
```json
{"code": 0, "message": "success", "data": {}}
```

## Database Models

- **User** - 用户 (username, password, nickname, avatar, email, status, roles)
- **Role** - 角色 (name, label, description, menus) - many2many with User and Menu
- **Menu** - 菜单 (name, title, path, component, icon, parent_id, sort, type)
- **Article** - 文章
- **Category** - 分类
- **Tag** - 标签
- **Comment** - 评论
- **FriendLink** - 友链
- **Message** - 留言
- **Page** - 页面
- **SiteConfig** - 站点配置
- **OperationLog** - 操作日志

## Authentication

- JWT (HS256), 72小时过期
- Claims: UserID (uint), Username (string), Role (string)
- 密码使用 bcrypt 加密 (`golang.org/x/crypto/bcrypt`)
- 角色: admin, editor, user

## Tech Stack

| 组件 | 技术 |
|------|------|
| 后端框架 | Go-Zero REST |
| ORM | GORM |
| 数据库 | MySQL |
| 缓存 | Redis |
| 认证 | JWT (golang-jwt/jwt/v5) + bcrypt |
| 日志 | Zap |
| 博客前端 | Vue3 + Naive UI + Pinia + TypeScript + Vite |
| 管理后台 | Vue3 + Element Plus + Pinia + TypeScript + Vite |

## Critical Implementation Rules

### Password Security
- 所有密码必须使用 bcrypt 加密存储，禁止明文
- 使用 `utils.HashPassword()` 加密，`utils.CheckPassword()` 验证
- bcrypt 哈希以 `$2a$` 或 `$2b$` 开头

### Database
- GORM AutoMigrate 自动建表
- SeedMenus 自动插入默认菜单
- initRoles 自动创建默认角色 (admin, editor, user)
- 软删除使用 `gorm.DeletedAt`

### API Conventions
- Blog API 前缀: `/blog-api/v1/`
- Admin API 前缀: `/admin-api/v1/`
- Admin API 使用 ve-admin-element 响应格式 (flag/code/data/msg)
- Blog API 使用旧格式 (code/message/data)

### Adding New Features
1. 在 `internal/model/` 定义 GORM 模型
2. 在 `internal/dao/` 创建 DAO
3. 在 `internal/types/` 定义请求/响应类型
4. 在 `internal/logic/` 实现业务逻辑
5. 在 `internal/handler/` 创建 HTTP handler
6. 在 `internal/routes/routes.go` 注册路由
7. 在 `internal/svc/service_context.go` 注入 DAO

## Default Admin Account

- 用户名: `admin`
- 邮箱: `lizimo@elian.net.cn`
- 通过 `go run cmd/test/main.go` 创建
