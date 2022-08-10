package route_handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func NewListHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		fmt.Println("handling the /list endpoint")
	}
}
