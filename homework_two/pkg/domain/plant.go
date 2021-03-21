package domain

import (
	"encoding/json"
)

type Plant struct {
	Plant_id                 int32   `bson:"_id,omitempty" json:"ID"`
	Plant_type_name          string  `bson:"plantTypeName,omitempty" json:"plantTypeName"`
	Plant_daily_rental_price float32 `bson:"plantRentalDailyPrice,omitempty" json:"plantRentalDailyPrice"`
	Plant_name               string  `bson:"plantName,omitempty" json:"plantName"`
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
