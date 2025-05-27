package repository

import "errors"

var (
	ErrorTxIdAlreadyExists = errors.New("Transaction with such ID already exists")
	ErrorTxIdNotFound = errors.New("No transaction with such ID exists")

	ErrorPermissionDenied = errors.New("Permission to view transaction for this user denied")
)
