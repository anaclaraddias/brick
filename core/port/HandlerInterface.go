package port

import "github.com/gin-gonic/gin"

type HandlerInterface interface {
	Handle(context *gin.Context)
}
