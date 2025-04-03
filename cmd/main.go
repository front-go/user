package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
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

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "user.create",
		GroupID: "123",
	})

	go func() {
		for {
			msg, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Fatalf("failed to read message: %v", err)
			}
			fmt.Printf("massage: %s", string(msg.Value))
		}
	}()

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
