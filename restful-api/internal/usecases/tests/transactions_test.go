package tests

import (
	"coinflow/coinflow-server/restful-api/internal/models"
	"coinflow/coinflow-server/restful-api/internal/repository/mocks"
	"coinflow/coinflow-server/restful-api/internal/usecases/service"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestTransactionsService_CommitAndGet(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)

	tsRepo := mocks.NewMockTransactionsRepo(ctrl)
	tsService := service.NewTransactionsService(tsRepo)

	ts := models.Transaction{
		UserId: uuid.New(),
		Type: "buy",
		Target: "Starbucks",
		Description: "Purchased latte",
		Category: "Restaurants",
		Cost: 400,
	}
	id := uuid.New()

	tsRepo.EXPECT().PostTransaction(&ts).Return(id, nil)
	tsRepo.EXPECT().GetTransaction(id).
		DoAndReturn(func(idArg uuid.UUID) (*models.Transaction, error) {
			newTs := ts
			newTs.Id = id
			newTs.Timestamp = time.Now()
			return &newTs, nil
		})

	id, err := tsService.PostTransaction(&ts)
	require.NoError(t, err)

	ret, err := tsService.GetTransaction(id)
	require.NoError(t, err)
	require.Equal(t, id, ret.Id)
	require.Equal(t, ts.Description, ret.Description)
	require.NotNil(t, ret.Timestamp)
}
