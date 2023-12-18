package models

import (
	"database/sql"
	"time"
)

type PolicyModel struct {
	Id            string
	Status        string
	StartDate     string
	EndDate       string
	CoverageLimit float64
	Value         float64
	UserId        string
	CreatedAt     time.Time
	UpdatedAt     sql.NullTime
}
