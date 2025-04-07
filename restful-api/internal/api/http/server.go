package http

import (
	"coinflow/coinflow-server/restful-api/internal/api/http/types"
	"coinflow/coinflow-server/restful-api/internal/usecases"
	"net/http"
)

type CoinflowServer struct {
    tsService usecases.TransactionsService
}

func NewCoinflowServer(tsService usecases.TransactionsService) *CoinflowServer {
    return &CoinflowServer{tsService: tsService}
}

func (s *CoinflowServer) GetTransactionHandler(w http.ResponseWriter, r *http.Request) {
    reqObj, err := types.CreateGetTransactionRequestObject(r)
    if err != nil {
        WriteError(w, err)
    }

    res, err := s.tsService.GetTransaction(reqObj.TsId)
    if err != nil {
        WriteError(w, err)
    }

    resp, err := types.CreateGetTransactionResponse(res)
    if err != nil {
        WriteError(w, err)
    }

    WriteResponse(w, http.StatusOK, resp)
}

func (s *CoinflowServer) PostTransactionHandler(w http.ResponseWriter, r *http.Request) {
    reqObj, err := types.CreatePostTransactionRequestObject(r)
    if err != nil {
        WriteError(w, err)
    }

    res, err := s.tsService.PostTransaction(reqObj.Ts)
    if err != nil {
        WriteError(w, err)
    }

    resp, err := types.CreatePostTransactionResponse(res)
    if err != nil {
        WriteError(w, err)
    }

    WriteResponse(w, http.StatusCreated, resp)
}
