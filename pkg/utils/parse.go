package utils

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func ParseAuthHeader(values []string, kind string) (string, error) {
	if len(values) == 0 {
		return "", ErrorIncorrectToken
	}

	content := strings.Split(values[0], " ")

	if len(content) != 2 || content[0] != kind {
		return "", ErrorIncorrectToken
	}

	return content[1], nil
}

func ParseStringToUuid(s string) (uuid.UUID, error) {
	id, err := uuid.Parse(s)

	if err != nil {
		return uuid.Nil, fmt.Errorf("%w: %s", ErrorInvalidId, err.Error())
	}

	return id, nil
}

func CheckJwtFormat(token string) error {
	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		return fmt.Errorf("incorrect token separators count: %d", len(parts))
	}

	for i, pt := range parts[0:2] {
		if _, err := base64.StdEncoding.DecodeString(pt); err != nil {
			return fmt.Errorf("failed to decode part #%d: %w", i + 1, err)
		}
	}

	return nil
}
