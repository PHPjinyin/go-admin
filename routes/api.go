package routes

import (
	"github.com/gin-gonic/gin"
	"jin-gin/app/controllers/app"
	"jin-gin/app/services"
	"jin-gin/middleware"
)

func SetApiGroupRoutes(routes *gin.RouterGroup) {
	routes.POST("/auth/register", app.Register)
	routes.POST("/auth/login", app.Login)
	authRoutes := routes.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRoutes.GET("/auth/info", app.AuthInfo)
	}

}
