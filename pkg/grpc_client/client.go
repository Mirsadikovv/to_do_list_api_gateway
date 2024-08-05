package grpc_client

import (
	"api_gateway/config"
	"api_gateway/genproto/admin_service"
	"api_gateway/genproto/task_service"
	"api_gateway/genproto/user_service"

	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClientI interface {
	User() user_service.UserServiceClient
	Admin() admin_service.AdminServiceClient
	Task() task_service.TaskServiceClient
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

	connTask, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.TaskServiceHost, cfg.TaskServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("user service dial host: %v port:%v err: %v",
			cfg.TaskServiceHost, cfg.TaskServicePort, err)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"user_service":  user_service.NewUserServiceClient(connUser),
			"admin_service": admin_service.NewAdminServiceClient(connUser),
			"task_service":  task_service.NewTaskServiceClient(connTask),
		},
	}, nil
}

func (g *GrpcClient) UserService() user_service.UserServiceClient {
	return g.connections["user_service"].(user_service.UserServiceClient)
}

func (g *GrpcClient) AdminService() admin_service.AdminServiceClient {
	return g.connections["admin_service"].(admin_service.AdminServiceClient)
}

func (g *GrpcClient) TaskService() task_service.TaskServiceClient {
	return g.connections["task_service"].(task_service.TaskServiceClient)
}
