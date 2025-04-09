package testutils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func SendRequest(t *testing.T, method string, url string, payload any) (*http.Response, error) {
    body, err := json.Marshal(payload)
    require.NoError(t, err)

    req, err := http.NewRequest(method, url, bytes.NewReader(body))
    require.NoError(t, err)

    cli := http.Client{}
    return cli.Do(req)
}

func DecodeResponse(t *testing.T, resp *http.Response) (map[string]any) {
    var decoded map[string]any

    err := json.NewDecoder(resp.Body).Decode(&decoded)
    require.NoError(t, err)

    return decoded
}
