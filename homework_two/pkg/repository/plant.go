package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"rentit/pkg/domain"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	redisKey = "app:plant"
	mongoDbName = "Plants"
	mongoCollectionName = "plant"
)

type PlantRepository struct {
	mongoClient *mongo.Client
	db          *sql.DB
	redis       *redis.Client
}

func NewPlantRepository(mongoClient *mongo.Client, db *sql.DB, redis *redis.Client) *PlantRepository {
	return &PlantRepository{
		mongoClient: mongoClient,
		db:          db,
		redis:       redis,
	}
}

func (r *PlantRepository) GetAll() ([]*domain.Plant, error) {
	log.Printf("received get all request")

	// checking cache
	cached, err := r.redis.Exists(context.Background(), redisKey).Result()

	if err != nil {
		log.Println("Error checking plants in cache")
	}

	if cached == 1 {
		log.Println("Retrieving plants from cache")

		// results are not in the original order, should be ok as it is not specified in the task
		if res, err := r.getAllRedis(); err == nil{
			return res, nil
		}
		log.Println("Failed to get plants from cache: " + err.Error())
	} else {
		log.Println("Plants not found in cache, querying DB")
	}

	// mongo
	results, err := r.getAllMongo()

	if err != nil{
		return nil, fmt.Errorf("Failed to get plants from Mongo, %v", err)
	}

	// postgres
	postgresResults, err := r.getAllPostgres()
	
	if err != nil {
		return nil, fmt.Errorf("Failed to get plants from Postgres, %v", err)
	}

	// combine mongo and postgres
	for _, elem := range postgresResults {
		results = append(results, elem)
	}

	for _, elem := range results{
		// cache the plant
		_, err := r.redis.HSetNX(context.Background(), redisKey, string(elem.Plant_id), elem).Result()

		if err != nil {
			log.Println(err.Error())
			log.Println("Failed to cache a plant, idk what to do about it..")
		}
	}

	return results, nil
}

func (r *PlantRepository) EstimateRental(queryStruct *domain.GetInfoQuery) (float32, error) {

	log.Printf("Received an estimation request")

	if res, err := r.estimatePricePostgres(queryStruct); err == nil {
		return res, nil
	}

	log.Println("Plant not in Postgres, checking Mongo")
	res, err := r.estimatePriceMongo(queryStruct)

	if err != nil{
		return 0, fmt.Errorf("Plant not found")
	}

	return res, nil
}

/*
*  It's not specified to keep bookings in Mongo, so they will all be in PG
*/
	
func (r *PlantRepository) AvailabilityCheck(queryStruct *domain.GetInfoQuery) (bool, error) {

	log.Println("Received an availability request")
	log.Println("Checking Postgres")

	plantId, err := r.getPlantIdPostgres(queryStruct.Plant_name)

	if err == nil{
		av, err := r.getAvailabilityPostgres(plantId, queryStruct)
		if err != nil{
			return false, fmt.Errorf("Failed to get availability, %v", err) 
		}
		return av, nil
	}

	log.Println(err.Error())
	log.Println("Checking Mongo")

	plantId, err = r.getPlantIdMongo(queryStruct.Plant_name)

	if err != nil{
		return false, fmt.Errorf("Couldn't get plant Id from either DB") 
	}

	av, err := r.getAvailabilityPostgres(plantId, queryStruct)
	if err != nil{
		return false, fmt.Errorf("Failed to get availability, %v", err) 
	}
	return av, nil
}

func (r *PlantRepository) getAllRedis() ([]*domain.Plant, error) {
	res, err := r.redis.HGetAll(context.Background(), redisKey).Result()

	if err != nil {
		return nil, fmt.Errorf("Failed to get plants from cache: %v", err)
	}

	plants := []*domain.Plant{}
	for _, stringValue := range res {
		b := &domain.Plant{}
		err := json.Unmarshal([]byte(stringValue), b)
		if err != nil {
			return nil, fmt.Errorf("Error decoding plant from cache: %v", err)
		}
		plants = append(plants, b)
	}
	return plants, nil
}

func (r *PlantRepository) getAllMongo() ([]*domain.Plant, error) {
	mongoList, err := r.mongoClient.Database(mongoDbName).Collection(mongoCollectionName).Find(context.Background(), bson.D{})
	results := []*domain.Plant{}

	if err != nil {
		return nil, fmt.Errorf("Mongo query error, %v", err)
	}

	for mongoList.Next(context.Background()) {

		elem := &domain.Plant{}
		err := mongoList.Decode(&elem)
		if err != nil {
			return nil, fmt.Errorf("Mongo iter error, %v", err)
		}
		results = append(results, elem)
	}
	mongoList.Close(context.Background())

	return results, nil
}

func (r *PlantRepository) getAllPostgres() ([]*domain.Plant, error) {
	query := "SELECT p.plant_id, p.plant_type_name, p.plant_daily_rental_price, p.plant_name FROM plant p;"
	rows, err := r.db.QueryContext(context.Background(), query)

	plants := make([]*domain.Plant, 0)
	for rows.Next() {
		p := &domain.Plant{}
		err := rows.Scan(&p.Plant_id, &p.Plant_type_name, &p.Plant_daily_rental_price, &p.Plant_name)
		if err != nil {
			return nil, fmt.Errorf("Error postgres scaning query, %v", err)
		}
		plants = append(plants, p)
	}

	err = rows.Close()
	if err != nil {
		return nil, fmt.Errorf("Could not close postrgres rows, %v", err)
	}

	return plants, nil
}

func (r *PlantRepository) estimatePricePostgres(queryStruct *domain.GetInfoQuery) (float32, error){
	query := "select p.plant_daily_rental_price * EXTRACT(DAY FROM ($3::timestamp - $2::timestamp)) from plant p WHERE p.plant_name ILIKE '%' || $1 || '%';"

	row := r.db.QueryRowContext(context.Background(), query, &queryStruct.Plant_name, &queryStruct.Start_date, &queryStruct.End_date)

	if row == nil {
		return 0, fmt.Errorf("error estimating rental, %v", row)
	}

	var estimation float32

	err := row.Scan(&estimation)

	if err != nil {
		return 0, fmt.Errorf("Error estimating rental, %v", err)
	}

	return estimation, nil
}

func (r *PlantRepository) estimatePriceMongo(queryStruct *domain.GetInfoQuery) (float32, error){

	duration := queryStruct.End_date.Sub(queryStruct.Start_date).Hours() / 24
	db :=  r.mongoClient.Database(mongoDbName).Collection(mongoCollectionName)

	result := db.FindOne(context.Background(), bson.M{"plantName": queryStruct.Plant_name})

	if result == nil {
		return 0, fmt.Errorf("Error estimating rental, couldn't find plant")
	}

	plant := domain.Plant{}
	if err := result.Decode(&plant); err != nil{
		return 0, fmt.Errorf("Error decoding result, %v", err)
	}


	return plant.Plant_daily_rental_price * float32(duration), nil
}

func (r *PlantRepository) getPlantIdPostgres(name string) (int, error){
	query := "SELECT p.plant_id FROM plant p WHERE p.plant_name LIKE $1;"
	row := r.db.QueryRowContext(context.Background(), query, name)

	if row == nil {
		return 0, fmt.Errorf("Id not in Postgres")
	}

	var id int
	err := row.Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("Error scanning row, %v", err)
	}

	return id, nil
}

func (r *PlantRepository) getPlantIdMongo(name string) (int, error){
	db :=  r.mongoClient.Database(mongoDbName).Collection(mongoCollectionName)

	result := db.FindOne(context.Background(), bson.M{"plantName": name})

	if result == nil {
		return 0, fmt.Errorf("Error checking for Id, couldn't find plant")
	}

	plant := domain.Plant{}
	if err := result.Decode(&plant); err != nil{
		return 0, fmt.Errorf("Error decoding result, %v", err)
	}


	return int(plant.Plant_id), nil
}

func (r *PlantRepository) getAvailabilityPostgres(id int, queryStruct *domain.GetInfoQuery) (bool, error){
	query :=
		`
	select CASE
	when exists(
			SELECT 1
			FROM booking b
			WHERE b.plant_id = $1
			AND (($2::timestamp >= b.start_date AND $3::timestamp <= b.end_date) OR
				($2::timestamp <= b.start_date AND $3::timestamp >= b.start_date) OR
				($2::timestamp <= b.end_date AND $3::timestamp >= b.end_date))
		)
	then false
	else true
	end`

	row := r.db.QueryRowContext(context.Background(), query, id, &queryStruct.Start_date, &queryStruct.End_date)

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
