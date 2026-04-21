package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

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

	// 加载配置
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 初始化日志
	log := logger.New(c.Log.Level)
	zap.ReplaceGlobals(log)

	// 初始化数据库
	db, err := initDB(&c)
	if err != nil {
		log.Fatal("Failed to connect database", zap.Error(err))
	}

	// 自动迁移
	_ = model.AutoMigrate(db)

	// 初始化 Redis
	rdb, err := initRedis(&c)
	if err != nil {
		log.Fatal("Failed to connect redis", zap.Error(err))
	}

	// 创建 ServiceContext
	svcCtx := svc.NewServiceContext(c, db, rdb, log)

	// 创建 HTTP 服务器
	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	// 全局 CORS 中间件
	server.Use(corsMiddleware)

	// 注册路由
	routes.RegisterHandlers(server, svcCtx)

	// 打印路由信息
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

	// 测试连接
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
