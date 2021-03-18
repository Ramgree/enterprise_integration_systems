package grpc

import (
	"rentit/pkg/domain"
	"rentit/protos"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)
type plantService interface {
	GetAll() ([]*domain.Plant, error)
	EstimateRental(name string, start_date time.Time, end_date time.Time) (float32, error)
	AvailabilityCheck(name string, start_date time.Time, end_date time.Time) (bool, error)
}

type rentitServiceServer struct {
	plantService plantService
	rentitServer protos.UnimplementedRentitServiceServer
}

func NewRentitServiceServer(pS plantService) *rentitServiceServer {
	return &rentitServiceServer{
		plantService: pS,
	}
}

func (s *rentitServiceServer) GetAll(req *emptypb.Empty, srv protos.RentitService_GetAllPlantsServer) (*protos.Plants, error) {

	var (
		plants []*protos.Plant
		//plant_id                 int32
		//plant_type_name          string
		//plant_daily_rental_price float32
		//plant_name               string
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
