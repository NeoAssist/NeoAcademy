package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// User model
type Account struct {
	gorm.Model
	Name           string `gorm:"size:255"`
	Email          string `gorm:"type:varchar(100);unique_index"`
	HashedPassword []byte
	Active         bool
}

func SayHi() {
	fmt.Println("Hi Fucker")
}
