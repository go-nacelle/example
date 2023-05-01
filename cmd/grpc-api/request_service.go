package main

import (
	"context"
	"example/internal"
	"example/proto"

	"github.com/garyburd/redigo/redis"
	nacelle "github.com/go-nacelle/nacelle/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RequestService struct {
	logger nacelle.Logger
	redis  redis.Conn
}

func NewRequestService(logger nacelle.Logger, redis redis.Conn) *RequestService {
	return &RequestService{
		logger: logger,
		redis:  redis,
	}
}

func (rs *RequestService) Get(ctx context.Context, r *proto.GetRequest) (*proto.GetResponse, error) {
	response, err := internal.GetResult(rs.redis, r.GetId())
	if err != nil {
		return nil, err
	}

	if response == nil {
		return nil, status.Error(codes.NotFound, "request was not performed")
	}

	return &proto.GetResponse{
		Body:  response.Body,
		Error: response.Error,
	}, nil
}

func (rs *RequestService) Queue(ctx context.Context, r *proto.QueueRequest) (*proto.QueueResponse, error) {
	id, err := internal.PublishWork(rs.redis, r.GetUrl())
	if err != nil {
		return nil, err
	}

	return &proto.QueueResponse{Id: id}, nil
}
