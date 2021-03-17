package test

import (
	"database/sql"
	"fmt"
	"rentit/pkg/repository"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

const (
	postgresConnection = "dbname=postgres host=localhost password=postgres user=postgres sslmode=disable port=5432"
	// What the fuck is this ARBITRARY date????????????????
	layout = "2006-01-02 15:04:05"
)

func TestGetAll(t *testing.T) {
	dbConn, err := sql.Open("postgres", postgresConnection)
	if err != nil {
		t.Error(err)
	}
	defer dbConn.Close()
	plantRepository := repository.NewPlantRepository(dbConn)

	vals, err := plantRepository.GetAll()

	if err != nil {
		t.Error(err)
	}

	for _, val := range vals {

		if val == nil {
			t.Error("Nil value")
		}
		fmt.Println(*val)
	}
}

func TestEstimateRental(t *testing.T) {
	dbConn, err := sql.Open("postgres", postgresConnection)
	if err != nil {
		t.Error(err)
	}
	defer dbConn.Close()
	plantRepository := repository.NewPlantRepository(dbConn)

	name := "excavator"
	var start_date time.Time
	start_date, _ = time.Parse(layout, "2020-01-01 00:00:00")
	var end_date time.Time
	end_date, _ = time.Parse(layout, "2020-01-03 00:00:00")

	vals, err := plantRepository.EstimateRental(name, start_date, end_date)

	if err != nil {
		t.Error(err)
	}

	if vals != 2500 {
		t.Error("Wrong availability response", vals)
	}

}

func TestAvailabilityCheck(t *testing.T) {
	dbConn, err := sql.Open("postgres", postgresConnection)
	if err != nil {
		t.Error(err)
	}
	defer dbConn.Close()
	plantRepository := repository.NewPlantRepository(dbConn)

	name := "road roller"
	var start_date time.Time
	start_date, _ = time.Parse(layout, "2020-01-01 00:00:00")
	var end_date time.Time
	end_date, _ = time.Parse(layout, "2020-01-03 00:00:00")

	vals, err := plantRepository.AvailabilityCheck(name, start_date, end_date)

	if err != nil {
		t.Error(err)
	}

	if vals != true {
		t.Error("Wrong availability response", vals)
	}

	start_date, _ = time.Parse(layout, "2021-10-19 00:00:00")
	end_date, _ = time.Parse(layout, "2021-10-21 00:00:00")

	vals, err = plantRepository.AvailabilityCheck(name, start_date, end_date)

	if err != nil {
		t.Error(err)
	}

	if vals != false {
		t.Error("wrong availability response", vals)
	}

}
