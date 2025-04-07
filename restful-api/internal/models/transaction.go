package models

import "time"

type Transaction struct {
    Id string               `json:"id"`
    UserId string           `json:"user_id"`
    Type string             `json:"type"`

	Target string           `json:"target"`
	Description string      `json:"description"`
	Category string         `json:"category"`
	Cost float64            `json:"cost"`

    Timestamp time.Time     `json:"timestamp"`
}
