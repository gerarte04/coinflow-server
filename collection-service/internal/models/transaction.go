package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id     uuid.UUID
	UserId uuid.UUID
	Type   string

	Target      string
	Description string
	Category    string
	Cost        float64

	Timestamp time.Time
}
