# Introduction

This service is a GO-lang written order management system (OMS), or more widely known as a Point of Sales service system, that is structured with a microservice architecture, with service to service calls using GRPC.

# Setup instructions

# Tools

- Go 1.22+
- Go workspaces for monorepo
- Golang cosmtrek/air for hot-reloading.
- gRPC for communication between services.
- RabbitMQ as message broker.
- Docker with docker compose.
- MongoDB as storage layer.
- Jaeger for service tracing.
- HashiCorp's Consul for service discovery.
- Stripe for payments.
