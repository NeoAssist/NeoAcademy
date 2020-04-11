package interfaces

import (
	model "github.com/NeoAssist/NeoAcademy/internal/pkg/database/model"
	uuid "github.com/satori/go.uuid"
)

// Store .
type Store interface {
	GetByID(uuid.UUID) (*model.Account, error)
	GetByEmail(string) (*model.Account, error)
	GetByUsername(string) (*model.Account, error)
	Create(*model.Account) error
	Update(*model.Account) error
}
