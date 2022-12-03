package connection

import (
	"log"
	"os"
	"strconv"

	"github.com/Zhiyenbek/users-main-service/config"
	"github.com/go-redis/redis/v7"
)

func NewRedis(cfg *config.RedisConf) (*redis.Client, error) {
	redis_URL, ok := os.LookupEnv("REDIS_TLS_URL")
	if !ok {
		log.Println("Couldn't get database url. Continuing with config")
		redis_URL = cfg.Host + ":" + strconv.Itoa(cfg.Port)
	} else {
		log.Println("Redis url: ", redis_URL)
	}

	client := redis.NewClient(&redis.Options{
		Addr: redis_URL,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
