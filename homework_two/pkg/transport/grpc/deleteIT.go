package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
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

func (s *rentitServiceServer) GetAll(context.Context,*empty.Empty) (*protos.Plants, error) {

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

	return &protos.Plants{
		Plants: plants,
	}, nil

}
