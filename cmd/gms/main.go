package main

import (
	"github.com/gin-gonic/gin"

	"gcms/internal/conf"
	"gcms/pkg/http"
)

func main() {
	cfg, err := conf.New()
	if err != nil {
		return
	}

	app, f, err := wireApp(cfg.Data, cfg.Security.JWT)
	defer f()
	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.DebugMode)
	http.Run(app, cfg.Server.Addr)

	// server.NewHTTPServer(cfg.Server, func(r *gin.Engine, userHandler *handler.UserHandler) {
	// 	router.RegisterUserRoute(r, userHandler)
	// })
}
