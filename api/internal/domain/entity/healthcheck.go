package entity

import "time"

// HealthCheck はヘルスチェックのドメインエンティティを表します
type HealthCheck struct {
	ServerTime time.Time
}

// NewHealthCheck は新しいHealthCheckエンティティを作成します
func NewHealthCheck() *HealthCheck {
	return &HealthCheck{
		ServerTime: time.Now(),
	}
}
