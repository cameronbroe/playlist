package internal

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

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
