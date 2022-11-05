package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Configs struct {
	App   *AppConfig
	DB    *DBConf
	Redis *RedisConf
}

type AppConfig struct {
	TimeOut time.Duration
	Port    int
}

type DBConf struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	SSLMode  string
	TimeOut  time.Duration
}

type RedisConf struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       int
}

func New() (*Configs, error) {
	vi := viper.New()
	vi.SetConfigType("yaml")
	vi.SetConfigName("configs")
	vi.AddConfigPath("config/")

	//databse default values
	vi.SetDefault("db.port", 5432)
	vi.SetDefault("db.ssl_mode", "disable")
	vi.SetDefault("db.user", "postgres")
	vi.SetDefault("db.password", "postgres")
	vi.SetDefault("db.timeout", 5)

	//app default values
	vi.SetDefault("app.port", 8080)
	vi.SetDefault("app.timeout", 60)

	//redis default values
	vi.SetDefault("redis.port", 6370)
	err := vi.ReadInConfig()

	if err != nil {
		return nil, fmt.Errorf("error while parsing config. %v", err)
	}
	dbHost := vi.GetString("db.host")
	if dbHost == "" {
		return nil, fmt.Errorf("error while parsing config. Database host not defined")
	}
	dbName := vi.GetString("db.db_name")
	if dbName == "" {
		return nil, fmt.Errorf("error while parsing config. Database name not defined")
	}

	redisHost := vi.GetString("redis.host")
	if redisHost == "" {
		return nil, fmt.Errorf("error while parsing config. Redis host not defined")
	}

	return &Configs{
		App: &AppConfig{
			TimeOut: time.Second * time.Duration(vi.GetInt("app.timeout")),
			Port:    vi.GetInt("app.port"),
		},
		DB: &DBConf{
			Host:     dbHost,
			Port:     vi.GetInt("db.port"),
			Username: vi.GetString("db.user"),
			Password: vi.GetString("db.password"),
			SSLMode:  vi.GetString("db.ssl_mode"),
			TimeOut:  time.Second * time.Duration(vi.GetInt("db.timeout")),
		},
		Redis: &RedisConf{
			Host: redisHost,
			DB:   vi.GetInt("redis.db"),
			Port: vi.GetInt("redis.port"),
		},
	}, nil
}
