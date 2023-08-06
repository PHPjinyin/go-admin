package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LoginHandle(ctx *gin.Context) {
	fmt.Println("Router end")
	a, err := ctx.Get("name")

	if err != false {

	} else {
		fmt.Println(a)
	}
	ctx.JSON(200, gin.H{
		"message": "success",
	})
}
