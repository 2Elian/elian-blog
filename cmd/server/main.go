package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"

	"elian-blog/internal/config"
	"elian-blog/internal/model"
	"elian-blog/internal/routes"
	"elian-blog/internal/svc"
	"elian-blog/pkg/logger"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var configFile = flag.String("f", "configs/config.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	log := logger.New(c.AppLog.Level)
	zap.ReplaceGlobals(log)

	db, err := initDB(&c)
	if err != nil {
		log.Fatal("Failed to connect database", zap.Error(err))
	}

	_ = model.AutoMigrate(db) // 自动建表
	model.SeedMenus(db)

	rdb, err := initRedis(&c)
	if err != nil {
		log.Fatal("Failed to connect redis", zap.Error(err))
	}
	// 依赖注入到容器中
	svcCtx := svc.NewServiceContext(c, db, rdb, log)
	initRoles(svcCtx)

	server := rest.MustNewServer(c.RestConf,
		rest.WithCors(),
		rest.WithNotFoundHandler(http.HandlerFunc(spaHandler)),
	)
	defer server.Stop()

	server.Use(corsMiddleware)

	routes.RegisterHandlers(server, svcCtx) // 注册所以路由

	server.PrintRoutes() // 打印路由信息

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	fmt.Printf("Blog:       http://localhost:%d/\n", c.Port)
	fmt.Printf("Admin:      http://localhost:%d/admin/\n", c.Port)
	fmt.Printf("Blog API:   http://localhost:%d/blog-api/v1/\n", c.Port)
	fmt.Printf("Admin API:  http://localhost:%d/admin-api/v1/\n", c.Port)

	server.Start()
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func initDB(c *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		c.Database.Username,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.DBName,
		c.Database.Charset,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(c.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.Database.MaxOpenConns)

	return db, nil
}

func initRedis(c *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port),
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}

func initRoles(svcCtx *svc.ServiceContext) {
	roles := []model.Role{
		{Name: "管理员", Label: "admin", Description: "系统管理员", Status: 1, Sort: 1},
		{Name: "编辑者", Label: "editor", Description: "内容编辑", Status: 1, Sort: 2},
		{Name: "普通用户", Label: "user", Description: "默认用户角色", Status: 1, Sort: 3},
	}
	for i := range roles {
		if err := svcCtx.RoleDao.CreateIfNotExist(&roles[i]); err != nil {
			fmt.Printf("初始化角色 %s 失败: %v\n", roles[i].Label, err)
		}
	}
}

// unused but kept for potential future use
var _ = filepath.Clean

// spaHandler serves both frontends and uploads with SPA fallback.
// Priority: API routes (handled by router) → uploads → admin → blog
func spaHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	// Uploaded files
	if strings.HasPrefix(path, "/uploads/") {
		w.Header().Set("Cache-Control", "public, max-age=86400")
		http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))).ServeHTTP(w, r)
		return
	}

	// Admin frontend (/admin/*)
	if strings.HasPrefix(path, "/admin") {
		serveSPA(w, r, "./web-admin/dist", path[len("/admin"):])
		return
	}

	// Blog frontend (everything else)
	serveSPA(w, r, "./web-blog/dist", path)
}

// serveSPA tries to serve a static file, falls back to index.html for SPA routing.
func serveSPA(w http.ResponseWriter, r *http.Request, root, relPath string) {
	// Normalize path
	if relPath == "" || relPath == "/" {
		relPath = "/index.html"
	}

	// Try exact file first
	filePath := filepath.Join(root, filepath.Clean(relPath))
	if info, err := os.Stat(filePath); err == nil && !info.IsDir() {
		// Cache hashed assets (Vite generates filenames like xxx-AbCd1234.js)
		if strings.Contains(relPath, "/assets/") {
			w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		}
		http.ServeFile(w, r, filePath)
		return
	}

	// SPA fallback: serve index.html for client-side routing
	indexPath := filepath.Join(root, "index.html")
	if _, err := os.Stat(indexPath); err != nil {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, indexPath)
}
