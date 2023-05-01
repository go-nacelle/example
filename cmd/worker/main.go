package main

import (
	"context"
	"example/internal"

	nacelle "github.com/go-nacelle/nacelle/v2"
	"github.com/go-nacelle/workerbase"
)

func setup(ctx context.Context, processes *nacelle.ProcessContainerBuilder, services *nacelle.ServiceContainer) error {
	processes.RegisterInitializer(internal.NewRedisInitializer(), nacelle.WithMetaName("redis"))
	processes.RegisterInitializer(internal.NewPubSubInitializer(), nacelle.WithMetaName("pubsub"))
	ws1 := NewWorkerSpec()
	ws2 := workerbase.NewWorker(ws1)
	processes.RegisterProcess(ws2, nacelle.WithMetaName("worker"))
	return nil
}

func main() {
	nacelle.NewBootstrapper("workerbase-example", setup).BootAndExit()
}
