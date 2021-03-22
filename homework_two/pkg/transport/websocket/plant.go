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

const(
	dateFormat = "2006-01-02"
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
	router.HandleFunc("/ws", h.WsQuery)
}

func (h* RebuildItService) WsQuery (w http.ResponseWriter, r *http.Request){
	c, err := h.wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("ws upgrade:", err)
		return
	}
	defer c.Close()

	// loop while connection is alive
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("ws read:", err)
			break
		}
		resp := h.executeQuery(message)
		err = c.WriteMessage(mt, resp)
		if err != nil {
			log.Println("ws write: ", err)
			break
		}
	}
}

func (h* RebuildItService) executeQuery (message []byte) (response []byte){
	var query domain.SocketQuery
	err := json.Unmarshal(message, &query)

	if err != nil {
		msg, _ := json.Marshal(map[string]string{"error": "Failed to decode message " + err.Error()})
		return msg
	}

	// get all 

	if query.Command == "all"{
		var msg []byte
		res, err := h.plantService.GetAll()
		if err != nil{
			msg, _ = json.Marshal(map[string]string{"error": "Failed to get all Plants " + err.Error()})
		}else{
			msg, _ = json.Marshal(res)
		}
		return msg
	}


	// Queries with params

	startDateStr := query.StartDate
	endDateStr := query.EndDate

	if startDateStr == nil || endDateStr == nil || query.Name == nil{
		msg, _ := json.Marshal(map[string]string{"error": "Query has missing fields"})
		return msg
	}

	startDate, dateErr1 := time.Parse(dateFormat, *startDateStr)
	endDate, dateErr2 := time.Parse(dateFormat, *endDateStr)

	if dateErr1 != nil {
		msg, _ := json.Marshal(map[string]string{"error": "Failed to parse start date, " + dateErr1.Error()})
		return msg
	}

	if dateErr2 != nil {
		msg, _ := json.Marshal(map[string]string{"error": "Failed to parse end date, " + dateErr2.Error()})
		return msg
	}

	var msg []byte

	switch command := query.Command; command {
	case "estimate":
		res, err := h.plantService.EstimateRental(*query.Name, startDate, endDate)

		if err != nil {
			msg, _ = json.Marshal(map[string]string{"error": "Failed to get estimate,  " + err.Error()})
			break
		}
		msg, _ = json.Marshal(map[string]float32{"price": res})
		break

	case "availability":
		res, err := h.plantService.AvailabilityCheck(*query.Name, startDate, endDate)

		if err != nil {
			msg, _ = json.Marshal(map[string]string{"error": "Failed to get estimate,  " + err.Error()})
			break
		}
		msg, _ = json.Marshal(map[string]bool{"isAvailable": res})
		break

	default:
		msg, _ = json.Marshal(map[string]string{"error": "Invalid query type"})
		break
	}

	return msg
}