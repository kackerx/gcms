//go:build wireinject
// +build wireinject

package gms

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
)

var domainSet = wire.NewSet(domain.NewUserDomainService)

var dataSet = wire.NewSet(data.NewData, data.NewUserRepo)

var serviceSet = wire.NewSet(service.NewUserService)

var handlerSet = wire.NewSet(handler.NewUserHandler)

var serverSet = wire.NewSet(server.NewHTTPServer)

var jwtSet = wire.NewSet(middleware.NewJwt)

func wireApp(confData *conf.Data, confJwt *conf.JWT) (*gin.Engine, func(), error) {
	panic(wire.Build(
		jwtSet,
		domainSet,
		dataSet,
		serverSet,
		serviceSet,
		handlerSet,
	))
}
