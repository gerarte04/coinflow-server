package utils

import (
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
