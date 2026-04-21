# Elian Blog

基于 Go (Gin) + Vue3 的全栈个人博客系统，采用前后端分离架构。

## 项目结构

```
elian-blog/
├── cmd/api/main.go          # 程序入口
├── configs/config.yaml       # 配置文件
├── internal/
│   ├── model/               # 数据模型 (GORM Model)
│   ├── dao/                 # 数据访问层 (数据库操作)
│   ├── service/             # 业务逻辑层
│   ├── controller/          # 控制器层 (HTTP 处理)
│   │   ├── admin/           # 后台管理接口
│   │   └── blog/            # 前台博客接口
│   ├── middleware/          # 中间件 (JWT/RBAC/CORS/限流/日志)
│   ├── router/              # 路由注册
│   └── utils/               # 工具函数 (JWT/密码加密)
├── pkg/
│   ├── config/              # 配置加载、DB/Redis 初始化
│   ├── logger/              # Zap 日志
│   └── response/            # 统一响应格式
└── web/                     # Vue3 前端
    ├── src/
    │   ├── api/             # API 请求封装
    │   ├── components/      # 公共组件
    │   ├── layouts/         # 布局组件
    │   ├── views/           # 页面视图
    │   ├── router/          # 路由配置
    │   ├── store/           # Pinia 状态管理
    │   └── styles/          # 全局样式
    └── vite.config.js       # Vite 配置 (含开发代理)
```

后端采用分层架构，请求流转路径：

```
HTTP Request → Router → Middleware → Controller → Service → DAO → Database
```

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

项目启动时会通过 GORM `AutoMigrate` 自动创建所有表，无需手动建表。

### 2. 修改配置

编辑 `configs/config.yaml`，填入你的数据库和 Redis 连接信息：

```yaml
database:
  host: 127.0.0.1
  port: 3306
  username: root
  password: your_password   # 改为你的 MySQL 密码
  dbname: elian_blog

redis:
  host: 127.0.0.1
  port: 6379
  password: ""

jwt:
  secret: "your-secret-key"  # 生产环境务必修改
  expire_hours: 72
```

### 3. 启动后端

```bash
cd elian-blog
go mod tidy
go run cmd/api/main.go
```

后端默认运行在 `http://localhost:8080`。

### 4. 启动前端

```bash
cd web
npm install
npm run dev
```

前端默认运行在 `http://localhost:3000`，开发模式下 API 请求会自动代理到后端。

### 5. 构建前端

```bash
cd web
npm run build
```

构建产物在 `web/dist/` 目录，可由 Nginx 或 Go 静态文件服务托管。

---

## 后端开发指南

### 项目分层说明

| 层级 | 目录 | 职责 |
|------|------|------|
| Model | `internal/model/` | 定义数据库表结构（GORM 模型） |
| DAO | `internal/dao/` | 封装数据库 CRUD 操作 |
| Service | `internal/service/` | 业务逻辑处理 |
| Controller | `internal/controller/` | 接收 HTTP 请求、参数校验、返回响应 |
| Router | `internal/router/` | 注册路由和依赖注入 |
| Middleware | `internal/middleware/` | JWT 认证、RBAC 鉴权、CORS、限流、操作日志 |

### 统一响应格式

所有接口返回统一 JSON 结构：

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

分页响应的 `data` 结构：

```json
{
  "list": [],
  "total": 100,
  "page": 1,
  "page_size": 10
}
```

`pkg/response/response.go` 提供的响应方法：

- `response.Ok(c, data)` — 成功，code=0
- `response.OkPage(c, list, total, page, pageSize)` — 分页成功
- `response.BadRequest(c, msg)` — 400 参数错误
- `response.Unauthorized(c, msg)` — 401 未认证
- `response.Forbidden(c, msg)` — 403 无权限
- `response.InternalError(c, msg)` — 500 服务器错误
- `response.Fail(c, httpCode, msg)` — 自定义状态码

### 路由结构

| 分组 | 前缀 | 中间件 | 说明 |
|------|------|--------|------|
| 博客前台 | `/blog-api/v1/` | 无（部分需 JWT） | 面向访客 |
| 后台管理 | `/admin-api/v1/` | JWT + RBAC + 操作日志 | 面向管理员 |

### 如何修改已有接口

以"文章列表增加按标题搜索"为例，需要修改以下文件：

**1. DAO 层** — `internal/dao/article.go`

在 `List` 方法中添加搜索条件：

```go
func (d *ArticleDao) List(page, pageSize int, status int, categoryID uint, tagID uint, keyword string) ([]model.Article, int64, error) {
    query := d.db.Model(&model.Article{})
    // ... 已有条件 ...
    if keyword != "" {
        query = query.Where("title LIKE ?", "%"+keyword+"%")
    }
    // ...
}
```

**2. Service 层** — `internal/service/article.go`

在 `ArticleQueryReq` 中添加字段，并传递给 DAO：

```go
type ArticleQueryReq struct {
    // ... 已有字段 ...
    Keyword string `json:"keyword" form:"keyword"`
}

func (s *ArticleService) List(req *ArticleQueryReq) ([]model.Article, int64, error) {
    // ...
    return s.articleDao.List(req.Page, req.PageSize, req.Status, req.CategoryID, req.TagID, req.Keyword)
}
```

**3. Controller 层** — `internal/controller/blog/article.go`

前端传来的 query 参数会通过 `ShouldBindQuery` 自动绑定到 `ArticleQueryReq`，无需额外修改 Controller。

如果需要修改返回字段或逻辑，直接编辑对应 Controller 方法即可。

### 如何新增一个完整功能模块

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
    Sort    int    `json:"sort" gorm:"default:0"`
}

func (Announcement) TableName() string { return "announcement" }
```

> `Model` 是公共基础结构体，已包含 `ID`、`CreatedAt`、`UpdatedAt`、`DeletedAt` 字段。

#### 第 2 步：注册自动迁移

编辑 `internal/model/migrate.go`，在 `AutoMigrate` 函数的 `AutoMigrate` 调用中加入：

```go
err := db.AutoMigrate(
    // ... 已有模型 ...
    &Announcement{},
)
```

#### 第 3 步：编写 DAO

创建 `internal/dao/announcement.go`：

```go
package dao

import (
    "elian-blog/internal/model"
    "gorm.io/gorm"
)

type AnnouncementDao struct {
    db *gorm.DB
}

func NewAnnouncementDao(db *gorm.DB) *AnnouncementDao {
    return &AnnouncementDao{db: db}
}

func (d *AnnouncementDao) Create(item *model.Announcement) error {
    return d.db.Create(item).Error
}

func (d *AnnouncementDao) Update(item *model.Announcement) error {
    return d.db.Save(item).Error
}

func (d *AnnouncementDao) Delete(id uint) error {
    return d.db.Delete(&model.Announcement{}, id).Error
}

func (d *AnnouncementDao) GetByID(id uint) (*model.Announcement, error) {
    var item model.Announcement
    err := d.db.First(&item, id).Error
    return &item, err
}

func (d *AnnouncementDao) List() ([]model.Announcement, error) {
    var items []model.Announcement
    err := d.db.Order("sort ASC, created_at DESC").Find(&items).Error
    return items, err
}
```

#### 第 4 步：编写 Service

创建 `internal/service/announcement.go`：

```go
package service

import (
    "elian-blog/internal/dao"
    "elian-blog/internal/model"
)

type AnnouncementService struct {
    dao *dao.AnnouncementDao
}

func NewAnnouncementService(dao *dao.AnnouncementDao) *AnnouncementService {
    return &AnnouncementService{dao: dao}
}

func (s *AnnouncementService) Create(title, content string, sort int) (*model.Announcement, error) {
    item := &model.Announcement{Title: title, Content: content, Sort: sort}
    return item, s.dao.Create(item)
}

func (s *AnnouncementService) Update(item *model.Announcement) error {
    return s.dao.Update(item)
}

func (s *AnnouncementService) Delete(id uint) error {
    return s.dao.Delete(id)
}

func (s *AnnouncementService) List() ([]model.Announcement, error) {
    return s.dao.List()
}
```

#### 第 5 步：编写 Controller

**前台** — 创建 `internal/controller/blog/announcement.go`：

```go
package blog

import (
    "elian-blog/internal/service"
    "elian-blog/pkg/response"
    "github.com/gin-gonic/gin"
)

type AnnouncementController struct {
    svc *service.AnnouncementService
}

func NewAnnouncementController(svc *service.AnnouncementService) *AnnouncementController {
    return &AnnouncementController{svc: svc}
}

func (ctrl *AnnouncementController) List(c *gin.Context) {
    items, err := ctrl.svc.List()
    if err != nil {
        response.InternalError(c, "获取公告失败")
        return
    }
    response.Ok(c, items)
}
```

**后台** — 创建 `internal/controller/admin/announcement.go`（包含完整 CRUD，可参考已有的 `category.go` 或 `role.go`）。

#### 第 6 步：注册路由

编辑 `internal/router/router.go`，完成依赖注入和路由绑定：

```go
func Setup(db *gorm.DB, rdb *redis.Client, cfg *config.Config, log *zap.Logger) *gin.Engine {
    // ... 已有 DAO 初始化 ...
    announcementDao := dao.NewAnnouncementDao(db)

    // ... 已有 Service 初始化 ...
    announcementSvc := service.NewAnnouncementService(announcementDao)

    // 传入 setupBlogRoutes 和 setupAdminRoutes
    setupBlogRoutes(r, /* ... */ announcementSvc, cfg)
    setupAdminRoutes(r, /* ... */ announcementSvc, opLogDao, cfg)
}
```

在 `setupBlogRoutes` 中添加前台路由：

```go
func setupBlogRoutes(r *gin.Engine, /* ... */ announcementSvc *service.AnnouncementService, cfg *config.Config) {
    // ...
    announcementCtrl := blogctrl.NewAnnouncementController(announcementSvc)
    api.GET("/announcements", announcementCtrl.List)
}
```

在 `setupAdminRoutes` 中添加后台路由：

```go
func setupAdminRoutes(r *gin.Engine, /* ... */ announcementSvc *service.AnnouncementService, /* ... */) {
    // ...
    adminAnnouncementCtrl := admin.NewAnnouncementController(announcementSvc)
    adminGroup.GET("/announcements", adminAnnouncementCtrl.List)
    adminGroup.POST("/announcements", adminAnnouncementCtrl.Create)
    adminGroup.PUT("/announcements/:id", adminAnnouncementCtrl.Update)
    adminGroup.DELETE("/announcements/:id", adminAnnouncementCtrl.Delete)
}
```

至此，后端新增模块完成。启动项目后数据库会自动建表，接口即可使用。

### 中间件说明

| 中间件 | 文件 | 作用 |
|--------|------|------|
| `JWTAuth` | `middleware/auth.go` | 从 `Authorization: Bearer <token>` 解析用户信息，写入 Context |
| `RBACAuth` | `middleware/rbac.go` | 校验用户角色是否为 admin 或 editor |
| `CORS` | `middleware/cors.go` | 处理跨域请求 |
| `RateLimit` | `middleware/ratelimit.go` | IP 级限流 |
| `Logger` | `middleware/cors.go` | 请求日志 |
| `OperationLog` | `middleware/operation_log.go` | 异步记录后台操作日志 |

---

## 前端开发指南

前端基于 Vue3 + Vite + Element Plus + Pinia，API 请求通过 Axios 封装。

### 目录说明

| 目录 | 说明 |
|------|------|
| `src/api/` | HTTP 封装和所有 API 函数 |
| `src/views/` | 页面组件 |
| `src/components/` | 可复用公共组件 |
| `src/layouts/` | 布局组件（Header + Footer） |
| `src/router/` | Vue Router 路由配置 |
| `src/store/` | Pinia 状态管理 |
| `src/styles/` | 全局 SCSS 变量和样式 |

### 开发代理

`vite.config.js` 已配置代理，前端开发时 API 请求自动转发：

```js
proxy: {
    '/blog-api': { target: 'http://localhost:8080' },
    '/admin-api': { target: 'http://localhost:8080' },
    '/uploads': { target: 'http://localhost:8080' }
}
```

### 如何修改已有页面

找到 `src/views/` 下对应文件直接编辑即可。

例如修改首页布局，编辑 `src/views/Home.vue`；修改文章列表样式，编辑 `src/views/Blog.vue` 或 `src/components/ArticleCard.vue`。

### 如何新增页面

以新增"公告页"为例：

#### 1. 添加 API 函数

编辑 `src/api/index.js`：

```js
export const getAnnouncements = () => http.get('/blog-api/v1/announcements')
```

#### 2. 创建页面视图

创建 `src/views/Announcement.vue`：

```vue
<template>
  <div class="announcement-page container">
    <h1>公告</h1>
    <div class="list">
      <div class="item" v-for="item in list" :key="item.id">
        <h3>{{ item.title }}</h3>
        <p>{{ item.content }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getAnnouncements } from '@/api'

const list = ref([])
onMounted(async () => {
  list.value = await getAnnouncements() || []
})
</script>
```

#### 3. 注册路由

编辑 `src/router/index.js`，在 `BlogLayout` 的 `children` 数组中添加：

```js
{ path: 'announcements', name: 'Announcement', component: () => import('@/views/Announcement.vue') }
```

#### 4. 添加导航入口（可选）

编辑 `src/components/Header.vue`，在 `<nav>` 中添加链接：

```html
<router-link to="/announcements">公告</router-link>
```

### 如何对接新的后端接口

在 `src/api/index.js` 中添加对应的函数即可：

```js
// GET 请求
export const getXXX = (params) => http.get('/blog-api/v1/xxx', { params })

// POST 请求
export const createXXX = (data) => http.post('/blog-api/v1/xxx', data)

// PUT 请求
export const updateXXX = (id, data) => http.put(`/blog-api/v1/xxx/${id}`, data)

// DELETE 请求
export const deleteXXX = (id) => http.delete(`/blog-api/v1/xxx/${id}`)
```

`http.js` 中的拦截器会自动处理：
- **请求拦截**：附带 JWT Token
- **响应拦截**：统一处理 `code !== 0` 的错误提示，401 时自动清除登录状态

### 状态管理

`src/store/user.js` — 用户登录状态、Token、用户信息
`src/store/site.js` — 站点配置（名称、描述等）

在组件中使用：

```js
import { useUserStore } from '@/store/user'
const userStore = useUserStore()

userStore.isLoggedIn    // 是否已登录
userStore.userInfo      // 用户信息对象
userStore.setToken(t)   // 保存 Token
userStore.logout()      // 退出登录
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
| GET | /articles/search | 文章搜索 | 否 |
| GET | /categories | 分类列表 | 否 |
| GET | /tags | 标签列表 | 否 |
| GET | /articles/:id/comments | 文章评论 | 否 |
| GET | /comments/recent | 最新评论 | 否 |
| GET | /friend-links | 友链列表 | 否 |
| GET | /messages | 留言列表 | 否 |
| GET | /pages | 页面列表 | 否 |
| GET | /pages/:slug | 页面详情（按 slug） | 否 |
| GET | /site/config | 站点配置 | 否 |
| GET | /user/info | 当前用户信息 | 是 |
| POST | /comments | 发表评论 | 是 |
| POST | /messages | 发表留言 | 是 |

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
| PUT | /roles/:id/menus | 角色菜单分配 | JWT+RBAC |
| GET/POST/PUT/DELETE | /menus | 菜单 CRUD | JWT+RBAC |

---

## 技术栈

**后端：** Go / Gin / GORM / MySQL / Redis / JWT / Zap

**前端：** Vue3 / Vite / Vue Router / Pinia / Element Plus / Axios / SCSS
