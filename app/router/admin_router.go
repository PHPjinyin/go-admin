package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin"
)

func AdminRoutes(router *gin.Engine) {
	adminRoute := router.Group("/admin")
	{
		adminRoute.GET("/login", admin.LoginHandle)
	}
}
