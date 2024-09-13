package main

import (
	"context"

	"github.com/hnsia/micro-oms/common/api"
)

type OrdersService interface {
	CreateOrder(context.Context) error
	ValidateOrder(context.Context, *api.CreateOrderRequest) error
}

type OrdersStore interface {
	Create(context.Context) error
}
