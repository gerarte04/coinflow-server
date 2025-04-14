package repository

import "errors"

var (
	ErrorTransactionKeyExists = errors.New("Transaction with such key already exists")
	ErrorTransactionKeyNotFound = errors.New("No transaction with such key exists")

	ErrorUserKeyExists = errors.New("User with such key already exists")
	ErrorUserKeyNotFound = errors.New("No user with such key exists")
	ErrorNoSuchCredExists = errors.New("No user with such credentials exists")
)
