package models

import (
	"github.com/jinzhu/gorm"
)

type BaseModel struct {
	gorm.Model
	ID string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
}
