package models

import "time"

type User struct {
    Id string
    Login string
    Password string

    Name string
    Email string
    RegisterTime time.Time
}
