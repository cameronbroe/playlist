package internal

import (
	"github.com/gin-gonic/gin"
)

func InstallRoutes(engine *gin.Engine, db *Database, decorator *SongDecorator) {
	engine.GET("/list", NewListHandler(db))
	engine.POST("/submit", NewSubmitHandler(db, decorator))
}
