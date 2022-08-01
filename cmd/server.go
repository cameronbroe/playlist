package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type App struct {
	server *gin.Engine
}

func buildApp() *App {
	app := new(App)

	app.server = gin.Default()

	app.server.GET("/list", func(c *gin.Context) {
		fmt.Println("Got a request to /list")
	})

	return app
}

func (app *App) Run() {
	app.server.Run()
}

func main() {
	app := buildApp()

	app.Run()
}
