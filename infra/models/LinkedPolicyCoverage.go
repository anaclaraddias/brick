package models

import "time"

type LinkedPolicyCoverageModel struct {
	Id         string
	CoverageId string
	PolicyId   string
	CreatedAt  time.Time
}
