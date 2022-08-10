package routes

import "github.com/gin-gonic/gin"

type RouteHandler interface {
	HandlerFunc func(context *gin.Context)
}