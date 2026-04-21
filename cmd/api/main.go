package main

import (
	"elian-blog/internal/router"
	"elian-blog/pkg/config"
	"elian-blog/pkg/logger"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	log := logger.New(cfg.Log.Level)
	zap.ReplaceGlobals(log)

	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect database", zap.Error(err))
	}

	rdb, err := config.InitRedis(cfg)
	if err != nil {
		log.Fatal("Failed to connect redis", zap.Error(err))
	}

	engine := router.Setup(db, rdb, cfg, log)

	go func() {
		addr := fmt.Sprintf(":%d", cfg.Server.Port)
		log.Info("Server starting", zap.String("addr", addr))
		if err := engine.Run(addr); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")
}
