package model

import (
	"time"
)

// Common represents meta data of entity
type Common struct {
	ID         uint64
	CreateTime time.Time
	UpdateTime time.Time
	DeleteTime time.Time
}
