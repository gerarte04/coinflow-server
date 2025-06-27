package tests

import (
	tu "coinflow/coinflow-server/pkg/testutils"
	"coinflow/coinflow-server/pkg/vars"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var (
	clfTimeout = time.Second * 15
	clfPeriod = time.Millisecond * 2500
)

func commitTx(t *testing.T, payload tu.Payload) (*http.Response, uuid.UUID) {
	url := fmt.Sprintf("%s%s", addr, CommitPath)
	resp, err := tu.SendRequest(t, cli, http.MethodPost, url, payload)
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, resp.StatusCode)

	decoded := tu.DecodeResponse(t, resp)
	require.Contains(t, decoded, "tx_id")

	decVal, ok := decoded["tx_id"].(string)
	require.True(t, ok)

	txId, err := uuid.Parse(decVal)
	require.NoError(t, err)

	return resp, txId
}

func getTxById(t *testing.T, txId string) (*http.Response, tu.Payload) {
	url := fmt.Sprintf("%s%s/%s", addr, TransactionPath, txId)
	resp, err := tu.SendRequest(t, cli, http.MethodGet, url, tu.Payload{})
	require.NoError(t, err)

	if resp.StatusCode != http.StatusOK {
		return resp, nil
	}

	return resp, tu.DecodeResponse(t, resp)
}

func getInPeriod(t *testing.T, begin time.Time, end time.Time) (*http.Response, []tu.Payload) {
	url := fmt.Sprintf("%s%s", addr, TransactionsInPeriodPath)
	resp, err := tu.SendRequest(t, cli, http.MethodPost, url, tu.Payload {
		"begin": begin.UTC().Format(vars.TimeLayout),
		"end": end.UTC().Format(vars.TimeLayout),
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

// Tests ---------------------------------------------------

func TestTransactions_CommitWithoutAutoCategory(t *testing.T) {
	tx := tu.GetPayloadCopy(t, exampleTx)

	_, txId := commitTx(t, tu.Payload{
		"tx": tx,
		"with_auto_category": false,
	})

	resp, decoded := getTxById(t, txId.String())
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Contains(t, decoded, "tx")

	gotTx, ok := decoded["tx"].(map[string]any)
	require.True(t, ok)

	tu.ValidateResult(t, gotTx, exampleTx,
		tu.ValidateOpt{Key: "timestamp", CheckValue: false},
		tu.ValidateOpt{Key: "id", CheckValue: true, Value: txId.String()},
	)
}

func TestTransactions_CommitWithAutoCategory(t *testing.T) {
	tx := tu.GetPayloadCopy(t, exampleTx)
	delete(tx, "category")

	_, txId := commitTx(t, tu.Payload{
		"tx": tx,
		"with_auto_category": true,
	})

	var resp *http.Response
	var gotTx tu.Payload
	endTime := time.Now().Add(clfTimeout)
	
	for time.Now().Before(endTime) {
		var decoded tu.Payload
		resp, decoded = getTxById(t, txId.String())
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Contains(t, decoded, "tx")

		var ok bool
		gotTx, ok = decoded["tx"].(map[string]any)
		require.True(t, ok)
		require.Contains(t, gotTx, "category")
		
		if gotTx["category"] != "other" {
			break
		}

		time.Sleep(clfPeriod)
	}

	tu.ValidateResult(t, gotTx, exampleTx,
		tu.ValidateOpt{Key: "timestamp", CheckValue: false},
		tu.ValidateOpt{Key: "id", CheckValue: true, Value: txId.String()},
		tu.ValidateOpt{Key: "category", CheckValue: true, Value: exampleTx["category"]},
	)
}

func TestTransactions_GetTxInPeriod(t *testing.T) {
	payload := tu.Payload{
		"tx": exampleTx,
		"with_auto_category": true,
	}

	sz := 6
	beginIdx, endIdx := 1, 4

	var beginTime, endTime time.Time
	ids := make([]string, 0)

	for i := 0; i < sz; i++ {
		startTime := time.Now()

		if i == beginIdx {
			beginTime = time.Now()
		} else if i == endIdx {
			endTime = time.Now()
		} else {
			_, id := commitTx(t, payload)
			
			if i > beginIdx && i < endIdx {
				ids = append(ids, id.String())
			}
		}

		time.Sleep(time.Second - time.Since(startTime))
	}

	resp, txs := getInPeriod(t, beginTime, endTime)

	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, 2, len(txs))

	for i, tx := range txs {
		tu.ValidateResult(t, tx, exampleTx,
			tu.ValidateOpt{Key: "timestamp", CheckValue: false},
			tu.ValidateOpt{Key: "id", CheckValue: true, Value: ids[i]},
			tu.ValidateOpt{Key: "category", Ignore: true},
		)
	}
}

func TestTransactions_WrongId(t *testing.T) {
	commitTx(t, tu.Payload{
		"tx": exampleTx,
		"with_auto_category": false,
	})

	resp, _ := getTxById(t, uuid.New().String())
	require.Equal(t, resp.StatusCode, http.StatusNotFound)
}

func TestTransactions_InvalidId(t *testing.T) {
	commitTx(t, tu.Payload{
		"tx": exampleTx,
		"with_auto_category": false,
	})

	resp, _ := getTxById(t, exampleInvalidId)
	require.Equal(t, resp.StatusCode, http.StatusBadRequest)
}
