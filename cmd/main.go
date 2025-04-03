package main

import (
	"fmt"
	"log"
	"net"

	"github.com/front-go/user/internal/config"
	"github.com/front-go/user/internal/repository"
	"github.com/front-go/user/internal/service"
	"github.com/front-go/user/pkg/user"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.MustLoad()
	db := repository.NewRepository(cfg)
	srv := service.NewService(db)
	grpcSrv := grpc.NewServer()

	user.RegisterUserServiceServer(grpcSrv, srv)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Service.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	err = grpcSrv.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
