package services

import (
	"errors"
	"jin-gin/app/common/request"
	"jin-gin/app/models"
	"jin-gin/global"
	"jin-gin/utils"
	"strconv"
)

type userService struct {
}

var UserService = new(userService)

// Register 注册
func (userService *userService) Register(params request.Register) (err error, user models.User) {
	var result = global.App.DB.Where("mobile = ? or username =?", params.Mobile, params.Name).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("手机号已存在或用户名已存在")
		return
	}
	user = models.User{Username: params.Name, Mobile: params.Mobile, Password: utils.BcryptDecode([]byte(params.Password))}
	err = global.App.DB.Create(&user).Error
	return
}

func (userService *userService) Login(params request.Login) (err error, user *models.User) {
	err = global.App.DB.Where("username = ?", params.Username).First(&user).Error
	if err != nil || !utils.BcryptCheck([]byte(params.Password), user.Password) {
		err = errors.New("用户名不存在或密码错误")
	}
	return
}

func (userService *userService) Info(ID string) (err error, user *models.User) {
	userId, err := strconv.Atoi(ID)
	err = global.App.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}
