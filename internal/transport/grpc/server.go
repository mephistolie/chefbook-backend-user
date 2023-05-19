package grpc

import (
	api "github.com/mephistolie/chefbook-backend-user/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-user/internal/transport/dependencies/service"
)

type UserServer struct {
	api.UnsafeUserServiceServer
	service service.Service
}

func NewServer(service service.Service) *UserServer {
	return &UserServer{service: service}
}
