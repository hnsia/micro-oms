package main

import (
	"context"
	"log"

	"github.com/hnsia/micro-oms/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	api.UnimplementedOrderServiceServer

	service OrdersService
}

func NewGRPCHandler(grpcServer *grpc.Server, service OrdersService) {
	handler := &grpcHandler{
		service: service,
	}
	api.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, req *api.CreateOrderRequest) (*api.Order, error) {
	log.Printf("New order received! Order %v", req)
	o := &api.Order{
		ID: "42",
	}
	return o, nil
}
