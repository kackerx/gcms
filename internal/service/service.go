package service

import (
	"gcms/internal/data"
	"gcms/internal/middleware"
	"gcms/pkg/log"
)

type Service struct {
	logger *log.Logger
	jwt    *middleware.JWT
	cache  *data.RedisCache
}

func NewService(jwt *middleware.JWT, cache *data.RedisCache, logger *log.Logger) *Service {
	return &Service{jwt: jwt, cache: cache, logger: logger}
}
