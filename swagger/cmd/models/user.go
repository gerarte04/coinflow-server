package models

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`

	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
