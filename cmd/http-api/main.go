package main

import (
	"context"
	"example/internal"

	"github.com/go-nacelle/httpbase"
	nacelle "github.com/go-nacelle/nacelle/v2"
)

func setup(ctx context.Context, processes *nacelle.ProcessContainerBuilder, services *nacelle.ServiceContainer) error {
	processes.RegisterInitializer(internal.NewRedisInitializer(), nacelle.WithMetaName("redis"))
	processes.RegisterProcess(httpbase.NewServer(NewServerInitializer()), nacelle.WithMetaName("http-server"))
	return nil
}

func main() {
	nacelle.NewBootstrapper("httpbase-example", setup).BootAndExit()
}
