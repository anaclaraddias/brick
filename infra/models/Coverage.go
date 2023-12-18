package models

import "time"

type CoverageModel struct {
	Id          string
	Name        string
	Description string
	RateValue   float64
	CreatedAt   time.Time
}
