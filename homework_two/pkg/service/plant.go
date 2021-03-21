package service

import (
	"rentit/pkg/domain"
	"time"
)

type plantRepository interface {
	GetAll() ([]*domain.Plant, error)
	EstimateRental(*domain.GetInfoQuery) (float32, error)
	AvailabilityCheck(*domain.GetInfoQuery) (bool, error)
}

type PlantService struct {
	plantRepository plantRepository
}

func NewPlantService(pR plantRepository) *PlantService {
	return &PlantService{
		plantRepository: pR,
	}
}

func (s *PlantService) GetAll() ([]*domain.Plant, error) {
	return s.plantRepository.GetAll()

}

func (s *PlantService) EstimateRental(name string, start_date time.Time, end_date time.Time) (float32, error) {
	queryStruct := domain.GetInfoQuery{
		Plant_name: name,
		Start_date: start_date,
		End_date:   end_date,
	}

	return s.plantRepository.EstimateRental(&queryStruct)
}

func (s *PlantService) AvailabilityCheck(name string, start_date time.Time, end_date time.Time) (bool, error) {
	queryStruct := domain.GetInfoQuery{
		Plant_name: name,
		Start_date: start_date,
		End_date:   end_date,
	}
	return s.plantRepository.AvailabilityCheck(&queryStruct)
}
