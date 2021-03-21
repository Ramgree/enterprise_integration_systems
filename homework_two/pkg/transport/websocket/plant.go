package websocket

import (
	"encoding/json"
	"net/http"
	"rentit/pkg/domain"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type plantService interface {
	GetAll() ([]*domain.Plant, error)
	EstimateRental(name string, start_date time.Time, end_date time.Time) (float32, error)
	AvailabilityCheck(name string, start_date time.Time, end_date time.Time) (bool, error)
}

type RebuildItService struct {
	plantService plantService
	wsUpgrader   websocket.Upgrader
}

func NewRebuildItHandler(pS plantService, ws websocket.Upgrader) *RebuildItService {
	return &RebuildItService{
		plantService: pS,
		wsUpgrader:   ws,
	}
}

func (h *RebuildItService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/plants", h.GetAll)
	// router.HandleFunc("/estimate", h.EstimateRental).Methods(http.MethodGet)
	// router.HandleFunc("/estimate", h.AvailabilityCheck).Methods(http.MethodGet)
}

func (h *RebuildItService) GetAll(w http.ResponseWriter, r *http.Request) {
	c, err := h.wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer c.Close()
	mt, _, err := c.ReadMessage()
	if err != nil {
		log.Error(err.Error())
		c.WriteMessage(mt, []byte(err.Error()))
		return
	}

	plants, err := h.plantService.GetAll()
	if err != nil {
		log.Error(err.Error())
		c.WriteMessage(mt, []byte(err.Error()))
		return
	}

	plantsList, err := json.Marshal(plants)
	err = c.WriteMessage(mt, plantsList)
	if err != nil {
		log.Println("write:", err)
	}
}
