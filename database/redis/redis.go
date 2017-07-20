package redis

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/mitchellh/mapstructure"
)

type RedisDB struct {
	Config 			*RedisConfig
	Redis 			*redis.Client
}

type RedisConfig struct {
	Address        	string
	Password 		string
	ConnectTimeout 	time.Duration
	ReadTimeout    	time.Duration
	WriteTimeout   	time.Duration
}

func NewRedis(cfg *RedisConfig) (*RedisDB, error) {

	client := redis.NewClient(&redis.Options{
		Addr:cfg.Address,
		Password:cfg.Password,
		DialTimeout:cfg.ConnectTimeout,
		ReadTimeout:cfg.ReadTimeout,
		WriteTimeout:cfg.WriteTimeout,

	})

	db := &RedisDB{
		Config:cfg,
		Redis:client,
	}

	return db, nil
}

func Creator(cfg interface{}) (interface{}, error) {
	var redisConfig RedisConfig
	err := mapstructure.WeakDecode(cfg, &redisConfig)
	if err != nil {
		return nil, err
	}
	redisConfig.ConnectTimeout = redisConfig.ConnectTimeout * time.Millisecond
	redisConfig.ReadTimeout = redisConfig.ReadTimeout * time.Millisecond
	redisConfig.WriteTimeout = redisConfig.WriteTimeout * time.Millisecond
	return NewRedis(&redisConfig)
}

func (this *RedisDB) Get(write bool) *redis.Client {
	return this.Redis
}