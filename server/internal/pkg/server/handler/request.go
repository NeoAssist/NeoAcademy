package handler

import (
	model "github.com/NeoAssist/NeoAcademy/internal/pkg/database/models"
	"github.com/labstack/echo/v4"
)

type accountCreateRequest struct {
	Account struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Active   bool   `json:"active"`
	} `json:"account"`
}

type accountUpdateRequest struct {
	Account struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Active   bool   `json:"active"`
	} `json:"account"`
}

func (r *accountCreateRequest) bind(context echo.Context, account *model.Account) error {

	if err := context.Bind(r); err != nil {
		return err
	}
	if err := context.Validate(r); err != nil {
		return err
	}

	account.Email = r.Account.Email
	account.Password = r.Account.Password
	account.Name = r.Account.Name
	account.Active = r.Account.Active

	return nil
}

func (r *accountUpdateRequest) bind(context echo.Context, account *model.Account) error {

	if err := context.Bind(r); err != nil {
		return err
	}
	if err := context.Validate(r); err != nil {
		return err
	}

	account.Email = r.Account.Email
	account.Password = r.Account.Password
	account.Name = r.Account.Name
	account.Active = r.Account.Active

	return nil
}

// type accountUpdateRequest struct {
// 	Account struct {
// 		Username string `json:"username"`
// 		Email    string `json:"email" validate:"email"`
// 		Password string `json:"password"`
// 		Bio      string `json:"bio"`
// 		Image    string `json:"image"`
// 	} `json:"account"`
// }

// func newAccountUpdateRequest() *accountUpdateRequest {
// 	return new(accountUpdateRequest)
// }

// func (r *accountUpdateRequest) populate(u *models.Account) {
// 	r.Account.Username = u.Username
// 	r.Account.Email = u.Email
// 	r.Account.Password = u.Password
// 	if u.Bio != nil {
// 		r.Account.Bio = *u.Bio
// 	}
// 	if u.Image != nil {
// 		r.Account.Image = *u.Image
// 	}
// }

// func (r *accountUpdateRequest) bind(c echo.Context, u *models.Account) error {
// 	if err := c.Bind(r); err != nil {
// 		return err
// 	}
// 	if err := c.Validate(r); err != nil {
// 		return err
// 	}
// 	u.Username = r.Account.Username
// 	u.Email = r.Account.Email
// 	if r.Account.Password != u.Password {
// 		h, err := u.HashPassword(r.Account.Password)
// 		if err != nil {
// 			return err
// 		}
// 		u.Password = h
// 	}
// 	u.Bio = &r.Account.Bio
// 	u.Image = &r.Account.Image
// 	return nil
// }

// type accountRegisterRequest struct {
// 	Account struct {
// 		Username string `json:"username" validate:"required"`
// 		Email    string `json:"email" validate:"required,email"`
// 		Password string `json:"password" validate:"required"`
// 	} `json:"user"`
// }

// func (r *accountRegisterRequest) bind(c echo.Context, u *models.Account) error {
// 	if err := c.Bind(r); err != nil {
// 		return err
// 	}
// 	if err := c.Validate(r); err != nil {
// 		return err
// 	}
// 	u.Username = r.Account.Username
// 	u.Email = r.Account.Email
// 	h, err := u.HashPassword(r.Account.Password)
// 	if err != nil {
// 		return err
// 	}
// 	u.Password = h
// 	return nil
// }
