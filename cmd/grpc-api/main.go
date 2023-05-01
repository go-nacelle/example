package main

import (
	"context"
	"example/internal"

	"github.com/go-nacelle/grpcbase"
	nacelle "github.com/go-nacelle/nacelle/v2"
)

func setup(ctx context.Context, processes *nacelle.ProcessContainerBuilder, services *nacelle.ServiceContainer) error {
	processes.RegisterInitializer(internal.NewRedisInitializer(), nacelle.WithMetaName("redis"))
	processes.RegisterProcess(grpcbase.NewServer(NewServerInitializer()), nacelle.WithMetaName("grpc-server"))
	return nil
}

func main() {
	nacelle.NewBootstrapper("grpcbase-example", setup).BootAndExit()
}
