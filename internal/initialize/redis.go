package initialize

import (
	"context"
	"fmt"

	"github.com/haidaqn/go-ecommerce-backend-api/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	redisSetting := global.Config.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisSetting.Host, redisSetting.Port),
		Password: redisSetting.Password, // no password set
		DB:       redisSetting.Db,       // use default DB
		PoolSize: 10,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Failed to connect to Redis", zap.Error(err))
		panic(err)
	}

	global.Redis = rdb
	global.Logger.Info("Connected to Redis", zap.String("host", redisSetting.Host), zap.Int("port", redisSetting.Port))
	// RedisExample()
}

func RedisExample() {
	// Example of setting a key in Redis
	err := global.Redis.Set(ctx, "example_key", "example_value11111111", 0).Err()
	if err != nil {
		global.Logger.Error("Failed to set key in Redis", zap.Error(err))
		return
	}

	// Example of getting a key from Redis
	val, err := global.Redis.Get(ctx, "example_key").Result()
	if err != nil {
		global.Logger.Error("Failed to get key from Redis", zap.Error(err))
		return
	}
	global.Logger.Info("Successfully got key from Redis", zap.String("key", "example_key"), zap.String("value", val))

}
