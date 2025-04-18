package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id uuid.UUID            `json:"id" swaggerignore:"true"`
	UserId uuid.UUID        `json:"user_id"`
	Type string             `json:"type"`

	Target string           `json:"target"`
	Description string      `json:"description"`
	Category string         `json:"category" swaggerignore:"true"`
	Cost float64            `json:"cost"`

	Timestamp time.Time     `json:"timestamp" swaggerignore:"true"`
}
