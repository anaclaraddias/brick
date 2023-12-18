package models

import (
	"database/sql"
	"time"
)

type VehicleModel struct {
	Id           string
	Brand        string
	Model        string
	Year         string
	Renavam      string
	LicensePlate string
	Value        float64
	Cargo        float64
	Height       float64
	Width        float64
	Length       float64
	Type         string
	CreatedAt    time.Time
	UpdatedAt    sql.NullTime
}
