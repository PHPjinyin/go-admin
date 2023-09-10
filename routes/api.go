package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SetApiGroupRoutes(routes *gin.RouterGroup) {
	routes.GET("ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})
	routes.GET("/test", func(c *gin.Context) {
		time.Sleep(15 * time.Second)
		c.String(http.StatusOK, "success")
	})

}
