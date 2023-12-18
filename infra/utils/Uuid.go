package utils

import "github.com/google/uuid"

type Uuid struct{}

func NewUuid() *Uuid {
	return &Uuid{}
}

func (uuidgenerator *Uuid) GenerateUuid() string {
	return uuid.New().String()
}
