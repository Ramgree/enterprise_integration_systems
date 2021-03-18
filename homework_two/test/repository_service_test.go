package test

import (
	"database/sql"
	"rentit/pkg/repository"
	"rentit/pkg/service"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

const (
	postgresConnection = "postgres://postgres:postgres@postgres:5432?sslmode=disable"
	// What the fuck is this ARBITRARY date????????????????
	layout = "2006-01-02 15:04:05"
)

func TestGetAllRepository(t *testing.T) {
	dbConn, err := sql.Open("postgres", postgresConnection)
	if err != nil {
		t.Error(err)
	}
	defer dbConn.Close()

	redisConn := redis.NewClient(&redis.Options{
		Addr:     redisURI,
		Password: redisPassword,
		DB:       redisDB,
	})

	defer redisConn.Close()
	plantRepository := repository.NewPlantRepository(dbConn, redisConn)
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
}

func TestEstimateRentalRepository(t *testing.T) {
	dbConn, err := sql.Open("postgres", postgresConnection)
	if err != nil {
		t.Error(err)
	}
	defer dbConn.Close()

	redisConn := redis.NewClient(&redis.Options{
		Addr:     redisURI,
		Password: redisPassword,
		DB:       redisDB,
	})

	defer redisConn.Close()

	plantRepository := repository.NewPlantRepository(dbConn, redisConn)
	plantService := service.NewPlantService(plantRepository)

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
		t.Error("Wrong availability response", vals)
	}

}

func TestAvailabilityCheckRepository(t *testing.T) {
	dbConn, err := sql.Open("postgres", postgresConnection)
	if err != nil {
		t.Error(err)
	}
	defer dbConn.Close()

	redisConn := redis.NewClient(&redis.Options{
		Addr:     redisURI,
		Password: redisPassword,
		DB:       redisDB,
	})

	defer redisConn.Close()

	plantRepository := repository.NewPlantRepository(dbConn, redisConn)
	plantService := service.NewPlantService(plantRepository)

	name := "road roller"
	var start_date time.Time
	start_date, _ = time.Parse(layout, "2020-01-01 00:00:00")
	var end_date time.Time
	end_date, _ = time.Parse(layout, "2020-01-03 00:00:00")

	vals, err := plantService.AvailabilityCheck(name, start_date, end_date)

	if err != nil {
		t.Error(err)
	}

	if vals != true {
		t.Error("Wrong availability response", vals)
	}

	start_date, _ = time.Parse(layout, "2021-10-19 00:00:00")
	end_date, _ = time.Parse(layout, "2021-10-21 00:00:00")

	vals, err = plantService.AvailabilityCheck(name, start_date, end_date)

	if err != nil {
		t.Error(err)
	}

	if vals != false {
		t.Error("wrong availability response", vals)
	}

}
