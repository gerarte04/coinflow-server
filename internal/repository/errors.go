package repository

import "errors"

var (
    ErrorTransactionKeyExists = errors.New("Transaction with such key already exists\n")
    ErrorTransactionKeyNotFound = errors.New("No transaction with such key exists\n")

    ErrorUserKeyExists = errors.New("User with such key already exists\n")
    ErrorUserKeyNotFound = errors.New("No user with such key exists\n")
    ErrorNoSuchCredExists = errors.New("No user with such credentials exists\n")
)
