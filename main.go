package main

import (
	"database/sql"
	"github.com/solethus/order-service/internal/client"
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

	// Set up connection to Inventory Service
	inventoryConn, err := grpc.Dial("inventory-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Logger.Fatalf("Failed to connect to Inventory Service: %v", err)
	}
	defer inventoryConn.Close()
	inventoryClient := client.NewInventoryClient(inventoryConn)

	// Set up connection to Dealership Service
	dealershipConn, err := grpc.Dial("dealership-service:50052", grpc.WithInsecure())
	if err != nil {
		log.Logger.Fatalf("Failed to connect to Dealership Service: %v", err)
	}
	defer dealershipConn.Close()
	dealershipClient := client.NewDealershipClient(dealershipConn)

	// Create repository, service, and server
	repo := repository.NewOrderRepository(db)
	svc := service.NewOrderService(repo, inventoryClient, dealershipClient)
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
