package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	ID           uuid.UUID  `gorm:"type:uuid;primary_key;"`
	Name         string     `gorm:"size:255"`
	AccountID    Account    `gorm:"foreignkey:ID"`
	AttachmentID Attachment `gorm:"foreignkey:ID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

// BeforeCreate will set a UUID rather than numeric ID.
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())

	return nil
}
