package interfaces

import (
	"github.com/NeoAssist/NeoAcademy/internal/pkg/database/models"
	uuid "github.com/satori/go.uuid"
)

type AttachmentsInterface interface {
	GetByID(uuid.UUID) (models.Attachment, error)
	Create(models.Attachment) error
}
