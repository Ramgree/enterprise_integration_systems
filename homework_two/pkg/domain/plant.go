package domain

import (
	"encoding/json"
)

type Plant struct {
	Plant_id                 int32
	Plant_type_name          string
	Plant_daily_rental_price float32
	Plant_name               string
}

func (t *Plant) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}

func (t *Plant) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	return nil
}