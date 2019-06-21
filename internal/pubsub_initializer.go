package internal

import (
	"github.com/garyburd/redigo/redis"
	"github.com/go-nacelle/nacelle"
)

type PubSubInitializer struct {
	Services nacelle.ServiceContainer `service:"services"`
	conn     redis.PubSubConn
}

func NewPubSubInitializer() nacelle.Initializer {
	return &PubSubInitializer{}
}

func (psi *PubSubInitializer) Init(config nacelle.Config) error {
	conn, err := dialFromConfig(config)
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
