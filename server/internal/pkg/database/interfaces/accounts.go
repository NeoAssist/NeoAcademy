package interfaces

import (
	"github.com/NeoAssist/NeoAcademy/internal/pkg/database/models"
	uuid "github.com/satori/go.uuid"
)

type AccountsInterface interface {
	GetByID(uuid.UUID) (models.Account, error)
	GetByEmail(string) (models.Account, error)
	Create(models.Account) error
	Update(models.Account) error
	Delete(models.Account) error
	Deactivate(models.Account) error
	Activate(models.Account) error
}
