package interfaces

import (
	"github.com/NeoAssist/NeoAcademy/internal/pkg/database/models"
	uuid "github.com/satori/go.uuid"
)

type UsersInterface interface {
	GetByID(uuid.UUID) (models.User, error)
	GetByAccountID(uuid.UUID) (models.User, error)
	Create(models.User) error
	Update(models.User) error
}
