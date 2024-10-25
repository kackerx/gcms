package server

import (
	"github.com/gin-gonic/gin"

	"gcms/internal/handler"
	"gcms/internal/middleware"
)

// todo: 函数注入的方式如何使用wire?
// func NewHTTPServer(c *conf.Server, createHandler func(router *gin.Engine, userHandler *handler.UserHandler)) {
// 	r := gin.Default()
//
// 	// 设置中间件
// 	setMiddlewares(r)
//
// 	// 注册路由
// 	createHandler(r)
//
// 	slog.Info("开始HTTPServer: ", c.Addr)
// 	if err := http.ListenAndServe(c.Addr, r); err != nil {
// 		panic(err)
// 	}
// }

const (
	rootPath = "/api"
)

func NewHTTPServer(
	jwt *middleware.JWT,
	userHandler *handler.UserHandler,
) *gin.Engine {
	r := gin.Default()

	// 设置中间件
	setMiddlewares(r)

	// 注册路由
	auth := r.Group(rootPath)
	auth.Use(middleware.StrictAuth(jwt, nil))

	registerUserPath(r, userHandler)

	return r
}

func setMiddlewares(r *gin.Engine) {
	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())
	// todo
}

func registerUserPath(r *gin.Engine, handler *handler.UserHandler) {
	userRoute := r.Group("/user")
	userRoute.GET("haha", handler.GetUser)
	userRoute.POST("/register", handler.Register)
}
