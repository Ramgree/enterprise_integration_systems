package main

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"rentit/protos"
)

// This client is useless actually
// We won't use it
// DestroyIT will make their own client based on our proto

func main() {

	conn, err := grpc.Dial("localhost:10001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial %v", err)
	}
	defer conn.Close()

	client := protos.NewRentitServiceClient(conn)
	request := empty.Empty{}
	resp, err := client.GetAllPlants(context.Background(), &request)
	if err != nil {
		log.Fatalf("Failed to get all plants")
	}
	fmt.Println(resp)

}