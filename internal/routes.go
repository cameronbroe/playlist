package internal

import "github.com/gin-gonic/gin"

func InstallRoutes(engine *gin.Engine) {
	engine.GET("/list", handleList)
}