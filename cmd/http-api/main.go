package main

import (
	"example/internal"

	"github.com/go-nacelle/httpbase"
	"github.com/go-nacelle/nacelle"
)

func setup(processes nacelle.ProcessContainer, services nacelle.ServiceContainer) error {
	processes.RegisterInitializer(internal.NewRedisInitializer(), nacelle.WithInitializerName("redis"))
	processes.RegisterProcess(httpbase.NewServer(NewServerInitializer()), nacelle.WithProcessName("http-server"))
	return nil
}

func main() {
	nacelle.NewBootstrapper("httpbase-example", setup).BootAndExit()
}
