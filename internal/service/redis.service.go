package service

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type IRedisService interface {
	Set(key string, value string, expiration int) bool
	Get(key string) (string, bool)
	Delete(key string) bool
}

type redisService struct {
	redisClient *redis.Client
}

func NewRedisService(redisClient *redis.Client) IRedisService {
	return &redisService{
		redisClient: redisClient,
	}
}

func (r *redisService) Set(key string, value string, expiration int) bool {
	ctx := context.Background()

	err := r.redisClient.Set(
		ctx,
		key,
		value,
		time.Duration(expiration)*time.Second,
	).Err()

	return err == nil
}

func (r *redisService) Get(key string) (string, bool) {
	ctx := context.Background()

	val, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		return "", false
	}

	return val, true
}

func (r *redisService) Delete(key string) bool {
	ctx := context.Background()

	err := r.redisClient.Del(ctx, key).Err()
	return err == nil
}
