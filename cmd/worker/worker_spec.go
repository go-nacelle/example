package main

import (
	"context"
	"io/ioutil"
	"net/http"

	"example/internal"

	"github.com/garyburd/redigo/redis"
	"github.com/go-nacelle/nacelle"
	"github.com/go-nacelle/workerbase"
)

type WorkerSpec struct {
	Logger nacelle.Logger   `service:"logger"`
	Redis  redis.Conn       `service:"redis"`
	PubSub redis.PubSubConn `service:"pubsub"`
}

func NewWorkerSpec() workerbase.WorkerSpec {
	return &WorkerSpec{}
}

func (ws *WorkerSpec) Init(config nacelle.Config) error {
	return nil
}

func (ws *WorkerSpec) Tick(ctx context.Context) error {
outer:
	for {
		switch payload := ws.PubSub.Receive().(type) {
		case redis.Message:
			request, err := internal.ParseRequest(payload.Data)
			if err != nil {
				ws.Logger.Warning("Failed to deserialize payload")
				continue outer
			}

			response, err := internal.SerializeResponse(ws.handleRequest(ctx, request.URL))
			if err != nil {
				return err
			}

			if _, err := ws.Redis.Do("SET", request.ID, response); err != nil {
				return err
			}

		case redis.Subscription:
			if payload.Count == 0 {
				return nil
			}

		case error:
			return payload
		}
	}
}

func (ws *WorkerSpec) handleRequest(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
