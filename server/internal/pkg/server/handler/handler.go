package handler

import "github.com/NeoAssist/NeoAcademy/internal/pkg/database/interfaces"

// Handler .
type Handler struct {
	accountInterface    interfaces.AccountsInterface
	userInterface       interfaces.UsersInterface
	attachmentInterface interfaces.AttachmentsInterface
}

// NewHandler Handle the requests
func NewHandler(accountInterface interfaces.AccountsInterface, userInterface interfaces.UsersInterface, attachmentInterface interfaces.AttachmentsInterface) *Handler {
	return &Handler{
		accountInterface:    accountInterface,
		userInterface:       userInterface,
		attachmentInterface: attachmentInterface,
	}
}
