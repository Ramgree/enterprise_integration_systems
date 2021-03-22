package domain

import "time"

type GetInfoQuery struct {
	Plant_name string
	Start_date time.Time
	End_date   time.Time
}

type SocketQuery struct {
	Command string
	Name *string
	StartDate *string
	EndDate *string
}