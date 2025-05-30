package tests

import (
	tu "coinflow/coinflow-server/pkg/testutils"
	"coinflow/coinflow-server/pkg/utils"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var (
	registered = make(map[string]uuid.UUID)
)

func register(t *testing.T, usr tu.Payload) uuid.UUID {
	if val, ok := registered[usr["login"].(string)]; ok {
		return val
	}

	resp, err := tu.SendRequest(t, cli, http.MethodPost, addr + RegisterPath, tu.Payload{
		"usr": usr,
	})

	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, resp.StatusCode)

	decoded := tu.DecodeResponse(t, resp)
	require.Contains(t, decoded, "user_id")

	decVal, ok := decoded["user_id"].(string)
	require.True(t, ok)

	userId, err := uuid.Parse(decVal)
	require.NoError(t, err)

	registered[usr["login"].(string)] = userId

	return userId
}

func getTokensFromResponse(t *testing.T, resp *http.Response) (string, string) {
	decoded := tu.DecodeResponse(t, resp)
	require.Contains(t, decoded, "access_token")
	require.Contains(t, decoded, "refresh_token")

	accToken, ok := decoded["access_token"].(string)
	require.True(t, ok)
	refrToken, ok := decoded["refresh_token"].(string)
	require.True(t, ok)

	require.NoError(t, utils.CheckJwtFormat(accToken))
	require.NoError(t, utils.CheckJwtFormat(refrToken))

	return accToken, refrToken
}

func login(t *testing.T, usr tu.Payload) (string, string) {
	resp, err := tu.SendRequest(t, cli, http.MethodPost, addr + LoginPath, tu.Payload{
		"login": usr["login"].(string),
		"password": usr["password"].(string),
	})

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	return getTokensFromResponse(t, resp)
}

func refresh(t *testing.T, token string) (string, string) {
	resp, err := tu.SendRequest(t, cli, http.MethodPost, addr + RefreshPath, tu.Payload{
		"refresh_token": token,
	})

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	return getTokensFromResponse(t, resp)
}

// Tests ---------------------------------------------------

func TestAuth_Register(t *testing.T) {
	register(t, exampleUser)
}

func TestAuth_Login(t *testing.T) {
	login(t, exampleUser)
}

func TestAuth_Refresh(t *testing.T) {
	_, ref := login(t, exampleUser)
	refresh(t, ref)
}

func TestAuth_DoubleRefresh(t *testing.T) {
	_, ref := login(t, exampleUser)
	refresh(t, ref)

	resp, err := tu.SendRequest(t, cli, http.MethodPost, addr + RefreshPath, tu.Payload{
		"refresh_token": ref,
	})

	require.NoError(t, err)
	require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}
