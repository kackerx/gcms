package cache

import (
	"context"
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	ctx := context.Background()
	cache := NewCache(nil)

	load, err := cache.Load(ctx, "uid:1", func(ctx context.Context) (any, error) {
		// return db.get("")
		return nil, nil
	})

	fmt.Println(load, err)
}
