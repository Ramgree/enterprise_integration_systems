package main

import (
	"database/sql"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"rentit/pkg/repository"
	"rentit/pkg/service"
	rentitGrpc "rentit/pkg/transport/grpc"
	"rentit/protos"

	_ "github.com/lib/pq"
)

const (
	port = 10001
	postgresConnection = "dbname=postgres host=localhost password=postgres user=postgres sslmode=disable port=5432"
)

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	dbConn, err := sql.Open("postgres", postgresConnection)
	if err != nil {
		log.Fatalf("Could not connect to postgres: %v", err)
	}
	plantRepository := repository.NewPlantRepository(dbConn)
	plantService := service.NewPlantService(plantRepository)

	rentitServiceServer := rentitGrpc.NewRentitServiceServer(plantService)
	protos.RegisterRentitServiceServer(grpcServer, rentitServiceServer)
	log.Printf("Serving at port: %v", port)
	grpcServer.Serve(lis)

}