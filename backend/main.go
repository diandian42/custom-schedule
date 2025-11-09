package main

import (
	"log"

	appConfig "custom-schedule/internal/config"
	"custom-schedule/internal/database"
	"custom-schedule/internal/repository"
	"custom-schedule/internal/router"
	"custom-schedule/internal/service"
)

func main() {
	// 加载配置
	cfg, err := appConfig.LoadConfig()
	if err != nil {
		log.Fatalf("load config failed: %v", err)
	}

	// 初始化数据库
	db, err := database.InitDatabase(&cfg.Database)
	if err != nil {
		log.Fatalf("init database failed: %v", err)
	}

	// 仓储 & 服务实例化（后续可替换为依赖注入框架）
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	// 初始化路由
	r := router.SetupRouter(userService)

	log.Printf("server starting on %s", cfg.Server.Addr())
	if err := r.Run(cfg.Server.Addr()); err != nil {
		log.Fatalf("server shutdown: %v", err)
	}
}
