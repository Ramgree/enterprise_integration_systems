package test

import (
	"context"
	"database/sql"
	"log"
	"rentit/pkg/repository"
	"rentit/pkg/service"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoConnection = "mongodb://localhost:27017/"
	postgresConnection = "dbname=postgres host=localhost password=postgres user=postgres sslmode=disable port=5432"
	// What the duck is this ARBITRARY date????????????????
	layout   = "2006-01-02 15:04:05"
	redisKey = "app:plant"
)

func TestGetAllService(t *testing.T) {
	dbConn, err := sql.Open("postgres", postgresConnection)
	if err != nil {
		t.Error(err)
	}
	defer dbConn.Close()

	mongoConn := options.Client().ApplyURI(mongoConnection)
	clientMongo, err := mongo.Connect(context.Background(), mongoConn)
	if err != nil {
		log.Fatal(err)
	}

	redisConn := redis.NewClient(&redis.Options{
		Addr:     redisURI,
		Password: redisPassword,
		DB:       redisDB,
	})

	defer redisConn.Close()
	plantRepository := repository.NewPlantRepository(clientMongo, dbConn, redisConn)
	plantService := service.NewPlantService(plantRepository)

	vals, err := plantService.GetAll()

	if err != nil {
		t.Error(err)
	}

	for _, val := range vals {

		if val == nil {
			t.Error("Nil value")
		}
		//fmt.Println(*val)
	}

	// checking cache
	cached, err := redisConn.Exists(context.Background(), redisKey).Result()

	if err != nil {
		t.Error("Checking cache failed")
	}

	if cached == 0 {
		t.Error("Plants not found in cache")
	}
}

func TestEstimateRentalService(t *testing.T) {
	dbConn, err := sql.Open("postgres", postgresConnection)
	if err != nil {
		t.Error(err)
	}
	defer dbConn.Close()

	mongoConn := options.Client().ApplyURI(mongoConnection)
	clientMongo, err := mongo.Connect(context.Background(), mongoConn)
	if err != nil {
		log.Fatal(err)
	}

	redisConn := redis.NewClient(&redis.Options{
		Addr:     redisURI,
		Password: redisPassword,
		DB:       redisDB,
	})

	defer redisConn.Close()

	plantRepository := repository.NewPlantRepository(clientMongo, dbConn, redisConn)
	plantService := service.NewPlantService(plantRepository)

	// from postgres
	name := "excavator"
	var start_date time.Time
	start_date, _ = time.Parse(layout, "2020-01-01 00:00:00")
	var end_date time.Time
	end_date, _ = time.Parse(layout, "2020-01-03 00:00:00")

	vals, err := plantService.EstimateRental(name, start_date, end_date)

	if err != nil {
		t.Error(err)
	}

	if vals != 2500 {
		t.Error("Wrong estimate response", vals)
	}

	// from mongo db
	name = "sweeper"
	vals, err = plantService.EstimateRental(name, start_date, end_date)

	if err != nil {
		t.Error(err)
	}

	if vals != 5000 {
		t.Error("Wrong estimate response", vals)
	}

}

func TestAvailabilityCheckService(t *testing.T) {
	dbConn, err := sql.Open("postgres", postgresConnection)
	if err != nil {
		t.Error(err)
	}
	defer dbConn.Close()

	mongoConn := options.Client().ApplyURI(mongoConnection)
	clientMongo, err := mongo.Connect(context.Background(), mongoConn)
	if err != nil {
		log.Fatal(err)
	}

	redisConn := redis.NewClient(&redis.Options{
		Addr:     redisURI,
		Password: redisPassword,
		DB:       redisDB,
	})

	defer redisConn.Close()

	plantRepository := repository.NewPlantRepository(clientMongo, dbConn, redisConn)
	plantService := service.NewPlantService(plantRepository)

	name := "crane"
	var start_date time.Time
	start_date, _ = time.Parse(layout, "2021-11-17 00:00:00")
	var end_date time.Time
	end_date, _ = time.Parse(layout, "2021-11-20 00:00:00")

	vals, err := plantService.AvailabilityCheck(name, start_date, end_date)

	if err != nil {
		t.Error(err)
	}

	if vals != false {
		t.Error("Wrong availability response", vals)
	}

	start_date, _ = time.Parse(layout, "2021-10-19 00:00:00")
	end_date, _ = time.Parse(layout, "2021-10-21 00:00:00")

	vals, err = plantService.AvailabilityCheck(name, start_date, end_date)

	if err != nil {
		t.Error(err)
	}

	if vals != true {
		t.Error("wrong availability response", vals)
	}

	// From mongo
	name = "sweeper"
	vals, err = plantService.AvailabilityCheck(name, start_date, end_date)

	if err != nil {
		t.Error(err)
	}

	if vals != true {
		t.Error("wrong availability response", vals)
	}
}
