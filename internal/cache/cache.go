package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"gcms/internal/conf"
)

type RedisCache struct {
	cache redis.Cmdable
}

func (c *RedisCache) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
	// TODO implement me
	panic("implement me")
}

func (c *RedisCache) Get(ctx context.Context, key string) (any, error) {
	// TODO implement me
	panic("implement me")
}

func (c *RedisCache) Delete(ctx context.Context, key string) error {
	// TODO implement me
	panic("implement me")
}

func (c *RedisCache) LoadAndDelete(ctx context.Context, key string) (any, error) {
	// TODO implement me
	panic("implement me")
}

func NewCache(c *conf.Redis) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr: c.Addr,
		DB:   c.DB,
		// Password: c.Password
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("redis error: %s", err.Error()))
	}

	return &RedisCache{rdb}
}

func (c *RedisCache) Load(ctx context.Context, key string, loadFn func(ctx context.Context) (any, error)) (any, error) {
	val, err := c.cache.Get(ctx, key).Result()
	if err == nil {
		return val, nil
	}
	if errors.Is(err, redis.Nil) {
		data, err := loadFn(ctx)
		if err != nil {
			return nil, err
		}

		if er := c.cache.Set(ctx, key, data, 10).Err(); er != nil {
			fmt.Println("回写缓存失败")
		}
		return data, nil
	}

	return nil, err
}
