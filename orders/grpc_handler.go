package main

import (
	"context"
	"log"

	"github.com/hnsia/micro-oms/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	api.UnimplementedOrderServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server) {
	handler := &grpcHandler{}
	api.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, req *api.CreateOrderRequest) (*api.Order, error) {
	log.Println("New order received!")
	o := &api.Order{
		ID: "42",
	}
	return o, nil
}
