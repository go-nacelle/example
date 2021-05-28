package internal

import (
	"context"

	"github.com/garyburd/redigo/redis"
	"github.com/go-nacelle/nacelle"
)

type PubSubInitializer struct {
	Services nacelle.ServiceContainer `service:"services"`
	conn     redis.PubSubConn
	config   *Config
}

func NewPubSubInitializer() nacelle.Initializer {
	return &PubSubInitializer{
		config: &Config{},
	}
}

func (psi *PubSubInitializer) RegisterConfiguration(registry nacelle.ConfigurationRegistry) {
	registry.Register(psi.config)
}

func (psi *PubSubInitializer) Init(ctx context.Context) error {
	conn, err := redis.DialURL(psi.config.RedisAddr)
	if err != nil {
		return err
	}

	pubsub := redis.PubSubConn{
		Conn: conn,
	}

	if err := pubsub.Subscribe("work"); err != nil {
		return err
	}

	psi.conn = pubsub
	return psi.Services.Set("pubsub", pubsub)
}

func (psi *PubSubInitializer) Finalize() error {
	psi.conn.Unsubscribe()
	psi.conn.Close()
	return nil
}
