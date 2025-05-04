package testutils

import (
	"slices"
	"testing"

	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/require"
)

type Payload map[string]any

func GetPayloadCopy(t *testing.T, payload Payload) Payload {
	var newPayload Payload
	require.NoError(t, copier.CopyWithOption(&newPayload, &payload, copier.Option{DeepCopy: true}))

	return newPayload
}

type ValidateOpt struct {
	Ignore bool
	Key string
	CheckValue bool
	Value any
}

func ValidateResult(t *testing.T, res Payload, payload Payload, opts ...ValidateOpt) {
	ignoreFields := make([]string, 0)

	for _, opt := range opts {
		if opt.Ignore {
			ignoreFields = append(ignoreFields, opt.Key)
		} else {
			require.Contains(t, res, opt.Key)
	
			if opt.CheckValue {
				require.Equal(t, opt.Value, res[opt.Key])
			}
		}
	}

	for k, v := range payload {
		if !slices.Contains(ignoreFields, k) {
			require.Contains(t, res, k)
			require.Equal(t, v, res[k])
		}
	}
}
