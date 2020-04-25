package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm" // not blank
	uuid "github.com/satori/go.uuid"
)

// Account model
type Account struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name      string    `gorm:"size:255"`
	Email     string    `gorm:"type:varchar(100)"`
	Password  string    `gorm:"not null"`
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// BeforeCreate will set a UUID rather than numeric ID.
func (account *Account) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	scope.SetColumn("ID", uuid)
	scope.SetColumn("Active", true)

	return nil
}
