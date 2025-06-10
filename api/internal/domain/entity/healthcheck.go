package entity

import "time"

// HealthCheck represents the domain entity for health check
type HealthCheck struct {
	ServerTime time.Time
}

// NewHealthCheck creates a new HealthCheck entity
func NewHealthCheck() *HealthCheck {
	return &HealthCheck{
		ServerTime: time.Now(),
	}
}
