package tests

import (
	tu "coinflow/coinflow-server/pkg/testutils"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func getUserById(t *testing.T, userId string) (*http.Response, tu.Payload) {
	url := fmt.Sprintf("%s%s/%s", addr, GetUserDataPath, userId)
	resp, err := tu.SendRequest(t, cli, http.MethodGet, url, tu.Payload{})
	require.NoError(t, err)

	if resp.StatusCode != http.StatusOK {
		return resp, nil
	}

	return resp, tu.DecodeResponse(t, resp)
}

// Tests ---------------------------------------------------

func TestUsers_GetUserData(t *testing.T) {
	userId := register(t, exampleUser)
	login(t, exampleUser)

	resp, decoded := getUserById(t, userId.String())
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Contains(t, decoded, "usr")

	gotUsr, ok := decoded["usr"].(map[string]any)
	require.True(t, ok)

	tu.ValidateResult(t, gotUsr, exampleUser,
		tu.ValidateOpt{Key: "password", Ignore: true},
		tu.ValidateOpt{Key: "id", Value: userId.String(), CheckValue: true},
		tu.ValidateOpt{Key: "registration_timestamp", CheckValue: false},
	)
}

func TestUsers_WrongId(t *testing.T) {
	register(t, exampleUser)
	login(t, exampleUser)

	resp, _ := getUserById(t, uuid.New().String())
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestUsers_InvalidId(t *testing.T) {
	register(t, exampleUser)
	login(t, exampleUser)

	resp, _ := getUserById(t, exampleInvalidId)
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
