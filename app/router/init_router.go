package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/config"
	"log"
)

func Init(conf *config.AdminConf) {
	router := gin.Default()
	AdminRoutes(router)
	err := router.Run(":" + conf.Port)
	if err != nil {
		panic(err.Error())
		log.Fatalln()
	}
}
