package tests

import (
	"coinflow/coinflow-server/pkg/testutils"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

const (
    CommitPath = "/commit"
    TransactionPath = "/transaction"
)

var (
    addr = fmt.Sprintf("http://%s:%s", os.Getenv("HTTP_HOST"), os.Getenv("HTTP_PORT"))

    exampleTsPayload = map[string]any{
        "user_id": uuid.New().String(),
        "type": "buy",
        "target": "Starbucks",
        "description": "Purchased latte",
        "category": "Restaurants",
        "cost": float64(400),
    }

    exampleInvalidId = "abcdefgh123"
)

func commitPayload(t *testing.T, payload any) (*http.Response, uuid.UUID) {
    url := fmt.Sprintf("%s%s", addr, CommitPath)
    resp, err := testutils.SendRequest(t, http.MethodPost, url, payload)
    require.NoError(t, err)
    require.Equal(t, resp.StatusCode, http.StatusCreated)

    decoded := testutils.DecodeResponse(t, resp)
    require.Contains(t, decoded, "ts_id")

    tsId, err := uuid.Parse(decoded["ts_id"].(string))
    require.NoError(t, err)

    return resp, tsId
}

func getById(t *testing.T, tsId string) (*http.Response, map[string]any) {
    url := fmt.Sprintf("%s%s/%s", addr, TransactionPath, tsId)
    resp, err := testutils.SendRequest(t, http.MethodGet, url, nil)
    require.NoError(t, err)

    if resp.StatusCode != http.StatusOK {
        return resp, nil
    }

    return resp, testutils.DecodeResponse(t, resp)
}

func TestTransactions_CommitAndGet(t *testing.T) {
    _, tsId := commitPayload(t, exampleTsPayload)

    resp, decoded := getById(t, tsId.String())
    require.Equal(t, resp.StatusCode, http.StatusOK)

    for k, v := range exampleTsPayload {
        require.Contains(t, decoded, k)
        require.Equal(t, v, decoded[k])
    }

    require.Contains(t, decoded, "timestamp")
    require.Contains(t, decoded, "id")
    require.Equal(t, tsId.String(), decoded["id"])
}

func TestTransactions_WrongId(t *testing.T) {
    commitPayload(t, exampleTsPayload)

    resp, _ := getById(t, uuid.New().String())
    require.Equal(t, resp.StatusCode, http.StatusNotFound)
}

func TestTransactions_InvalidId(t *testing.T) {
    commitPayload(t, exampleTsPayload)

    resp, _ := getById(t, exampleInvalidId)
    require.Equal(t, resp.StatusCode, http.StatusBadRequest)
}
