package handler

import (
	"net/http"

	model "github.com/NeoAssist/NeoAcademy/internal/pkg/database/models"
	errors "github.com/NeoAssist/NeoAcademy/internal/pkg/errors"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

// CreateAccount .
func (h *Handler) CreateAccount(context echo.Context) error {
	var account model.Account

	req := &accountCreateRequest{}

	if err := req.bind(context, &account); err != nil {
		return context.JSON(http.StatusUnprocessableEntity, errors.NewError(err))
	}

	err := h.accountStore.Create(&account)

	if err != nil {
		return context.JSON(http.StatusUnprocessableEntity, errors.NewError(err))
	}

	return context.JSON(http.StatusCreated, newAccountCreateResponse(&account))
}

// UpdateAccount .
func (h *Handler) UpdateAccount(context echo.Context) error {
	id, err := uuid.FromString(context.Param("id"))
	account, err := h.accountStore.GetByID(id)

	req := &accountUpdateRequest{}

	if err := req.bind(context, account); err != nil {
		return context.JSON(http.StatusUnprocessableEntity, errors.NewError(err))
	}

	h.accountStore.Update(account)

	if err != nil {
		return context.JSON(http.StatusUnprocessableEntity, errors.NewError(err))
	}

	return context.JSON(http.StatusOK, newAccountUpdateResponse(account))
}
