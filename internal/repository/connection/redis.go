package connection

import (
	"strconv"

	"github.com/Zhiyenbek/users-main-service/config"
	"github.com/go-redis/redis/v7"
)

func NewRedis(cfg *config.RedisConf) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Host + ":" + strconv.Itoa(cfg.Port),
		DB:   cfg.DB,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
