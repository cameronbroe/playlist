package internal

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func NewListHandler(db *Database) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		log.Println("getting list of songs in database")
		playedSongs, err := db.GetListOfPlayedSongs()
		if err != nil {
			ctx.JSON(500, map[string]string{
				"error": fmt.Sprintf("%s", err),
			})
		}
		ctx.JSON(200, playedSongs)
	}
}
