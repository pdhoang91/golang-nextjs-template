package health

import "time"

type HealthStatus struct {
	Status      string
	Service     string
	Environment string
	Timestamp   time.Time
}
