package types

import (
	"fmt"

	"github.com/google/uuid"
)

func ParseStringToTransactionId(s string) (string, error) {
    if _, err := uuid.Parse(s); err != nil {
        return "", fmt.Errorf("%w: %s", ErrorInvalidId, err.Error())
    }

    return s, nil
}
