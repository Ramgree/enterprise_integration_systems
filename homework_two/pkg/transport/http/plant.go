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
	dateFormat = "2006-01-02"
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
	router.HandleFunc("/plants", h.getAllPlants).Methods(http.MethodGet)
	router.HandleFunc("/estimate", h.EstimateRental).Methods(http.MethodGet)
	router.HandleFunc("/availability", h.AvailabilityCheck).Methods(http.MethodGet)
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
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *PlantHandler) EstimateRental(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	startDateStr := r.URL.Query().Get("from")
	endDateStr := r.URL.Query().Get("to")

	startDate, dateErr1 := time.Parse(dateFormat, startDateStr)
	endDate, dateErr2 := time.Parse(dateFormat, endDateStr)

	if dateErr1 != nil{
		log.Error(dateErr1.Error())
		http.Error(w, dateErr1.Error(), http.StatusBadRequest)
		return
	}

	if dateErr2 != nil{
		log.Error(dateErr2.Error())
		http.Error(w, dateErr2.Error(), http.StatusBadRequest)
		return
	}

	price, err := h.plantService.EstimateRental(name, startDate, endDate)

	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
    
	// write success response
	w.WriteHeader(http.StatusOK)
	res := map[string]float32{"price": price}
	err = json.NewEncoder(w).Encode(res)
	
	if err != nil {
        log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func (h *PlantHandler) AvailabilityCheck(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	startDateStr := r.URL.Query().Get("from")
	endDateStr := r.URL.Query().Get("to")

	startDate, dateErr1 := time.Parse(dateFormat, startDateStr)
	endDate, dateErr2 := time.Parse(dateFormat, endDateStr)

	if dateErr1 != nil{
		log.Error(dateErr1.Error())
		http.Error(w, dateErr1.Error(), http.StatusBadRequest)
		return
	}

	if dateErr2 != nil{
		log.Error(dateErr2.Error())
		http.Error(w, dateErr2.Error(), http.StatusBadRequest)
		return
	}

	isAvailable, err := h.plantService.AvailabilityCheck(name, startDate, endDate)

	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
    
	// write success response
	w.WriteHeader(http.StatusOK)

	res := map[string]bool{"isAvailable": isAvailable}
	err = json.NewEncoder(w).Encode(res)
	
	if err != nil {
        log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}