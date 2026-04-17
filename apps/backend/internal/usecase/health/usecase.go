package health

import (
	"context"
	"time"

	domain "github.com/your-org/fullstack-template/apps/backend/internal/domain/health"
)

type HealthUsecase interface {
	Get(ctx context.Context) (domain.HealthStatus, error)
}

type usecase struct {
	serviceName string
	environment string
}

func NewHealthUsecase(serviceName, environment string) HealthUsecase {
	return &usecase{
		serviceName: serviceName,
		environment: environment,
	}
}

func (u *usecase) Get(_ context.Context) (domain.HealthStatus, error) {
	return domain.HealthStatus{
		Status:      domain.StatusOK,
		Service:     u.serviceName,
		Environment: u.environment,
		Timestamp:   time.Now().UTC(),
	}, nil
}
