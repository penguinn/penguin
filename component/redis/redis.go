package redis

import (
	"github.com/go-redis/redis"
	"errors"
)

var (
	redisTemp *redis.Client
)

type ReidsConfig struct {
	Address        	string
	Password 		string
}

type RedisComponent struct {}

func (RedisComponent) Init(ops ...interface{}) (err error) {
	if len(ops) == 0 {
		return errors.New("param error")
	}
	cfg := ops[0].(*ReidsConfig)
	client := redis.NewClient(&redis.Options{
		Addr:cfg.Address,
		Password:cfg.Password,
	})
	redisTemp = client
	return nil
}

func Get() *redis.Client {
	return redisTemp
}