package test

import (
	"context"
	"database/sql"
	"fmt"
	"rentit/pkg/repository"
	"rentit/pkg/service"
	rentitGrpc "rentit/pkg/transport/grpc"
	"testing"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/emptypb"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener
/*
func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	plantRepository := repository.NewPlantRepository(dbConn)
	plantService := service.NewPlantService(plantRepository)

	server := rentitGrpc.NewRentitServiceServer(plantService)
	protos.RegisterRentitServiceServer(s, &server)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

 */

func TestGetAllGrpc(t *testing.T) {

	var req *emptypb.Empty
	dbConn, err := sql.Open("postgres", postgresConnection)
	if err != nil {
		t.Error(err)
	}
	defer dbConn.Close()
	plantRepository := repository.NewPlantRepository(dbConn)
	plantService := service.NewPlantService(plantRepository)

	s := rentitGrpc.NewRentitServiceServer(plantService)
	resp, err := s.GetAll(context.Background(), req)

	if err != nil {
		t.Error("shit", err)
	}

	fmt.Println(resp)

}
