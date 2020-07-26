package handler

import (
	model "github.com/NeoAssist/NeoAcademy/internal/pkg/database/models"
)

type accountResponse struct {
	Account struct {
		Name   string `json:"username"`
		Email  string `json:"email"`
		Active bool   `json:"active"`
	} `json:"account"`
}

func newAccountCreateResponse(account *model.Account) *accountResponse {
	r := new(accountResponse)
	r.Account.Name = account.Name
	r.Account.Email = account.Email
	r.Account.Active = account.Active

	return r
}

func newAccountUpdateResponse(account *model.Account) *accountResponse {
	r := new(accountResponse)
	r.Account.Name = account.Name
	r.Account.Email = account.Email
	r.Account.Active = account.Active

	return r
}
