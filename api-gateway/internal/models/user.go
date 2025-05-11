package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id 						uuid.UUID	`swaggerignore:"true"`
	Login 					string		`json:"login"`	
	Password 				string		`json:"password"`

	Name 					string		`json:"name"`
	Email 					string		`json:"email"`
	Phone 					string		`json:"phone"`

	RegistrationTimestamp 	time.Time	`swaggerignore:"true"`
}
