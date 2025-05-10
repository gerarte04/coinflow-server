package repository

import "errors"

var (
	ErrorUserIdNotFound = errors.New("No user with such id found")
	ErrorUserLoginNotFound = errors.New("No user with such login found")
	ErrorWrongPassword = errors.New("Wrong password")
	ErrorUserCredAlreadyExists = errors.New("User with such login or email already exists")
)
