package app

import (
	"github.com/gin-gonic/gin"
	"jin-gin/app/common/request"
	"jin-gin/app/common/response"
	"jin-gin/app/services"
)

// Login 登录
func Login(c *gin.Context) {
	var form request.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Login(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		tokenData, err, _ := services.JwtServices.CreateToken(services.AppGuardName, user)
		if err != nil {
			response.BusinessFail(c, err.Error())
			return
		}
		response.Success(c, tokenData)
	}
}

func AuthInfo(c *gin.Context) {
	err, user := services.UserService.Info(c.Keys["authId"].(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, user)
}
