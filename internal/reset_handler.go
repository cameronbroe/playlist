package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func NewResetHandler(db *Database) func(ctx *gin.Context) {
	staticApiSecret := os.Getenv("API_SECRET")

	return func(ctx *gin.Context) {
		if staticApiSecret != "" {
			bearerToken := ctx.GetHeader("Authorization")
			if bearerToken != fmt.Sprintf("Bearer %s", staticApiSecret) {
				ctx.JSON(http.StatusUnauthorized, map[string]string{
					"error": "you are not authorized to reset what I listen to, silly hacker",
				})
				return
			}
		}

		log.Println("resetting playlist")
		err := db.ResetPlaylist()
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
