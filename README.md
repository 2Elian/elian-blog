# Elian Blog

基于 Go-Zero + Vue3 的全栈个人博客系统，采用前后端分离架构。

## 项目结构

```
elian-blog/
├── cmd/
│   └── server/main.go          # Go-Zero 服务入口
│   └── api/main.go             # (旧版 Gin 入口，保留参考)
├── configs/config.yaml         # Go-Zero 配置文件
├── internal/
│   ├── config/                 # 配置结构体
│   ├── model/                  # 数据模型 (GORM Model)
│   ├── dao/                    # 数据访问层
│   ├── svc/                    # ServiceContext 依赖注入
│   ├── handler/                # HTTP 处理器
│   ├── logic/                  # 业务逻辑层
│   │   ├── blog/               # 前台业务逻辑
│   │   └── admin/              # 后台管理逻辑
│   ├── types/                  # 请求/响应类型定义
│   ├── middleware/             # 中间件 (JWT/RBAC/CORS)
│   ├── routes/                 # 路由注册
│   └── utils/                  # 工具函数 (JWT/密码加密)
│   ├── controller/             # (旧版 Gin Controller，保留参考)
│   ├── service/                # (旧版 Gin Service，保留参考)
│   └── router/                 # (旧版 Gin Router，保留参考)
├── pkg/
│   ├── config/                 # 配置加载
│   ├── logger/                 # Zap 日志
│   └── response/               # 统一响应格式
├── web-blog/                   # 博客前台 (Vue3 + Naive UI)
│   └── src/
│       ├── api/                # API 请求封装
│       ├── components/         # 公共组件
│       ├── layouts/            # 布局组件
│       ├── views/              # 页面视图
│       ├── router/             # 路由配置
│       ├── stores/             # Pinia 状态管理
│       └── styles/             # 全局样式
├── web-admin/                  # 后台管理 (Vue3 + Element Plus)
│   └── src/
│       ├── api/                # API 请求封装
│       ├── layouts/            # 布局组件
│       ├── views/              # 管理页面
│       ├── router/             # 路由配置
│       ├── stores/             # Pinia 状态管理
│       └── styles/             # 全局样式
└── web/                        # (旧版前端，保留参考)
```

## 架构说明

### 后端架构 (Go-Zero)

采用 **Go-Zero rest** 框架，Handler-Logic-DAO 分层模式：

```
HTTP Request → Router → Middleware → Handler → Logic → DAO → Database
```

| 层级 | 目录 | 职责 |
|------|------|------|
| Handler | `internal/handler/` | 解析请求参数，调用 Logic，返回响应 |
| Logic | `internal/logic/` | 业务逻辑处理，调用 DAO |
| DAO | `internal/dao/` | 数据库 CRUD 操作 |
| Model | `internal/model/` | GORM 数据模型定义 |
| Types | `internal/types/` | 请求/响应结构体定义 |
| Middleware | `internal/middleware/` | JWT 认证、RBAC 鉴权、CORS |
| Routes | `internal/routes/` | 路由注册与中间件绑定 |
| ServiceContext | `internal/svc/` | 依赖注入容器 |

### 前端架构

| 项目 | 目录 | 技术栈 | 说明 |
|------|------|--------|------|
| 博客前台 | `web-blog/` | Vue3 + Naive UI + Pinia + TypeScript | 面向访客的博客展示 |
| 后台管理 | `web-admin/` | Vue3 + Element Plus + Pinia + TypeScript | 面向管理员的内容管理 |

---

## 快速开始

### 环境要求

- Go 1.22+
- Node.js 18+
- MySQL 8.0+
- Redis 7.0+

### 1. 准备数据库

创建 MySQL 数据库：

```sql
CREATE DATABASE elian_blog CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
```

项目启动时会通过 GORM `AutoMigrate` 自动创建所有表。

### 2. 修改配置

编辑 `configs/config.yaml`：

```yaml
Name: elian-blog-api
Host: 0.0.0.0
Port: 8080
Database:
  Host: 127.0.0.1
  Port: 3306
  Username: root
  Password: your_password    # 改为你的 MySQL 密码
  DBName: elian_blog
Redis:
  Host: 127.0.0.1
  Port: 6379
  Password: ""
JWT:
  Secret: "change-this-in-production"  # 生产环境务必修改
  ExpireHours: 72
```

### 3. 启动后端

```bash
cd elian-blog
go mod tidy
go run cmd/server/main.go
```

后端运行在 `http://localhost:8080`。

### 4. 启动博客前台

```bash
cd web-blog
npm install
npm run dev
```

前台运行在 `http://localhost:3000`。

### 5. 启动后台管理

```bash
cd web-admin
npm install
npm run dev
```

后台运行在 `http://localhost:3001` (或其他端口，见控制台输出)。

### 6. 构建前端

```bash
# 博客前台
cd web-blog && npm run build

# 后台管理
cd web-admin && npm run build
```

构建产物在各自项目的 `dist/` 目录。

---

## 后端开发指南

### 统一响应格式

所有接口返回统一 JSON 结构：

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

分页响应：

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [],
    "total": 100,
    "page": 1,
    "page_size": 10
  }
}
```

### 路由结构

| 分组 | 前缀 | 中间件 | 说明 |
|------|------|--------|------|
| 博客前台 | `/blog-api/v1/` | 无（部分需 JWT） | 面向访客 |
| 后台管理 | `/admin-api/v1/` | JWT + RBAC | 面向管理员 |

### 如何新增功能模块

以新增"公告（Announcement）"模块为例：

#### 第 1 步：定义 Model

创建 `internal/model/announcement.go`：

```go
package model

type Announcement struct {
    Model
    Title   string `json:"title" gorm:"size:200;not null"`
    Content string `json:"content" gorm:"type:text;not null"`
    Status  int    `json:"status" gorm:"default:1;comment:0-隐藏 1-显示"`
}
```

在 `internal/model/migrate.go` 的 `AutoMigrate` 中注册：

```go
db.AutoMigrate(&Announcement{})
```

#### 第 2 步：编写 DAO

创建 `internal/dao/announcement.go`：

```go
package dao

type AnnouncementDao struct { db *gorm.DB }

func NewAnnouncementDao(db *gorm.DB) *AnnouncementDao {
    return &AnnouncementDao{db: db}
}

func (d *AnnouncementDao) Create(item *model.Announcement) error {
    return d.db.Create(item).Error
}

func (d *AnnouncementDao) List() ([]model.Announcement, error) {
    var items []model.Announcement
    return items, d.db.Order("created_at DESC").Find(&items).Error
}
```

#### 第 3 步：编写 Logic

创建 `internal/logic/blog/announcement_logic.go`：

```go
package blog

type AnnouncementLogic struct { svc *svc.ServiceContext }

func NewAnnouncementLogic(svc *svc.ServiceContext) *AnnouncementLogic {
    return &AnnouncementLogic{svc}
}

func (l *AnnouncementLogic) List() ([]model.Announcement, error) {
    return l.svc.AnnouncementDao.List()
}
```

#### 第 4 步：编写 Handler

创建 `internal/handler/blog_handler.go` 中添加：

```go
func ListAnnouncements(svc *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        logic := blog.NewAnnouncementLogic(svc)
        items, err := logic.List()
        if err != nil {
            fail(w, "获取失败")
            return
        }
        ok(w, items)
    }
}
```

#### 第 5 步：注册路由

在 `internal/routes/routes.go` 中添加：

```go
// 博客路由
blogRoutes := []rest.Route{
    // ... 已有路由
    {Method: http.MethodGet, Path: "/announcements", Handler: handler.ListAnnouncements(svc)},
}
server.AddRoutes(blogRoutes, rest.WithPrefix("/blog-api/v1"))
```

#### 第 6 步：注册依赖注入

在 `internal/svc/service_context.go` 中添加：

```go
type ServiceContext struct {
    // ... 已有 DAO
    AnnouncementDao *dao.AnnouncementDao
}

func NewServiceContext(c config.Config) *ServiceContext {
    return &ServiceContext{
        // ... 已有初始化
        AnnouncementDao: dao.NewAnnouncementDao(db),
    }
}
```

### 中间件说明

| 中间件 | 文件 | 作用 |
|--------|------|------|
| `JWTAuthMiddleware` | `middleware/auth_gozero.go` | JWT Token 解析，写入 Context |
| `RBACAuthMiddleware` | `middleware/auth_gozero.go` | 角色校验（admin/editor） |
| `CORSMiddleware` | `middleware/cors_gozero.go` | 跨域处理 |

---

## 前端开发指南

### 博客前台 (web-blog)

#### 设计特点

- **Shoka 风格**：温暖的粉橙渐变色系
- **Hero 区域**：全屏渐变背景 + 打字动画
- **两栏布局**：主内容区 + 粘性侧边栏
- **文章卡片**：悬停动画 + 斜切效果

#### 主要页面

| 页面 | 文件 | 路径 |
|------|------|------|
| 首页 | `views/Home.vue` | `/` |
| 博客列表 | `views/Blog.vue` | `/blog` |
| 文章详情 | `views/Article.vue` | `/article/:id` |
| 归档 | `views/Archive.vue` | `/archive` |
| 标签 | `views/Tags.vue` | `/tags` |
| 友链 | `views/Friends.vue` | `/friends` |
| 留言 | `views/Messages.vue` | `/messages` |
| 关于 | `views/About.vue` | `/about` |

#### 如何新增页面

1. 在 `src/views/` 创建 Vue 组件
2. 在 `src/router/index.ts` 注册路由
3. 在 `src/components/AppHeader.vue` 添加导航链接（可选）

### 后台管理 (web-admin)

#### 设计特点

- **左侧边栏布局**：可折叠导航菜单
- **CRUD 模式**：列表 + 搜索 + 新增/编辑弹窗
- **主题切换**：亮色/暗色模式

#### 主要页面

| 页面 | 文件 | 功能 |
|------|------|------|
| Dashboard | `views/Dashboard.vue` | 统计仪表盘 |
| 文章管理 | `views/article/ArticleList.vue` | 文章 CRUD |
| 分类管理 | `views/category/CategoryList.vue` | 分类 CRUD |
| 标签管理 | `views/tag/TagList.vue` | 标签 CRUD |
| 用户管理 | `views/user/UserList.vue` | 用户管理 |
| 评论管理 | `views/comment/CommentList.vue` | 评论审核 |
| 友链管理 | `views/friendlink/FriendLinkList.vue` | 友链 CRUD |
| 页面管理 | `views/page/PageList.vue` | 自定义页面 |
| 角色管理 | `views/role/RoleList.vue` | 角色 CRUD |
| 菜单管理 | `views/menu/MenuList.vue` | 菜单 CRUD |
| 站点配置 | `views/site/SiteConfig.vue` | 站点设置 |

#### API 请求封装

`src/api/http.ts` 配置了 Axios 拦截器：

- **请求拦截**：自动附带 JWT Token
- **响应拦截**：统一处理错误提示，401 自动清除登录状态

在 `src/api/index.ts` 中添加新接口：

```ts
// GET 请求
export const getXXX = (params?: any) => http.get('/blog-api/v1/xxx', { params })

// POST 请求
export const createXXX = (data: any) => http.post('/admin-api/v1/xxx', data)
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
| POST | /comments | 发表评论 | JWT |
| POST | /messages | 发表留言 | JWT |

### 后台管理 `/admin-api/v1/`

| 方法 | 路径 | 说明 | 认证 |
|------|------|------|------|
| POST | /login | 管理员登录 | 否 |
| GET/POST/PUT/DELETE | /articles | 文章 CRUD | JWT+RBAC |
| GET/POST/PUT/DELETE | /categories | 分类 CRUD | JWT+RBAC |
| GET/POST/PUT/DELETE | /tags | 标签 CRUD | JWT+RBAC |
| GET/DELETE | /users | 用户管理 | JWT+RBAC |
| GET/PUT/DELETE | /comments | 评论管理 | JWT+RBAC |
| GET/POST/PUT/DELETE | /friend-links | 友链 CRUD | JWT+RBAC |
| GET/DELETE | /messages | 留言管理 | JWT+RBAC |
| GET/POST/PUT/DELETE | /pages | 页面 CRUD | JWT+RBAC |
| GET/PUT | /site/config | 站点配置 | JWT+RBAC |
| GET/POST/PUT/DELETE | /roles | 角色 CRUD | JWT+RBAC |
| GET/POST/PUT/DELETE | /menus | 菜单 CRUD | JWT+RBAC |

---

## 技术栈

| 组件 | 技术 |
|------|------|
| 后端框架 | Go-Zero rest |
| ORM | GORM |
| 数据库 | MySQL |
| 缓存 | Redis |
| 认证 | JWT |
| 日志 | Zap |
| 博客前端 | Vue3 + Naive UI + Pinia + TypeScript + Vite |
| 管理后台 | Vue3 + Element Plus + Pinia + TypeScript + Vite |

---

## 参考项目

- [ve-blog-golang](https://github.com/ve-weiyi/ve-blog-golang) — Go-Zero 微服务架构参考
- [ve-blog-naive](https://github.com/ve-weiyi/ve-blog-naive) — Naive UI 前端设计参考
- [ve-admin-element](https://github.com/ve-weiyi/ve-admin-element) — Element Plus 后台设计参考