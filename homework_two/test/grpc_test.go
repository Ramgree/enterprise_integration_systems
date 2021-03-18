package test

import (
	"context"
	"database/sql"
	"fmt"
	"rentit/pkg/repository"
	"rentit/pkg/service"
	rentitGrpc "rentit/pkg/transport/grpc"
	"testing"
	"google.golang.org/protobuf/types/known/emptypb"
)

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
