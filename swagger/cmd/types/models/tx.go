package models

type Transaction struct {
	Type string             `json:"type"`
	Target string           `json:"target"`
	Description string      `json:"description"`
	Category string         `json:"category"`
	Cost float64            `json:"cost"`
}
