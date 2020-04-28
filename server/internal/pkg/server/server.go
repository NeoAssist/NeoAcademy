package server

import (
	"fmt"

	"github.com/NeoAssist/NeoAcademy/internal/pkg/database"
	accounts "github.com/NeoAssist/NeoAcademy/internal/pkg/database/store"
	handler "github.com/NeoAssist/NeoAcademy/internal/pkg/server/handler"
	router "github.com/NeoAssist/NeoAcademy/internal/pkg/server/router"
	environment "github.com/NeoAssist/NeoAcademy/internal/utils"
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

	address := fmt.Sprintf("%v:%v", environment.GetEnv("API_HOST"), environment.GetEnv("API_PORT"))

	r.Logger.Fatal(r.Start(address))
}
