package main

import (
	"fmt"
	"net/http"
	repository "rentit/pkg/repository"
	"rentit/pkg/service"
	httpTransport "rentit/pkg/transport/http"
	rebuildItWS "rentit/pkg/transport/websocket"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"database/sql"
	"net"
	rentitGrpc "rentit/pkg/transport/grpc"
	"rentit/protos"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	logLevel           = "debug"
	httpServicePort    = 8080
	grpcServicePort    = 10001
	wsServicePort      = 8081
	postgresConnection = "dbname=postgres host=postgres password=postgres user=postgres sslmode=disable port=5432"
	redisURI           = "redis:6379"
	redisPassword      = "" // no password set
	redisDB            = 0  // use default DB
)

func main() {
	// begin setup
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		panic(err)
	}
	log.SetLevel(level)

	// construct application
	redisConn := redis.NewClient(&redis.Options{
		Addr:     redisURI,
		Password: redisPassword,
		DB:       redisDB,
	})
	dbConn, err := sql.Open("postgres", postgresConnection)
	plantRepository := repository.NewPlantRepository(dbConn, redisConn)
	plantService := service.NewPlantService(plantRepository)

	// setup http server
	plantHTTPHandler := httpTransport.NewPlantHandler(plantService)
	httpRouter := mux.NewRouter()

	plantHTTPHandler.RegisterRoutes(httpRouter)

	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", httpServicePort),
		Handler: httpRouter,
	}
	log.Info("Serving HTTP (BuildIT) on port ", httpServicePort)

	go func() {
		err = httpSrv.ListenAndServe()
		if err != nil {
			log.Fatalf("Could not start http server")
		}
	}()

	// setup WS server
	websocketRouter := mux.NewRouter()
	websocketHandler := rebuildItWS.NewRebuildItHandler(plantService)
	websocketHandler.RegisterRoutes(websocketRouter)

	log.Info("Serving WebSocket (ReBuildIT) on port ", wsServicePort)
	wsSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", wsServicePort),
		Handler: websocketRouter,
	}

	go func() {
		log.Info("Starting ws")
		err = wsSrv.ListenAndServe()
		if err != nil {
			log.Fatalf("Could not start server")
		}
	}()

	// setup gRPC server

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", grpcServicePort))
	if err != nil {
		log.Fatalf("Failed to listen to gRPC port: %v", err)
	}
	grpcServer := grpc.NewServer()
	rentitServiceServer := rentitGrpc.NewRentitServiceServer(plantService)
	protos.RegisterRentitServiceServer(grpcServer, rentitServiceServer)

	log.Infof("Serving gRPC (DestroyIT) on port: %v", grpcServicePort)
	// make sure to run this on a separate thread when adding WS

	grpcServer.Serve(lis)
	log.Infof("Stopped application")
}
