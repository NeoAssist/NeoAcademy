package models

import (
	"time"
)

// BaseModel .
type BaseModel struct {
	ID        []byte
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
