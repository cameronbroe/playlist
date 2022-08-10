package main

import (
	"github.com/cameronbroe/music.cameronbroe.com/internal"
	"github.com/gin-gonic/gin"
)

type App struct {
	server *gin.Engine
}

func buildApp() *App {
	app := new(App)

	app.server = gin.Default()
	internal.InstallRoutes(app.server)

	return app
}

func (app *App) Run() {
	app.server.Run()
}

func main() {
	app := buildApp()

	app.Run()
}
