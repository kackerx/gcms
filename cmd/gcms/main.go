package main

import (
	"github.com/gin-gonic/gin"

	"gcms/internal/conf"
	"gcms/pkg/http"
	"gcms/pkg/log"
)

func main() {
	cfg, err := conf.New()
	if err != nil {
		return
	}

	logger := log.NewLog(cfg)

	app, f, err := wireApp(cfg.Data, cfg.Security.JWT, logger)
	defer f()
	if err != nil {
		panic(err)
	}

	// 初始化缓存, 还是注入吧, 全局变量使用感觉不太好
	// cache.InitCache(cfg.Data)

	gin.SetMode(gin.DebugMode)
	http.Run(app, cfg.Server.Addr)

	// server.NewHTTPServer(cfg.Server, func(r *gin.Engine, userHandler *handler.UserHandler) {
	// 	router.RegisterUserRoute(r, userHandler)
	// })
}
