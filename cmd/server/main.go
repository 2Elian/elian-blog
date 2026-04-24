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

	_ = model.AutoMigrate(db)
	model.SeedMenus(db)

	rdb, err := initRedis(&c)
	if err != nil {
		log.Fatal("Failed to connect redis", zap.Error(err))
	}

	svcCtx := svc.NewServiceContext(c, db, rdb, log)
	initRoles(svcCtx)

	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	server.Use(corsMiddleware)

	routes.RegisterHandlers(server, svcCtx)

	// Register uploads file server - match /uploads/:subdir/:year/:month/:file
	// This covers the pattern: /uploads/misc/2026/04/file.jpg
	uploadsHandler := http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads")))
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/uploads/:a/:b/:c/:d",
		Handler: uploadsHandler.ServeHTTP,
	})

	server.PrintRoutes()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	fmt.Printf("Blog API: http://localhost:%d/blog-api/v1/\n", c.Port)
	fmt.Printf("Admin API: http://localhost:%d/admin-api/v1/\n", c.Port)

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
var _ = os.Stat
var _ = strings.HasPrefix
