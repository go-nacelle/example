package main

import (
	"context"
	"example/proto"

	"github.com/garyburd/redigo/redis"
	"github.com/go-nacelle/grpcbase"
	nacelle "github.com/go-nacelle/nacelle/v2"
	"google.golang.org/grpc"
)

type ServerInitializer struct {
	Logger nacelle.Logger `service:"logger"`
	Redis  redis.Conn     `service:"redis"`
}

func NewServerInitializer() grpcbase.ServerInitializer {
	return &ServerInitializer{}
}

func (si *ServerInitializer) Init(ctx context.Context, server *grpc.Server) error {
	proto.RegisterRequestServiceServer(server, NewRequestService(si.Logger, si.Redis))
	return nil
}
