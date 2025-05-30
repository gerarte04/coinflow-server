package testutils

import (
	"encoding/json"
	"net/http"
	"testing"

	"coinflow/coinflow-server/pkg/http/request"

	"github.com/stretchr/testify/require"
)

func SendRequest(t *testing.T, cli *http.Client, method string, url string, payload any) (*http.Response, error) {
	req := request.NewRequest(method, url).WithBody(payload)
	require.NoError(t, req.Err())

	return cli.Do(req.Http())
}

func DecodeResponse(t *testing.T, resp *http.Response) (map[string]any) {
	var decoded map[string]any

	err := json.NewDecoder(resp.Body).Decode(&decoded)
	require.NoError(t, err)

	return decoded
}
