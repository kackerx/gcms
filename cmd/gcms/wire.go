//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"gcms/internal/conf"
	"gcms/internal/data"
	"gcms/internal/domain"
	"gcms/internal/handler"
	"gcms/internal/middleware"
	"gcms/internal/server"
	"gcms/internal/service"
	"gcms/pkg/log"
)

var domainSet = wire.NewSet(domain.NewUserDomainService)

var dataSet = wire.NewSet(data.NewData, data.NewUserRepo, data.NewCache)

var serviceSet = wire.NewSet(service.NewService, service.NewUserService)

var handlerSet = wire.NewSet(handler.NewHandler, handler.NewUserHandler)

var serverSet = wire.NewSet(server.NewHTTPServer)

var jwtSet = wire.NewSet(middleware.NewJwt)

var logSet = wire.NewSet(log.NewLog)

func wireApp(confData *conf.Data, confJwt *conf.JWT, logger *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		jwtSet,
		domainSet,
		dataSet,
		serverSet,
		serviceSet,
		handlerSet,
	))
}
