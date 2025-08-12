package tests

import (
	tu "coinflow/coinflow-server/pkg/testutils"
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

	resp, err := tu.SendRequest(t, cli, http.MethodPost, addr + RegisterPath, usr)

	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, resp.StatusCode)

	decoded := tu.DecodeResponse(t, resp)
	tu.ValidateResult(t, decoded, usr, 
		tu.ValidateOpt{Key: "registration_timestamp", CheckValue: false},
		tu.ValidateOpt{Key: "id", CheckValue: false},
		tu.ValidateOpt{Key: "password", Ignore: true},
	)

	userId, err := uuid.Parse(decoded["id"].(string))
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

	// require.NoError(t, utils.CheckJwtFormat(accToken))
	// require.NoError(t, utils.CheckJwtFormat(refrToken))

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
