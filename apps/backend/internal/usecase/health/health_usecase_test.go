package health

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/your-org/fullstack-template/apps/backend/internal/constants"
)

func TestHealthUsecaseGet(t *testing.T) {
	uc := NewHealthUsecase(constants.TestHealthServiceName, constants.TestHealthEnvironment)

	status, err := uc.Get(context.Background())

	require.NoError(t, err)
	require.Equal(t, constants.TestHealthStatus, status.Status)
	require.Equal(t, constants.TestHealthServiceName, status.Service)
	require.Equal(t, constants.TestHealthEnvironment, status.Environment)
	require.False(t, status.Timestamp.IsZero())
}
