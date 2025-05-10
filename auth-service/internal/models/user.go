package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id uuid.UUID
	Login string
	Password string

	Name string
	Email string
	Phone string

	RegisterTimestamp time.Time
}
