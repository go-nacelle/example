package internal

import (
	"context"

	"github.com/garyburd/redigo/redis"
	"github.com/go-nacelle/nacelle"
)

type RedisInitializer struct {
	Services *nacelle.ServiceContainer `service:"services"`
	conn     redis.Conn
	config   *Config
}

type Config struct {
	RedisAddr string `env:"redis_addr" required:"true"`
}

func NewRedisInitializer() nacelle.Initializer {
	return &RedisInitializer{
		config: &Config{},
	}
}

func (ri *RedisInitializer) RegisterConfiguration(registry nacelle.ConfigurationRegistry) {
	registry.Register(ri.config)
}

func (ri *RedisInitializer) Init(ctx context.Context) error {
	conn, err := redis.DialURL(ri.config.RedisAddr)
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
