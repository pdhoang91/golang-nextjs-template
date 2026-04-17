package dto

import (
	"time"

	domain "github.com/your-org/fullstack-template/apps/backend/internal/domain/health"
)

type HealthResponse struct {
	Status      string    `json:"status"`
	Service     string    `json:"service"`
	Environment string    `json:"environment"`
	Timestamp   time.Time `json:"timestamp"`
}

func ToHealthResponse(status domain.HealthStatus) HealthResponse {
	return HealthResponse{
		Status:      status.Status,
		Service:     status.Service,
		Environment: status.Environment,
		Timestamp:   status.Timestamp,
	}
}
