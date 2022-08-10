package internal

import (
	"github.com/cameronbroe/music.cameronbroe.com/internal/route_handlers"
	"github.com/gin-gonic/gin"
)

func InstallRoutes(engine *gin.Engine) {
	engine.GET("/list", route_handlers.NewListHandler())
}
