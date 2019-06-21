package main

import (
	"example/internal"
	"github.com/go-nacelle/nacelle"
	"github.com/go-nacelle/workerbase"
)

func setup(processes nacelle.ProcessContainer, services nacelle.ServiceContainer) error {
	processes.RegisterInitializer(internal.NewRedisInitializer(), nacelle.WithInitializerName("redis"))
	processes.RegisterInitializer(internal.NewPubSubInitializer(), nacelle.WithInitializerName("pubsub"))
	processes.RegisterProcess(workerbase.NewWorker(NewWorkerSpec()), nacelle.WithProcessName("worker"))
	return nil
}

func main() {
	nacelle.NewBootstrapper("workerbase-example", setup).BootAndExit()
}
