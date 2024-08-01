package grpc_client

import (
	"api_gateway/config"
	"api_gateway/genproto/admin_service"
	"api_gateway/genproto/user_service"

	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClientI interface {
	Student() user_service.StudentServiceClient
	Admin() admin_service.AdminServiceClient
}

type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

func New(cfg config.Config) (*GrpcClient, error) {
	connUser, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("user service dial host: %v port:%v err: %v",
			cfg.UserServiceHost, cfg.UserServicePort, err)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"user_service":  user_service.NewStudentServiceClient(connUser),
			"admin_service": admin_service.NewAdminServiceClient(connUser),
		},
	}, nil
}

func (g *GrpcClient) StudentService() user_service.StudentServiceClient {
	return g.connections["user_service"].(user_service.StudentServiceClient)
}

func (g *GrpcClient) AdminService() admin_service.AdminServiceClient {
	return g.connections["admin_service"].(admin_service.AdminServiceClient)
}
