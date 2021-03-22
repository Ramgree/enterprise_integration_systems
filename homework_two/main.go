package main

import (
	"context"
	"fmt"
	"net/http"
	repository "rentit/pkg/repository"
	"rentit/pkg/service"
	rentitHttp "rentit/pkg/transport/http"
	rebuildItWS "rentit/pkg/transport/websocket"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"

	"database/sql"
	"net"
	rentitGrpc "rentit/pkg/transport/grpc"
	"rentit/protos"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	logLevel           = "debug"
	httpServicePort    = 8080
	grpcServicePort    = 10001
	wsServicePort      = 8081
	postgresConnection = "dbname=postgres host=postgres password=postgres user=postgres sslmode=disable port=5432"
	mongoURI 		   = "mongodb://mongo:27017/"
	redisURI           = "redis:6379"
	redisPassword      = ""
	redisDB            = 0  
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

	mongoConn := options.Client().ApplyURI(mongoURI)
	clientMongo, err := mongo.Connect(context.Background(), mongoConn)
	if err != nil {
		log.Fatal(err)
	}

	plantRepository := repository.NewPlantRepository(clientMongo, dbConn, redisConn)
	plantService := service.NewPlantService(plantRepository)

	// setup http server
	plantHTTPHandler := rentitHttp.NewPlantHandler(plantService)
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
	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	websocketHandler := rebuildItWS.NewRebuildItHandler(plantService, upgrader)
	websocketHandler.RegisterRoutes(websocketRouter)

	log.Info("Serving WebSocket (ReBuildIT) on port ", wsServicePort)
	wsSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", wsServicePort),
		Handler: websocketRouter,
	}

	go func() {
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

	grpcServer.Serve(lis)
	log.Infof("Stopped application")
}
