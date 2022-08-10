package routes

import "github.com/gin-gonic/gin"

type ListHandler struct {}

func (handler *ListHandler) HandlerFunc(context *gin.Context)  {
	fmt.Println("handling the list route")
}