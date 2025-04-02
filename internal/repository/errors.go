package repository

import "errors"

var (
    ErrorTransactionKeyExists = errors.New("Transaction with such key already exists\n")
    ErrorTransactionKeyNotFound = errors.New("No transaction with such key exists\n")
)
