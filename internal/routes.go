package internal

import (
	"github.com/gin-gonic/gin"
)

func InstallRoutes(engine *gin.Engine, db *Database) {
	engine.GET("/list", NewListHandler(db))
}
