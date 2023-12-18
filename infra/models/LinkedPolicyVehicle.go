package models

import "time"

type LinkedPolicyVehicleModel struct {
	Id        string
	VehicleId string
	PolicyId  string
	CreatedAt time.Time
}
