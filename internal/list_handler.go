package internal

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewListHandler(db *Database) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		log.Println("getting list of songs in database")
		playedSongs, err := db.GetListOfPlayedSongs()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]string{
				"error": fmt.Sprintf("%s", err),
			})
		}
		ctx.JSON(http.StatusOK, playedSongs)
	}
}
