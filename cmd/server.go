package main

import (
	"log"

	"github.com/cameronbroe/music.cameronbroe.com/internal"
	"github.com/gin-gonic/gin"
)

type App struct {
	server    *gin.Engine
	db        *internal.Database
	decorator *internal.SongDecorator
}

func buildApp() *App {
	app := new(App)

	app.decorator = new(internal.SongDecorator)

	app.db = internal.InitializeDatabase()
	err := app.db.EnsureDatabaseExists()
	if err != nil {
		panic(err)
	}
	log.Printf("database existence has been ensured")

	app.server = gin.Default()
	internal.InstallRoutes(app.server, app.db, app.decorator)

	return app
}

func (app *App) Run() {
	app.server.Run()
}

func main() {
	app := buildApp()

	app.Run()
}
