package interfaces

import (
	model "github.com/NeoAssist/NeoAcademy/internal/pkg/database/model"
)

// Store .
type Store interface {
	GetByID(uint) (*model.Account, error)
	GetByEmail(string) (*model.Account, error)
	GetByUsername(string) (*model.Account, error)
	Create(*model.Account) error
	Update(*model.Account) error
}
