package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"rentit/pkg/repository"
	"rentit/pkg/service"
	rentitGrpc "rentit/pkg/transport/grpc"
	"rentit/protos"

	"google.golang.org/grpc"

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
	// add redis as second argument if this is ever used
	plantRepository := repository.NewPlantRepository(dbConn, nil)
	plantService := service.NewPlantService(plantRepository)

	rentitServiceServer := rentitGrpc.NewRentitServiceServer(plantService)
	protos.RegisterRentitServiceServer(grpcServer, rentitServiceServer)
	log.Printf("Serving at port: %v", port)
	grpcServer.Serve(lis)

}