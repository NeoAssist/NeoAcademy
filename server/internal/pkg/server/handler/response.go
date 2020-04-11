package handler

// import (
// 	model "github.com/NeoAssist/NeoAcademy/internal/pkg/database/model"
// 	middleware "github.com/NeoAssist/NeoAcademy/internal/pkg/server/router/middleware"
// )

// type accountResponse struct {
// 	Account struct {
// 		Username string  `json:"username"`
// 		Email    string  `json:"email"`
// 		Bio      *string `json:"bio"`
// 		Image    *string `json:"image"`
// 		Token    string  `json:"token"`
// 	} `json:"user"`
// }

// func newAccountResponse(u *model.Account) *accountResponse {
// 	r := new(accountResponse)
// 	r.User.Username = u.Username
// 	r.User.Email = u.Email
// 	r.User.Bio = u.Bio
// 	r.User.Image = u.Image
// 	r.User.Token = middleware.GenerateJWT(u.ID)
// 	return r
// }
