package models

import "time"

type Transaction struct {
    Id string
    UserId string
    Type string

	Target string
	Description string
	Category string
	Cost float64

    Timestamp time.Time
}
