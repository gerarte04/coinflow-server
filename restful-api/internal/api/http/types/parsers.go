package types

import (
	"fmt"

	"github.com/google/uuid"
)

func ParseStringToTransactionId(s string) (uuid.UUID, error) {
	id, err := uuid.Parse(s)

	if err != nil {
		return uuid.Nil, fmt.Errorf("%w: %s", ErrorInvalidId, err.Error())
	}

	return id, nil
}
