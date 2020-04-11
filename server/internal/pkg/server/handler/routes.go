package handler

import (
	middleware "github.com/NeoAssist/NeoAcademy/internal/pkg/server/router/middleware"
	"github.com/labstack/echo/v4"
)

// Register .
func (h *Handler) Register(v1 *echo.Group) {
	//jwtMiddleware := middleware.JWT(middleware.JWTSecret)

	// need to implement user authentication
	// guestUsers := v1.Group("/users")
	// guestUsers.POST("", h.SignUp)
	// guestUsers.POST("/login", h.Login)

	// user := v1.Group("/user", jwtMiddleware)
	// user.GET("", h.CurrentUser)
	// user.PUT("", h.UpdateUser)

	accounts := v1.Group("/accounts", middleware.JWTWithConfig(
		middleware.JWTConfig{
			SigningKey: middleware.JWTSecret,
		},
	))

	accounts.POST("", h.CreateAccount)
	accounts.PUT("", h.UpdateAccount)
}
