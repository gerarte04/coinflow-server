package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id uuid.UUID            `json:"id" db:"id" swaggerignore:"true"`
	UserId uuid.UUID        `json:"user_id" db:"user_id"`

	Type string             `json:"type" db:"type"`
	Target string           `json:"target" db:"target"`
	Description string      `json:"description" db:"description"`
	Category string         `json:"category" db:"category" swaggerignore:"true"`
	Cost float64            `json:"cost" db:"cost"`

	Timestamp time.Time     `json:"timestamp" db:"timestamp" swaggerignore:"true"`
}
