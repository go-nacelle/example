package main

import (
	"example/internal"

	"github.com/go-nacelle/config"
	"github.com/go-nacelle/httpbase"
	"github.com/go-nacelle/nacelle"
	"github.com/go-nacelle/process"
)

func setup(processes *nacelle.ProcessContainerBuilder, services *nacelle.ServiceContainer) error {
	// TODO - rename
	processes.RegisterInitializer(internal.NewRedisInitializer(), process.WithMetaName("redis"))
	processes.RegisterProcess(httpbase.NewServer(NewServerInitializer(), httpbase.WithTagModifiers(config.NewEnvTagPrefixer("example"))), process.WithMetaName("http-server"))
	return nil
}

func main() {
	nacelle.NewBootstrapper("httpbase-example", setup).BootAndExit()
}
