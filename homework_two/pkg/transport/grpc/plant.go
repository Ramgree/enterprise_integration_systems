package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
	"rentit/pkg/domain"
	"rentit/protos"
	"time"
)



type plantService interface {
	GetAll() ([]*domain.Plant, error)
	EstimateRental(name string, start_date time.Time, end_date time.Time) (float32, error)
	AvailabilityCheck(name string, start_date time.Time, end_date time.Time) (bool, error)
}

type rentitServiceServer struct {
	plantService plantService
}

func NewRentitServiceServer(pS plantService) *rentitServiceServer {
	return &rentitServiceServer{
		plantService: pS,
	}
}

func (s *rentitServiceServer) GetAllPlants(context.Context, *empty.Empty) (*protos.GetAllPlantsResponse, error) {

	var (
		plants []*protos.Plant
	)
	rows, err := s.plantService.GetAll()

	if err != nil {
		return nil, err
	}

	for _, val := range rows {

		plants = append(plants, &protos.Plant{
			PlantId:               val.Plant_id,
			PlantTypeName:         val.Plant_type_name,
			PlantDailyRentalPrice: val.Plant_daily_rental_price,
			PlantName:             val.Plant_name,
		},
		)
	}

	return &protos.GetAllPlantsResponse{
		Plants: plants,
	}, nil

}

func (s *rentitServiceServer) EstimateRental(ctx context.Context, request *protos.EstimateRentalRequest) (*protos.EstimateRentalResponse, error) {
	var (
		response *protos.EstimateRentalResponse
		start_date time.Time
		end_date time.Time
		err error
	)

	start_date, err = ptypes.Timestamp(request.StartDate)
	if err != nil {
		log.Printf("Failed to convert proto start_date timestamp to time.Time %v", err)
		return nil, err
	}
	end_date, err = ptypes.Timestamp(request.EndDate)
	if err != nil {
		log.Printf("Failed to oncvert proto end_date timestamp to time.Time %v", err)
		return nil, err
	}

	row, err := s.plantService.EstimateRental(request.Name, start_date, end_date)

	if err != nil {
		log.Printf("Failed to estimate rental %v", err)
	}

	response = &protos.EstimateRentalResponse{PriceEstimation: row}

	return response,
		nil
}

func (s *rentitServiceServer) AvailabilityCheck(ctx context.Context, request *protos.AvailabilityCheckRequest) (*protos.AvailabilityCheckResponse, error) {
	var (
		response *protos.AvailabilityCheckResponse
		start_date time.Time
		end_date time.Time
		err error
	)

	start_date, err = ptypes.Timestamp(request.StartDate)
	if err != nil {
		log.Printf("Failed to convert proto start_date timestamp to time.Time %v", err)
		return nil, err
	}
	end_date, err = ptypes.Timestamp(request.EndDate)
	if err != nil {
		log.Printf("Failed to oncvert proto end_date timestamp to time.Time %v", err)
		return nil, err
	}

	row, err := s.plantService.AvailabilityCheck(request.Name, start_date, end_date)

	if err != nil {
		log.Printf("Failed to estimate rental %v", err)
	}

	response = &protos.AvailabilityCheckResponse{Available: row}

	return response,
		nil
}