package main

import (
	"fmt"
	"net/http"
	repository "rentit/pkg/repository"
	"rentit/pkg/service"
	httpTransport "rentit/pkg/transport/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"database/sql"

	_ "github.com/lib/pq"
)

const (
	logLevel        = "debug"
	httpServicePort = 8080
	postgresConnection = "dbname=postgres host=postgres password=postgres user=postgres sslmode=disable port=5432"

)

func main() {
	// begin setup
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		panic(err)
	}
	log.SetLevel(level)

	log.Info("Start server on port ", httpServicePort)

	// construct application
	dbConn, err := sql.Open("postgres", postgresConnection)
	plantRepository := repository.NewPlantRepository(dbConn)
	plantService := service.NewPlantService(plantRepository)
	plantHTTPHandler := httpTransport.NewPlantHandler(plantService)
	httpRouter := mux.NewRouter()

	plantHTTPHandler.RegisterRoutes(httpRouter)

	// setup http server
	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", httpServicePort),
		Handler: httpRouter,
	}

	err = httpSrv.ListenAndServe()
	if err != nil {
		log.Fatalf("Could not start server")
	}

	log.Infof("Stoped server")
}
