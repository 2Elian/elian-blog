# Elian Blog

基于 Go-Zero + Vue3 的全栈个人博客系统，采用前后端分离架构。

## 项目结构

```
elian-blog/
├── cmd/
│   └── server/main.go          # 服务入口
├── configs/config.yaml         # 配置文件
├── internal/
│   ├── config/                 # 配置结构体
│   ├── model/                  # 数据模型 (GORM)
│   ├── dao/                    # 数据访问层
│   ├── svc/                    # ServiceContext 依赖注入
│   ├── handler/                # HTTP 处理器
│   ├── logic/
│   │   ├── blog/               # 前台业务逻辑
│   │   └── admin/              # 后台管理逻辑
│   ├── types/                  # 请求/响应类型
│   ├── middleware/             # JWT / RBAC / CORS
│   ├── routes/                 # 路由注册
│   └── utils/                  # JWT / 密码加密
├── pkg/
│   ├── logger/                 # Zap 日志
│   └── response/               # 旧版响应格式（blog API 用）
├── web-blog/                   # 博客前台 (Vue3 + Naive UI)
└── web-admin/                  # 后台管理 (Vue3 + Element Plus + ve-admin-element)
```

## 架构说明

### 后端 (Go-Zero)

```
HTTP Request → Router → Middleware → Handler → Logic → DAO → Database
```

| 层级 | 目录 | 职责 |
|------|------|------|
| Handler | `internal/handler/` | 解析请求，调用 Logic，返回响应 |
| Logic | `internal/logic/` | 业务逻辑，调用 DAO |
| DAO | `internal/dao/` | 数据库 CRUD |
| Model | `internal/model/` | GORM 模型定义 |
| Types | `internal/types/` | 请求/响应结构体 |
| Middleware | `internal/middleware/` | JWT 认证、RBAC 鉴权 |

### 前端

| 项目 | 目录 | 技术栈 |
|------|------|--------|
| 博客前台 | `web-blog/` | Vue3 + Naive UI + Pinia + TypeScript |
| 后台管理 | `web-admin/` | Vue3 + Element Plus + Pinia + TypeScript (基于 ve-admin-element) |

---

## 快速开始

### 环境要求

- Go 1.22+
- Node.js 18+
- MySQL 8.0+
- Redis 7.0+

### 1. 准备数据库

```sql
CREATE DATABASE elian_blog CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
```

启动时 GORM `AutoMigrate` 自动建表，`SeedMenus` 自动插入默认菜单。

### 2. 配置

编辑 `configs/config.yaml`，修改数据库密码和 JWT 密钥。

### 3. 启动后端

```bash
go mod tidy
go run cmd/server/main.go
```

后端运行在 `http://localhost:8080`。

### 4. 启动前端

```bash
# 博客前台
cd web-blog && pnpm install && pnpm dev    # localhost:3000

# 后台管理
cd web-admin && pnpm install && pnpm dev   # localhost:3001
```

### 5. 创建管理员账号

```bash
go run cmd/test/main.go
```

默认管理员：`admin / admin123`。

---

## API 响应格式

### Admin API（ve-admin-element 格式）

```json
{
  "flag": 0,
  "code": 200,
  "data": {},
  "msg": "success",
  "trace_id": ""
}
```

错误响应：

```json
{
  "flag": 0,
  "code": 401,
  "data": null,
  "msg": "未登录",
  "trace_id": ""
}
```

### Blog API（旧格式）

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

---

## API 接口一览

### 博客前台 `/blog-api/v1/`

| 方法 | 路径 | 说明 | 认证 |
|------|------|------|------|
| POST | /login | 登录 | 否 |
| POST | /register | 注册 | 否 |
| GET | /articles | 文章列表 | 否 |
| GET | /articles/:id | 文章详情 | 否 |
| GET | /categories | 分类列表 | 否 |
| GET | /tags | 标签列表 | 否 |
| GET | /articles/:id/comments | 文章评论 | 否 |
| GET | /friend-links | 友链列表 | 否 |
| GET | /messages | 留言列表 | 否 |
| GET | /pages | 页面列表 | 否 |
| GET | /site/config | 站点配置 | 否 |
| GET | /user/info | 用户信息 | JWT |
| POST | /comments | 发表评论 | JWT |
| POST | /messages | 发表留言 | JWT |

### 后台管理 `/admin-api/v1/`

| 方法 | 路径 | 说明 | 认证 |
|------|------|------|------|
| POST | /login | 登录 | 否 |
| GET | /logout | 登出 | JWT |
| GET | /user/get_user_info | 用户信息 | JWT |
| GET | /user/get_user_menus | 用户菜单 | JWT |
| GET | /user/get_user_roles | 用户角色 | JWT |
| GET | /user/get_user_apis | 用户权限 | JWT |
| POST | /article/find_article_list | 文章列表 | JWT |
| POST | /article/get_article | 文章详情 | JWT |
| POST | /article/add_article | 新增文章 | JWT |
| PUT | /article/update_article | 更新文章 | JWT |
| DELETE | /article/delete_article | 删除文章 | JWT |
| POST | /category/find_category_list | 分类列表 | JWT |
| POST | /category/add_category | 新增分类 | JWT |
| PUT | /category/update_category | 更新分类 | JWT |
| DELETE | /category/deletes_category | 删除分类 | JWT |
| POST | /tag/find_tag_list | 标签列表 | JWT |
| POST | /tag/add_tag | 新增标签 | JWT |
| PUT | /tag/update_tag | 更新标签 | JWT |
| DELETE | /tag/deletes_tag | 删除标签 | JWT |
| POST | /comment/find_comment_back_list | 评论列表 | JWT |
| PUT | /comment/update_comment_status | 更新评论状态 | JWT |
| DELETE | /comment/deletes_comment | 删除评论 | JWT |
| POST | /friend/find_friend_list | 友链列表 | JWT |
| POST | /friend/add_friend | 新增友链 | JWT |
| PUT | /friend/update_friend | 更新友链 | JWT |
| DELETE | /friend/deletes_friend | 删除友链 | JWT |
| POST | /message/find_message_list | 留言列表 | JWT |
| PUT | /message/update_message_status | 更新留言状态 | JWT |
| DELETE | /message/deletes_message | 删除留言 | JWT |
| POST | /page/find_page_list | 页面列表 | JWT |
| POST | /page/add_page | 新增页面 | JWT |
| PUT | /page/update_page | 更新页面 | JWT |
| DELETE | /page/delete_page | 删除页面 | JWT |
| POST | /role/find_role_list | 角色列表 | JWT |
| POST | /role/add_role | 新增角色 | JWT |
| PUT | /role/update_role | 更新角色 | JWT |
| DELETE | /role/deletes_role | 删除角色 | JWT |
| POST | /role/find_role_resources | 角色资源 | JWT |
| PUT | /role/update_role_menus | 更新角色菜单 | JWT |
| POST | /menu/find_menu_list | 菜单列表 | JWT |
| POST | /menu/add_menu | 新增菜单 | JWT |
| PUT | /menu/update_menu | 更新菜单 | JWT |
| DELETE | /menu/deletes_menu | 删除菜单 | JWT |
| POST | /account/find_account_list | 用户列表 | JWT |
| PUT | /account/update_account_status | 更新用户状态 | JWT |
| GET | /admin | 首页统计 | JWT |
| GET | /admin/get_website_config | 网站配置 | JWT |
| PUT | /admin/update_website_config | 更新配置 | JWT |
| POST | /upload/upload_file | 文件上传 | JWT |

---

## 技术栈

| 组件 | 技术 |
|------|------|
| 后端框架 | Go-Zero rest |
| ORM | GORM |
| 数据库 | MySQL |
| 缓存 | Redis |
| 认证 | JWT (golang-jwt/jwt/v5) |
| 日志 | Zap |
| 博客前端 | Vue3 + Naive UI + Pinia + TypeScript + Vite |
| 管理后台 | Vue3 + Element Plus + UnoCSS + Pinia + TypeScript + Vite |

---

## 参考项目

- [ve-admin-element](https://github.com/ve-weiyi/ve-admin-element) — 后台管理前端基础框架
- [ve-blog-golang](https://github.com/ve-weiyi/ve-blog-golang) — Go-Zero 后端架构参考
- [ve-blog-naive](https://github.com/ve-weiyi/ve-blog-naive) — 博客前台设计参考
