package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"rentit/pkg/domain"
)

type PlantRepository struct {
	db *sql.DB
}

func NewPlantRepository(db *sql.DB) *PlantRepository {
	return &PlantRepository{
		db: db,
	}
}

func (r *PlantRepository) GetAll() ([]*domain.Plant, error) {
	log.Printf("received get all request")
	query := "SELECT p.plant_id, pt.plant_type_name, p.plant_daily_rental_price, p.plant_name FROM plant p LEFT JOIN plant_type pt ON pt.plant_type_id = p.plant_type_id;"
	rows, err := r.db.QueryContext(context.Background(), query)

	if err != nil {
		return nil, fmt.Errorf("Error getting all plants, %v", err)
	}

	plants := make([]*domain.Plant, 0)
	for rows.Next() {
		p := &domain.Plant{}
		err := rows.Scan(&p.Plant_id, &p.Plant_type_name, &p.Plant_daily_rental_price, &p.Plant_name)
		if err != nil {
			return nil, fmt.Errorf("error scaning query, %v", err)
		}
		plants = append(plants, p)
	}

	err = rows.Close()
	if err != nil {
		return nil, fmt.Errorf("could not close rows, %v", err)
	}

	return plants, nil

}

func (r *PlantRepository) EstimateRental(queryStruct *domain.GetInfoQuery) (float32, error) {

	log.Printf("Received an estimation request")
	query := "select p.plant_daily_rental_price * EXTRACT(DAY FROM ($3::timestamp - $2::timestamp)) from plant p WHERE p.plant_name ILIKE '%' || $1 || '%';"

	row := r.db.QueryRowContext(context.Background(), query, &queryStruct.Plant_name, &queryStruct.Start_date, &queryStruct.End_date)

	if row == nil {
		return 0, fmt.Errorf("error estimating rental, %v", row)
	}

	var estimation float32

	err := row.Scan(&estimation)

	if err != nil {
		return 0, fmt.Errorf("error estimating rental, %v", err)
	}

	return estimation, nil

}

func (r *PlantRepository) AvailabilityCheck(queryStruct *domain.GetInfoQuery) (bool, error) {

	log.Printf("Received an availability request")
	query :=
		`
	select CASE
    when exists(
            SELECT 1
            FROM booking b
            WHERE b.plant_id = (SELECT p.plant_id FROM plant p WHERE p.plant_name ILIKE '%' || $1 || '%')
              AND (($2::timestamp >= b.start_date AND $3::timestamp <= b.end_date) OR
                   ($2::timestamp <= b.start_date AND $3::timestamp >= b.start_date) OR
                   ($2::timestamp <= b.end_date AND $3::timestamp >= b.end_date))
        )
    then false
    else true
	end`

	row := r.db.QueryRowContext(context.Background(), query, &queryStruct.Plant_name, &queryStruct.Start_date, &queryStruct.End_date)

	if row == nil {
		return false, fmt.Errorf("error checking for availability, %v", row)
	}

	var availability bool

	err := row.Scan(&availability)

	if err != nil {
		return false, fmt.Errorf("error checking for availability, %v", err)
	}

	return availability, nil

}
