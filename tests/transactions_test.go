package tests

import (
	tu "coinflow/coinflow-server/pkg/testutils"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

const (
	CommitPath = "/commit"
	TransactionPath = "/transaction/id"
	TransactionsInPeriodPath = "/transaction/period"
)

var (
	addr = fmt.Sprintf("http://%s:%s", os.Getenv("HTTP_HOST"), os.Getenv("HTTP_PORT"))

	exampleTx = tu.Payload{
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

	TimeLayout = "02/01/2006 15:04:05 -0700"
)

func commitTx(t *testing.T, payload tu.Payload) (*http.Response, uuid.UUID) {
	url := fmt.Sprintf("%s%s", addr, CommitPath)
	resp, err := tu.SendRequest(t, http.MethodPost, url, payload)
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, http.StatusCreated)

	decoded := tu.DecodeResponse(t, resp)
	require.Contains(t, decoded, "tx_id")

	decVal, ok := decoded["tx_id"].(string)
	require.True(t, ok)

	txId, err := uuid.Parse(decVal)
	require.NoError(t, err)

	return resp, txId
}

func getById(t *testing.T, txId string) (*http.Response, tu.Payload) {
	url := fmt.Sprintf("%s%s/%s", addr, TransactionPath, txId)
	resp, err := tu.SendRequest(t, http.MethodGet, url, nil)
	require.NoError(t, err)

	if resp.StatusCode != http.StatusOK {
		return resp, nil
	}

	return resp, tu.DecodeResponse(t, resp)
}

func getInPeriod(t *testing.T, begin time.Time, end time.Time) (*http.Response, []tu.Payload) {
	url := fmt.Sprintf("%s%s", addr, TransactionsInPeriodPath)
	resp, err := tu.SendRequest(t, http.MethodPost, url, tu.Payload {
		"begin": begin.UTC().Format(TimeLayout),
		"end": end.UTC().Format(TimeLayout),
	})
	require.NoError(t, err)

	if resp.StatusCode != http.StatusOK {
		return resp, nil
	}

	decoded := tu.DecodeResponse(t, resp)
	require.Contains(t, decoded, "txs")

	list, ok := decoded["txs"].([]any)
	require.True(t, ok)

	txs := make([]tu.Payload, 0)

	for _, v := range list {
		tx, ok := v.(map[string]any)
		require.True(t, ok)

		txs = append(txs, tx)
	}

	return resp, txs
}



func TestTransactions_CommitWithoutAutoCategory(t *testing.T) {
	payload := tu.GetPayloadCopy(t, exampleTx)
	payload["with_auto_category"] = false

	_, txId := commitTx(t, payload)

	resp, decoded := getById(t, txId.String())
	require.Equal(t, resp.StatusCode, http.StatusOK)

	tu.ValidateResult(t, decoded, exampleTx,
		tu.ValidateOpt{Key: "timestamp", CheckValue: false},
		tu.ValidateOpt{Key: "id", CheckValue: true, Value: txId.String()},
	)
}

func TestTransactions_CommitWithAutoCategory(t *testing.T) {
	payload := tu.GetPayloadCopy(t, exampleTx)
	payload["with_auto_category"] = true
	delete(payload, "category")

	_, txId := commitTx(t, payload)

	var resp *http.Response
	var decoded tu.Payload
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

	tu.ValidateResult(t, decoded, exampleTx,
		tu.ValidateOpt{Key: "timestamp", CheckValue: false},
		tu.ValidateOpt{Key: "id", CheckValue: true, Value: txId.String()},
		tu.ValidateOpt{Key: "category", CheckValue: true, Value: exampleTx["category"]},
	)
}

func TestTransactions_GetTxInPeriod(t *testing.T) {
	offset := time.Millisecond * 500

	beginTime := time.Now()
	endTime := time.Now().Add(time.Millisecond * 3000 + offset)

	ids := make([]string, 0)

	for time.Now().Before(endTime) {
		startIterTime := time.Now()

		_, txId := commitTx(t, exampleTx)
		ids = append(ids, txId.String())

		time.Sleep(time.Second - time.Now().Sub(startIterTime))
	}

	resp, txs := getInPeriod(t, beginTime.Add(offset), endTime.Add(-2 * offset))

	require.Equal(t, resp.StatusCode, http.StatusOK)
	require.Equal(t, len(txs), 2)

	for i, tx := range txs {
		tu.ValidateResult(t, tx, exampleTx,
			tu.ValidateOpt{Key: "timestamp", CheckValue: false},
			tu.ValidateOpt{Key: "id", CheckValue: true, Value: ids[i + 1]},
			tu.ValidateOpt{Key: "category", Ignore: true},
		)
	}
}

func TestTransactions_WrongId(t *testing.T) {
	commitTx(t, exampleTx)

	resp, _ := getById(t, uuid.New().String())
	require.Equal(t, resp.StatusCode, http.StatusNotFound)
}

func TestTransactions_InvalidId(t *testing.T) {
	commitTx(t, exampleTx)

	resp, _ := getById(t, exampleInvalidId)
	require.Equal(t, resp.StatusCode, http.StatusBadRequest)
}
