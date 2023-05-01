package internal

import (
	"context"

	"github.com/garyburd/redigo/redis"
	nacelle "github.com/go-nacelle/nacelle/v2"
)

type RedisInitializer struct {
	Services *nacelle.ServiceContainer `service:"services"`
	Config   *nacelle.Config           `service:"config"`
	conn     redis.Conn
}

type Config struct {
	RedisAddr string `env:"redis_addr" required:"true"`
}

func NewRedisInitializer() nacelle.Initializer {
	return &RedisInitializer{}
}

func (ri *RedisInitializer) Init(ctx context.Context) error {
	conn, err := dialFromConfig(ri.Config)
	if err != nil {
		return err
	}

	ri.conn = conn
	return ri.Services.Set("redis", conn)
}

func (ri *RedisInitializer) Finalize() error {
	ri.conn.Close()
	return nil
}

func dialFromConfig(config *nacelle.Config) (redis.Conn, error) {
	redisConfig := &Config{}
	if err := config.Load(redisConfig); err != nil {
		return nil, err
	}

	conn, err := redis.DialURL(redisConfig.RedisAddr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
