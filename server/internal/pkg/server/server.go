package server

import (
	"github.com/NeoAssist/NeoAcademy/internal/pkg/database"
	accounts "github.com/NeoAssist/NeoAcademy/internal/pkg/database/store"
	handler "github.com/NeoAssist/NeoAcademy/internal/pkg/server/handler"
	router "github.com/NeoAssist/NeoAcademy/internal/pkg/server/router"
)

// RunServer Initialize a new server
func RunServer() {
	r := router.InitializeRoutes()
	v1 := r.Group("/api/v1")

	d := database.New()
	database.AutoMigrate(d)

	as := accounts.NewAccountStore(d)
	h := handler.NewHandler(as)
	h.Register(v1)
	r.Logger.Fatal(r.Start("127.0.0.1:8080"))
}
