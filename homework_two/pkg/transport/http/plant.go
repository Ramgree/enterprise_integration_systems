package http

import (
	"encoding/json"
	"net/http"
	"rentit/pkg/domain"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const (
	logLevel        = "debug"
	httpServicePort = 8080
)

type plantService interface {
	GetAll() ([]*domain.Plant, error)
	EstimateRental(name string, start_date time.Time, end_date time.Time) (float32, error)
	AvailabilityCheck(name string, start_date time.Time, end_date time.Time) (bool, error)
}

type PlantHandler struct {
	plantService plantService
}


func NewPlantHandler(pS plantService) *PlantHandler {
	return &PlantHandler{
		plantService: pS,
	}
}

func (h *PlantHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/plant", h.getAllPlants).Methods(http.MethodGet)
}

func (h *PlantHandler) getAllPlants(w http.ResponseWriter, _ *http.Request) {
	plants, err := h.plantService.GetAll()
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&plants)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}