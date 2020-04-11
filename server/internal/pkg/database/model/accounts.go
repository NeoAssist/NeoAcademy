package models

import (
	_ "github.com/jinzhu/gorm" // not blank
	_ "github.com/satori/go.uuid"
)

// Account model
type Account struct {
	BaseModel
	Name     string `gorm:"size:255"`
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password []byte
	Active   bool
}
