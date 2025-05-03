package repository

import "errors"

var (
	ErrorTxIdAlreadyExists = errors.New("Transaction with such ID already exists")
	ErrorTxIdNotFound = errors.New("No transaction with such ID exists")

	ErrorUserIdAlreadyExists = errors.New("User with such ID already exists")
	ErrorUserIdNotFound = errors.New("No user with such ID exists")
	ErrorNoSuchCredExists = errors.New("No user with such credentials exists")
)
