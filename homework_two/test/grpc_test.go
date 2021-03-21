package test

import (
	"context"
	"database/sql"
	"log"
	"net"
	"rentit/pkg/repository"
	"rentit/pkg/service"
	rentitGrpc "rentit/pkg/transport/grpc"
	"rentit/protos"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GRPC is awesome in the sense that it provides a mock server to test, without requiring the real server to be running
const (
	redisURI        = "localhost:6379"
	redisPassword   = "" // no password set
	redisDB         = 0  // use default DB
	bufSize = 1024 * 1024
)

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	dbConn, err := sql.Open("postgres", postgresConnection)
	if err != nil {
		log.Fatalf("Could not connect to postgres: %v", err)
	}

	redisConn := redis.NewClient(&redis.Options{
		Addr:     redisURI,
		Password: redisPassword,
		DB:       redisDB,
	})
	
	plantRepository := repository.NewPlantRepository(dbConn, redisConn)
	plantService := service.NewPlantService(plantRepository)

	rentitServiceServer := rentitGrpc.NewRentitServiceServer(plantService)
	protos.RegisterRentitServiceServer(s, rentitServiceServer)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}

	}()

}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGetAllGrpc(t *testing.T) {

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Errorf("fail to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := protos.NewRentitServiceClient(conn)
	resp, err := client.GetAllPlants(ctx, &emptypb.Empty{})
	if err != nil {
		t.Errorf("GetAll failed: %v", err)
	}
	if len(resp.Plants) != 8 {
		t.Errorf("GetAll got the wrong number of plants: %v", resp)
	}

}

func TestEstimateRentalGrpc(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Errorf("fail to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := protos.NewRentitServiceClient(conn)
	name := "excavator"
	var start_date time.Time
	start_date, _ = time.Parse(layout, "2020-01-01 00:00:00")
	start_timestamp, _ := ptypes.TimestampProto(start_date)
	var end_date time.Time
	end_date, _ = time.Parse(layout, "2020-01-03 00:00:00")
	end_timestamp, _ := ptypes.TimestampProto(end_date)

	resp, err := client.EstimateRental(ctx, &protos.EstimateRentalRequest{
		Name:      name,
		StartDate: start_timestamp,
		EndDate:   end_timestamp,
	})

	if err != nil {
		t.Errorf("Failed to estimate rental %v", err)
	}

	if resp.PriceEstimation != 2500 {
		t.Errorf("Rental estimation calculation is wrong %v", resp)
	}

	log.Printf("Resp: %v", resp)

}

func TestAvailabilityCheckGrpc(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Errorf("fail to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := protos.NewRentitServiceClient(conn)
	name := "excavator"
	var start_date time.Time
	start_date, _ = time.Parse(layout, "2020-01-01 00:00:00")
	start_timestamp, _ := ptypes.TimestampProto(start_date)
	var end_date time.Time
	end_date, _ = time.Parse(layout, "2020-01-03 00:00:00")
	end_timestamp, _ := ptypes.TimestampProto(end_date)

	resp, err := client.AvailabilityCheck(ctx, &protos.AvailabilityCheckRequest{
		Name:      name,
		StartDate: start_timestamp,
		EndDate:   end_timestamp,
	})

	if err != nil {
		t.Errorf("Failed to check for availability %v", err)
	}

	if resp.Available != true {
		t.Errorf("Availability not as expected %v", resp.Available)
	}

	name = "road roller"
	start_date, _ = time.Parse(layout, "2021-10-19 00:00:00")
	end_date, _ = time.Parse(layout, "2021-10-21 00:00:00")
	start_timestamp, _ = ptypes.TimestampProto(start_date)
	end_timestamp, _ = ptypes.TimestampProto(end_date)

	resp, err = client.AvailabilityCheck(ctx, &protos.AvailabilityCheckRequest{
		Name:      name,
		StartDate: start_timestamp,
		EndDate:   end_timestamp,
	})

	if err != nil {
		t.Errorf("Failed to check for availability %v", err)
	}

	if resp.Available != false {
		t.Errorf("Availability not as expected %v", resp.Available)
	}



	log.Printf("Resp: %v", resp)

}