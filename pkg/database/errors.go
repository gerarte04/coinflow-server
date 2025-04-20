package database

import "errors"

var (
	ErrorUniqueViolation = errors.New("Unique violation")
	ErrorUndocumented = errors.New("Undocumented error")
)
