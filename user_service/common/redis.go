package common

import (
	"context"
	"fmt"
	"log"
	"time"
	"user_service/utils"

	redis "github.com/redis/go-redis/v9"
)

type RedisTemplate struct {
    RedisClient *redis.Client // Redis 客户端
}


var rdb *redis.Client
var ctx = context.Background()


func InitRedis() *redis.Client {
	config := utils.AppConfig.Redis
	host := config.Host
	port := config.Port
	log.Println("redis host:", host)
	password := config.Password
	db := config.DB

	rdb = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port, // Redis 地址
		Password: password,          // Redis 密码，如果没有密码可以留空
		DB:       db,                // 使用默认数据库
	})

	// 测试连接
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis successfully")

	return rdb
}

func GetRedisClient() *redis.Client {
	if rdb == nil {
		utils.LoadConfig()
		rdb = InitRedis()
	}
	return rdb
}

func (r *RedisTemplate) GetValue(key string) (string, error) {
	if r.RedisClient == nil {
		return "", fmt.Errorf("redis client is not initialized")
	}

	return r.RedisClient.Get(ctx, key).Result()
}

func (r *RedisTemplate) SetValue(key string, value string, ex time.Duration) error {
	if r.RedisClient == nil {
		return fmt.Errorf("redis client is not initialized")
	}
	
	return r.RedisClient.Set(ctx, key, value, ex).Err()
}