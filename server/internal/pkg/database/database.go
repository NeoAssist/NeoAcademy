package database

import (
	"fmt"

	model "github.com/NeoAssist/NeoAcademy/internal/pkg/database/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // indirect
)

// New .
func New() *gorm.DB {
	db, err := gorm.Open("postgres", "host=10.10.10.2 port=5432 user=neoacademy dbname=neoacademy password=neoacademy sslmode=disable")
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db
}

// AutoMigrate .
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.Account{},
	)
}
