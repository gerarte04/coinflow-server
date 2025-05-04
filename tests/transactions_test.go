package tests

import (
	"coinflow/coinflow-server/pkg/testutils"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/require"
)

const (
	CommitPath = "/commit"
	TransactionPath = "/transaction/id"
)

var (
	addr = fmt.Sprintf("http://%s:%s", os.Getenv("HTTP_HOST"), os.Getenv("HTTP_PORT"))

	exampleTxPayload = map[string]any{
		"user_id": uuid.New().String(),
		"type": "purchase",
		"target": "Coffee Point",
		"description": "Purchased latte and croissant",
		"category": "food",
		"cost": float64(400),
	}

	exampleInvalidId = "abcdefgh123"

	clfTimeout = time.Second * 15
	clfPeriod = time.Millisecond * 2500
)

func commitPayload(t *testing.T, payload any) (*http.Response, uuid.UUID) {
	url := fmt.Sprintf("%s%s", addr, CommitPath)
	resp, err := testutils.SendRequest(t, http.MethodPost, url, payload)
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, http.StatusCreated)

	decoded := testutils.DecodeResponse(t, resp)
	require.Contains(t, decoded, "tx_id")

	txId, err := uuid.Parse(decoded["tx_id"].(string))
	require.NoError(t, err)

	return resp, txId
}

func getById(t *testing.T, txId string) (*http.Response, map[string]any) {
	url := fmt.Sprintf("%s%s/%s", addr, TransactionPath, txId)
	resp, err := testutils.SendRequest(t, http.MethodGet, url, nil)
	require.NoError(t, err)

	if resp.StatusCode != http.StatusOK {
		return resp, nil
	}

	return resp, testutils.DecodeResponse(t, resp)
}

type ValidateOpt struct {
	Key string
	CheckValue bool
	Value any
}

func validateResult(t *testing.T, res map[string]any, payload map[string]any, opts ...ValidateOpt) {
	for k, v := range payload {
		require.Contains(t, res, k)
		require.Equal(t, v, res[k])
	}

	for _, opt := range opts {
		require.Contains(t, res, opt.Key)

		if opt.CheckValue {
			require.Equal(t, opt.Value, res[opt.Key])
		}
	}
}

func TestTransactions_CommitWithoutAutoCategory(t *testing.T) {
	var payload map[string]any
	require.NoError(t, copier.CopyWithOption(&payload, &exampleTxPayload, copier.Option{DeepCopy: true}))
	payload["with_auto_category"] = false

	_, txId := commitPayload(t, payload)

	resp, decoded := getById(t, txId.String())
	require.Equal(t, resp.StatusCode, http.StatusOK)

	validateResult(t, decoded, exampleTxPayload,
		ValidateOpt{Key: "timestamp", CheckValue: false},
		ValidateOpt{Key: "id", CheckValue: true, Value: txId.String()},
	)
}

func TestTransactions_CommitWithAutoCategory(t *testing.T) {
	var payload map[string]any
	require.NoError(t, copier.CopyWithOption(&payload, &exampleTxPayload, copier.Option{DeepCopy: true}))
	payload["with_auto_category"] = true
	delete(payload, "category")

	_, txId := commitPayload(t, payload)

	var resp *http.Response
	var decoded map[string]any
	endTime := time.Now().Add(clfTimeout)
	
	for time.Now().Before(endTime) {
		resp, decoded = getById(t, txId.String())
		require.Equal(t, resp.StatusCode, http.StatusOK)
		require.Contains(t, decoded, "category")
		
		if decoded["category"] != "other" {
			break
		}

		time.Sleep(clfPeriod)
	}

	validateResult(t, decoded, exampleTxPayload,
		ValidateOpt{Key: "timestamp", CheckValue: false},
		ValidateOpt{Key: "id", CheckValue: true, Value: txId.String()},
		ValidateOpt{Key: "category", CheckValue: true, Value: exampleTxPayload["category"]},
	)
}

func TestTransactions_WrongId(t *testing.T) {
	commitPayload(t, exampleTxPayload)

	resp, _ := getById(t, uuid.New().String())
	require.Equal(t, resp.StatusCode, http.StatusNotFound)
}

func TestTransactions_InvalidId(t *testing.T) {
	commitPayload(t, exampleTxPayload)

	resp, _ := getById(t, exampleInvalidId)
	require.Equal(t, resp.StatusCode, http.StatusBadRequest)
}
