package handler

import (
	accounts "github.com/NeoAssist/NeoAcademy/internal/pkg/database/interface"
)

// Handler .
type Handler struct {
	accountStore accounts.Store
}

// NewHandler Handle the requests
func NewHandler(as accounts.Store) *Handler {
	return &Handler{
		accountStore: as,
	}
}
