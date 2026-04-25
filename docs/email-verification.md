# 邮箱验证功能 - 技术方案

## 一、需求背景

- 注册时验证邮箱真实性
- 后续支持找回密码功能
- 重要操作二次确认

## 二、数据库变更

### User 表新增字段

| 字段 | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| `email_verified` | `tinyint(1)` | `0` | 邮箱是否已验证 |

```sql
ALTER TABLE `users` ADD COLUMN `email_verified` tinyint(1) NOT NULL DEFAULT 0 COMMENT '邮箱是否验证';
```

GORM 模型变更 (`internal/model/user.go`)：

```go
type User struct {
    // ... 已有字段 ...
    EmailVerified bool `json:"email_verified" gorm:"default:0;comment:邮箱是否验证"`
}
```

## 三、SMTP 配置

### 3.1 配置文件 (`configs/config.yaml`)

```yaml
Email:
  Host: smtp.example.com       # SMTP 服务器地址
  Port: 587                    # SMTP 端口 (TLS 一般 587, SSL 一般 465)
  Username: your@email.com     # 发件邮箱账号
  Password: your_auth_code     # 邮箱授权码（非登录密码）
  From: your@email.com         # 显示的发件人地址
  FromName: Elian Blog         # 发件人名称
  UseSSL: false                # 是否使用 SSL
```

### 3.2 配置结构体 (`internal/config/config.go`)

```go
type EmailConfig struct {
    Host     string `json:"Host"`
    Port     int    `json:"Port"`
    Username string `json:"Username"`
    Password string `json:"Password"`
    From     string `json:"From"`
    FromName string `json:"FromName"`
    UseSSL   bool   `json:"UseSSL"`
}
```

在 `Config` 结构体中添加：

```go
type Config struct {
    // ... 已有字段 ...
    Email EmailConfig `json:"Email"`
}
```

## 四、邮件发送工具

新建文件 `internal/utils/email.go`：

```go
package utils

import (
    "fmt"
    "net/smtp"
    "strings"

    "elian-blog/internal/config"
)

type EmailSender struct {
    config config.EmailConfig
}

func NewEmailSender(c config.EmailConfig) *EmailSender {
    return &EmailSender{config: c}
}

// SendVerificationCode 发送验证码邮件
func (s *EmailSender) SendVerificationCode(toEmail string, code string) error {
    subject := "Elian Blog - 邮箱验证码"
    body := fmt.Sprintf(`
        <div style="padding: 20px; background: #f5f7fa; border-radius: 8px;">
            <h2 style="color: #303133;">邮箱验证</h2>
            <p style="color: #606266; font-size: 14px;">您的验证码为：</p>
            <div style="font-size: 32px; font-weight: bold; color: #409eff; letter-spacing: 8px; padding: 16px 0;">%s</div>
            <p style="color: #909399; font-size: 12px;">验证码有效期为 5 分钟，请勿泄露给他人。</p>
        </div>
    `, code)

    return s.sendHTML(toEmail, subject, body)
}

func (s *EmailSender) sendHTML(to string, subject string, body string) error {
    from := s.config.From
    headers := make(map[string]string)
    headers["From"] = fmt.Sprintf("%s <%s>", s.config.FromName, from)
    headers["To"] = to
    headers["Subject"] = subject
    headers["MIME-Version"] = "1.0"
    headers["Content-Type"] = "text/html; charset=UTF-8"

    msg := ""
    for k, v := range headers {
        msg += fmt.Sprintf("%s: %s\r\n", k, v)
    }
    msg += "\r\n" + body

    addr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
    auth := smtp.PlainAuth("", s.config.Username, s.config.Password, s.config.Host)

    return smtp.SendMail(addr, auth, from, []string{to}, []byte(msg))
}
```

**更优方案**：使用第三方库 `gopkg.in/gomail.v2`，支持 SSL 和更灵活的配置。

## 五、验证码逻辑

### 5.1 验证码存储（Redis）

```
Key:    email_verify:{email}
Value:  验证码（6位数字）
TTL:    300 秒（5 分钟）
```

冷却机制（防止频繁发送）：

```
Key:    email_cooldown:{email}
Value:  1
TTL:    60 秒（1 分钟内不能重复发送）
```

### 5.2 验证码生成与发送

新建文件 `internal/logic/blog/email_logic.go`：

```go
package blog

import (
    "context"
    "crypto/rand"
    "fmt"
    "time"

    "elian-blog/internal/svc"
    "elian-blog/internal/types"

    "github.com/zeromicro/go-zero/core/logx"
)

type EmailLogic struct {
    svcCtx *svc.ServiceContext
}

func NewEmailLogic(svcCtx *svc.ServiceContext) *EmailLogic {
    return &EmailLogic{svcCtx: svcCtx}
}

// SendVerifyCode 发送验证码
func (l *EmailLogic) SendVerifyCode(ctx context.Context, req *types.SendVerifyCodeReq) error {
    // 检查冷却
    cooldownKey := fmt.Sprintf("email_cooldown:%s", req.Email)
    exists, _ := l.svcCtx.Redis.Exists(ctx, cooldownKey).Result()
    if exists == 1 {
        return fmt.Errorf("发送太频繁，请稍后再试")
    }

    // 生成 6 位验证码
    code := generateCode(6)

    // 存储到 Redis（5分钟有效）
    verifyKey := fmt.Sprintf("email_verify:%s", req.Email)
    l.svcCtx.Redis.Set(ctx, verifyKey, code, 5*time.Minute)

    // 设置冷却
    l.svcCtx.Redis.Set(ctx, cooldownKey, "1", time.Minute)

    // 发送邮件
    return l.svcCtx.EmailSender.SendVerificationCode(req.Email, code)
}

// VerifyEmail 验证邮箱
func (l *EmailLogic) VerifyEmail(ctx context.Context, req *types.VerifyEmailReq) error {
    verifyKey := fmt.Sprintf("email_verify:%s", req.Email)
    storedCode, err := l.svcCtx.Redis.Get(ctx, verifyKey).Result()
    if err != nil {
        return fmt.Errorf("验证码已过期，请重新发送")
    }

    if storedCode != req.Code {
        return fmt.Errorf("验证码错误")
    }

    // 更新用户邮箱验证状态
    userID, _ := ctx.Value("user_id").(uint)
    if err := l.svcCtx.UserDao.UpdateEmailVerified(userID, true); err != nil {
        return fmt.Errorf("更新验证状态失败")
    }

    // 删除验证码
    l.svcCtx.Redis.Del(ctx, verifyKey)

    return nil
}

func generateCode(length int) string {
    b := make([]byte, length)
    rand.Read(b)
    for i := range b {
        b[i] = '0' + (b[i] % 10)
    }
    return string(b)
}
```

## 六、API 接口设计

### 6.1 发送验证码

```
POST /blog-api/v1/email/send
```

请求体：
```json
{
    "email": "user@example.com"
}
```

响应：
```json
{
    "code": 0,
    "message": "验证码已发送"
}
```

### 6.2 验证邮箱

```
POST /blog-api/v1/email/verify
```

请求体：
```json
{
    "email": "user@example.com",
    "code": "123456"
}
```

响应：
```json
{
    "code": 0,
    "message": "验证成功"
}
```

### 6.3 类型定义 (`internal/types/email.go`)

```go
package types

type SendVerifyCodeReq struct {
    Email string `json:"email" validate:"required,email"`
}

type VerifyEmailReq struct {
    Email string `json:"email" validate:"required,email"`
    Code  string `json:"code" validate:"required,len=6"`
}
```

## 七、ServiceContext 变更

`internal/svc/service_context.go` 新增：

```go
type ServiceContext struct {
    // ... 已有字段 ...
    EmailSender *utils.EmailSender
}

func NewServiceContext(c config.Config, db *gorm.DB, rdb *redis.Client, log *zap.Logger) *ServiceContext {
    return &ServiceContext{
        // ... 已有字段 ...
        EmailSender: utils.NewEmailSender(c.Email),
    }
}
```

## 八、路由注册

在 `internal/routes/` 中注册新路由：

```go
// 邮箱验证（需要登录）
blogGroup := server.Group("/blog-api/v1", jwtMiddleware)
{
    // ... 已有路由 ...
    blogGroup.Post("/email/send", handler.SendVerifyCodeHandler(svcCtx))
    blogGroup.Post("/email/verify", handler.VerifyEmailHandler(svcCtx))
}
```

## 九、DAO 层新增

`internal/dao/user.go` 新增方法：

```go
func (d *UserDao) UpdateEmailVerified(userID uint, verified bool) error {
    return d.db.Model(&model.User{}).Where("id = ?", userID).
        Update("email_verified", verified).Error
}
```

## 十、实施步骤

1. 数据库添加 `email_verified` 字段 + GORM 模型更新
2. 配置文件和配置结构体添加 Email 配置
3. 实现 `internal/utils/email.go` 邮件发送工具
4. 实现 `internal/logic/blog/email_logic.go` 验证码逻辑
5. 添加 `internal/types/email.go` 类型定义
6. ServiceContext 注入 EmailSender
7. 注册路由和 Handler
8. DAO 层添加 `UpdateEmailVerified` 方法
9. 前台添加发送验证码和验证的 UI 组件

## 十一、注意事项

- 邮箱授权码不是登录密码，需要到邮箱服务商后台获取
- QQ 邮箱：设置 → 账户 → POP3/SMTP 服务 → 生成授权码
- 163 邮箱：设置 → POP3/SMTP/IMAP → 开启并获取授权码
- Gmail：需要使用应用专用密码
- 验证码长度建议 6 位，有效期 5 分钟
- 发送间隔限制 60 秒，防止滥用
- 生产环境建议使用邮件服务（SendGrid、阿里邮件推送等）而非直接 SMTP
