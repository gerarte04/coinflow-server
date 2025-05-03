// package tests

// import (
// 	pkgConfig "coinflow/coinflow-server/pkg/config"
// 	"coinflow/coinflow-server/storage-service/internal/models"
// 	"coinflow/coinflow-server/storage-service/internal/repository/mocks"
// 	"coinflow/coinflow-server/storage-service/internal/usecases/service"
// 	"testing"
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/require"
// 	"go.uber.org/mock/gomock"
// )

// func TestTransactionsService_CommitAndGet(t *testing.T) {
// 	t.Parallel()

// 	ctrl := gomock.NewController(t)

// 	txRepo := mocks.NewMockTransactionsRepo(ctrl)
// 	txService, err := service.NewTransactionsService(txRepo, pkgConfig.GrpcConfig{})
// 	require.NoError(t, err)

// 	id := uuid.New()
// 	tx := models.Transaction{
// 		UserId: uuid.New(),
// 		Type: "buy",
// 		Target: "Starbucks",
// 		Description: "Purchased latte",
// 		Category: "Restaurants",
// 		Cost: 400,
// 	}

// 	txRepo.EXPECT().PostTransaction(&tx).Return(id, nil)
// 	txRepo.EXPECT().GetTransaction(id).
// 		DoAndReturn(func(idArg uuid.UUID) (*models.Transaction, error) {
// 			newTx := tx
// 			newTx.Id = id
// 			newTx.Timestamp = time.Now()
// 			return &newTx, nil
// 		})

// 	id, err = txService.PostTransaction(&tx)
// 	require.NoError(t, err)

//		ret, err := txService.GetTransaction(id)
//		require.NoError(t, err)
//		require.Equal(t, id, ret.Id)
//		require.Equal(t, tx.Description, ret.Description)
//		require.NotNil(t, ret.Timestamp)
//	}
package tests