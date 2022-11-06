package internal

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewSubmitHandler(db *Database, decorator *SongDecorator) func(ctx *gin.Context) {
  return func(ctx *gin.Context) {
    log.Println("submitting a new song to the database")
    var newSong PlayedSong
    ctx.BindJSON(&newSong)
		log.Println("decorating song with streaming service links")
		decorator.DecoratePlayedSong(&newSong)
    log.Printf("saving new song to database: %+v\n", newSong)
    err := db.SubmitPlayedSong(newSong)
    if err != nil {
    	ctx.JSON(http.StatusInternalServerError, map[string]string{
    		"error": fmt.Sprintf("%s", err),
    	})
    } else {
    	ctx.JSON(http.StatusOK, map[string]string{
    		"status": "success",
    	})
    }
  }
}
