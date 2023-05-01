package internal

import (
	"context"

	"github.com/garyburd/redigo/redis"
	nacelle "github.com/go-nacelle/nacelle/v2"
)

type PubSubInitializer struct {
	Services *nacelle.ServiceContainer `service:"services"`
	Config   *nacelle.Config           `service:"config"`
	conn     redis.PubSubConn
}

func NewPubSubInitializer() nacelle.Initializer {
	return &PubSubInitializer{}
}

func (psi *PubSubInitializer) Init(ctx context.Context) error {
	conn, err := dialFromConfig(psi.Config)
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
