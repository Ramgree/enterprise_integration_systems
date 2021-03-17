package domain

import "time"

type GetInfoQuery struct {
	Plant_name string
	Start_date time.Time
	End_date   time.Time
}
