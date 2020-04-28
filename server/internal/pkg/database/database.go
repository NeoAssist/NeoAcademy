package database

import (
	"fmt"

	model "github.com/NeoAssist/NeoAcademy/internal/pkg/database/model"
	environment "github.com/NeoAssist/NeoAcademy/internal/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // indirect
)

// New .
func New() *gorm.DB {
	dbHost := environment.GetEnv("DATABASE_HOST")
	dbPort := environment.GetEnv("DATABASE_PORT")
	dbUser := environment.GetEnv("DATABASE_USER")
	dbName := environment.GetEnv("DATABASE_NAME")
	dbPassword := environment.GetEnv("DATABASE_PASSWORD")

	connectionString := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword)

	db, err := gorm.Open("postgres", connectionString)
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
