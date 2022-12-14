package internal

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func validateSong(song *PlayedSong) bool {
	return !(len(song.Artist) == 0 || len(song.Album) == 0 || len(song.Title) == 0)
}

func NewSubmitHandler(db *Database, decorator *SongDecorator) func(ctx *gin.Context) {
	staticApiSecret := os.Getenv("API_SECRET")

	return func(ctx *gin.Context) {
		if staticApiSecret != "" {
			bearerToken := ctx.GetHeader("Authorization")
			if bearerToken != fmt.Sprintf("Bearer %s", staticApiSecret) {
				ctx.JSON(http.StatusUnauthorized, map[string]string{
					"error": "you are not authorized to post what I listen to, silly hacker",
				})
				return
			}
		}

		log.Println("validating submitted song")
		var newSong PlayedSong
		ctx.BindJSON(&newSong)
		if !validateSong(&newSong) {
			ctx.JSON(http.StatusBadRequest, map[string]string{
				"error": "song is not valid, so ignoring saving to database",
			})
			return
		}
		
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
