package main

import (
	"database/sql"
	"net"

	pb "github.com/solethus/shared-proto/proto/order"
	"google.golang.org/grpc"

	"github.com/solethus/order-service/internal/repository"
	"github.com/solethus/order-service/internal/server"
	"github.com/solethus/order-service/internal/service"
	log "github.com/solethus/order-service/pkg/pkg/logger"
)

func main() {
	log.InitLogging()

	// Set up database connection
	db, err := sql.Open("postgres", "postgres://username:password@localhost/orders?sslmode=disable")
	if err != nil {
		log.Logger.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create repository, service, and server
	repo := repository.NewOrderRepository(db)
	svc := service.NewOrderService(repo)
	srv := server.NewOrderServer(svc)

	// Create gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, srv)

	// Start listening
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Logger.Fatalf("Failed to listen: %v", err)
	}

	log.Logger.Info("Server listening on :50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Logger.Fatalf("Failed to serve: %v", err)
	}
}
